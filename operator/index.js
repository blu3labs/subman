"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const ethers_1 = require("ethers");
const axios_1 = __importDefault(require("axios"));
const converter_1 = require("./converter/converter");
const subman_1 = require("./subman/subman");
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config();
const API_URL = process.env.API_URL || "";
function getExecuteableSubscriptions() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const response = yield axios_1.default.get(`${API_URL}/executeables`);
            let subs = [];
            if (response.status !== 200) {
                console.error("Error getting requests1", response.status);
                return subs;
            }
            if (response.data) {
                if (!Array.isArray(response.data)) {
                    console.error("Error getting requests2", response.data);
                    return subs;
                }
                for (const sub of response.data) {
                    subs.push((0, converter_1.ConvertToSubscription)(sub));
                }
            }
            return subs;
        }
        catch (error) {
            console.error("Error getting requests3", error);
            return [];
        }
    });
}
function main() {
    return __awaiter(this, void 0, void 0, function* () {
        console.log("Starting validator");
        const pvKey = process.env.PV_KEY || "";
        const wallet = new ethers_1.Wallet(pvKey);
        const walletAddress = wallet.address;
        while (true) {
            const before = Date.now();
            const subs = yield getExecuteableSubscriptions();
            console.log("subs", subs);
            if (subs.length > 0) {
                for (const sub of subs) {
                    const address = (0, subman_1.GetSubManAddress)(sub.ChainID);
                    const abi = (0, subman_1.GetSubManABI)();
                    const providerUrl = (0, subman_1.GetRPC)(sub.ChainID);
                    const provider = (0, ethers_1.getDefaultProvider)(providerUrl);
                    const walletClient = wallet.connect(provider);
                    const ctr = new ethers_1.Contract(address, abi, walletClient);
                    try {
                        const args = (0, converter_1.ConvertToTxArgs)(sub);
                        const tx = yield ctr.processPayment(args.subPayment, args.signature, walletAddress);
                        yield tx.wait();
                        console.log("tx success", tx.hash);
                    }
                    catch (error) {
                        console.error("tx error", error);
                    }
                }
            }
            const after = Date.now();
            const diff = after - before;
            const duration = 3000 - diff;
            if (duration > 0) {
                yield new Promise((resolve) => setTimeout(resolve, duration));
            }
        }
    });
}
main()
    .then(() => process.exit(0))
    .catch((error) => {
    console.error(error);
    process.exit(1);
});

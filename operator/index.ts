import ethers from "ethers";
import { Contract, Wallet, getDefaultProvider } from "ethers";
import axios from "axios";
import { ConvertToSubscription, ConvertToTxArgs, Subscription } from "./converter/converter";
import { GetSubManAddress, GetRPC, GetSubManABI } from "./subman/subman";

const API_URL = process.env.API_URL || "";

async function getExecuteableSubscriptions() {
  try {
    const response = await axios.get(`${API_URL}/executeables`);
    let subs: Subscription[] = [];
    if (response.status !== 200) {
      console.error("Error getting requests", response.status);
      return subs;
    }
    if (response.data) {
      if (!Array.isArray(response.data)) {
        console.error("Error getting requests", response.data);
        return subs;
      }
      for (const sub of response.data) {
        subs.push(ConvertToSubscription(sub));
      }
    }
    return subs;
  } catch (error) {
    console.error("Error getting requests", error);
    return [];
  }
}

async function main() {
  console.log("Starting validator");
  const pvKey = process.env.PV_KEY || "";
  const wallet = new Wallet(pvKey);
  const walletAddress = wallet.address;
  while (true) {
    const before = Date.now();
    const subs = await getExecuteableSubscriptions();
    console.log(`Found ${subs.length} subscriptions to execute`);
    if (subs.length > 0) {
      for (const sub of subs) {
        const address = GetSubManAddress(sub.ChainID);
        const abi = GetSubManABI();
        const providerUrl = GetRPC(sub.ChainID);
        const provider = getDefaultProvider(providerUrl);
        const walletClient = new Wallet(pvKey, provider);
        const ctr = new Contract(address, abi, wallet);
        try {
          const args = ConvertToTxArgs(sub);
          const tx = await ctr.executeSubscription(args.subPayment, args.signature, walletAddress);
          await tx.wait();
          console.log("tx success", tx.hash);
        } catch (error) {
          console.error("tx error", error);
        }
      }
    }
    const after = Date.now();
    const diff = after - before;
    const duration = 3000 - diff;
    if (duration > 0) {
      await new Promise((resolve) => setTimeout(resolve, duration));
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

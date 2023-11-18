package chains

import (
	"fmt"
	"log"
	"os"
	types "subman/types"

	"github.com/joho/godotenv"
)

var Chains = []types.Chain{}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	BASE_GOERLI_URL, success := os.LookupEnv("BASE_GOERLI_URL")
	if !success {
		panic("BASE_GOERLI_URL not found")
	}
	Chains = append(Chains, types.Chain{
		ChainID:  84531,
		Name:     "Base Goerli",
		RPC:      BASE_GOERLI_URL,
		Address:  "0x609165A00BC01405a48F044A29d9497030F8c2B7",
		Interval: 5,
		Confirms: 3,
	})
	Chains = append(Chains, types.Chain{
		ChainID:  534351,
		Name:     "Scroll Sepolia",
		RPC:      "https://rpc.ankr.com/scroll_sepolia_testnet",
		Address:  "0x609165A00BC01405a48F044A29d9497030F8c2B7",
		Interval: 5,
		Confirms: 3,
	})
}

func GetChain(chainId uint64) (types.Chain, error) {
	for _, chain := range Chains {
		if chain.ChainID == chainId {
			return chain, nil
		}
	}
	err := fmt.Errorf("chain not found")
	return types.Chain{}, err
}

package chains

import (
	"fmt"
	types "subman/types"
)

var Chains = []types.Chain{}

func init() {
	Chains = append(Chains, types.Chain{
		ChainID:  84531,
		Name:     "Base Goerli",
		RPC:      "https://base-goerli.public.blastapi.io",
		Address:  "0x000000",
		Interval: 5,
		Confirms: 3,
	})
	Chains = append(Chains, types.Chain{
		ChainID:  534351,
		Name:     "Scroll Sepolia",
		RPC:      "https://rpc.ankr.com/scroll_sepolia_testnet",
		Address:  "0x000000",
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

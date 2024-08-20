package main

import (
	"fmt"
	"log"

	"github.com/CornMars2020/bav"
)

func main() {
	addresses := []string{
		// mainnet
		"bc1qsh2f075nhs5ec6vzael7av30c24llwxrmdp7wkh6dtxnjgxfuk9sawzaxh", // p2wsh
		"bc1pwpc4py0rmkrxuy3gs2qxukfszr9tzrc4gznhc26r0t3f5h2qdmfsjggztf", // taproot
		"bc1qf7egt4sdvf44p4rlv2m29xvtnrwx6jzvy56zyw",                     // p2wpkh
		"3BxYc4apMpmiiKw4jNxAKvzxjWSFcfbaxH",                             // p2sh-p2wpkh
		"1KzdJzSAapfcrSxDx4HgQawENxtyTHbnen",                             // p2pkh

		"",

		// testnet3
		// "" // p2wsh
		"tb1pwpc4py0rmkrxuy3gs2qxukfszr9tzrc4gznhc26r0t3f5h2qdmfs9q7d3x", // taproot
		"tb1qf7egt4sdvf44p4rlv2m29xvtnrwx6jzvwjp3la",                     // p2wpkh
		"2N3WkfoWqyHH4v7ZcQWa2wszDwreRSSKDpX",                            // p2sh-p2wpkh
		"mzWac3X9Pr6sdZRqfdG4EW9ZExVgNEyMnA",                             // p2pkh

		"",

		// signet
		// "" // p2wsh
		"sb1pwpc4py0rmkrxuy3gs2qxukfszr9tzrc4gznhc26r0t3f5h2qdmfsz352gr", // taproot
		"sb1qf7egt4sdvf44p4rlv2m29xvtnrwx6jzvt425lz",                     // p2wpkh
		"rfqinskp97T7DUPHctFmYj7o12qZW5h4T8",                             // p2sh-p2wpkh
		"SgHdLqDKKBrpNkjgVVGkxV5o2k8QAXFomx",                             // p2pkh

		"",

		// regtest
		// "" // p2wsh
		"bcrt1pwpc4py0rmkrxuy3gs2qxukfszr9tzrc4gznhc26r0t3f5h2qdmfsge5tyu", // taproot
		"bcrt1qf7egt4sdvf44p4rlv2m29xvtnrwx6jzvvmcug5",                     // p2wpkh
		"2N3WkfoWqyHH4v7ZcQWa2wszDwreRSSKDpX",                              // p2sh-p2wpkh (same as testnet)
		"mzWac3X9Pr6sdZRqfdG4EW9ZExVgNEyMnA",                               // p2pkh (same as testnet)
	}

	for _, addr := range addresses {
		if addr == "" {
			fmt.Println("")
			continue
		}
		info := bav.GetAddressInfo(addr)
		log.Printf("Address: %s, Type: %s, Network: %s, IsBench32: %t\n", info.Address, info.Type, info.Network, info.IsBench32)
	}
}

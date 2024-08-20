package bav
 
import (
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

type AddressType string

const (
	P2PKH                AddressType = "P2PKH"       // legacy
	P2SH_P2WPKH          AddressType = "P2SH-P2WPKH" // nested segwit
	P2WPKH               AddressType = "P2WPKH"      // native segwit
	P2TR                 AddressType = "P2TR"        // taproot
	P2SH                 AddressType = "P2SH"        // unlock by script hash
	P2WSH                AddressType = "P2WSH"       // multisig
	UNKNOWN_ADDRESS_TYPE AddressType = "Unknown"
)

type NetworkType string

const (
	MainNet              NetworkType = "mainnet"
	TestNet              NetworkType = "testnet"
	Signet               NetworkType = "signet"
	Regtest              NetworkType = "regtest"
	UNKNOWN_NETWORK_TYPE NetworkType = "Unknown"
)

type AddressInfo struct {
	IsBench32 bool        `json:"is_bench32" mapstructure:"is_bench32"`
	Address   string      `json:"address" mapstructure:"address"`
	Network   NetworkType `json:"network" mapstructure:"network"`
	Type      AddressType `json:"type" mapstructure:"type"`
}

func GetNetworkType(addr string) (btcutil.Address, NetworkType) {
	var network NetworkType

	address, err := btcutil.DecodeAddress(addr, &chaincfg.MainNetParams)
	if err == nil && address.EncodeAddress() == addr {
		if strings.HasPrefix(addr, "bcrt") {
			network = Regtest
		} else if strings.HasPrefix(addr, "bc") {
			network = MainNet
		} else if strings.HasPrefix(addr, "tb") {
			network = TestNet
		} else if strings.HasPrefix(addr, "sb") {
			network = Signet
		} else {
			network = MainNet
		}
		return address, network
	}

	address, err = btcutil.DecodeAddress(addr, &chaincfg.TestNet3Params)
	if err == nil && address.EncodeAddress() == addr {
		network = TestNet
		return address, network
	}

	address, err = btcutil.DecodeAddress(addr, &chaincfg.SimNetParams)
	if err == nil && address.EncodeAddress() == addr {
		network = Signet
		return address, network
	}

	address, err = btcutil.DecodeAddress(addr, &chaincfg.RegressionNetParams)
	if err == nil && address.EncodeAddress() == addr {
		network = Regtest
		return address, network
	}

	network = UNKNOWN_NETWORK_TYPE
	return address, network
}

func GetAddressInfo(addr string) AddressInfo {
	address, network := GetNetworkType(addr)

	var addrType AddressType

	switch address.(type) {
	case *btcutil.AddressPubKeyHash:
		addrType = P2PKH
	case *btcutil.AddressWitnessPubKeyHash:
		addrType = P2WPKH
	case *btcutil.AddressWitnessScriptHash:
		addrType = P2WSH
	case *btcutil.AddressTaproot:
		addrType = P2TR
	case *btcutil.AddressScriptHash:
		addrType = P2SH
		// TODO: support P2SH-P2WPKH
		// addrType = P2SH_P2WPKH
	default:
		addrType = UNKNOWN_ADDRESS_TYPE
	}

	return AddressInfo{
		IsBench32: strings.HasPrefix(addr, "bc") || strings.HasPrefix(addr, "tb") || strings.HasPrefix(addr, "sb"),
		Address:   addr,
		Network:   network,
		Type:      addrType,
	}
}

func Validate(addr string, defaultNetwork NetworkType) bool {
	addressInfo := GetAddressInfo(addr)

	if addressInfo.Type == "" ||
		addressInfo.Type == UNKNOWN_ADDRESS_TYPE ||
		addressInfo.Network == "" ||
		addressInfo.Network == UNKNOWN_NETWORK_TYPE {
		return false
	}

	if defaultNetwork != UNKNOWN_NETWORK_TYPE && defaultNetwork != "" {
		if addressInfo.Network != defaultNetwork {
			return false
		}
	}

	return true
}

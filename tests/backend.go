package tests

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"os"
)

var (
	TestEthProviderUrl  = os.Getenv("TEST_ETH_PROVIDER_URL")
	TestProviderUrl     = os.Getenv("TEST_PROVIDER_URL")
	defaultAlloc        = big.NewInt(1e15)
	defaultGenesisAlloc = core.GenesisAlloc{
		Account0:  core.GenesisAccount{Balance: defaultAlloc},
		Account1:  core.GenesisAccount{Balance: defaultAlloc},
		Account2:  core.GenesisAccount{Balance: defaultAlloc},
		Account3:  core.GenesisAccount{Balance: defaultAlloc},
		Account4:  core.GenesisAccount{Balance: defaultAlloc},
		Account5:  core.GenesisAccount{Balance: defaultAlloc},
		Account6:  core.GenesisAccount{Balance: defaultAlloc},
		Account7:  core.GenesisAccount{Balance: defaultAlloc},
		Account8:  core.GenesisAccount{Balance: defaultAlloc},
		Account9:  core.GenesisAccount{Balance: defaultAlloc},
		Account10: core.GenesisAccount{Balance: defaultAlloc},
	}
)

func NewFakeEthBackend() *backends.SimulatedBackend {
	defaultAlloc.SetString("1000000000000000000000", 10)
	return backends.NewSimulatedBackend(defaultGenesisAlloc, 4712388)
}

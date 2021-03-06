package core

import (
	"bytes"
	"encoding/gob"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rovergulf/chain/core/types"
	"github.com/rovergulf/chain/params"
	"math/big"
)

// Genesis represents BlockChain initialization state
// and provides its root state for new nodes initialization
type Genesis struct {
	ChainId     *big.Int       `json:"chain_id" yaml:"chain_id"`
	GenesisTime int64          `json:"genesis_time" yaml:"genesis_time"`
	NetherPrice uint64         `json:"nether_limit" yaml:"nether_limit"`
	Nonce       uint64         `json:"nonce" yaml:"nonce"`
	Coinbase    common.Address `json:"coinbase" yaml:"coinbase"`
	Symbol      string         `json:"symbol" yaml:"symbol"`
	Units       string         `json:"units" yaml:"units"`
	ParentHash  common.Hash    `json:"parent_hash" yaml:"parent_hash"`
	Alloc       genesisAlloc   `json:"alloc" yaml:"alloc"`
	ExtraData   []byte         `json:"extra_data,omitempty" yaml:"extra_data,omitempty"`
}

// DevNetGenesis returns default Genesis for development and testing network
func DevNetGenesis() *Genesis {
	return &Genesis{
		ChainId:     big.NewInt(params.OpenDevNetworkId),
		GenesisTime: 1625422671,
		NetherPrice: 21000,
		Nonce:       0,
		Coinbase:    common.HexToAddress("0x01"),
		Symbol:      "Nether",
		Units:       "Wei", // in favor of Etherium native denomination
		ParentHash:  common.Hash{},
		Alloc:       developerNetAlloc(),
		ExtraData:   []byte{},
	}
}

// DefaultMainNetGenesis returns default Genesis for main Rovergulf BlockChain Network
func DefaultMainNetGenesis() *Genesis {
	return &Genesis{
		ChainId:     big.NewInt(params.MainNetworkId),
		GenesisTime: 1625422671,
		NetherPrice: 21000,
		Nonce:       0,
		Coinbase:    common.HexToAddress("0x3c0b3b41a1e027d3E759612Af08844f1cca0DdE3"),
		Symbol:      "Coin",   // Rovergulf Coin
		Units:       "Nether", // is like it powered by atoms or quantum
		Alloc:       defaultMainNetAlloc(),
		ExtraData:   []byte{},
	}
}

// Serialize binary encodes genesis
func (g Genesis) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(g); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

// Deserialize decodes binary value to genesis
func (g *Genesis) Deserialize(data []byte) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(g)
}

func (g *Genesis) ToBlock() (*types.Block, error) {
	var txs []*types.SignedTx

	for addr := range g.Alloc {
		alloc := g.Alloc[addr]
		tx, err := types.NewTransaction(g.Coinbase, addr, alloc.Balance, 0, g.ExtraData)
		if err != nil {
			return nil, err
		}
		tx.Time = g.GenesisTime

		txs = append(txs, &types.SignedTx{Transaction: tx})
	}

	header := types.BlockHeader{
		PrevHash:  g.ParentHash,
		Number:    g.Nonce,
		Timestamp: g.GenesisTime,
		Coinbase:  g.Coinbase,
	}

	b := types.NewBlock(header, txs)
	hash, err := b.Hash()
	if err != nil {
		return nil, err
	}
	b.BlockHeader.BlockHash = common.BytesToHash(hash)

	txHash, err := b.HashTransactions()
	if err != nil {
		return nil, err
	}
	b.TxHash = common.BytesToHash(txHash)

	return b, nil
}

func genesisByNetworkId(networkId *big.Int) *Genesis {
	switch networkId.Int64() {
	case params.OpenDevNetworkId:
		return DevNetGenesis()
	case params.MainNetworkId:
		return DefaultMainNetGenesis()
	default:
		return DefaultMainNetGenesis()
	}
}

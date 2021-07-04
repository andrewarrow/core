package main

import (
	"fmt"
)

func main() {
	gb := MsgBitCloutBlock{}
	fmt.Println("gb", gb)
}

type MsgBitCloutBlock struct {
	Header            *MsgBitCloutHeader
	Txns              []*MsgBitCloutTxn
	BlockProducerInfo *BlockProducerInfo
}

type BlockProducerInfo struct {
	PublicKey []byte
	Signature string
}

type BitCloutOutput struct {
	PublicKey   []byte
	AmountNanos uint64
}

type UtxoKey struct {
	TxID  BlockHash
	Index uint32
}

type MsgBitCloutTxn struct {
	TxInputs    []*UtxoKey
	TxOutputs   []*BitCloutOutput
	TxnMeta     string //BitCloutTxnMetadata
	PublicKey   []byte
	ExtraData   map[string][]byte
	Signature   string
	TxnTypeJSON uint64
}

type BlockHash [32]byte

type MsgBitCloutHeader struct {
	Version               uint32
	PrevBlockHash         *BlockHash
	TransactionMerkleRoot *BlockHash
	TstampSecs            uint64
	Height                uint64
	Nonce                 uint64
	ExtraNonce            uint64
}

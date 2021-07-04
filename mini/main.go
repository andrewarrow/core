package main

import (
	"fmt"
)

func main() {
	aMerkleRoot := BlockHash{1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		1, 2}
	gb := MsgBitCloutBlock{
		Header: &MsgBitCloutHeader{
			Version:               0,
			PrevBlockHash:         &BlockHash{},
			TransactionMerkleRoot: &aMerkleRoot,
			TstampSecs:            uint64(1610948544),
			Height:                uint64(0),
			Nonce:                 uint64(0),
		}}
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

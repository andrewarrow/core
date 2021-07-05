package main

import (
	"encoding/json"
	"fmt"
	"math/big"
)

func main() {
	SeedBalances := []*BitCloutOutput{
		{
			PublicKey:   []byte{1, 2, 3, 4, 5},
			AmountNanos: uint64(177204235393800),
		},
		{
			PublicKey:   []byte{6, 7, 8, 9, 10},
			AmountNanos: uint64(2662447206500),
		},
	}
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
		},
		Txns: []*MsgBitCloutTxn{
			{
				TxInputs:  []*BitCloutInput{},
				TxOutputs: SeedBalances,
				TxnMeta:   "They came here, to the New World. World 2.0, version 1776.",
			},
		},
	}
	fmt.Println("gb", gb)
	diffTarget := &aMerkleRoot
	blockHash := &aMerkleRoot
	genesisNode := NewBlockNode(
		nil, // Parent
		blockHash,
		0, // Height
		diffTarget,
		big.NewInt(1),
		gb.Header, // Header
		11111,     // Status
	)
	fmt.Println("gn", genesisNode)
	PutBestHash(blockHash)
	PutBlock(&gb)
	PutHeightHashToNodeInfo(genesisNode)
	//DbPutNanosPurchased(uint64(6000000000000000))
	//DbPutGlobalParamsEntry(InitialGlobalParamsEntry)
	utxoView := NewUtxoView()
	for index, txOutput := range SeedBalances {
		outputKey := UtxoKey{
			TxID:  BlockHash{},
			Index: uint32(index),
		}
		utxoEntry := UtxoEntry{
			AmountNanos: txOutput.AmountNanos,
			PublicKey:   txOutput.PublicKey,
			BlockHeight: 0,
			UtxoType:    UtxoTypeOutput,
			UtxoKey:     &outputKey,
		}

		utxoView._addUtxo(&utxoEntry)
	}
	fmt.Println(utxoView)

}

func PutBestHash(bh *BlockHash) {
	// prefix 3
	fmt.Println("PutBestHash")
	key := "3"
	b, _ := json.Marshal(bh)
	value := string(b)
	fmt.Println(key)
	fmt.Println(value)
}

func PutBlock(bitcloutBlock *MsgBitCloutBlock) {
	b, _ := json.Marshal(bitcloutBlock.Header)
	key := string(b)
	b, _ = json.Marshal(bitcloutBlock)
	value := string(b)
	fmt.Println("PutBlock")
	fmt.Println(key)
	fmt.Println(value)
}

func PutHeightHashToNodeInfo(node *BlockNode) {
	m := map[string]interface{}{}
	m["Height"] = node.Height
	m["Hash"] = node.Hash
	b, _ := json.Marshal(m)
	key := string(b)
	b, _ = json.Marshal(node)
	value := string(b)
	fmt.Println("PutHeightHashToNodeInfo")
	fmt.Println(key)
	fmt.Println(value)
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

type BitCloutInput UtxoKey

type MsgBitCloutTxn struct {
	TxInputs    []*BitCloutInput
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

type BlockStatus uint32

type BlockNode struct {
	Parent           *BlockNode
	Hash             *BlockHash
	Height           uint32
	DifficultyTarget *BlockHash
	CumWork          *big.Int
	Header           *MsgBitCloutHeader
	Status           BlockStatus
}

func NewBlockNode(
	parent *BlockNode,
	hash *BlockHash,
	height uint32,
	difficultyTarget *BlockHash,
	cumWork *big.Int,
	header *MsgBitCloutHeader,
	status BlockStatus) *BlockNode {

	return &BlockNode{
		Parent:           parent,
		Hash:             hash,
		Height:           height,
		DifficultyTarget: difficultyTarget,
		CumWork:          cumWork,
		Header:           header,
		Status:           status,
	}
}

type OperationType uint

const (
	UtxoTypeOutput       UtxoType      = 0
	OperationTypeAddUtxo OperationType = 0
)

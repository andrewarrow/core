package main

import (
	"encoding/json"
)

var BestHash *BlockHash
var KeyVal = map[string]string{}

func DbGetBestHash() *BlockHash {
	return BestHash
}

func PutBestHash(bh *BlockHash) {
	BestHash = bh
}

func PutBlock(bitcloutBlock *MsgBitCloutBlock) {
	b, _ := json.Marshal(bitcloutBlock.Header)
	key := string(b)
	b, _ = json.Marshal(bitcloutBlock)
	value := string(b)
	KeyVal[key] = value
}

func PutHeightHashToNodeInfo(node *BlockNode) {
	m := map[string]interface{}{}
	m["Height"] = node.Height
	m["Hash"] = node.Hash
	b, _ := json.Marshal(m)
	key := string(b)
	b, _ = json.Marshal(node)
	value := string(b)
	KeyVal[key] = value
}

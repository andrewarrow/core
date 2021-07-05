package main

var BestHash *BlockHash

func DbGetBestHash() *BlockHash {
	return BestHash
}

func PutBestHash(bh *BlockHash) {
	BestHash = bh
}

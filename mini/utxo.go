package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func NewUtxoView() *UtxoView {

	aMerkleRoot := BlockHash{1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		1, 2}
	view := UtxoView{
		TipHash: &aMerkleRoot, //DbGetBestHash(),
	}
	view._ResetViewMappingsAfterFlush()

	return &view
}

type UtxoKey struct {
	TxID  BlockHash
	Index uint32
}

type UtxoType uint8

type UtxoEntry struct {
	AmountNanos uint64
	PublicKey   []byte
	BlockHeight uint32
	UtxoType    UtxoType
	isSpent     bool
	UtxoKey     *UtxoKey
}

type UtxoView struct {
	NumUtxoEntries     uint64
	UtxoKeyToUtxoEntry map[UtxoKey]*UtxoEntry
	NanosPurchased     uint64
	USDCentsPerBitcoin uint64
	//MessageKeyToMessageEntry              map[MessageKey]*MessageEntry
	//PostHashToPostEntry                   map[BlockHash]*PostEntry
	//PublicKeyToPKIDEntry                  map[PkMapKey]*PKIDEntry
	//PKIDToPublicKey                       map[PKID]*PKIDEntry
	//ProfilePKIDToProfileEntry             map[PKID]*ProfileEntry
	//ProfileUsernameToProfileEntry         map[UsernameMapKey]*ProfileEntry
	TipHash *BlockHash
}

func (bav *UtxoView) _addUtxo(utxoEntryy *UtxoEntry) (*UtxoOperation, error) {
	utxoEntryCopy := *utxoEntryy
	if utxoEntryCopy.UtxoKey == nil {
		return nil, fmt.Errorf("_addUtxo: utxoEntry must have utxoKey set")
	}
	if utxoEntryCopy.isSpent {
		return nil, fmt.Errorf("_addUtxo: UtxoEntry being added has isSpent = true")
	}
	if err := bav._setUtxoMappings(&utxoEntryCopy); err != nil {
		return nil, errors.Wrapf(err, "_addUtxo: ")
	}

	bav.NumUtxoEntries++

	return &UtxoOperation{
		Type:  OperationTypeAddUtxo,
		Key:   utxoEntryCopy.UtxoKey,
		Entry: &utxoEntryCopy,
	}, nil
}
func (bav *UtxoView) _setUtxoMappings(utxoEntry *UtxoEntry) error {
	if utxoEntry.UtxoKey == nil {
		return fmt.Errorf("_setUtxoMappings: utxoKey missing for utxoEntry %+v", utxoEntry)
	}

	bav.UtxoKeyToUtxoEntry[*utxoEntry.UtxoKey] = utxoEntry

	return nil
}

type UtxoOperation struct {
	Type                   OperationType
	Entry                  *UtxoEntry
	Key                    *UtxoKey
	PrevNanosPurchased     uint64
	PrevUSDCentsPerBitcoin uint64
	//PrevPostEntry              *PostEntry
	//PrevParentPostEntry        *PostEntry
	//PrevGrandparentPostEntry   *PostEntry
	//PrevRecloutedPostEntry     *PostEntry
	//PrevProfileEntry           *ProfileEntry
	//PrevLikeEntry              *LikeEntry
	PrevLikeCount uint64
	//PrevDiamondEntry           *DiamondEntry
	//PrevRecloutEntry           *RecloutEntry
	PrevRecloutCount uint64
	//PrevCoinEntry              *CoinEntry
	//PrevTransactorBalanceEntry *BalanceEntry
	//PrevCreatorBalanceEntry    *BalanceEntry
	FounderRewardUtxoKey *UtxoKey
	//PrevSenderBalanceEntry     *BalanceEntry
	//PrevReceiverBalanceEntry   *BalanceEntry
	//PrevGlobalParamsEntry      *GlobalParamsEntry
	//PrevForbiddenPubKeyEntry   *ForbiddenPubKeyEntry
}

func (bav *UtxoView) _ResetViewMappingsAfterFlush() {
	bav.UtxoKeyToUtxoEntry = make(map[UtxoKey]*UtxoEntry)
	bav.NumUtxoEntries = 0 //GetUtxoNumEntries(bav.Handle)
	bav.NanosPurchased = 0 //DbGetNanosPurchased(bav.Handle)
}

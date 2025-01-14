package inx

import (
	iotago "github.com/iotaledger/iota.go/v3"
)

// nolint:revive,stylecheck // this name is auto generated
func NewBlockId(blockID iotago.BlockID) *BlockId {
	id := &BlockId{
		Id: make([]byte, iotago.BlockIDLength),
	}
	copy(id.Id, blockID[:])

	return id
}

func NewBlockIds(blockIDs iotago.BlockIDs) []*BlockId {
	result := make([]*BlockId, len(blockIDs))
	for i := range blockIDs {
		result[i] = NewBlockId(blockIDs[i])
	}

	return result
}

func NewBlockWithBytes(blockID iotago.BlockID, data []byte) *Block {
	m := &Block{
		BlockId: NewBlockId(blockID),
		Block: &RawBlock{
			Data: make([]byte, len(data)),
		},
	}
	copy(m.Block.Data, data)

	return m
}

// nolint:revive,stylecheck // this name is auto generated
func NewTransactionId(transactionID iotago.TransactionID) *TransactionId {
	id := &TransactionId{
		Id: make([]byte, iotago.TransactionIDLength),
	}
	copy(id.Id, transactionID[:])

	return id
}

// nolint:revive,stylecheck // this name is auto generated
func NewOutputId(outputID iotago.OutputID) *OutputId {
	id := &OutputId{
		Id: make([]byte, iotago.OutputIDLength),
	}
	copy(id.Id, outputID[:])

	return id
}

// nolint:revive,stylecheck // this name is auto generated
func NewMilestoneId(milestoneID iotago.MilestoneID) *MilestoneId {
	id := &MilestoneId{
		Id: make([]byte, iotago.MilestoneIDLength),
	}
	copy(id.Id, milestoneID[:])

	return id
}

func NewMilestoneInfo(milestoneID iotago.MilestoneID, index uint32, timestamp uint32) *MilestoneInfo {
	return &MilestoneInfo{
		MilestoneId:        NewMilestoneId(milestoneID),
		MilestoneIndex:     index,
		MilestoneTimestamp: timestamp,
	}
}

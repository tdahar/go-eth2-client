// Code generated by fastssz. DO NOT EDIT.
// Hash: 6e158ecacc56184121edf05b279968e58e6bfcd243916f960406221c95eba882
package bellatrix

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadHeader object to a target array
func (e *ExecutionPayloadHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(536)

	// Field (0) 'ParentHash'
	dst = append(dst, e.ParentHash[:]...)

	// Field (1) 'FeeRecipient'
	dst = append(dst, e.FeeRecipient[:]...)

	// Field (2) 'StateRoot'
	dst = append(dst, e.StateRoot[:]...)

	// Field (3) 'ReceiptsRoot'
	dst = append(dst, e.ReceiptsRoot[:]...)

	// Field (4) 'LogsBloom'
	dst = append(dst, e.LogsBloom[:]...)

	// Field (5) 'PrevRandao'
	dst = append(dst, e.PrevRandao[:]...)

	// Field (6) 'BlockNumber'
	dst = ssz.MarshalUint64(dst, e.BlockNumber)

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, e.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, e.GasUsed)

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, e.Timestamp)

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	dst = append(dst, e.BaseFeePerGas[:]...)

	// Field (12) 'BlockHash'
	dst = append(dst, e.BlockHash[:]...)

	// Field (13) 'TransactionsRoot'
	dst = append(dst, e.TransactionsRoot[:]...)

	// Field (10) 'ExtraData'
	if len(e.ExtraData) > 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, e.ExtraData...)

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 536 {
		return ssz.ErrSize
	}

	tail := buf
	var o10 uint64

	// Field (0) 'ParentHash'
	copy(e.ParentHash[:], buf[0:32])

	// Field (1) 'FeeRecipient'
	copy(e.FeeRecipient[:], buf[32:52])

	// Field (2) 'StateRoot'
	copy(e.StateRoot[:], buf[52:84])

	// Field (3) 'ReceiptsRoot'
	copy(e.ReceiptsRoot[:], buf[84:116])

	// Field (4) 'LogsBloom'
	copy(e.LogsBloom[:], buf[116:372])

	// Field (5) 'PrevRandao'
	copy(e.PrevRandao[:], buf[372:404])

	// Field (6) 'BlockNumber'
	e.BlockNumber = ssz.UnmarshallUint64(buf[404:412])

	// Field (7) 'GasLimit'
	e.GasLimit = ssz.UnmarshallUint64(buf[412:420])

	// Field (8) 'GasUsed'
	e.GasUsed = ssz.UnmarshallUint64(buf[420:428])

	// Field (9) 'Timestamp'
	e.Timestamp = ssz.UnmarshallUint64(buf[428:436])

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 536 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	copy(e.BaseFeePerGas[:], buf[440:472])

	// Field (12) 'BlockHash'
	copy(e.BlockHash[:], buf[472:504])

	// Field (13) 'TransactionsRoot'
	copy(e.TransactionsRoot[:], buf[504:536])

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) SizeSSZ() (size int) {
	size = 536

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadHeader object with a hasher
func (e *ExecutionPayloadHeader) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	hh.PutBytes(e.ParentHash[:])

	// Field (1) 'FeeRecipient'
	hh.PutBytes(e.FeeRecipient[:])

	// Field (2) 'StateRoot'
	hh.PutBytes(e.StateRoot[:])

	// Field (3) 'ReceiptsRoot'
	hh.PutBytes(e.ReceiptsRoot[:])

	// Field (4) 'LogsBloom'
	hh.PutBytes(e.LogsBloom[:])

	// Field (5) 'PrevRandao'
	hh.PutBytes(e.PrevRandao[:])

	// Field (6) 'BlockNumber'
	hh.PutUint64(e.BlockNumber)

	// Field (7) 'GasLimit'
	hh.PutUint64(e.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(e.GasUsed)

	// Field (9) 'Timestamp'
	hh.PutUint64(e.Timestamp)

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(e.ExtraData)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
	}

	// Field (11) 'BaseFeePerGas'
	hh.PutBytes(e.BaseFeePerGas[:])

	// Field (12) 'BlockHash'
	hh.PutBytes(e.BlockHash[:])

	// Field (13) 'TransactionsRoot'
	hh.PutBytes(e.TransactionsRoot[:])

	hh.Merkleize(indx)
	return
}

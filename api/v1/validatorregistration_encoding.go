package v1

import (
	"time"

	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ValidatorRegistration object
func (v *ValidatorRegistration) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the ValidatorRegistration object to a target array
func (v *ValidatorRegistration) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'FeeRecipient'
	if len(v.FeeRecipient) != 20 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, v.FeeRecipient[:]...)

	// Field (1) 'GasLimit'
	dst = ssz.MarshalUint64(dst, v.GasLimit)

	// Field (2) 'Timestamp'
	dst = ssz.MarshalUint64(dst, uint64(v.Timestamp.Unix()))

	// Field (3) 'Pubkey'
	if len(v.Pubkey) != 48 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, v.Pubkey[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ValidatorRegistration object
func (v *ValidatorRegistration) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 84 {
		return ssz.ErrSize
	}

	// Field (0) 'FeeRecipient'
	copy(v.FeeRecipient[:], buf[0:20])

	// Field (1) 'GasLimit'
	v.GasLimit = ssz.UnmarshallUint64(buf[20:28])

	// Field (2) 'Timestamp'
	v.Timestamp = time.Unix(int64(ssz.UnmarshallUint64(buf[28:36])), 0)

	// Field (3) 'Pubkey'
	copy(v.Pubkey[:], buf[36:84])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ValidatorRegistration object
func (v *ValidatorRegistration) SizeSSZ() (size int) {
	size = 84
	return
}

// HashTreeRoot ssz hashes the ValidatorRegistration object
func (v *ValidatorRegistration) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the ValidatorRegistration object with a hasher
func (v *ValidatorRegistration) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'FeeRecipient'
	if len(v.FeeRecipient) != 20 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(v.FeeRecipient[:])

	// Field (1) 'GasLimit'
	hh.PutUint64(v.GasLimit)

	// Field (2) 'Timestamp'
	hh.PutUint64(uint64(v.Timestamp.Unix()))

	// Field (3) 'Pubkey'
	if len(v.Pubkey) != 48 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(v.Pubkey[:])

	hh.Merkleize(indx)
	return
}

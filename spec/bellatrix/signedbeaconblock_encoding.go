// Code generated by fastssz. DO NOT EDIT.
// Hash: ecb9e3c38db0485fde133a10fb94c4de735236d4a5e9305d1e6efe2953cb285b
package bellatrix

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedBeaconBlock object
func (s *SignedBeaconBlock) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedBeaconBlock object to a target array
func (s *SignedBeaconBlock) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(100)

	// Offset (0) 'Message'
	dst = ssz.WriteOffset(dst, offset)
	if s.Message == nil {
		s.Message = new(BeaconBlock)
	}
	offset += s.Message.SizeSSZ()

	// Field (1) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Field (0) 'Message'
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedBeaconBlock object
func (s *SignedBeaconBlock) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 100 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Message'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 100 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Signature'
	copy(s.Signature[:], buf[4:100])

	// Field (0) 'Message'
	{
		buf = tail[o0:]
		if s.Message == nil {
			s.Message = new(BeaconBlock)
		}
		if err = s.Message.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBeaconBlock object
func (s *SignedBeaconBlock) SizeSSZ() (size int) {
	size = 100

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(BeaconBlock)
	}
	size += s.Message.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the SignedBeaconBlock object
func (s *SignedBeaconBlock) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedBeaconBlock object with a hasher
func (s *SignedBeaconBlock) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package abi

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *PieceInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Size (abi.PaddedPieceSize) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Size))); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	return nil
}

func (t *PieceInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Size (abi.PaddedPieceSize) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Size = PaddedPieceSize(extra)

	}
	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	return nil
}

func (t *SectorID) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Miner (abi.ActorID) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Miner))); err != nil {
		return err
	}

	// t.Number (abi.SectorNumber) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Number))); err != nil {
		return err
	}

	return nil
}

func (t *SectorID) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (abi.ActorID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Miner = ActorID(extra)

	}
	// t.Number (abi.SectorNumber) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Number = SectorNumber(extra)

	}
	return nil
}

func (t *SectorInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.SealProof (abi.RegisteredSealProof) (int64)
	if t.SealProof >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SealProof))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SealProof)-1)); err != nil {
			return err
		}
	}

	// t.SectorNumber (abi.SectorNumber) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorNumber))); err != nil {
		return err
	}

	// t.SealedCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.SealedCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.SealedCID: %w", err)
	}

	return nil
}

func (t *SectorInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.SealProof (abi.RegisteredSealProof) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SealProof = RegisteredSealProof(extraI)
	}
	// t.SectorNumber (abi.SectorNumber) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.SectorNumber = SectorNumber(extra)

	}
	// t.SealedCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SealedCID: %w", err)
		}

		t.SealedCID = c

	}
	return nil
}

func (t *SealVerifyInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{136}); err != nil {
		return err
	}

	// t.SealProof (abi.RegisteredSealProof) (int64)
	if t.SealProof >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SealProof))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SealProof)-1)); err != nil {
			return err
		}
	}

	// t.SectorID (abi.SectorID) (struct)
	if err := t.SectorID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.DealIDs ([]abi.DealID) (slice)
	if len(t.DealIDs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.DealIDs was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.DealIDs)))); err != nil {
		return err
	}
	for _, v := range t.DealIDs {
		if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, uint64(v)); err != nil {
			return err
		}
	}

	// t.Randomness (abi.SealRandomness) (slice)
	if len(t.Randomness) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Randomness was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Randomness)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Randomness); err != nil {
		return err
	}

	// t.InteractiveRandomness (abi.InteractiveSealRandomness) (slice)
	if len(t.InteractiveRandomness) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.InteractiveRandomness was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.InteractiveRandomness)))); err != nil {
		return err
	}
	if _, err := w.Write(t.InteractiveRandomness); err != nil {
		return err
	}

	// t.Proof ([]uint8) (slice)
	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Proof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Proof); err != nil {
		return err
	}

	// t.SealedCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.SealedCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.SealedCID: %w", err)
	}

	// t.UnsealedCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.UnsealedCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.UnsealedCID: %w", err)
	}

	return nil
}

func (t *SealVerifyInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 8 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.SealProof (abi.RegisteredSealProof) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SealProof = RegisteredSealProof(extraI)
	}
	// t.SectorID (abi.SectorID) (struct)

	{

		if err := t.SectorID.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.SectorID: %w", err)
		}

	}
	// t.DealIDs ([]abi.DealID) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.DealIDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.DealIDs = make([]DealID, extra)
	}

	for i := 0; i < int(extra); i++ {

		maj, val, err := cbg.CborReadHeader(br)
		if err != nil {
			return xerrors.Errorf("failed to read uint64 for t.DealIDs slice: %w", err)
		}

		if maj != cbg.MajUnsignedInt {
			return xerrors.Errorf("value read for array t.DealIDs was not a uint, instead got %d", maj)
		}

		t.DealIDs[i] = DealID(val)
	}

	// t.Randomness (abi.SealRandomness) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Randomness: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Randomness = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Randomness); err != nil {
		return err
	}
	// t.InteractiveRandomness (abi.InteractiveSealRandomness) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.InteractiveRandomness: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.InteractiveRandomness = make([]byte, extra)
	if _, err := io.ReadFull(br, t.InteractiveRandomness); err != nil {
		return err
	}
	// t.Proof ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Proof = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Proof); err != nil {
		return err
	}
	// t.SealedCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SealedCID: %w", err)
		}

		t.SealedCID = c

	}
	// t.UnsealedCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.UnsealedCID: %w", err)
		}

		t.UnsealedCID = c

	}
	return nil
}

func (t *PoStProof) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.PoStProof (abi.RegisteredPoStProof) (int64)
	if t.PoStProof >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.PoStProof))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.PoStProof)-1)); err != nil {
			return err
		}
	}

	// t.ProofBytes ([]uint8) (slice)
	if len(t.ProofBytes) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.ProofBytes was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.ProofBytes)))); err != nil {
		return err
	}
	if _, err := w.Write(t.ProofBytes); err != nil {
		return err
	}
	return nil
}

func (t *PoStProof) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PoStProof (abi.RegisteredPoStProof) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.PoStProof = RegisteredPoStProof(extraI)
	}
	// t.ProofBytes ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.ProofBytes: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.ProofBytes = make([]byte, extra)
	if _, err := io.ReadFull(br, t.ProofBytes); err != nil {
		return err
	}
	return nil
}

func (t *WindowPoStVerifyInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.Randomness (abi.PoStRandomness) (slice)
	if len(t.Randomness) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Randomness was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Randomness)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Randomness); err != nil {
		return err
	}

	// t.Proofs ([]abi.PoStProof) (slice)
	if len(t.Proofs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Proofs was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Proofs)))); err != nil {
		return err
	}
	for _, v := range t.Proofs {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.ChallengedSectors ([]abi.SectorInfo) (slice)
	if len(t.ChallengedSectors) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.ChallengedSectors was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.ChallengedSectors)))); err != nil {
		return err
	}
	for _, v := range t.ChallengedSectors {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Prover (abi.ActorID) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Prover))); err != nil {
		return err
	}

	return nil
}

func (t *WindowPoStVerifyInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Randomness (abi.PoStRandomness) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Randomness: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Randomness = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Randomness); err != nil {
		return err
	}
	// t.Proofs ([]abi.PoStProof) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Proofs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Proofs = make([]PoStProof, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v PoStProof
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Proofs[i] = v
	}

	// t.ChallengedSectors ([]abi.SectorInfo) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.ChallengedSectors: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.ChallengedSectors = make([]SectorInfo, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v SectorInfo
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.ChallengedSectors[i] = v
	}

	// t.Prover (abi.ActorID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Prover = ActorID(extra)

	}
	return nil
}

func (t *WinningPoStVerifyInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.Randomness (abi.PoStRandomness) (slice)
	if len(t.Randomness) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Randomness was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Randomness)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Randomness); err != nil {
		return err
	}

	// t.Proofs ([]abi.PoStProof) (slice)
	if len(t.Proofs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Proofs was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Proofs)))); err != nil {
		return err
	}
	for _, v := range t.Proofs {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.ChallengedSectors ([]abi.SectorInfo) (slice)
	if len(t.ChallengedSectors) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.ChallengedSectors was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.ChallengedSectors)))); err != nil {
		return err
	}
	for _, v := range t.ChallengedSectors {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Prover (abi.ActorID) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Prover))); err != nil {
		return err
	}

	return nil
}

func (t *WinningPoStVerifyInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Randomness (abi.PoStRandomness) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Randomness: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Randomness = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Randomness); err != nil {
		return err
	}
	// t.Proofs ([]abi.PoStProof) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Proofs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Proofs = make([]PoStProof, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v PoStProof
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Proofs[i] = v
	}

	// t.ChallengedSectors ([]abi.SectorInfo) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.ChallengedSectors: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.ChallengedSectors = make([]SectorInfo, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v SectorInfo
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.ChallengedSectors[i] = v
	}

	// t.Prover (abi.ActorID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Prover = ActorID(extra)

	}
	return nil
}
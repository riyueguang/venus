// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package dispatch

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufSimpleParams = []byte{129}

func (t *SimpleParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufSimpleParams); err != nil {
		return err
	}

	// t.Name (string) (string)
	if len(t.Name) > 8192 {
		return xerrors.Errorf("Value in field t.Name was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Name))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Name)); err != nil {
		return err
	}
	return nil
}

func (t *SimpleParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SimpleParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Name (string) (string)

	{
		sval, err := cbg.ReadStringWithMax(cr, 8192)
		if err != nil {
			return err
		}

		t.Name = string(sval)
	}
	return nil
}

package types

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/big"

	"github.com/huandu/xstrings"
	"github.com/itering/scale.go/utiles"
	"github.com/itering/scale.go/utiles/uint128"
	"github.com/shopspring/decimal"
)

type U8 struct {
	ScaleDecoder
}

func (u *U8) Process() {
	u.Value = u.GetNextU8()
}

func (u *U8) Encode(value interface{}) string {
	var i int
	switch v := value.(type) {
	case int:
		i = v
	case decimal.Decimal:
		i = int(v.IntPart())
	case uint32:
		i = int(v)
	case int64:
		i = int(v)
	case uint16:
		i = int(v)
	case float64:
		i = int(v)
	}
	return utiles.U8Encode(i)
}

type U16 struct {
	Reader io.Reader
	ScaleDecoder
}

func (u *U16) Process() {
	buf := &bytes.Buffer{}
	u.Reader = buf
	_, _ = buf.Write(u.NextBytes(2))
	c := make([]byte, 2)
	_, _ = u.Reader.Read(c)
	u.Value = binary.LittleEndian.Uint16(c)
}

func (u *U16) Encode(value interface{}) string {
	var u16 uint16
	switch v := value.(type) {
	case int:
		u16 = uint16(v)
	case decimal.Decimal:
		u16 = uint16(v.IntPart())
	case uint32:
		u16 = uint16(v)
	case int64:
		u16 = uint16(v)
	case uint16:
		u16 = v
	case float64:
		u16 = uint16(v)
	}
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, u16)
	return utiles.BytesToHex(bs)
}

type U32 struct {
	Reader io.Reader
	ScaleDecoder
}

func (u *U32) Process() {
	buf := &bytes.Buffer{}
	u.Reader = buf
	_, _ = buf.Write(u.NextBytes(4))
	c := make([]byte, 4)
	_, _ = u.Reader.Read(c)
	u.Value = binary.LittleEndian.Uint32(c)
}

func (u *U32) Encode(value interface{}) string {
	var u32 uint32
	switch v := value.(type) {
	case int:
		u32 = uint32(v)
	case decimal.Decimal:
		u32 = uint32(v.IntPart())
	case uint32:
		u32 = v
	case float64:
		u32 = uint32(v)
	}
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, u32)
	return utiles.BytesToHex(bs)
}

type U64 struct {
	ScaleDecoder
	Reader io.Reader
}

func (u *U64) Process() {
	buf := &bytes.Buffer{}
	u.Reader = buf
	_, _ = buf.Write(u.NextBytes(8))
	c := make([]byte, 8)
	_, _ = u.Reader.Read(c)
	u.Value = binary.LittleEndian.Uint64(c)
}

func (u *U64) Encode(value interface{}) string {
	var u64 uint64
	switch v := value.(type) {
	case int:
		u64 = uint64(v)
	case decimal.Decimal:
		u64 = uint64(v.IntPart())
	case uint32:
		u64 = uint64(v)
	case uint64:
		u64 = v
	case float64:
		u64 = uint64(v)
	}
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, u64)
	return utiles.BytesToHex(bs)
}

type U128 struct {
	ScaleDecoder
}

func (u *U128) Process() {
	elementBytes := u.NextBytes(16)
	if len(elementBytes) < 16 {
		elementBytes = utiles.HexToBytes(xstrings.LeftJustify(utiles.BytesToHex(elementBytes), 32, "0"))
	}
	u.Value = uint128.FromBytes(elementBytes).String()
}

func (u *U128) Encode(value interface{}) string {
	var bigInt *big.Int
	switch v := value.(type) {
	case *big.Int:
		bigInt = v
	case decimal.Decimal:
		bigInt = v.BigInt()
	case string:
		bigInt = decimal.RequireFromString(v).BigInt()
	case int64:
		bigInt = big.NewInt(v)
	case int:
		bigInt = big.NewInt(int64(v))
	default:
		panic("U128 value not decimal or big int")
	}
	bs := make([]byte, 16)
	u128 := uint128.FromBig(bigInt)
	u128.PutBytes(bs)
	return utiles.BytesToHex(bs)
}

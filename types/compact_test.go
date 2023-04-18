package types

import (
	"testing"

	"github.com/Innonminate/scale.go/types/scaleBytes"
	"github.com/Innonminate/scale.go/utiles"
	"github.com/stretchr/testify/assert"
)

func decode(encodedVal string, typeDescr string) interface{} {
	decoder := ScaleDecoder{}
	decoder.Init(scaleBytes.ScaleBytes{Data: utiles.HexToBytes(encodedVal)}, nil)
	return decoder.ProcessAndUpdateData(typeDescr)
}


// uint8  : 0 to 255 
// uint16 : 0 to 65535 
// uint32 : 0 to 4294967295 
// uint64 : 0 to 18446744073709551615 
// int8   : -128 to 127 
// int16  : -32768 to 32767 
// int32  : -2147483648 to 2147483647 
// int64  : -9223372036854775808 to 9223372036854775807

func TestDecodeCompactBase(t *testing.T) {
	raw1 := "00"
	raw2 := "04"
	raw3 := "a8"
	raw4 := "fc"
	raw5 := "0101"
	raw6 := "1501"
	raw7 := "fd03"
	raw8 := "fd07"
	raw9 := "fdff"
	raw10 := "02000100"
	raw11 := "feff0300"
	raw12 := "feffffff"
	raw13 := "0300000040"
	raw14 := "03ffffffff"
	raw15 := "13ffffffffffffffff"
	assert.EqualValues(t, 0, int(decode(raw1, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 1, int(decode(raw2, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 42, int(decode(raw3, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 63, int(decode(raw4, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 64, int(decode(raw5, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 69, int(decode(raw6, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 255, int(decode(raw7, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 511, int(decode(raw8, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 16383, int(decode(raw9, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 16384, int(decode(raw10, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 65535, int(decode(raw11, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 1073741823, int(decode(raw12, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 1073741824, int(decode(raw13, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 4294967295, decode(raw14, "Compact<u32>").(int))
	assert.EqualValues(t, uint64(18446744073709551615), decode(raw15, "Compact<u64>").(uint64))
}

func TestEncodeCompactBase(t *testing.T) {
	val1 := 0
	val2 := 1
	val3 := 42
	val4 := 63
	val5 := 64
	val6 := 69
	val7 := 255
	val8 := 511
	val9 := 16383
	val10 := 16384
	val11 := 65535
	val12 := 1073741823
	val13 := 1073741824
	val14 := 4294967295
	val15 := uint64(18446744073709551615)
	assert.EqualValues(t, "00", Encode("Compact<u8>", val1))
	assert.EqualValues(t, "04", Encode("Compact<u8>", val2))
	assert.EqualValues(t, "a8", Encode("Compact<u8>", val3))
	assert.EqualValues(t, "fc", Encode("Compact<u8>", val4))
	assert.EqualValues(t, "0101", Encode("Compact<u8>", val5))
	assert.EqualValues(t, "1501", Encode("Compact<u8>", val6))
	assert.EqualValues(t, "fd03", Encode("Compact<u8>", val7))
	assert.EqualValues(t, "fd07", Encode("Compact<u16>", val8))
	assert.EqualValues(t, "fdff", Encode("Compact<u16>", val9))
	assert.EqualValues(t, "02000100", Encode("Compact<u16>", val10))
	assert.EqualValues(t, "feff0300", Encode("Compact<u16>", val11))
	assert.EqualValues(t, "feffffff", Encode("Compact<u16>", val12))
	assert.EqualValues(t, "0300000040", Encode("Compact<u16>", val13))
	assert.EqualValues(t, "03ffffffff", Encode("Compact<u32>", val14))
	assert.EqualValues(t, "13ffffffffffffffff", Encode("Compact<u64>", val15))
}
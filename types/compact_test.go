package types

import (
	"testing"

	"github.com/Innonminate/scale.go/types/scaleBytes"
	"github.com/Innonminate/scale.go/utiles"
	"github.com/stretchr/testify/assert"
	"github.com/Innonminate/scale.go/types"
)

func decode(encodedVal string, typeDescr string) interface{} {
	decoder := types.ScaleDecoder{}
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
	raw10 := "feff0300"
	raw11 := "03ffffffff"
	raw12 := "13ffffffffffffffff"
	assert.EqualValues(t, 0, int(decode(raw1, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 1, int(decode(raw2, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 42, int(decode(raw3, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 63, int(decode(raw4, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 64, int(decode(raw5, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 69, int(decode(raw6, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 255, int(decode(raw7, "Compact<u8>").(uint64)))
	assert.EqualValues(t, 511, int(decode(raw8, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 16383, int(decode(raw9, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 65535, int(decode(raw10, "Compact<u16>").(uint16)))
	assert.EqualValues(t, 4294967295, decode(raw11, "Compact<u32>").(int))
	assert.EqualValues(t, uint64(18446744073709551615), decode(raw12, "Compact<u64>").(uint64))
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
	val10 := 65535
	val11 := 4294967295
	val12 := uint64(18446744073709551615)
	assert.EqualValues(t, "00", types.Encode("Compact<u8>", val1))
	assert.EqualValues(t, "04", types.Encode("Compact<u8>", val2))
	assert.EqualValues(t, "a8", types.Encode("Compact<u8>", val3))
	assert.EqualValues(t, "fc", types.Encode("Compact<u8>", val4))
	assert.EqualValues(t, "0101", types.Encode("Compact<u8>", val5))
	assert.EqualValues(t, "1501", types.Encode("Compact<u8>", val6))
	assert.EqualValues(t, "fd03", types.Encode("Compact<u8>", val7))
	
	assert.EqualValues(t, "fd07", types.Encode("Compact<u16>", val8))
	assert.EqualValues(t, "fdff", types.Encode("Compact<u16>", val9))
	assert.EqualValues(t, "feff0300", types.Encode("Compact<u16>", val10))
	assert.EqualValues(t, "03ffffffff", types.Encode("Compact<u32>", val11))
	assert.EqualValues(t, "13ffffffffffffffff", types.Encode("Compact<u64>", val12))
}
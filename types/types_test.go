package types_test

import (
	"encoding/json"
	"fmt"
	"github.com/itering/scale.go/source"
	"github.com/itering/scale.go/types"
	"github.com/itering/scale.go/utiles"
	"reflect"
	"strings"
	"testing"
)

func TestCompactU64(t *testing.T) {
	raw := "10"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	r := m.ProcessAndUpdateData("Compact<U64>").(uint64)
	if r != 4 {
		t.Errorf("Test TestCompactU64 Process fail, decode return %d", r)
	}
}

func TestRawBabePreDigest(t *testing.T) {
	raw := "0x02020000008b86750900000000"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	m.ProcessAndUpdateData("RawBabePreDigest")
}

func TestRawBabePreDigestVRF(t *testing.T) {
	raw := "0x030000000099decc0f0000000040a523a6fdd15ef7ffb2956689b828185b4d60cfac789f64d1b6f26257ebbe543349f8ceae602875c705a59b156af586c7cf907df5c8d5b541fa755638e32b07b02bfb5e7549fb88aa1f32da93519c67275e999da1cd58ec168c80b30e5b4d05"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	m.ProcessAndUpdateData("RawBabePreDigest")
}

func TestSet_Process(t *testing.T) {
	types.RuntimeType{}.Reg()
	types.RegCustomTypes(map[string]source.TypeStruct{
		"CustomSet": {
			Type:      "set",
			ValueList: []string{"Value1", "Value2", "Value3", "Value4", "Value5"},
		},
	})
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes("0x03000000")}, nil)
	r := m.ProcessAndUpdateData("CustomSet")
	if strings.Join(r.([]string), "") != "Value1Value2" {
		t.Errorf("Test TestSet_Process Process fail, decode return %v", r.([]string))
	}
}

// 0x025ed0b2 Compact<Balance>
func TestCompactBalance(t *testing.T) {
	raw := "0x025ed0b2"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	m.ProcessAndUpdateData("Compact<Balance>")
}

// 0xe52d2254c67c430a0000000000000000 Balance
func TestBalance(t *testing.T) {
	raw := "0xe52d2254c67c430a0000000000000000"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	c := m.ProcessAndUpdateData("Balance")
	fmt.Println(c)
}

//
func TestRegistration(t *testing.T) {
	raw := "0x04010000000200a0724e180900000000000000000000000d505552455354414b452d30310e507572655374616b65204c74641b68747470733a2f2f7777772e707572657374616b652e636f6d2f000000000d40707572657374616b65636f"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	r := m.ProcessAndUpdateData("Registration<BalanceOf>")
	rb, _ := json.Marshal(r)
	fmt.Println(string(rb))
}

func TestInt(t *testing.T) {
	raw := "0x2efb"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	r := m.ProcessAndUpdateData("i16")
	rb, _ := json.Marshal(r)
	fmt.Println(string(rb))
}

func TestBoolArray(t *testing.T) {
	raw := "0x00000100"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	r := m.ProcessAndUpdateData("Approvals")
	c := []interface{}{false, false, true, false}
	if !reflect.DeepEqual(c, r.([]interface{})) {
		t.Errorf("Test TestBoolArray Process fail, decode return %v", r)
	}
}

func TestReferendumInfo(t *testing.T) {
	raw := "0x00004e0c00295ce46278975a53b855188482af699f7726fbbeac89cf16a1741c4698dcdbc90080970600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	r := m.ProcessAndUpdateData("ReferendumInfo<BlockNumber, Hash, BalanceOf>")
	c := map[string]interface{}{
		"Ongoing": map[string]interface{}{
			"delay":        432000,
			"end":          806400,
			"proposalHash": "0x295ce46278975a53b855188482af699f7726fbbeac89cf16a1741c4698dcdbc9",
			"tally":        map[string]interface{}{"ayes": "0", "nays": "0", "turnout": "0"}, "threshold": "SuperMajorityApprove",
		}}
	if !reflect.DeepEqual(utiles.ToString(c), utiles.ToString(r)) {
		t.Errorf("Test TestReferendumInfo Process fail, decode return %v", r.(map[string]interface{}))
	}
}

func TestEthereumAccountId(t *testing.T) {
	raw := "0x4119b2e6c3cb618f4f0B93ac77f9Beec7ff02887"
	fmt.Println(len(utiles.HexToBytes(raw)))
	m := types.ScaleDecoder{}
	m.Init(types.ScaleBytes{Data: utiles.HexToBytes(raw)}, nil)
	r := m.ProcessAndUpdateData("EthereumAccountId")
	if r.(string) != "0x4119b2e6c3Cb618F4f0B93ac77f9BeeC7FF02887" {
		t.Errorf("Test TestEthereumAccountId Process fail, decode return %v", r)
	}
}

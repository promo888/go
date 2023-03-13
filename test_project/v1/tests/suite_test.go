package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	types "xcoin.com/v1/types"
)

func TestPlan(t *testing.T) {
	//data := []byte(`{"name":"John","age":30}`)
	types.Echo()
	// hdr := &types.Header{
	// 	Hash:      "testHash",
	// 	RootHash:  "testRootHash",
	// 	Timestamp: 123456789,
	// 	Signature: []byte("testSig"),
	// }

	// hdr := &types.Header{}
	// fmt.Printf("hdr: %v", hdr)

	// trf := &types.Transfer{
	// 	Timestamp:   time.Now().Unix(),
	// 	From:        "testFrom",
	// 	To:          "testTo",
	// 	Asset:       "testAsset",
	// 	Amount:      123.456,
	// 	Fee:         0.001,
	// 	Description: "testDescription",
	// }
	// fmt.Printf("trf: %v", trf)

	// data, _ := types.MarshalMsgPack(trf)
	// tx := &types.Message{
	// 	Header: *hdr,
	// 	Data:   data, //[]byte("testData"),
	// }
	// dataObj, _ := types.UnmarshalMsgPack(data)
	// fmt.Printf("trf: %+v\n", dataObj)
	// fmt.Printf("tx: %+v\n", tx)

	newtx, _ := types.NewTransaction("testFrom", "testTo", "testAsset", 123.456, 0.001, "testDescription")
	objType := reflect.TypeOf(newtx)
	fmt.Printf("objType: %v\n", objType)
	assert.Equal(t, "*v1.Transaction", objType.String())

	fmt.Printf("tx: %+v\n", newtx)
	var trx types.Transaction = *newtx
	fmt.Printf("Transaction: %+v\n", trx)
	unpackedTx, _ := types.UnmarshalMsgPack(newtx.Message.Data)
	fmt.Printf("unmarshalled tx data: %+v\n", unpackedTx)
}

func BenchmarkSerializeDeserialize(b *testing.B) {
	trf := &types.Transfer{
		Timestamp:   time.Now().Unix(),
		From:        "testFrom",
		To:          "testTo",
		Asset:       "testAsset",
		Amount:      123.456,
		Fee:         0.001,
		Description: "testDescription",
	}
	//fmt.Printf("trf: %v", trf)

	data, _ := types.MarshalMsgPack(trf)
	_, _ = types.UnmarshalMsgPack(data)

}

package utils

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"sync"
)

/*json  操作*/
var jsonApi jsoniter.API
var oJson sync.Once
func DataByJsonByte(params interface{}) []byte {
	oJson.Do(func() {
		jsonApi = jsoniter.ConfigCompatibleWithStandardLibrary
	})
	by ,err := jsonApi.Marshal(params)
	if err != nil {
		logrus.Error("DataByJsonStr err ", err)
	}
	return by
}

func StrJsonByData(str string,data interface{}) {
	oJson.Do(func() {
		jsonApi = jsoniter.ConfigCompatibleWithStandardLibrary
	})
	err := jsonApi.Unmarshal([]byte(str), data)
	if err != nil {
		logrus.Error("StrJsonByData err ", err)
	}
}

func ByteJsonByData(by []byte,data interface{}) error {
	oJson.Do(func() {
		jsonApi = jsoniter.ConfigCompatibleWithStandardLibrary
	})
	err := jsonApi.Unmarshal(by, data)
	if err != nil {
		return err
	}
	return nil
}


//解出固定格式uint32
func Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

//写入固定格式的uint32
func PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}
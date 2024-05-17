package common

import (
	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
)

var jsn = jsoniter.ConfigCompatibleWithStandardLibrary

func Copy(src, dst interface{}) error {
	return copier.Copy(dst, src)
}

func MarshalJSON(dst interface{}) ([]byte, error) {
	return jsn.Marshal(dst)
}

func UnmarshalJSON(src []byte, dst interface{}) error {
	return jsn.Unmarshal(src, dst)
}

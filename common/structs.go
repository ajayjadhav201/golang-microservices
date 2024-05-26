package common

import (
	"reflect"

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

// CopyJSON copies data from one struct to another struct
// for proper data copying the struct must implement omitempty in json tag
func CopyJSON(src interface{}, dest interface{}) error {
	//
	return nil
}

// MergeStructs merges non-empty fields from src into dest
func MergeStructs[T any](dest, src *T) {
	destVal := reflect.ValueOf(dest).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		if !isEmptyValue(srcField) {
			destVal.Field(i).Set(srcField)
		}
	}
}

// isEmptyValue checks if a reflect.Value is empty (zero value)
func isEmptyValue(v reflect.Value) bool {
	Println("ajaj kind: ", v.Kind())
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

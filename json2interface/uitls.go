package json2interface

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

func getType(data interface{}) ObjectType {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return Value
	case reflect.Slice, reflect.Array:
		return Array
	case reflect.Map:
		return Hash
	default:
		return Invalid
	}
}

func getNumberKind(f float64) reflect.Kind {
	decimals := 10000
	shift := float64(decimals) * f
	num := int(shift)
	if num%decimals == 0 {
		return reflect.Int
	}
	return reflect.Float64
}

func getFilePath(fileName string) string {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "data", fileName)
	return path
}

func GetJsonFileDataBytes(fileName string) ([]byte, error) {
	path := getFilePath(fileName)
	return ioutil.ReadFile(path)
}

func GetJsonFileDataAsInterface(fileName string) (interface{}, error) {
	dataBytes, err := GetJsonFileDataBytes(fileName)
	var res interface{}
	if err = json.Unmarshal(dataBytes, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func WriteInterfaceObject2JsonFile(data interface{}, fileName string) error {
	path := getFilePath(fileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	//dataBytes, err := json.MarshalIndent(data, "", " ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	fmt.Printf("interface data written to file %s\n", fileName)
	return nil
}

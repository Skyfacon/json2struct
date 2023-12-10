package json2struct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

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
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	fmt.Printf("interface data written to file %s\n", fileName)
	return nil
}

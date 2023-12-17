package json2interface

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ObjectType uint

const (
	Invalid ObjectType = iota
	Value
	Hash
	Array
)

const (
	rootID            = "$"
	DefaultStructName = "data"
)

type structure struct {
	prefix   string         // 表示上一层级的树前缀
	name     string         // 表示本层级的名称标识
	data     interface{}    // 表示本层级的值
	kind     reflect.Kind   // 表示本层级的类型
	children []*structure   // 表示子层级的内容
	record   map[string]int //根据key快速对children进行索引，找到对应key的child
}

func parseJsonSchema(dataBytes []byte) (*structure, error) {
	var input interface{}
	if err := json.Unmarshal(dataBytes, &input); err != nil {
		return nil, err
	}
	rootStructure := &structure{
		prefix:   rootID,
		name:     DefaultStructName,
		data:     input,
		children: make([]*structure, 0),
		record:   make(map[string]int),
	}
	rootStructure.travel()
	return rootStructure, nil
}

func ParseFromFile(fileName string) (*structure, error) {
	dataBytes, err := GetJsonFileDataBytes(fileName)
	if err != nil {
		return nil, err
	}
	rootStructure, err := parseJsonSchema(dataBytes)
	rootStructure.ShowSchema(" ")
	return rootStructure, err
}

func (s *structure) travel() {
	data := s.data
	switch getType(data) {
	case Value:
		v := reflect.ValueOf(data)
		kind := v.Kind()
		// 对于 json 而言， 30 和 30.0都是数值类型，是等价的
		if kind == reflect.Float64 {
			kind = getNumberKind(v.Float())
		}
		s.kind = kind

	case Array:
		list, _ := data.([]interface{})
		s.kind = reflect.Slice // 这个地方统一用slice应该可以涵盖所有的 json 数组类型吧
		// 如果取出所有的元素放入children中
		//prePath := fmt.Sprintf("%s[]", s.prefix)
		for i, oneItem := range list {
			prePath := fmt.Sprintf("%s.[%d]", s.prefix, i)
			name := fmt.Sprintf("[%d]", i)
			curr := &structure{
				prefix:   prePath,
				name:     name,
				data:     oneItem,
				children: make([]*structure, 0),
				record:   make(map[string]int),
			}
			s.children = append(s.children, curr)
			s.record[name] = len(s.children) - 1
			curr.travel()
		}

	case Hash:
		s.kind = reflect.Map
		mapValues := data.(map[string]interface{})
		for k, v := range mapValues {
			curr := &structure{
				prefix:   s.prefix + "." + k,
				name:     k,
				data:     v,
				children: make([]*structure, 0),
				record:   make(map[string]int),
			}
			s.children = append(s.children, curr)
			s.record[k] = len(s.children) - 1
			curr.travel()
		}
	case Invalid:
		return
	}
}

// Get 如果涉及到数组，path该如何填写？  $.person.contacts[2].type
func (s *structure) Get(path string) (*structure, string) {
	cpath := strings.Split(path, ".")[1:]
	curr := s
	for _, sub := range cpath {
		if index, ok := curr.record[sub]; ok {
			curr = curr.children[index]
		} else {
			return nil, "Not Found"
		}
	}
	return curr, curr.kind.String()
}

func (s *structure) GetVal() interface{} {
	return s.data
}

// Set 底层的 val 改变后，是否顶层 data 里包含的 interface{} 也会改变呢？if not， 还需要从顶层到底层进行重新的包装？
func (s *structure) Set(path string, val interface{}) error {
	var v *structure
	var typ string
	v, typ = s.Get(path)
	if typ == "Not Found" {
		return errors.New(typ)
	}
	v.data = val
	return nil
}

func (s *structure) ShowSchema(indentation string) {
	fmt.Printf(indentation + s.prefix + "(" + s.kind.String() + ")" + "\n")
	for _, v := range s.children {
		v.ShowSchema(indentation + indentation)
	}
}

func (s *structure) ToJsonFile(fileName string) {
	err := WriteInterfaceObject2JsonFile(s.data, fileName)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}

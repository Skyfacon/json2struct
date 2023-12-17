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
	prefix   string       // 表示上一层级的树前缀
	name     string       // 表示本层级的名称标识
	data     interface{}  // 表示本层级的值
	kind     reflect.Kind // 表示本层级的类型
	children []*structure // 表示子层级的内容
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
	rootStructure.showSchema(" ")
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
		prePath := fmt.Sprintf("%s[]", s.prefix)
		if len(list) == 0 {
			return
		}
		oneItem := list[0]
		curr := &structure{
			prefix:   prePath,
			name:     s.name + "[]", // name 用什么比较好？
			data:     oneItem,
			children: make([]*structure, 0),
		}
		s.children = append(s.children, curr)
		curr.travel()
	case Hash:
		s.kind = reflect.Map
		mapValues := data.(map[string]interface{})
		for k, v := range mapValues {
			curr := &structure{
				prefix:   s.prefix + "." + k,
				name:     k,
				data:     v,
				children: make([]*structure, 0),
			}
			s.children = append(s.children, curr)
			curr.travel()
		}
	case Invalid:
		return
	}
}

// 如果涉及到数组，path该如何填写？  $.person.contacts[2].type
func (s *structure) get(path string) (*interface{}, string) {
	cpath := strings.Split(path, ".")
	if len(cpath) == 1 {
		if cpath[0] == s.name {
			return &s.data, s.kind.String()
		} else {
			return nil, "Not Found"
		}
	}
	next := cpath[1]
	for _, child := range s.children {
		if child.name == next {
			return child.get(strings.Join(cpath[1:], "."))
		}
	}
	return nil, "Not Found"
}

// Set 底层的 val 改变后，是否顶层 data 里包含的 interface{} 也会改变呢？if not， 还需要从顶层到底层进行重新的包装？
func (s *structure) Set(path string, val interface{}) error {
	var v *interface{}
	var typ string
	v, typ = s.get(path)
	if typ == "Not Found" {
		return errors.New(typ)
	}
	v = &val
	if v == nil {
		return errors.New("nil value assign")
	}
	return nil
}

func (s *structure) showSchema(indentation string) {
	fmt.Printf(indentation + s.prefix + "(" + s.kind.String() + ")" + "\n")
	for _, v := range s.children {
		v.showSchema(indentation + indentation)
	}
}

func (s *structure) ToJsonFile(fileName string) {
	err := WriteInterfaceObject2JsonFile(s.data, fileName)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}

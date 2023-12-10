package json2struct

import (
	"fmt"
	"reflect"
)

func generateStringArrayAsInterface(len int) interface{} {
	res := make([]string, len)
	for i := 0; i < len; i++ {
		res[i] = fmt.Sprintf("value%d", i)
	}
	return interface{}(res)
}

func generateIntArrayAsInterface(len int) interface{} {
	res := make([]int, len)
	for i := 0; i < len; i++ {
		res[i] = i
	}
	return interface{}(res)
}

func generateHashMapOfDifferentValues() map[string]interface{} {
	record := make(map[string]interface{})
	record["stringArray"] = generateStringArrayAsInterface(3)
	record["intArray"] = generateIntArrayAsInterface(3)
	record["string"] = "string"
	record["int"] = 1
	record["float"] = 1.0
	record["bool"] = true

	nestedMap := make(map[string]interface{})
	nestedMap["name"] = "张三"
	nestedMap["age"] = 18
	nestedMap["male"] = true
	nestedMap["school"] = []string{"清华", "北大"}
	nestedMap["score"] = []float64{1.0, 2.0, 3.0}

	record["nested"] = nestedMap

	return record
}

func TestWriteInterface2file() {
	data := generateHashMapOfDifferentValues()
	err := WriteInterfaceObject2JsonFile(data, "testWriteInterface2file.json")
	if err != nil {
		panic(err)
	}
}

func TestShowDifferentValues() {
	data, err := GetJsonFileDataAsInterface("testWriteInterface2file.json")
	if err != nil {
		panic(err)
	}
	data1 := data.(map[string]interface{})
	boolVal := data1["bool"]
	fmt.Printf("v:%v,type:%v", boolVal, reflect.TypeOf(boolVal))
	fmt.Println()

	intVal := data1["int"]
	fmt.Printf("v:%v,type:%v", intVal, reflect.TypeOf(intVal))
	fmt.Println()

	floatVal := data1["float"]
	fmt.Printf("v:%v,type:%v", floatVal, reflect.TypeOf(floatVal))
	fmt.Println()

	stringVal := data1["string"]
	fmt.Printf("v:%v,type:%v", stringVal, reflect.TypeOf(stringVal))
	fmt.Println()

	stringArrayVal := data1["stringArray"]
	fmt.Printf("v:%v,type:%v", stringArrayVal, reflect.TypeOf(stringArrayVal))
	fmt.Println()

	intArrayVal := data1["intArray"]
	fmt.Printf("v:%v,type:%v", intArrayVal, reflect.TypeOf(intArrayVal))
	fmt.Println()

	nestedMap := data1["nested"].(map[string]interface{})
	name := nestedMap["name"]
	fmt.Printf("v:%v,type:%v", name, reflect.TypeOf(name))
	fmt.Println()

	age := nestedMap["age"]
	fmt.Printf("v:%v,type:%v", age, reflect.TypeOf(age))
	fmt.Println()

	male := nestedMap["male"]
	fmt.Printf("v:%v,type:%v", male, reflect.TypeOf(male))
	fmt.Println()

	school := nestedMap["school"]
	fmt.Printf("v:%v,type:%v", school, reflect.TypeOf(school))
	fmt.Println()

	score := nestedMap["score"]
	fmt.Printf("v:%v,type:%v", score, reflect.TypeOf(score))
	fmt.Println()

}

func TestModifyInterfaceValue() {
	data, err := GetJsonFileDataAsInterface("testWriteInterface2file.json")
	if err != nil {
		panic(err)
	}

	data1 := data.(map[string]interface{})
	// bool 字段包含的值为 bool 类型，转为 false
	data1["bool"] = false
	// int 字段值由1 转为 2
	data1["int"] = 2
	// float 字段值由1.0 转为 2.0
	data1["float"] = 2.0
	// string 字段包含的值为 string 类型，转为 STRING
	data1["string"] = "STRING"
	// stringArray 字段包含的值为 []string 类型，转为 []string{"Value0, Value1, Value2"}
	data1["stringArray"] = []string{"Value0", "Value1", "Value2"}
	// intArray 字段包含的值为 []int 类型，转为 []int{3, 4, 5}
	data1["intArray"] = []int{3, 4, 5}
	//
	nested := data1["nested"].(map[string]interface{})
	nested["age"] = 20
	nested["male"] = false
	nested["school"] = []string{"北航", "北理"}
	nested["score"] = []float64{4.1, 5.2, 6.3}

	//fmt.Printf("%v\n", data1)

	err = WriteInterfaceObject2JsonFile(data1, "modifyInterfaceInfo.json")
	if err != nil {
		panic(err)
	}
}

package json2interface

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_structure_get(t *testing.T) {
	st, err := ParseFromFile("data1.json")
	if err != nil {
		t.Error(err)
	}
	paths := []string{"$.person.age", "$.person.isStudent"}
	res := []struct {
		val interface{}
		typ string
	}{
		{val: 30, typ: reflect.Int.String()},
		{val: false, typ: reflect.Bool.String()},
	}
	for i, path := range paths {
		val, typ := st.get(path)
		fmt.Printf("val: %v, typ: %v\n", val, typ)
		fmt.Printf("res_val: %v, res_typ: %v\n", res[i].val, res[i].typ)
		if val != res[i].val {
			t.Error(val, res[i].val)
		}
		if typ != res[i].typ {
			t.Error(typ, res[i].typ)
		}
	}
}

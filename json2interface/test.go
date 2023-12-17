package json2interface

import (
	"reflect"
)

type Rf struct {
	Name string
	kind reflect.Kind
	val  interface{}
}

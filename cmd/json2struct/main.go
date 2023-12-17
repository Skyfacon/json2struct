package main

import (
	"flag"
	"fmt"
	"github.com/yudppp/json2struct"
	"github.com/yudppp/json2struct/json2interface"
)

var (
	debug     = flag.Bool("debug", false, "Set debug mode")
	omitempty = flag.Bool("omitempty", false, "Set omitempty mode")
	short     = flag.Bool("short", false, "Set short struct name mode")
	local     = flag.Bool("local", false, "Use local struct mode")
	example   = flag.Bool("example", false, "Use example tag mode")
	prefix    = flag.String("prefix", "", "Set struct name prefix")
	suffix    = flag.String("suffix", "", "Set struct name suffix")
	name      = flag.String("name", json2struct.DefaultStructName, "Set struct name")
)

func main() {
	root, err := json2interface.ParseFromFile("data1.json")
	if err != nil {
		return
	}
	paths := []string{"$.person.age", " $.company.employees.[1].name"}
	for _, path := range paths {
		s, typ := root.Get(path)
		fmt.Printf("type:%s,value:%v\n", typ, s.GetVal())
	}
	root.Set(paths[0], 40)
	root.Set(paths[1], "xulikai")
	for _, path := range paths {
		s, typ := root.Get(path)
		fmt.Printf("type:%s,value:%v\n", typ, s.GetVal())
	}
	root.ToJsonFile("data1Mod1.json")
	// clear
	root.Clear()
	paths = []string{"$.person.hobbies.[0]", "$.person.isStudent"}
	root.Set(paths[0], "跳舞")
	root.Set(paths[1], true)
	root.ToJsonFile("data1Mod2.json")
}

package main

import (
	"flag"
	"github.com/yudppp/json2struct"
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
	//flag.Parse()
	//json2struct.SetDebug(*debug)
	//opt := json2struct.Options{
	//	UseOmitempty:   *omitempty,
	//	UseShortStruct: *short,
	//	UseLocal:       *local,
	//	UseExample:     *example,
	//	Prefix:         *prefix,
	//	Suffix:         *suffix,
	//	Name:           strings.ToLower(*name),
	//}
	//parsed, err := json2struct.Parse(os.Stdin, opt)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(parsed)
	//json2struct.GenerateRandomInterfaceObject()

	//opt := json2struct.Options{
	//	UseOmitempty:   false,
	//	UseShortStruct: true,
	//	Name:           strings.ToLower(*name),
	//}
	//json2struct.SetDebug(true)
	//out, err := json2struct.ParseFromFile("data1.json", opt)
	//if err != nil {
	//	return
	//}
	//fmt.Println(out)

	//json2struct.TestWriteInterface2file()

	//json2struct.TestShowDifferentValues()

	json2struct.TestModifyInterfaceValue()

}

package goapack

import "fmt"
import "os"
import "strings"
import "reflect"
import "strconv"
import "runtime"

func Aerr(err error, fatal bool) {
	if err != nil {
		fmt.Printf("=== It broke with error: ===\n%s\n=== error end ===\n", err)
		if fatal { 
			os.Exit(1)
		} else {
			fmt.Printf("=== Error, but not exiting. ===\n")
		}
	}
	return
}

func ACycleStruct(obj_as_string string) {
	fmt.Printf("Hello from ACycleStruct\n")
	sliced := strings.Split( obj_as_string, " " )
	for c := 0; c < len(sliced); c++ {
		fmt.Printf("Slice element %v is: %v\n", c, sliced[c])
	}
	return
}

//empty struct just use to test if params for reflect are the same
//type i.e. structs too
type is_struct struct {
	foo int
}

func Dumper(obj interface{}) {
	t := reflect.TypeOf(obj)
	k := t.Kind()
	//is_test is ian empty struct just use to check if obj 
	//is struct too and
	//if we can print it
	t_test := is_struct{1}
	if k != reflect.TypeOf(t_test).Kind() {
		fmt.Printf("Dumper: Not Dumping. Not a struct. Kind is:%v\n",k)
		return
	}

	f := t.NumField()
	v := reflect.ValueOf(obj)
	fmt.Printf("Type:%v, Kind:%v, NumFields:%v\n", t, k, f)
	for c:=0; c<f; c++ {
		fmt.Printf("Dumper: Field name:%v, value:%v\n", t.Field(c).Name, v.Field(c) )
		//fmt.Printf("Value%+v\n", reflect.ValueOf(obj) )
	}
	 //as_string := fmt.Sprintf("%+v\n", t)
	//ACycleStruct( as_string )
}

func Apadding(obj interface{}, tl ...int) string {
	ostr := "Unknown type"
	t := reflect.TypeOf(obj)
	k := t.Kind()
	padding := 10 // default padding
	if len(tl) > 0 {
		padding = tl[0]
	}
	

	if k == reflect.String {
		ostr = "string"
	} else if k == reflect.Int || k == reflect.Int64 {
		ostr = strconv.FormatInt( reflect.ValueOf(obj).Int(), 10 )
	} else if k == reflect.Float32 || k == reflect.Float64 {
		//v := strconv.FormatFloat( reflect.ValueOf(obj).Float(),'f',2,32 )
		//fmt.Printf("=======:%v\n", v)
		ostr = strconv.FormatFloat( reflect.ValueOf(obj).Float(),'f',2,32 )
	} else {
		fmt.Printf("Unknown Kind of is:%v\n", k)
	}
	return doApadding(ostr, padding)
}

func doApadding(istr string, padding int) string {
	l := len(istr)
	add := padding - l
	pad := " "
	for c := 0; c < add; c++ {
		pad = pad + " "
	}
	//fmt.Printf("String: %v, Lenght: %v, padding:%v\n", istr, l, padding)
	return pad+istr
}
// not mine but useful
// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
	s := fmt.Sprintf("Alc=%v KiB", bToKb(m.Alloc))
//	s += fmt.Sprintf(" ToAlc = %v KiB", bToKb(m.TotalAlloc))
	s += fmt.Sprintf(" HeapUse=%v KiB", bToKb(m.HeapInuse))
	s += fmt.Sprintf(" Sys=%v KiB", bToKb(m.Sys))
	s += fmt.Sprintf(" HeapSys=%v KiB", bToKb(m.HeapSys))
	s += fmt.Sprintf(" StackSys = %v KiB", bToKb(m.StackSys))
	s += fmt.Sprintf(" nGC = %v\n", m.NumGC)
	return s
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func bToKb(b uint64) uint64 {
	return b / 1024
}

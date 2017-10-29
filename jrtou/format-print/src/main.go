package main

import (
	"fmt"
)

type Value struct {
	String   string
	Integer  int
	Float    float32
	Boolean  bool
	Array    []string
	ArrayInt []int
	Byte     byte
	Rune     rune
	Complex  complex64
	AsciiA   int
}

func newValue() Value {
	return Value{
		String:   "Hello World",
		Integer:  8080,
		Float:    99.99,
		Boolean:  true,
		Array:    []string{"Golang", "C", "C++"},
		ArrayInt: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		Byte:     0x07,
		Rune:     0040007, //G
		Complex:  complex(1.2, 2),
		AsciiA:   97,
	}
}

func main() {
	value := newValue() //初始化數值

	/* about print*/
	print(value)

	/* about println*/
	println(value)

	/* about printf*/

	//general
	PrintfValue1(value) //%v
	PrintfValue2(value) //%+v
	PrintfValue3(value) //%#v

	//bool
	PrintfBool(value) //%t

	//type
	PrintfType(value) //%T

	//print %
	PrintfPercent() //%%

	//string
	printfString1(value) //%s
	printfString2(value) //%a
	printfString3(value) //%x
	printfString4(value) //%X

	//integer
	printfInteger1(value) //%b
	printfInteger2(value) //%c
	printfInteger3(value) //%d
	printfInteger4(value) //%o
	printfInteger5(value) //%q
	printfInteger6(value) //%x
	printfInteger7(value) //%X
	printfInteger8(value) //%U

	//float
	printfFloat1(value) //%b
	printfFloat2(value) //%e
	printfFloat3(value) //%E
	printfFloat4(value) //%f
	printfFloat5(value) //%g
	printfFloat6(value) //%G

	//point
	printfPoint(value) //%p
}

//print
func print(value Value) {
	fmt.Print("[ Use Print ]")
	fmt.Print("print value = ", value)
	fmt.Print("\n")                                   // 使用 \n 換行
	fmt.Print(value.String, "\t", value.String, "\n") //使用 \t 做一個 tab 空行 最後用 \n 換行
}

//println
func println(value Value) {
	fmt.Println("\n[ Use Print ]")
	fmt.Println("println value = ", value)
	fmt.Println("println string = ", value.String)
	fmt.Println("println int = ", value.Integer)
	fmt.Println("println boolean = ", value.Boolean)
	fmt.Println("println float = ", value.Float)
	fmt.Println("println byte = ", value.Byte)
	fmt.Println("println rune = ", value.Rune)
	fmt.Println("println complex = ", value.Complex)
	fmt.Println("println AsciiA = ", value.AsciiA)
}

// general
//v
func PrintfValue1(value Value) {
	fmt.Println("\n[ Use v ]")
	fmt.Printf("value = %v\n", value)
	fmt.Printf("string = %v\n", value.String)
	fmt.Printf("int = %v\n", value.Integer)
	fmt.Printf("boolean = %v\n", value.Boolean)
	fmt.Printf("array = %v\n", value.Array)
	fmt.Printf("array int = %v\n", value.ArrayInt)
	fmt.Printf("float = %v\n", value.Float)
	fmt.Printf("byte = %v\n", value.Byte)
	fmt.Printf("rune = %v\n", value.Rune)
	fmt.Printf("complex = %v\n", value.Complex)
}

//+v
func PrintfValue2(value Value) {
	fmt.Println("\n[ Use +v ]")
	fmt.Printf("value = %+v\n", value)
	fmt.Printf("string = %+v\n", value.String)
	fmt.Printf("int = %+v\n", value.Integer)
	fmt.Printf("boolean = %+v\n", value.Boolean)
	fmt.Printf("array = %+v\n", value.Array)
	fmt.Printf("array int = %+v\n", value.ArrayInt)
	fmt.Printf("float = %+v\n", value.Float)
	fmt.Printf("byte = %+v\n", value.Byte)
	fmt.Printf("rune = %+v\n", value.Rune)
	fmt.Printf("complex = %+v\n", value.Complex)
}

//#v
func PrintfValue3(value Value) {
	fmt.Println("\n[ Use #v ]")
	fmt.Printf("value = %#v\n", value)
	fmt.Printf("string = %#v\n", value.String)
	fmt.Printf("int = %#v\n", value.Integer)
	fmt.Printf("boolean = %#v\n", value.Boolean)
	fmt.Printf("array = %#v\n", value.Array)
	fmt.Printf("array int = %#v\n", value.ArrayInt)
	fmt.Printf("float = %#v\n", value.Float)
	fmt.Printf("byte = %#v\n", value.Byte)
	fmt.Printf("rune = %#v\n", value.Rune)
	fmt.Printf("complex = %#v\n", value.Complex)
}

func PrintfBool(value Value) {
	fmt.Println("\n[ Use t ]")
	fmt.Printf("bool = %t", value.Boolean)
}

//%T
func PrintfType(value Value) {
	fmt.Println("\n[ Use T ]")
	fmt.Printf("struct = %t\n", value)
	fmt.Printf("string = %t\n", value.String)
	fmt.Printf("int = %t\n", value.Integer)
	fmt.Printf("string array = %t\n", value.Array)
	fmt.Printf("int arrary = %t\n", value.ArrayInt)
	fmt.Printf("float = %t\n", value.Float)
	fmt.Printf("bool = %t\n", value.Boolean)
	fmt.Printf("compex = %t\n", value.Complex)
	fmt.Printf("rune = %t\n", value.Rune)
	fmt.Printf("byte = %t\n", value.Byte)
}

//%%
func PrintfPercent() {
	fmt.Printf("\n[ Use %% ]\n")
	fmt.Printf("%%\n")
}

// string
//s
func printfString1(value Value) {
	fmt.Println("\n[ Use s ]")
	fmt.Printf("value = %s\n", value)
	fmt.Printf("string = %s\n", value.String)
	fmt.Printf("array = %s\n", value.Array)
}

//q
func printfString2(value Value) {
	fmt.Println("\n[ Use q ]")
	fmt.Printf("value = %q\n", value)
	fmt.Printf("string = %q\n", value.String)
	fmt.Printf("array = %q\n", value.Array)
}

//x
func printfString3(value Value) {
	fmt.Println("\n[ Use x ]")
	fmt.Printf("value = %x\n", value)
	fmt.Printf("string = %x\n", value.String)
	fmt.Printf("array = %x\n", value.Array)
}

//X
func printfString4(value Value) {
	fmt.Println("\n[ Use X ]")
	fmt.Printf("value = %X\n", value)
	fmt.Printf("string = %X\n", value.String)
	fmt.Printf("array = %X\n", value.Array)
}

//integer
//b
func printfInteger1(value Value) {
	fmt.Println("\n[ Use b ]")
	fmt.Printf("int = %b\n", value.Integer)
	fmt.Printf("array int = %b\n", value.ArrayInt)
	fmt.Printf("Assci A = %b\n", value.AsciiA)

}

//c
func printfInteger2(value Value) {
	fmt.Println("\n[ Use c ]")
	fmt.Printf("int = %c\n", value.Integer)
	fmt.Printf("array int = %c\n", value.ArrayInt)
	fmt.Printf("Assci A = %c\n", value.AsciiA)
}

//d
func printfInteger3(value Value) {
	fmt.Println("\n[ Use d ]")
	fmt.Printf("int = %d\n", value.Integer)
	fmt.Printf("array int = %d\n", value.ArrayInt)
	fmt.Printf("Assci A = %d\n", value.AsciiA)
}

//o
func printfInteger4(value Value) {
	fmt.Println("\n[ Use o ]")
	fmt.Printf("int = %o\n", value.Integer)
	fmt.Printf("array int = %o\n", value.ArrayInt)
	fmt.Printf("Assci A = %o\n", value.AsciiA)
}

//q
func printfInteger5(value Value) {
	fmt.Println("\n[ Use q ]")
	fmt.Printf("int = %q\n", value.Integer)
	fmt.Printf("array int = %q\n", value.ArrayInt)
	fmt.Printf("Assci A = %q\n", value.AsciiA)
}

//x
func printfInteger6(value Value) {
	fmt.Println("\n[ Use x ]")
	fmt.Printf("int = %x\n", value.Integer)
	fmt.Printf("array int = %x\n", value.ArrayInt)
	fmt.Printf("Assci A = %x\n", value.AsciiA)
}

//X
func printfInteger7(value Value) {
	fmt.Println("\n[ Use X]")
	fmt.Printf("int = %X\n", value.Integer)
	fmt.Printf("array int = %X\n", value.ArrayInt)
	fmt.Printf("Assci A = %X\n", value.AsciiA)

}

//U
func printfInteger8(value Value) {
	fmt.Println("\n[ Use U ]")
	fmt.Printf("int = %U\n", value.Integer)
	fmt.Printf("array int = %U\n", value.ArrayInt)
	fmt.Printf("Assci A = %U\n", value.AsciiA)
}

//float

//b
func printfFloat1(value Value) {
	fmt.Println("\n[ Use b ]")
	fmt.Printf("float = %b\n", value.Float)
	fmt.Printf("compex = %b\n", value.Complex)
}

//e
func printfFloat2(value Value) {
	fmt.Println("\n[ Use e ]")
	fmt.Printf("float = %e\n", value.Float)
	fmt.Printf("compex = %e\n", value.Complex)
}

//E
func printfFloat3(value Value) {
	fmt.Println("\n[ Use E ]")
	fmt.Printf("float = %E\n", value.Float)
	fmt.Printf("compex = %E\n", value.Complex)
}

//f
func printfFloat4(value Value) {
	fmt.Println("\n[ Use f ]")
	fmt.Printf("float = %f\n", value.Float)
	fmt.Printf("compex = %f\n", value.Complex)
}

//g
func printfFloat5(value Value) {
	fmt.Println("\n[ Use g ]")
	fmt.Printf("float = %g\n", value.Float)
	fmt.Printf("compex = %g\n", value.Complex)
}

//G
func printfFloat6(value Value) {
	fmt.Println("\n[ Use G ]")
	fmt.Printf("float = %G\n", value.Float)
	fmt.Printf("compex = %G\n", value.Complex)
}

//point
//p
func printfPoint(value Value) {
	fmt.Println("\n[ Use p ]")
	fmt.Printf("value point = %p\n", &value)
}

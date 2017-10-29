package main

import (
	"fmt"
)

type Value struct {
	String  string
	Integer int
	Float   float32
	Boolean bool
	Array   []string
	Byte    byte
	Rune    rune
	Complex complex64
	AsciiA  int
}

func newValue() Value {
	return Value{
		String:  "Hello World",
		Integer: 8080,
		Float:   99.99,
		Boolean: true,
		Array:   []string{"Golang", "C", "C++"},
		Byte:    0x07,
		Rune:    0040007, //G
		Complex: complex(1.2, 2),
		AsciiA:  97,
	}
}
func main() {
	value := newValue() //初始化數值

	//print(value)

	//println(value)

	//printfString(value)

	//printfString2(value)

	//printfDecimal(value)

	//printfBinary(value)

	//printfOctal(value)

	//printfHexadecimal(value)

	//printfFloat(value)

	printfScientificMark(value)

	//printfBoolean(value)

	//printfUnicode(value)

	//printfUnicodeFormat(value)

	//printfType(value)

	//PrintfValue(value)
}

//print 打印完不換行
func print(value Value) {
	fmt.Print("[ Use Print ]")
	fmt.Print("print w = ", value)
	fmt.Print("\n")                                   // 使用 \n 換行
	fmt.Print(value.String, "\t", value.String, "\n") //使用 \t 做一個 tab 空行 最後用 \n 換行
}

//println 打印完換行
func println(value Value) {
	fmt.Println("\n[ Use Print ]")
	fmt.Println("println w = ", value)
	fmt.Println("println string = ", value.String)
	fmt.Println("println int = ", value.Integer)
	fmt.Println("println boolean = ", value.Boolean)
	fmt.Println("println float = ", value.Float)
	fmt.Println("println byte = ", value.Byte)
	fmt.Println("println rune = ", value.Rune)
	fmt.Println("println complex = ", value.Complex)
	fmt.Println("println AsciiA = ", value.AsciiA)
}

//printf 字定義打印格式

//s 文字
func printfString(value Value) {
	fmt.Println("\n[ Use s Printf string]")

	fmt.Printf("value = %s\n", value)
	fmt.Printf("string = %s\n", value.String)
	fmt.Printf("int = %s\n", value.Integer)
	fmt.Printf("boolean = %s\n", value.Boolean)
	fmt.Printf("array = %s\n", value.Array)
	fmt.Printf("float = %s\n", value.Float)
	fmt.Printf("byte = %s\n", value.Byte)
	fmt.Printf("rune = %s\n", value.Rune)
	fmt.Printf("complex = %s\n", value.Complex)
	fmt.Printf("AsciiA = %s\n", value.AsciiA)
}

//q 文字加上引號
func printfString2(value Value) {
	fmt.Println("\n[ Use q Printf string]")

	fmt.Printf("value = %q\n", value)
	fmt.Printf("string = %q\n", value.String)
	fmt.Printf("int = %q\n", value.Integer)
	fmt.Printf("boolean = %q\n", value.Boolean)
	fmt.Printf("array = %q\n", value.Array)
	fmt.Printf("float = %q\n", value.Float)
	fmt.Printf("byte = %q\n", value.Byte)
	fmt.Printf("rune = %q\n", value.Rune)
	fmt.Printf("complex = %q\n", value.Complex)
	fmt.Printf("AsciiA = %q\n", value.AsciiA)
}

//d 十進位
func printfDecimal(value Value) {
	fmt.Println("\n[ Use d Printf Decimal]")

	fmt.Printf("value = %d\n", value)
	fmt.Printf("string = %d\n", value.String)
	fmt.Printf("int = %d\n", value.Integer)
	fmt.Printf("boolean = %d\n", value.Boolean)
	fmt.Printf("array = %d\n", value.Array)
	fmt.Printf("float = %d\n", value.Float)
	fmt.Printf("byte = %d\n", value.Byte)
	fmt.Printf("rune = %d\n", value.Rune)
	fmt.Printf("complex = %d\n", value.Complex)
	fmt.Printf("AsciiA = %d\n", value.AsciiA)
}

//b 二進位
func printfBinary(value Value) {
	fmt.Println("\n[ Use b Printf Binary]")

	fmt.Printf("value = %b\n", value)
	fmt.Printf("string = %b\n", value.String)
	fmt.Printf("int = %b\n", value.Integer)
	fmt.Printf("boolean = %b\n", value.Boolean)
	fmt.Printf("array = %b\n", value.Array)
	fmt.Printf("float = %b\n", value.Float)
	fmt.Printf("byte = %b\n", value.Byte)
	fmt.Printf("rune = %b\n", value.Rune)
	fmt.Printf("complex = %b\n", value.Complex)
	fmt.Printf("AsciiA = %b\n", value.AsciiA)
}

//o 八進位
func printfOctal(value Value) {
	fmt.Println("\n[ Use o Printf Octal]")

	fmt.Printf("value = %o\n", value)
	fmt.Printf("string = %o\n", value.String)
	fmt.Printf("int = %o\n", value.Integer)
	fmt.Printf("boolean = %o\n", value.Boolean)
	fmt.Printf("array = %o\n", value.Array)
	fmt.Printf("float = %o\n", value.Float)
	fmt.Printf("byte = %o\n", value.Byte)
	fmt.Printf("rune = %o\n", value.Rune)
	fmt.Printf("complex = %o\n", value.Complex)
	fmt.Printf("AsciiA = %o\n", value.AsciiA)
}

//x 十六進位
func printfHexadecimal(value Value) {
	fmt.Println("\n[ Use x Printf Hexadecimal]")

	fmt.Printf("value = %x\n", value)
	fmt.Printf("string = %x\n", value.String)
	fmt.Printf("int = %x\n", value.Integer)
	fmt.Printf("boolean = %x\n", value.Boolean)
	fmt.Printf("array = %x\n", value.Array)
	fmt.Printf("float = %x\n", value.Float)
	fmt.Printf("byte = %x\n", value.Byte)
	fmt.Printf("rune = %x\n", value.Rune)
	fmt.Printf("complex = %x\n", value.Complex)
	fmt.Printf("AsciiA = %x\n", value.AsciiA)

	fmt.Println("\n[ Use X Printf Hexadecimal]")

	fmt.Printf("value = %X\n", value)
	fmt.Printf("string = %X\n", value.String)
	fmt.Printf("int = %X\n", value.Integer)
	fmt.Printf("boolean = %X\n", value.Boolean)
	fmt.Printf("array = %X\n", value.Array)
	fmt.Printf("float = %X\n", value.Float)
	fmt.Printf("byte = %X\n", value.Byte)
	fmt.Printf("rune = %X\n", value.Rune)
	fmt.Printf("complex = %X\n", value.Complex)
	fmt.Printf("AsciiA = %X\n", value.AsciiA)
}

//f 浮點數
func printfFloat(value Value) {
	fmt.Println("\n[ Use f Printf float]")

	fmt.Printf("value = %f\n", value)
	fmt.Printf("string = %f\n", value.String)
	fmt.Printf("int = %f\n", value.Integer)
	fmt.Printf("boolean = %f\n", value.Boolean)
	fmt.Printf("array = %f\n", value.Array)
	fmt.Printf("float = %f\n", value.Float)
	fmt.Printf("byte = %f\n", value.Byte)
	fmt.Printf("rune = %f\n", value.Rune)
	fmt.Printf("complex = %f\n", value.Complex)
	fmt.Printf("AsciiA = %f\n", value.AsciiA)
}

//e 科學記號
func printfScientificMark(value Value) {
	fmt.Println("\n[ Use e Printf Scientific Mark]")

	fmt.Printf("value = %e\n", value)
	fmt.Printf("string = %e\n", value.String)
	fmt.Printf("int = %e\n", value.Integer)
	fmt.Printf("boolean = %e\n", value.Boolean)
	fmt.Printf("array = %e\n", value.Array)
	fmt.Printf("float = %e\n", value.Float)
	fmt.Printf("byte = %e\n", value.Byte)
	fmt.Printf("rune = %e\n", value.Rune)
	fmt.Printf("complex = %e\n", value.Complex)
	fmt.Printf("AsciiA = %e\n", value.AsciiA)

	fmt.Println("\n[ Use E Printf Scientific Mark]")

	fmt.Printf("value = %E\n", value)
	fmt.Printf("string = %E\n", value.String)
	fmt.Printf("int = %E\n", value.Integer)
	fmt.Printf("boolean = %E\n", value.Boolean)
	fmt.Printf("array = %E\n", value.Array)
	fmt.Printf("float = %E\n", value.Float)
	fmt.Printf("byte = %E\n", value.Byte)
	fmt.Printf("rune = %E\n", value.Rune)
	fmt.Printf("complex = %E\n", value.Complex)
	fmt.Printf("AsciiA = %E\n", value.AsciiA)
}

//t boolean
func printfBoolean(value Value) {
	fmt.Println("\n[ Use t Printf float]")

	fmt.Printf("value = %t\n", value)
	fmt.Printf("string = %t\n", value.String)
	fmt.Printf("int = %t\n", value.Integer)
	fmt.Printf("boolean = %t\n", value.Boolean)
	fmt.Printf("array = %t\n", value.Array)
	fmt.Printf("float = %t\n", value.Float)
	fmt.Printf("byte = %t\n", value.Byte)
	fmt.Printf("rune = %t\n", value.Rune)
	fmt.Printf("complex = %t\n", value.Complex)
	fmt.Printf("AsciiA = %t\n", value.AsciiA)
}

//c unicode 對應碼
func printfUnicode(value Value) {
	fmt.Println("\n[ Use c Printf unicode]")

	fmt.Printf("value = %c\n", value)
	fmt.Printf("string = %c\n", value.String)
	fmt.Printf("int = %c\n", value.Integer)
	fmt.Printf("boolean = %c\n", value.Boolean)
	fmt.Printf("array = %c\n", value.Array)
	fmt.Printf("float = %c\n", value.Float)
	fmt.Printf("byte = %c\n", value.Byte)
	fmt.Printf("rune = %c\n", value.Rune)
	fmt.Printf("complex = %c\n", value.Complex)
	fmt.Printf("AsciiA = %c\n", value.AsciiA)
}

//U unicode 格式
func printfUnicodeFormat(value Value) {
	fmt.Println("\n[ Use U Printf unicode format]")

	fmt.Printf("value = %U\n", value)
	fmt.Printf("string = %U\n", value.String)
	fmt.Printf("int = %U\n", value.Integer)
	fmt.Printf("boolean = %U\n", value.Boolean)
	fmt.Printf("array = %U\n", value.Array)
	fmt.Printf("float = %U\n", value.Float)
	fmt.Printf("byte = %U\n", value.Byte)
	fmt.Printf("rune = %U\n", value.Rune)
	fmt.Printf("complex = %U\n", value.Complex)
	fmt.Printf("AsciiA = %U\n", value.AsciiA)
}

//T 打印類型
func printfType(value Value) {
	fmt.Println("\n[ Use T Printf type]")

	fmt.Printf("value = %T\n", value)
	fmt.Printf("string = %T\n", value.String)
	fmt.Printf("int = %T\n", value.Integer)
	fmt.Printf("boolean = %T\n", value.Boolean)
	fmt.Printf("array = %T\n", value.Array)
	fmt.Printf("float = %T\n", value.Float)
	fmt.Printf("byte = %T\n", value.Byte)
	fmt.Printf("rune = %T\n", value.Rune)
	fmt.Printf("complex = %T\n", value.Complex)
	fmt.Printf("AsciiA = %T\n", value.AsciiA)
}

//v 打印結構
func PrintfValue(value Value) {
	fmt.Println("\n[ Use v Printf content]")

	fmt.Printf("value = %v\n", value)
	fmt.Printf("string = %v\n", value.String)
	fmt.Printf("int = %v\n", value.Integer)
	fmt.Printf("boolean = %v\n", value.Boolean)
	fmt.Printf("array = %v\n", value.Array)
	fmt.Printf("float = %v\n", value.Float)
	fmt.Printf("byte = %v\n", value.Byte)
	fmt.Printf("rune = %v\n", value.Rune)
	fmt.Printf("complex = %v\n", value.Complex)
	fmt.Printf("AsciiA = %v\n", value.AsciiA)

	fmt.Println("\n[ Use +v Printf Type + content]")

	fmt.Printf("value = %+v\n", value)
	fmt.Printf("string = %+v\n", value.String)
	fmt.Printf("int = %+v\n", value.Integer)
	fmt.Printf("boolean = %+v\n", value.Boolean)
	fmt.Printf("array = %+v\n", value.Array)
	fmt.Printf("float = %+v\n", value.Float)
	fmt.Printf("byte = %+v\n", value.Byte)
	fmt.Printf("rune = %+v\n", value.Rune)
	fmt.Printf("complex = %+v\n", value.Complex)
	fmt.Printf("AsciiA = %+v\n", value.AsciiA)

	fmt.Println("\n[ Use #v Printf struceure type]")

	fmt.Printf("value = %#v\n", value)
	fmt.Printf("string = %#v\n", value.String)
	fmt.Printf("int = %#v\n", value.Integer)
	fmt.Printf("boolean = %#v\n", value.Boolean)
	fmt.Printf("array = %#v\n", value.Array)
	fmt.Printf("float = %#v\n", value.Float)
	fmt.Printf("byte = %#v\n", value.Byte)
	fmt.Printf("rune = %#v\n", value.Rune)
	fmt.Printf("complex = %#v\n", value.Complex)
	fmt.Printf("AsciiA = %#v\n", value.AsciiA)
}

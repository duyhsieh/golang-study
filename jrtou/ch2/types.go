package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("*** Number ***")
	//int + int
	fmt.Println("int : 1 + 1 = ", 1+1)

	//float + float
	fmt.Println("float : 1 + 1 = ", 1.0+1.0)
	fmt.Println("float : 1.01 - 0.99 = ", 1.01-0.99)

	a := 1.01
	b := 0.99
	fmt.Println("float : a - b = ", a-b)

	//運算符
	fmt.Println("\n*** Operators ***")
	c := 10
	d := 3
	fmt.Println(" - : c - d = ", c-d)
	fmt.Println(" + : c + d = ", c+d)
	fmt.Println(" * : c * d = ", c*d)
	fmt.Println(" / : c / d = ", c/d)
	fmt.Println(" % : c % d = ", c%d)

	fmt.Println("\n *** String ***")
	world := "Hello World"
	fmt.Println("string len = ", len(world))
	fmt.Println("Get first charatcrt : ", world[0])
	fmt.Println("Concatenates string :", "Hi,"+"you")

	fmt.Println("\n*** Boolean ***")
	fmt.Println("T && T :", true && true)
	fmt.Println("T && F :", true && false)
	fmt.Println("T || T :", true || true)
	fmt.Println("T || F :", true || false)
	fmt.Println("!T:", !true)
	fmt.Println("!F:", !false)

	fmt.Println("\n*** Exercises ***")
	exercises2()
}

func exercises2() {
	digit := "11111111"
	sum := 0.0
	for index, value := range digit {
		tmp, err := strconv.ParseFloat(string(value), 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += tmp * math.Pow(2, float64(len(digit)-1-index))
	}
	fmt.Println("what’s the largest eight-digit number : ", sum)
}

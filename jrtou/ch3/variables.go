package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter Number then double to output")
	var input float64

	fmt.Scan(&input)

	output := input * 2
	fmt.Println("ouput = ", output)

	fmt.Println("\n*** exercises ***")
	exercises1()
	exercises2()
	exercises5()
	exercises6()
}

func exercises1() {
	a := 100
	var b int64 = 100
	fmt.Println("exercises1 :\na = ", a, "\tb = ", b)
}

func exercises2() {
	x := 5
	x += 1
	fmt.Println("\nexercises2 :\nx := 5 , x+=1 => ", x)
}

func exercises5() {
	fmt.Println("exercises5 :\nThis is Fahrenheit into Celsius")
	fmt.Println("Enter Fahrenheit")
	var temperature float64
	fmt.Scan(&temperature)
	fmt.Printf("Input Fahrenheit = %f into Celsius = %f\n", temperature, (temperature-32)*5/9)
}

func exercises6() {
	fmt.Println("exercises6 :\n feet into meters")
	fmt.Println("Enter feet")
	var feet float64
	fmt.Scan(&feet)
	fmt.Printf("feet = %f convert to meters = %f\n", feet, feet*0.3048)
}

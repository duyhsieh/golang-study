package main

import (
	"fmt"
)

func main() {
	forStatement()

	ifStatement()

	switchStatement()

	exercises()
}

// about for
func forStatement() {
	fmt.Println("[ For Statement ]")

	i := 1
	fmt.Println("do for i <= 10")
	for i <= 10 {
		fmt.Println("i = ", i)
		i++
	}
	fmt.Println("")

	fmt.Println("do for j := 1; j <= 10 ; j++")
	for j := 1; j <= 10; j++ {
		fmt.Println("j = ", j)
	}
	fmt.Println("")
}

// about if
func ifStatement() {
	fmt.Println("[ If Statement ]")
	fmt.Println("for i = 1 ; i <=10 ; i++\n\tif i%2 == 0 output true : even false : odd")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	fmt.Println("")
}

func switchStatement() {
	fmt.Println("[ Switch Statement ]")
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("ONE")
		case 2:
			fmt.Println("TWO")
		case 3:
			fmt.Println("THREE")
		case 4:
			fmt.Println("FOUR")
		case 5:
			fmt.Println("FIVE")
		default:
			fmt.Println("i > 5")
		}
	}
	fmt.Println("")
}

func exercises() {
	fmt.Println("[ Exercises ]")
	fmt.Println("exe 1")
	i := 10
	if i > 10 {
		fmt.Println("big")
	} else {
		fmt.Println("small")
	}
	fmt.Println("")

	fmt.Println("exe 2")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, ",")
		}
	}
	fmt.Println("")

	fmt.Println("exe 3")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("num = ", i, "\tFizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("num = ", i, "\tFizz")
		} else if i%5 == 0 {
			fmt.Println("num = ", i, "\tBuzz")
		} else {
			fmt.Println("num = ", i)
		}
	}
}

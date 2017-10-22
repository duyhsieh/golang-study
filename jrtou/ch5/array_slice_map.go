package main

import (
	"fmt"
)

func main() {
	arrays()

	slices()

	maps()

	exercises()
}

func arrays() {
	fmt.Println("[ Array ]")
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println("")

	x[0] = 20
	x[1] = 30
	x[2] = 40
	x[3] = 50
	x[4] = 60
	var total int = 0
	for index, value := range x {
		fmt.Println("index = ", index, "\tvalue = ", value)
		total += value
	}
	fmt.Println(total, "/", len(x), " = ", total/len(x))
	fmt.Println("")
}

func slices() {
	fmt.Println("[ Slice ]")
	var x []float64
	var y []float64

	x = make([]float64, 5)
	y = make([]float64, 5, 10)

	fmt.Println("x = ", x, "\t len = ", len(x), "\t capacity = ", cap(x))
	fmt.Println("y = ", y, "\t len = ", len(y), "\t capacity = ", cap(y))

	for i := 0; i < 5; i++ {
		x[i] = float64(i + 1)
	}
	fmt.Println("set x value = ", x)

	y = x
	fmt.Println("y = x , y = ", y, "\t len = ", len(y), "\t cap = ", cap(y))

	z := make([]float64, 10)
	for i := 0; i < 10; i++ {
		z[i] = float64(i + 1)
	}
	fmt.Println("make z = ", z)
	y = z[2:8]
	fmt.Println("y = z[2:8] , y = ", y, "\t len = ", len(y), "\t cap = ", cap(y))
	y = z[3:5]
	fmt.Println("y = z[3:5] , y = ", y, "\t len = ", len(y), "\t cap = ", cap(y))
	fmt.Println("")

	fmt.Println("[ Array Append ]")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1 = ", slice1, "\tslice2 = ", slice2)
	fmt.Println("")

	fmt.Println("[ Array COPY ]")
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	fmt.Println("slice3 = ", slice3, " len = ", len(slice3), "\tslice4 = ", slice4, " len = ", len(slice4))
	fmt.Println("after copy  copy(slice4, slice3)")
	copy(slice4, slice3)
	fmt.Println("slice3 = ", slice3, "\tslice4 = ", slice4)
	fmt.Println("")
}

func maps() {
	fmt.Println("[ Map ]")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	for key, value := range elements {
		fmt.Println("key = ", key, "\tvalue = ", value)
	}
	fmt.Println("")

	fmt.Println("key = Li value = ", elements["Li"])
	fmt.Println("")

	fmt.Println("get key = AAA")
	value, ok := elements["AAA"]
	fmt.Println("value = ", value, "is has = ", ok)
	fmt.Println("")
}

func exercises() {
	fmt.Println("[ Exercises ]")

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var min, tmp = x[0], 0
	for _, value := range x {
		tmp = value
		if min > tmp {
			min = tmp
		}
	}
	fmt.Println("min = ", min)
}

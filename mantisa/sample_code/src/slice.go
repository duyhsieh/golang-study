package main

import "fmt"

func main() {
    /*
    slice data structure: 1. pointer to base array 2. length 3. capacity
    slice can be generated from existing array. The slice will alter base array, so use with care.
    For dynamic array, just use make to generate default base array. Do not generate from existing array and try to appned it, 
    since when capacity exceeds, a new array must be reallocated and the slice is detached from old base array, which is bad performance.
    */

    // testing dynamic arrays
    fmt.Println("Testing dynamic arrays")
    var x = make([]int, 2, 4)
    fmt.Printf("init: x is %v, capacity=%v, len=%v\n", x, cap(x), len(x) ) // can be used as dynamic array
    y := x[:3]
    z := x[2:4]
    x[0]=1
    y[1]=2
    z[0] = 3
    all := x[:4]

    fmt.Printf("x is %v, capacity=%v, len=%v\n", x, cap(x), len(x) ) // can be used as dynamic array
    fmt.Printf("y is %v, capacity=%v, len=%v\n", y, cap(y), len(y) ) 
    fmt.Printf("z is %v, capacity=%v, len=%v\n", z, cap(z), len(z) ) 
    fmt.Printf("all is %v, capacity=%v, len=%v\n", all, cap(all), len(all) ) 

    fmt.Println("Testing sub array by slices")
    // accessing sub arrays
    var b = [...]int{1,2,3,4} 
    s := b[2:]

    fmt.Printf("slice s is %v , capacity is %v, length is %v, address is %p\n", s, cap(s), len(s), &s)
    s[1]*=10 // also modify base array
    fmt.Printf("after operate on slice : slice = %v, base array = %v\n", s, b)

    s = append(s, 5)  // capacity exceed limit, create a new base array(copied from original base), and detach forever with original array
    fmt.Printf("slice s after append is %v , capacity is %v, length is %v, address is %p\n", s, cap(s), len(s), &s)
    s[1]*=10
    fmt.Printf("after operate on slice again: slice = %v, detached array = %v\n", s, b)
    s = append(s, 6)  // capacity exceed limit, create a new base array(copied from original base), and detach forever with original array
    fmt.Printf("slice s after append is %v , capacity is %v, length is %v, address is %p\n", s, cap(s), len(s), &s)
    s = append(s, 7)  // capacity exceed limit, create a new base array(copied from original base), and detach forever with original array
    fmt.Printf("slice s after append is %v , capacity is %v, length is %v, address is %p\n", s, cap(s), len(s), &s)


}

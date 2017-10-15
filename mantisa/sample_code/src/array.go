package main

import ("fmt"; "reflect")

func main() {
//    var a[5] int
    var a1 [3]int
    var a2 = [4]int{1,2,3,4}
    var b = [...]int{1,2,3,4} // let compiler count for you
    fmt.Printf("a1 is %v a2 is %v\n", a1, a2)
    fmt.Println("b is", b)

    fmt.Println(reflect.TypeOf(a1))
    fmt.Println(reflect.TypeOf(a2))


    // traverse by iteration
    for i, v := range b { 
        fmt.Printf("index %d of array b is %d\n", i, v)
    }

    // traverse by loop
    fmt.Println("------------")
    for i:=0; i < len(b); i++ {
        fmt.Printf("index %d of array b is %d\n", i, b[i] )
    }
    

}

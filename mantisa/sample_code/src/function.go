package main

import ("fmt")

func main() {
    r1, r2 := f1()
    fmt.Println("f1 return value=", r1, r2)
    
    fmt.Println("add(1,2)=", variadic_add(1,2))
    fmt.Println("add(1,2,3)=", variadic_add(1,2,3))
    x:=[]int{1,2,3}
    variadic_dynamic(1, 0.2, x)
    variadic_default2(3,4)
    variadic_default2(3,4,5,6)

}
func f1()(ReturnDataName1 int, ReturnDataName2 int) {
    ReturnDataName1 = 3
    //ReturnDataName2 = 1 default = 0
    return
}
func variadic_add(args ...int) int {
    fmt.Println("============ variadic_add ==============")
    total:=0
    for k, v := range args {
        fmt.Printf("key %v=%v,",k ,v)
        total+=v
    }
    fmt.Println()
    return total
}

func variadic_dynamic(args ...interface{} ) {
    fmt.Println("============= variadic_dynami ============= ")
    for k,v := range args {
        fmt.Printf("key=%v,value=%v,",k,v)
    }
    fmt.Println()
}

func variadic_default2(r1 int, r2 int, args ...int) {
    fmt.Println("=========== variadic_default2 =============")
    fmt.Println("r1 is ", r1, " r2 is ", r2)
    for k,v := range args {
        fmt.Printf("key=%v,value=%v,",k,v)
    }
    fmt.Println()
}


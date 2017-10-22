package main
//import ("fmt"; "math/rand")
import ("fmt" )

func main() {
    g1 := MakeNumberGeneratorClosure()
    g2 := MakeNumberGeneratorClosure()
    fmt.Println(g1(),g1(),g2(),g2() )
    //fmt.Println(g1(),g2(),g1(),g2() )
    foo(g1)
    fmt.Println(g1())
}

func MakeNumberGeneratorClosure() func() int {
    i:=0
    return func() int {
        i+=1
        return i 
    }
}

func foo(funObj func() int) {
    funObj()
}

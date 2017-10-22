package main
import ("fmt" )

func main() {
    x := 0
    foo(&x)
    fmt.Println(x)
    bar(x)
    fmt.Println(x)

    y:=new(int)
    foo(y)
    fmt.Println(*y)
}

func foo(x *int) {
    *x+=1

}


func bar(x int) {
    x+=1

}



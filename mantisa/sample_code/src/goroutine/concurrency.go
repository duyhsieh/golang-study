package main
import ("fmt";"time")

func main() {
    foo()
    time.Sleep(time.Second)
    bar()
    var input string
    fmt.Scanln(&input)
}

func foo() {
    var i=0  // this may not necessary allocated on Stack, since compiler will infer whether it is referenced after function exits.
    fmt.Printf("foo %v Adr=%v\n", i, &i)
    for i=0;i<5;i++ {
        go func() {
            fmt.Printf("foo goroutine %v Ard=%v\n", i, &i)  // note: they will not print what you are thinking ! 
        }()
    }
    fmt.Println("Fin foo")
}


func bar() {
    var i=0 
    fmt.Printf("bar %v Adr=%v\n", i, &i)
    for i=0;i<5;i++ {
        go func(x int) {
            fmt.Printf("bar goroutine %v Ard=%v\n", x, &x) 
        }(i)
    }
    fmt.Println("Fin bar")
}



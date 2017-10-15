package main
import "fmt"

func main() {
//    var a[5] int
    x:=1

    if x == 1 {
        fmt.Println(1)
    } else if x == 2 {
        fmt.Println(2)
    } else {
        fmt.Println("x")
    }

    var i = 1
    switch i {
        case 1: fmt.Println(i)
        case 2: fmt.Println(i)
        case 3: fmt.Println(i)
        default: fmt.Println("default") // will not run
        break
    }
    
    // infinite loop 
    for {
        i+=1
        if i >= 10000 {
            fmt.Println("break infinite loop")
            break
        }
    }

}

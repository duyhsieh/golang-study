package main
import "fmt"
import "mylib"
import "mylib/sublib"

func main() {
    fmt.Println(mylib.Sum(1,2) )
    fmt.Println(mylib.Multiply(2,3) )
    fmt.Println(sublib.Inc(7) )
}

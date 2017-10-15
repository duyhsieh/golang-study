package main
import "fmt"

func main() {
//    var a[5] int
    var a = 1
    var b = 0.5
    var c = [4]int{1,2,3,4}
    var d = "Hello World"

    fmt.Printf("a=%d b=%f c=%v d=%s\n", a,b, c, d)
    fmt.Println("---------")
    fmt.Print(a, b,c,d, "\n")

    var b1 = [...]int{1,2,3,4} // let compiler count for you
    fmt.Printf("Address of b1 is %p\n", &b1)


}

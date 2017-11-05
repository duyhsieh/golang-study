package main
import ("fmt";"time"; "strconv")

func main() {
    var c1 chan string = make(chan string)
    var c2 chan string = make(chan string)
    go sender(1, c2)
    time.Sleep(time.Millisecond)
    go sender(0, c1)
    go receiver(c1,c2)
    var input string
    fmt.Scanln(&input)
}

func sender(no int, c chan string) {
    for { 
        c<- "Hello" + strconv.Itoa(no)
        time.Sleep(2000*time.Millisecond)
    }
}

func receiver(c1 chan string, c2 chan string) {
    for {
        select {
        case msg1 := <-c1:  // won't just block in c1 ; c2 can be triggered first
            fmt.Println("C1 got", msg1)
        case msg2 := <-c2:
            fmt.Println("C2 got", msg2)
        case <- time.After(time.Second):
            fmt.Println("Timeout")
        //default:
        //    break
            //fmt.Println("use default for non blocking channel multiplex")
        //default:
        //    break
            //i+=1
            //time.Sleep(time.Second)
        }
    }
}



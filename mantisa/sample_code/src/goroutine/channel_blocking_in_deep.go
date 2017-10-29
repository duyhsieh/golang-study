package main
//import ("fmt"; "time";"os")
import ("fmt"; "time"; "sync")
// if consume goroutine goes first, consume blocking reading ; produce sends and release consume's block, and produce may then go to to send next one
// and then it blocks at sending.
// when produce-consume pair of consume-producer pair is released blocking, next round either produce/consume can begin first;
// so the pattern will be pccp, pcpc, cppc, cpcp

/*
example output: 
Consuming
Producing: 0
Produced : 0  // at this time consume blocking is released;but execution does not switch to consume immediately, so "Consumed: 0 is printed later than its real happen time"
Producing: 1
Consumed: 0
Consuming
Consumed: 1
Produced : 1
Producing: 2

*/

func main() {
    var lock sync.RWMutex
    var c chan int = make(chan int)
    var buf chan int = make(chan int, 40)
    var synChan chan int = make(chan int)
    
    //go produce(c)
    //time.Sleep(100*time.Millisecond)
    //go consume(c)

    go sender(c, synChan, buf, &lock)
    time.Sleep(100*time.Millisecond)
    go receiver(c, synChan, buf, &lock)
    time.Sleep(time.Second) // wait for complete
    
    for i:=0; i < 20; i++ {
        fmt.Println(<-buf)
    }
}

func produce(c chan int) {
    for i:=0; i<10; i++ {
        //fmt.Println("Producing:", i)
        //os.Stdout.Sync()
        c <- i
        fmt.Println("Produced :", i)
        time.Sleep(10*time.Millisecond)
    }
}
func consume(c chan int) {
    for {
        //fmt.Println("Consuming")
        x := <-c
        //<-c
        fmt.Println("Consumed:", x)
        //os.Stdout.Sync()
        time.Sleep(10*time.Millisecond)
    }
}

// note: if sender is Blocking, receiver side will unblock sender and keep execution
//(will not switch back to sender immediately, since it does  not block). So receiver will write history in buf first. 
func sender(c chan int, ack chan int, buf chan int, lock *sync.RWMutex) {
    for i:=0; i<10; i++ {
        c <- i
        lock.Lock()
        buf <- i // to record behavior sequence
        lock.Unlock()
        x := <-ack
        if x != i {
            panic("Error")
        }
    }
}
func receiver(c chan int, ack chan int, buf chan int, lock *sync.RWMutex) {
    for {
        x := <-c
        //time.Sleep(time.Millisecond) if unmark this, sender will write history first
        lock.Lock()
        buf <- -x-1 // to record behavior sequence 
        lock.Unlock()
        ack<-x
    }
}


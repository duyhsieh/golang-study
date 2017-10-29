package main
import ("fmt"; "sync")

func main() {
    UnSafe()
    Safe()
}

func UnSafe() {
    var wg sync.WaitGroup
    var counter int = 0
    for i:=0; i<1000; i++ {
        wg.Add(1)
        go Foo1(&wg, &counter)
    }
    wg.Wait()
    fmt.Println("UnSafe:counter=", counter)
}

func Safe() {
    var lock sync.RWMutex
    var wg sync.WaitGroup
    var counter int = 0
    for i:=0; i<1000; i++ {
        wg.Add(1)
        go Foo2(&wg, &lock, &counter)
    }
    wg.Wait()
    fmt.Println("Safe:counter=", counter)
}

func Foo1(wg *sync.WaitGroup, counter *int) {  // race conditon ! 
    defer wg.Done()
    for i:=0; i<10000; i++ {
        add(counter)
    }
}

func Foo2(wg *sync.WaitGroup, lock *sync.RWMutex, counter *int) { // safe , but not efficient (only for demo)
    defer wg.Done()
    for i:=0; i<10000; i++ {
        lock.Lock()
        add(counter)
        lock.Unlock()
    }
}

func add(i *int) {
    *i+=1
}

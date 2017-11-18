package main
import ("testing"; "sync")
//import "fmt"

type DATA uint32
//type DATA uint64

type Foo struct {
    X DATA
    Y DATA
}

func BenchmarkFalseSharing(b* testing.B) {
    var wg sync.WaitGroup
	var f Foo
	for i:=0; i< 100; i++ {
		go sum_a(&f, &wg)
		go inc_b(&f, &wg)
		wg.Add(2)
	}
	wg.Wait()
}
func sum_a(f *Foo, wg *sync.WaitGroup) {
	var s DATA = 0
	for i := 0; i < 1000000; i++ {
        s += f.X
	}
	wg.Done()
}

func inc_b(f *Foo, wg *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
        f.Y+=1
	}
	wg.Done()
}


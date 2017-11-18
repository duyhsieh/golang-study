package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sync"
	//"runtime"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	//session, err := mgo.Dial("localhost?socketTimeoutMS=1&maxPoolSize=1&waitQueueMultiple=1&waitQueueTimeoutMS=1")
    var wg sync.WaitGroup
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetPoolLimit(5)
	//session.SetSyncTimeout(time.Millisecond)
	//session.SetSocketTimeout(time.Millisecond)
	//runtime.GOMAXPROCS(2)
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	for i:=0; i<100; i++ {
		wg.Add(1)
		//go testWrite(session, i, &wg)
		go testWrite(CloneSession(session), i, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
	var dummy string
	fmt.Scanln(&dummy)
}

func testWrite(s *mgo.Session, no int , wg *sync.WaitGroup) {
	defer wg.Done()
	defer s.Close()
	for i:=0; i < 100; i++ {
		c:=s.DB("test").C("people")
		err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			log.Fatal(err)
		}

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("Phone:", result.Phone)
		time.Sleep(10*time.Millisecond)

	}
	fmt.Printf("goroutine %v fin\n", no)
}

// clone will create new Socket if there is no reusable sockets in pool. If you don't clone, new sockets won't be created automatically.
// created socket upper limit is set by SetPoolLimit function.
// However some helper connections will be set by driver to understand cluster info so there may be more connections than limit connected to DB. 
// These helper connections are not counted into connection pool limit. 
func CloneSession(session *mgo.Session) *mgo.Session {
	return session.Clone()
}

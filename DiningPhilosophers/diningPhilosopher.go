package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var philLimit int
var eating sync.WaitGroup
var xMutex sync.RWMutex

type CSticks struct {
	sync.Mutex
}

type Phil struct {
	leftCS  *CSticks
	rightCS *CSticks
}

func (p Phil) eat(philosopher int, eatingRequest chan int) {

	for i := 1; i <= 3; i++ {
		chopPicked := rand.Intn(5)
		eatingRequest <- 1
		if chopPicked < 5 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}
		fmt.Printf("Starting to eat %d  \n", philosopher+1)
		fmt.Printf("Finisihing to eat %d  \n", philosopher+1)

		if chopPicked < 5 {
			p.rightCS.Unlock()
			p.leftCS.Unlock()
		} else {
			p.leftCS.Unlock()
			p.rightCS.Unlock()
		}
		eating.Done()
		xMutex.Lock()
		philLimit -= 1
		xMutex.Unlock()
	}
	wg.Done()
}

func hostRoutine(eatingRequest chan int) {

	for {
		<-eatingRequest
		eating.Add(1)
		xMutex.Lock()
		philLimit += 1
		if philLimit > 2 {
			xMutex.Unlock()
			eating.Wait()
		} else {
			xMutex.Unlock()
		}
	}
}

func main() {
	eatingRequest := make(chan int, 1)
	ChopSticks := make([]*CSticks, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(CSticks)
	}
	philos := make([]*Phil, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Phil{leftCS: ChopSticks[i], rightCS: ChopSticks[(i+1)%5]}
	}
	wg.Add(5)
	go hostRoutine(eatingRequest)
	for i := 0; i < 5; i++ {
		go philos[i].eat(i, eatingRequest)
	}
	wg.Wait()
	fmt.Println("All Philosophers have eaten!!!")
}

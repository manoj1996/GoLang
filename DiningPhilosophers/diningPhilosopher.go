package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

type CSticks struct {
	sync.Mutex
}

type Phil struct {
	leftCS *CSticks
	rightCS *CSticks
	timesEaten int
	philNum int
}

func (p Phil) eat(){
	for p.timesEaten < 3 {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", p.philNum)
		p.timesEaten = p.timesEaten + 1
		fmt.Printf("finishing eating %d\n", p.philNum)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func main(){
	ChopSticks := make([]*CSticks, 5)
	for i := 0; i < 5; i++{
		ChopSticks[i] = new(CSticks)
	}
	philos := make([]*Phil, 5)
	for i := 0; i < 5; i++{
		philos[i] = &Phil{leftCS:ChopSticks[i], rightCS:ChopSticks[(i+1)%5], timesEaten:0, philNum:i+1}
	}
	for i := 0; i < 5; i++{
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}


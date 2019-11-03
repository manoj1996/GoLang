package main
import (
	"fmt"
	"strings"
	"time"
)
func decrementValue(v *int) {
	*v = *v - 1
}

func printPattern(v *int) {
	for *v > 0 {
		s := strings.Repeat("*", *v)
		fmt.Printf(s+"\n")
		decrementValue(v)
	}
}

func printValue(v *int){
	i := *v
	for i > 0{
		fmt.Printf("Value pointed by v: %d\n", *v)
		i = i - 1
	}
}


func main(){
	var v *int
	a := 5
	v = &a
	go printPattern(v)
	go printValue(v)
	fmt.Printf("In Main\n")
	time.Sleep(10*1000)
}


//Race conditions are where 2 threads are accessing memory at the same time
//one of which is writing
//Race conditions occur because of unsynchronized access to shared memory.
//Here two go routines are trying to access the shared memory in an unsynchronized way.
//Output each time is different based on the order of execution.

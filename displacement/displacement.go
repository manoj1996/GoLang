package main

import (
	"fmt"
	"math"
)

func main(){
	var acc, v0, s0 float64
	fmt.Printf("Enter Value of acceleration : ")
	_, _ = fmt.Scan(&acc)
	fmt.Printf("Enter the value of initial velocity : ")
	_, _ = fmt.Scan(&v0)
	fmt.Printf("Enter the value of initial displacement : ")
	_, _ = fmt.Scan(&s0)

	fn := GenDisplaceFn(acc, v0, s0)
	fmt.Println("Displacement after 3 seconds:", fn(3))
	fmt.Println("Displacement after 5 seconds:", fn(5))

}

func GenDisplaceFn(a, v0, s0 float64) func (t float64) float64 {
	f := func(t float64) float64{
		s := (0.5) * a * (math.Pow(t, 2)) + v0*t + s0
		return s
	}
	return f
}


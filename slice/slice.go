package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var flag int = 0
	var s string
	sl := make([]int, 3)
	count := 0
	for flag == 0 {
		fmt.Printf("Enter an integer to be added to the slice : ")
		_, _ = fmt.Scan(&s)
		if s == "X" {
			flag = 1
		} else {
			sInt, _ := strconv.Atoi(s)
			if count < 3 {
				sl[count] = sInt
				count++
				sort.Ints(sl[0:count])
				fmt.Printf("Slice sorted -> ")
				printSlice(sl[0:count], count)
			} else {
				sl = append(sl, sInt)
				count++
				sort.Ints(sl)
				fmt.Printf("Slice sorted -> ")
				printSlice(sl, count)
			}
		}
	}
}

func printSlice(s []int, l int){
	i := 0
	for i < l {
		fmt.Printf("%d ", s[i])
		i++
	}
	fmt.Printf("\n")
}
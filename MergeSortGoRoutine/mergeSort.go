package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func sortSubArray(arr []int, ch chan []int){
	sort.Ints(arr)
	fmt.Printf("SubArray sorted : %v\n", arr)
	ch <- arr
	wg.Done()
}
func mergeTwoArray(arr1 []int, arr2 []int, ch chan []int){
	i := 0
	j := 0
	res := make([]int, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2){
		if arr1[i] <= arr2[j]{
			res = append(res, arr1[i])
			i += 1
		} else{
			res = append(res, arr2[j])
			j += 1
		}
	}
	for i < len(arr1){
		res = append(res, arr1[i])
		i += 1
	}
	for j < len(arr2){
		res = append(res, arr2[j])
		j += 1
	}
	ch <- res
	wg.Done()
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a series of numbers to sort separated by space :")
	line, err := reader.ReadString('\n')
	if err != nil{
		fmt.Printf("Error occurred!!")
	}
	stringIntegers := strings.Split(strings.TrimSpace(line), " ")
	fmt.Printf("String integers:%v\n", stringIntegers)
	subArray1 := make([]int, 0)
	subArray2 := make([]int, 0)
	subArray3 := make([]int, 0)
	subArray4 := make([]int, 0)
	var il = 0
	for _, v := range stringIntegers{
		i, _ := strconv.Atoi(v)
		if il%4 == 0{
			subArray1 = append(subArray1, i)
		} else if il%4 == 1{
			subArray2 = append(subArray2, i)
		} else if il%4 == 2{
			subArray3 = append(subArray3, i)
		} else{
			subArray4 = append(subArray4, i)
		}
		il += 1
	}

	ch1 := make(chan []int, 1)
	ch2 := make(chan []int, 1)
	ch3 := make(chan []int, 1)
	ch4 := make(chan []int, 1)
	wg.Add(1)
	go sortSubArray(subArray1, ch1)
	wg.Add(1)
	go sortSubArray(subArray2, ch2)
	wg.Add(1)
	go sortSubArray(subArray3, ch3)
	wg.Add(1)
	go sortSubArray(subArray4, ch4)
	wg.Wait()
	s1 := <- ch1
	s2 := <- ch2
	s3 := <- ch3
	s4 := <- ch4



	ch5 := make(chan []int, 1)
	ch6 := make(chan []int, 1)
	wg.Add(1)
	go mergeTwoArray(s1, s2, ch5)
	wg.Add(1)
	go mergeTwoArray(s3, s4, ch6)
	wg.Wait()
	s12 := <- ch5
	s34 := <- ch6

	ch7 := make(chan []int, 1)
	wg.Add(1)
	go mergeTwoArray(s12, s34, ch7)
	wg.Wait()
	sortedIntegers := <- ch7
	//for _, i := range sortedIntegers{
	//	fmt.Printf("%d  ", i)
	//}
	fmt.Printf("Final Sorted list : %v\n", sortedIntegers[len(sortedIntegers)-len(stringIntegers):])
}

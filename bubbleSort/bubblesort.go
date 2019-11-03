package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Printf("Enter a sequence of upto 10 integers (space separated) : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Numbers := scanner.Text()
	stringIntegers := strings.Split(Numbers, " ")
	integers := make([]int, 0)
	for _, v := range stringIntegers{
		i, _ := strconv.Atoi(v)
		integers = append(integers, i)
	}
	BubbleSort(integers)
	sortedNumbers := make([]string, 0)
	for _, v := range integers{
		i := strconv.Itoa(v)
		sortedNumbers = append(sortedNumbers, i)
	}
	fmt.Printf("Array in sorted order -> ")
	fmt.Printf(strings.Join(sortedNumbers, " < ") + "\n")
}

func BubbleSort(arr []int){
	i := 0
	for i < len(arr){
		j := 0
		for j < len(arr) - i - 1{
			if arr[j] > arr[j+1]{
				Swap(arr, j)
			}
			j++
		}
		i++
	}
}

func Swap(arr []int, pos int){
	arr[pos], arr[pos+1] = arr[pos+1], arr[pos]
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main(){
	fmt.Printf("Enter the file name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	People := make([]Person, 0)
	fd, _ := os.Open(fileName)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, " ")
		localPerson := Person{fname:names[0], lname:names[1]}
		People = append(People, localPerson)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Iterating through struct:\n")
	for _, v := range People{
		fmt.Printf("First Name : %s  --- Last Name : %s\n", v.fname, v.lname)
	}
}

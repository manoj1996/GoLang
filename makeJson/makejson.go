package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main(){
	fmt.Printf("Enter name : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Enter address : ")
	scanner.Scan()
	address := scanner.Text()

	Person := make(map[string] string)

	Person["name"] = name
	Person["address"] = address

	jObject, _ := json.Marshal(Person)

	fmt.Printf("Person details as a JSON object: %s\n", jObject)

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

type Cow struct {
	Name string
}

type Bird struct {
	Name string
}

type Snake struct {
	Name string
}

func (cow Cow) GetName() string {
	return cow.Name
}

func (cow Cow) Eat() {
	fmt.Printf("%s\n", "grass")
}

func (cow Cow) Move() {
	fmt.Printf("%s\n", "walk")
}

func (cow Cow) Speak() {
	fmt.Printf("%s\n", "moo")
}

func (bird Bird) GetName() string {
	return bird.Name
}

func (bird Bird) Eat() {
	fmt.Printf("%s\n", "worms")
}

func (bird Bird) Move() {
	fmt.Printf("%s\n", "fly")
}

func (bird Bird) Speak() {
	fmt.Printf("%s\n", "peep")
}

func (snake Snake) GetName() string {
	return snake.Name
}

func (snake Snake) Eat() {
	fmt.Printf("%s\n", "mice")
}

func (snake Snake) Move() {
	fmt.Printf("%s\n", "slither")
}

func (snake Snake) Speak() {
	fmt.Printf("%s\n", "hsss")
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	userAnimals := make(map[string] Animal)
	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inp := strings.Split(text, " ")
		request := strings.ToLower(strings.TrimSpace(inp[0]))
		animalName := strings.TrimSpace(inp[1])
		if strings.Compare(request, "newanimal") == 0 {
			newAnimalCreated := true
			_, ok := userAnimals[animalName]
			if ok == true {
				fmt.Printf("Animal with the same name exists\n")
				continue
			}
			animalType := strings.ToLower(strings.TrimSpace(inp[2]))
			switch animalType {
			case "cow" : userAnimals[animalName] = Cow{Name:animalName}
			case "bird" : userAnimals[animalName] = Bird{Name:animalName}
			case "snake" : userAnimals[animalName] = Snake{Name:animalName}
			default:
				fmt.Printf("Invalid animal type specified")
				newAnimalCreated = false
			}
			if newAnimalCreated{
				fmt.Printf("Created it!\n")
			}
		} else if strings.Compare(request, "query") == 0 {
			_, ok := userAnimals[animalName]
			if ok == false{
				fmt.Printf("Invalid animal name queried - animal doesn't exist")
				continue
			}
			animalInfo := strings.ToLower(strings.TrimSpace(inp[2]))
			switch animalInfo {
			case "eat" : userAnimals[animalName].Eat()
			case "move" : userAnimals[animalName].Move()
			case "speak" : userAnimals[animalName].Speak()
			default:
				fmt.Printf("Invalid information requested\n")
			}
		}
	}
}



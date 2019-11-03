package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Animal struct {
	food string
	locomotion string
	noise string
}


func (animal *Animal) Eat() {
	fmt.Printf("Food eaten - %s\n", animal.food)
}

func (animal *Animal) Move() {
	fmt.Printf("Locomotion method - %s\n", animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Printf("Spoken sound - %s\n", animal.noise)
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	cow := Animal{food:"grass", locomotion:"walk", noise:"moo"}
	bird := Animal{food:"worms", locomotion:"fly", noise:"peep"}
	snake := Animal{food:"mice", locomotion:"slither", noise:"hsss"}
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		inp := strings.Split(text, " ")
		animalName := strings.ToLower(strings.TrimSpace(inp[0]))
		animalInfo := strings.ToLower(strings.TrimSpace(inp[1]))
		switch animalName {
		case "cow":
			switch animalInfo {
			case "eat" : cow.Eat()
			case "move" : cow.Move()
			case "speak" : cow.Speak()
			default:
				fmt.Printf("Invalid animal characteristic\n")
			}
		case "bird":
			switch animalInfo {
			case "eat" : bird.Eat()
			case "move" : bird.Move()
			case "speak" : bird.Speak()
			default:
				fmt.Printf("Invalid animal characteristic\n")
			}
		case "snake":
			switch animalInfo {
			case "eat" : snake.Eat()
			case "move" : snake.Move()
			case "speak" : snake.Speak()
			default:
				fmt.Printf("Invalid animal characteristic\n")
			}
		default:
			fmt.Printf("Invalid animal\n")

		}
	}
}
package main
import "fmt"

func main(){
  var num float64
  fmt.Printf("Enter a Floating point number : ")
  fmt.Scan(&num)
  var numInt = int(num)
  fmt.Printf("Truncated Integer version : %d\n",numInt)
}

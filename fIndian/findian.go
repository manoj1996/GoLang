package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"

)
func main(){

    fmt.Printf("Enter a string : ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    t := scanner.Text()
    var sLower = strings.ToLower(t)
    if strings.HasPrefix(sLower, "i"){
        if strings.HasSuffix(sLower, "n"){
            if strings.Contains(sLower, "a"){
                fmt.Printf("Found!\n")
            } else {
                fmt.Printf("Not Found!\n")
            }
        } else {
            fmt.Printf("Not Found!\n")
        }
    } else {
        fmt.Printf("Not Found!\n")
    }
}

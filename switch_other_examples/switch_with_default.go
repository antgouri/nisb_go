package main

import "fmt"
import "time"

func main() {

fmt.Println("Usage of switch")

switch time.Now().Weekday() {
case time.Saturday:
    fmt.Println("Today is Saturday.")
case time.Sunday:
    fmt.Println("Today is Sunday.")
default:
    fmt.Println("Today is a weekday.")
}

}

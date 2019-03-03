package main

import "fmt"
import "time"

func main() {

switch hour := time.Now().Hour(); { // missing expression(condition) means "true"
case hour < 12:
    fmt.Println("Good morning!")
case hour < 17:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}

}

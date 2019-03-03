package main

import "fmt"

func main() {
    fmt.Print("Enter your name: ")
    var name string
    fmt.Scanln(&name)
    fmt.Print("Your name is " +name)
}


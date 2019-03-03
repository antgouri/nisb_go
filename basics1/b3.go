package main

import "fmt"

func main() {
    fmt.Print("Enter your name: ")
    var name string
    fmt.Scanln(&name)
    fmt.Print("Your name is " +name)
    fmt.Print("\n")
    fmt.Print("Enter your age:")		
    var age int
    fmt.Scanln(&age)
    fmt.Println("Your age is ", age)
}


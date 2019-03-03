package main

import "fmt"

func main() {
    fmt.Print("Enter the number: ")
    var num int
    fmt.Scanln(&num)
    if num % 2 == 0 {
	fmt.Println("Even number")
	} else {
	fmt.Println("Odd number")
	}
}


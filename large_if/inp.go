package main

import "fmt"

func main() {
	fmt.Println("Largest of 2 numbers impl.")

	var a,b int
	
	fmt.Println("Enter the value of a")
	fmt.Scanln(&a)

	fmt.Println("Enter the value of b")
	fmt.Scanln(&b)
	
	if a > b{
	fmt.Println( a, " is largest")
	} else {
	fmt.Println( b, " is largest")
	} 
}

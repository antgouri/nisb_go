package main

import (
"fmt" 
"math" 
"strings"
)

func main() {
   greetings :=  []string{"Hello","world!"}   
   fmt.Println(strings.Join(greetings, " "))

   // The next 2 lines will print the value as 1024
   //fmt.Println(math.Exp2(10))
   //fmt.Println(math.Pow(2,10))
}

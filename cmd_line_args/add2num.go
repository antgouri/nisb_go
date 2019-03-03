package main

import "os"
import "strconv"
import "fmt"

func main() {


    arg1, err := strconv.Atoi(os.Args[1])
    arg2, err1 := strconv.Atoi(os.Args[2])

    //panic(err) is a better way to call...
    if err == nil && err1 == nil {
	fmt.Println("The sum is")
        var sum int
        sum = arg1+arg2
        fmt.Println(sum)
    }
    		
}


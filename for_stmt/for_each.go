package main

import "fmt"

func main() {

//With range 

strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}

//With maps

m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
for k, v := range m {
    fmt.Println(k, v)
}

//With channels

ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()
for n := range ch {
    fmt.Println(n)
}

}

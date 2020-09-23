package main

import "fmt"

func main() {
    someFunc := func (nums ...int) {
        fmt.Println("begin printing")
        fmt.Println(len(nums))

        for _, i := range nums {
            fmt.Println(i)
        }
        fmt.Println("done printing")
    }
    someFunc(1, 2, 3, 4, 5)
    someFunc(1, 2, 3)
}

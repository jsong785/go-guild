package main

import "fmt"

type Dog struct {
    name string
}

func (d *Dog) bark() string {
    return "bark" }

func (d Dog) yell() string {
    return "yell"
}

func main() {
    barkFunc := (*Dog).bark
    yellFunc := (Dog).yell

    var d = Dog{"mugsy"}
    fmt.Println(barkFunc(&d))
    fmt.Println(yellFunc(d))
}

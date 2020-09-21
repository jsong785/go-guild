package main

import (
    "fmt"
    "sync"
)

type something int

func (s something) DoSomething() int {
    return 99
}

type Example interface {
	AddTwo(x int) int
}

type Impl struct {
	x int
}

func (i *Impl) AddTwo(x int) int {
	return i.x + x + 2
}

func Initial() {
    var i Impl
    i.x = 3
    fmt.Println(i.AddTwo(3))

    var e Example = &i
    fmt.Println(e.AddTwo(3))

    var array1 [3]string
    array2 := [3]string{"a", "b", "c"}
    array1 = array2
    fmt.Println(array1)
    array2[1] = "z"
    fmt.Println(array1)

    slice := make([]int, 3, 5)
    sliceCopy := slice;
    slice[0] = 1
    slice[1] = 2
    slice[2] = 3
    slice2 := slice[1:3]
    slice3 := slice[1:2:3]
    //sliceCopyBefore := slice3
    slice3 = append(slice3, 100)
    sliceCopyBefore := slice3
    slice3 = append(slice3, 100)
    slice3 = append(slice3, 100)
    slice3 = append(slice3, 100)
    slice3 = append(slice3, 100)
    fmt.Println(slice3)
    sliceCopyAdd := slice3
    fmt.Println(sliceCopyAdd)
    slice3 = append(slice3, 99)

    fmt.Println("gotcha")
    fmt.Println(slice3)
    fmt.Println(sliceCopyBefore)
    fmt.Println(sliceCopyAdd)
    fmt.Println("gotcha")


    slice2 = append(slice2, 5)
    slice2 = append(slice2, 6)
    slice2 = append(slice2, 7)
    fmt.Println(slice)
    fmt.Println(slice2)

    slice = append(slice, 1)
    slice = append(slice, 1)
    slice = append(slice, 1)
    slice = append(slice, 1)
    slice = append(slice, 1)
    fmt.Println(slice)
    fmt.Println(sliceCopy)

    fmt.Println("new")

    multi := [][]int{{1}, {3, 4}}
    fmt.Println(multi)
    test := multi[0]
    multi[0] = append(multi[0], 1)
    multi[0] = append(multi[0], 1)
    multi[0] = append(multi[0], 1)
    fmt.Println(test)
    fmt.Println(multi)

    fmt.Println("new")

    val := something(100).DoSomething()
    fmt.Println(val)
}

type InterfaceTest interface {
    ReturnOne() int
}

type Person struct {
    name string `username`
}
func(*Person) ReturnOne() int {
    return 1;
}

type Car struct {
    Person
    model string `model`
}
func(Car) ReturnOne() int {
    return 11;
}

func main() {
    var x func(int) int;
    x = func(x int) int {
        return x+1;
    }
    fmt.Println("whateverett")
    fmt.Println(x(2))
    //p := Person{ "adam" }
    //c := Car{ Person{"adam"}, "bmw" }
    //fmt.Println(Notify(&c))
    //fmt.Println((&c).ReturnOne())
    //fmt.Println(Notify(c))
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        for _, a := range []string{"a", "b", "c", "d"} {
            fmt.Println(a)
        }
    }()

    go func() {
        defer wg.Done()
        for _, a := range []string{"a", "b", "c", "d"} {
            fmt.Println(a)
        }
    }()

    fmt.Println("waiting to finish")
    defer wg.Wait()
    fmt.Println("done")
}

func Notify(i InterfaceTest) int {
    return i.ReturnOne()
}



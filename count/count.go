package main

import (
    "fmt"
    "sync"
)

type Something struct {
    x int
}

func getMeAChannel(x interface{}) chan []interface{} {
    c := make(chan []interface{})
    go func() {
        defer close(c)
        list := make([]interface{}, 5)
        list[0] = x
        list[1] = x
        list[2] = x
        list[3] = x
        list[4] = x
        c <- list
    }()
    return c
}

func main() {
    comms := make(chan interface{})

    go func(c chan<- interface{}) {
        defer close(c)
        for _, i := range [5]int{1, 2, 3, 4, 5} {
            c <- i
        }
    }(comms)

    for c := range comms {
        fmt.Println(c)
    }
    fmt.Println("finished")

    waitSyncComms := sync.WaitGroup{}
    syncComms := make(chan interface{})

    waitSyncComms.Add(1)
    go func() {
        defer waitSyncComms.Done()
        <-syncComms
        fmt.Println("finished syncComs")
    }()

    close(syncComms)
    waitSyncComms.Wait()
    fmt.Println("done")

    cc := getMeAChannel(1)
    fmt.Println(<-cc)
}


package main

import (
    "fmt"
    "sync"
)

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
}


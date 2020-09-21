package main

import (
	"fmt"
)

func first(lettersChannel chan<- []string) {
	buf := make([]string, 0)
	for _, letter := range []string{"a", "b", "c", "d"} {
		buf = append(buf, letter)
	}
	lettersChannel <- buf
}

func second(lettersChannel chan []string) {
	results := <-lettersChannel
	for _, letter := range []string{"e", "f", "g"} {
		results = append(results, letter)
	}
	lettersChannel <- results
}

func main() {
	channel := make(chan []string)
        defer close(channel)

	go first(channel)
	go second(channel)

	fmt.Println("waiting")
	fmt.Println("getting results")
	fmt.Println(<-channel)
	fmt.Println("done")
}

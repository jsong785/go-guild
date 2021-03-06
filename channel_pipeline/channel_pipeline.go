package main

import (
	"fmt"
	"strconv"
	"time"
)

type Channel <-chan interface{}
type ChannelFunc func(Channel, Channel) Channel

func Bind(done Channel, input Channel, f ...ChannelFunc) Channel {
	if input == nil {
		return nil
	}
	if len(f) == 1 {
		return f[0](done, input)
	}
	return Bind(done, f[0](done, input), f[1:]...)
}

func multiplyPipe(factor int) ChannelFunc {
	return func(done Channel, input Channel) Channel {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for x := range OrDone(done, input) {
				c <- x.(int) * factor
			}
		}()
		return c
	}
}

func addPipe(add int) ChannelFunc {
	return func(done Channel, input Channel) Channel {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for x := range OrDone(done, input) {
				c <- x.(int) + add
			}
		}()
		return c
	}
}

func multiplyPipeSlow(factor int) ChannelFunc {
	return func(done Channel, input Channel) Channel {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for x := range OrDone(done, input) {
				time.Sleep(1 * time.Second)
				c <- x.(int) * factor
			}
		}()
		return c
	}
}

func addPipeSlow(add int) ChannelFunc {
	return func(done Channel, input Channel) Channel {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for x := range OrDone(done, input) {
				time.Sleep(1 * time.Second)
				c <- x.(int) + add
			}
		}()
		return c
	}
}

func toString(done Channel, input Channel) Channel {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for x := range OrDone(done, input) {
			c <- strconv.Itoa(x.(int))
		}
	}()
	return c
}

func test1() {
	done := make(chan interface{})
	defer close(done)

	msg := make(chan interface{})
	go func() {
		defer close(msg)
		for _, i := range []int{1, 2, 3, 4, 5} {
			msg <- i
		}
	}()

	// 2 4 6 8 10
	// 3 5 7 9 11
	// 13 15 17 19 21
	res := Bind(done, msg, multiplyPipe(2), addPipe(1), addPipe(10), toString)
	for o := range res {
		fmt.Println(o)
	}
}

func slowtest() {
	done := make(chan interface{})
	defer close(done)

	msg := make(chan interface{})
	go func() {
		defer close(msg)
		for _, i := range []int{1, 2, 3, 4, 5} {
			msg <- i
		}
	}()

	// 2 4 6 8 10
	// 3 5 7 9 11
	// 13 15 17 19 21
	res := Bind(done, msg, multiplyPipe(2), addPipeSlow(1), addPipeSlow(10), toString)
	for o := range res {
		fmt.Println(o)
	}
}

func faninouttest() {
	done := make(chan interface{})
	defer close(done)

	msg := make(chan interface{})
	go func() {
		defer close(msg)
		for _, i := range []int{1, 2, 3, 4, 5} {
			msg <- i
		}
	}()

	// 2 4 6 8 10
	// 3 5 7 9 11
	// 13 15 17 19 21
	res := Bind(done, msg, multiplyPipe(2), FanInAdaptor(5, addPipeSlow(1)), FanInAdaptor(5, addPipeSlow(10)), toString)
	for o := range res {
		fmt.Println(o)
	}
}

func main() {
	test1()

	start := time.Now()
	slowtest()
	fmt.Println(time.Since(start))

	start = time.Now()
	faninouttest()
	fmt.Println(time.Since(start))
}

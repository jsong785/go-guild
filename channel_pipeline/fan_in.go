package main

import "sync"

func OrDone(done Channel, c Channel) Channel {
	output := make(chan interface{})

	go func() {
		defer close(output)
		for {
			select {
			case <-done:
				return
			case res, ok := <-c:
				if ok == false {
					return
				}
				select {
				case output <- res:
				case <-done:
				}
			}
		}
	}()

	return output
}

func FanInAdaptor(n int, f ChannelFunc) ChannelFunc {
	return func(done Channel, c Channel) Channel {
		output := make(chan interface{})

		wg := sync.WaitGroup{}
		multiplex := func(cur Channel) {
			defer wg.Done()
			for i := range OrDone(done, cur) {
				output <- i
			}
		}

		for i := 0; i < n; i++ {
			wg.Add(1)
			go multiplex(f(done, c))
		}

		go func() {
			wg.Wait()
			close(output)
		}()
		return output
	}
}

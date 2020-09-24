package main

//import sync

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

/*
func FanInAdaptor(n int, f ChannelFunc) ChannelFunc {
    fanChannels := make([]Channel, n)
    for i, _ := range fanChannels {
        fanChannels[i] = make(Channel)
    }

    closeFanChannels := func() {
        for i, _ := range fanChannels {
            close(fanChannels[i])
        }
    }

    wg := sync.WaitGroup{}
    wg.Add(n)

    return func(done <-chan interface{}, c Channel) Channel {
        defer closeFanChannels()
        for i, _ := range fanChannels {
            fanChannels[i] = f(done, c)
        }

        output := make(Channel)

        for _, fc := range fanChannels {
            go func() {
                output<- fc
            }
        }
        return output
    }
}
*/

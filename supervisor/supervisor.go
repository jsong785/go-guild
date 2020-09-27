package main

import (
    "fmt"
    "time"
)

// a supervisor is:
// 1. a decorator with a done. it decorates that function with its' own done, so it can kill that thread while startin ga new one if it is  unresponsive or too slow. This is or'ed with the don eof course (incase another thread wants to kill the supervisor)
// this wrapped function with a montitor funciton, it montirs
// 1. if the decoracted funciton bein watche dhits. reset the timer and go
//2 .for each supplised pulse interval. send a heartbeat
// 3. if a timeout (set in decorate hits.. death! kill and go

func orDone(done ...<-chan interface{}) (<-chan interface{}) {
    var doneCombine func(done ...<-chan interface{}) (<-chan interface{})
    doneCombine = func(done ...<-chan interface{}) (<-chan interface{}) {
        if len(done) == 0 {
            return nil
        } else if len(done) == 1 {
            return done[0]
        }

        current := make(chan interface{})
        go func () {
            defer close(current)
            select {
            case <-done[0]:
            case <-done[1]:
            case <-doneCombine(append(done[2:], current)...):
            }
        }()
        return current
    }
    return doneCombine(done...)
}
type SuperviseFunc func(done <-chan interface{}, timeout time.Duration) (<-chan interface{})

func Supervise(done <-chan interface{}, timeout time.Duration, f SuperviseFunc) SuperviseFunc {
    return func(done <-chan interface{}, heartBeat time.Duration) (<-chan interface{}) {
        hb := make(chan interface{})

        go func() {
            watchedDone := make(chan interface{})
            watchedHeartBeat := f(orDone(watchedDone, done), timeout/2)

            pulse := time.Tick(heartBeat)
            for {
                timeoutSignal := time.After(timeout)

                select{
                    case <-watchedHeartBeat: // continue, everything is normal
                case <-pulse:
                    select {
                    case hb <- struct{}{}:
                    default:
                    }
                case <-timeoutSignal:
                    fmt.Println("timed out wee")
                    close(watchedDone)
                    watchedDone = make(chan interface{})
                    watchedHeartBeat = f(orDone(watchedDone, done), timeout/2)
                case <-done:
                    return
                }
            }
        }()

        return hb
    }
}

func main() {
}

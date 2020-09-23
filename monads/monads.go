package main

import (
    "errors"
    "fmt"
    "strconv"
)

type ErrorMonad struct {
    val interface{}
    err error
}
type ErrorMonadFunc func(val interface{}) ErrorMonad

func bind(val ErrorMonad, f ErrorMonadFunc) ErrorMonad {
    if val.err != nil {
        return ErrorMonad{ err: val.err };
    }
    return f(val.val)
}

func AddOne(x interface{}) ErrorMonad {
    number := x.(int)
    if number < 1 {
        return ErrorMonad{ err: errors.New("number must be greater than 0") }
    }
    return ErrorMonad{ val: number+1 }
}

func ToString(number interface{}) ErrorMonad {
    return ErrorMonad{ val: strconv.Itoa(number.(int)) }
}

func main() {
    fmt.Println(bind(AddOne(1), AddOne))
    fmt.Println(bind(AddOne(-1), AddOne))
    fmt.Println(bind(AddOne(100), AddOne))

    fmt.Println(bind(AddOne(1), ToString))
    fmt.Println(bind(AddOne(-1), ToString))
    fmt.Println(bind(AddOne(100), ToString))
}


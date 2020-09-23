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

func bind(val ErrorMonad, f ...ErrorMonadFunc) ErrorMonad {
    if val.err != nil {
        return ErrorMonad{ err: val.err };
    }
    if len(f) == 1 {
        return f[0](val.val)
    }
    return bind(f[0](val.val), f[1:]...)
}

func AddOne(x interface{}) ErrorMonad {
    number := x.(int)
    if number < 1 {
        return ErrorMonad{ err: errors.New("number must be greater than 0") }
    }
    return ErrorMonad{ val: number+1 }
}

func SubtractOne(x interface{}) ErrorMonad {
    return ErrorMonad{ val: x.(int) - 1 };
}

func ToString(number interface{}) ErrorMonad {
    return ErrorMonad{ val: strconv.Itoa(number.(int)) }
}

func main() {
    // success
    fmt.Println(bind(AddOne(1), AddOne))
    fmt.Println(bind(AddOne(100), AddOne, AddOne, ToString))
    fmt.Println(bind(AddOne(1), SubtractOne, AddOne, ToString));

    // errors
    fmt.Println(bind(AddOne(1), SubtractOne, SubtractOne, AddOne));
    fmt.Println(bind(AddOne(1), SubtractOne, SubtractOne, AddOne, ToString));

    fmt.Println(bind(AddOne(-1), AddOne))
    fmt.Println(bind(AddOne(-1), AddOne, AddOne, ToString))

}


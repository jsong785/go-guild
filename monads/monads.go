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

func Bind(val ErrorMonad, f ...ErrorMonadFunc) ErrorMonad {
    if val.err != nil {
        return ErrorMonad{ err: val.err };
    }
    if len(f) == 1 {
        return f[0](val.val)
    }
    return Bind(f[0](val.val), f[1:]...)
}

func addOne(x interface{}) ErrorMonad {
    number := x.(int)
    if number < 1 {
        return ErrorMonad{ err: errors.New("number must be greater than 0") }
    }
    return ErrorMonad{ val: number+1 }
}

func subtractOne(number interface{}) ErrorMonad {
    return ErrorMonad{ val: number.(int) - 1 };
}

func toString(number interface{}) ErrorMonad {
    return ErrorMonad{ val: strconv.Itoa(number.(int)) }
}

func main() {
    // success
    fmt.Println(Bind(addOne(1), addOne))
    fmt.Println(Bind(addOne(100), addOne, addOne, toString))
    fmt.Println(Bind(addOne(1), subtractOne, addOne, toString));

    // errors
    fmt.Println(Bind(addOne(1), subtractOne, subtractOne, addOne));
    fmt.Println(Bind(addOne(1), subtractOne, subtractOne, addOne, toString));

    fmt.Println(Bind(addOne(-1), addOne))
    fmt.Println(Bind(addOne(-1), addOne, addOne, toString))

}


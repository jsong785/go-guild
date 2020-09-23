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
type ErrorMonadFunc func(interface{}) ErrorMonad

func Bind(val ErrorMonad, f ...ErrorMonadFunc) ErrorMonad {
	if val.err != nil {
		return ErrorMonad{err: val.err}
	}
	if len(f) == 1 {
		return f[0](val.val)
	}
	return Bind(f[0](val.val), f[1:]...)
}

func addOnePositiveOnly(x interface{}) ErrorMonad {
	number := x.(int)
	if number < 0 {
		return ErrorMonad{err: errors.New("number must be positive")}
	}
	return ErrorMonad{val: number + 1}
}

func subtractOne(number interface{}) ErrorMonad {
	return ErrorMonad{val: number.(int) - 1}
}

func toString(number interface{}) ErrorMonad {
	return ErrorMonad{val: strconv.Itoa(number.(int))}
}

func main() {
	// success
	fmt.Println(Bind(addOnePositiveOnly(0), addOnePositiveOnly))
	fmt.Println(Bind(addOnePositiveOnly(100), addOnePositiveOnly, addOnePositiveOnly, toString))
	fmt.Println(Bind(addOnePositiveOnly(0), subtractOne, addOnePositiveOnly, toString))

	// errors
	fmt.Println(Bind(addOnePositiveOnly(0), subtractOne, subtractOne, addOnePositiveOnly))
	fmt.Println(Bind(addOnePositiveOnly(0), subtractOne, subtractOne, addOnePositiveOnly, toString))

	fmt.Println(Bind(addOnePositiveOnly(-1), addOnePositiveOnly))
	fmt.Println(Bind(addOnePositiveOnly(-1), addOnePositiveOnly, addOnePositiveOnly, toString))

}

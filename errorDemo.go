package main

import (
	"fmt"
	"errors"
)

func main() {
	// basic errors
	for _, v := range []int{200, 404} {
		if r, e := basicErrFunc(v); e != nil {
			fmt.Println("work fail.", e)
		} else {
			fmt.Println("work succ.", r)
		}
	}

	// self-defined errors
	for _, v := range []int{200, 404} {
		if r, e := selfErrFunc(v); e != nil {
			fmt.Println("work fail.", e)
		} else {
			fmt.Println("work succ.", r)
		}
	}

	//If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error type via type assertion.
	_, e := selfErrFunc(500)
	if err, ok := e.(*selfError); ok {
		fmt.Println(err.arg, err.prob)
	}
}

/**
 * By convention, errors are the last return value and have type error, a built-in interface.
 * errors.New constructs a basic error value with the given error message.
 * A nil value in the error position indicates that there was no error.
 */
func basicErrFunc(arg int) (int, error) {
	if arg == 404 {
		return -1, errors.New("not found")
	} else if arg == 500{
		return -1, errors.New("system error")
	} else if arg >= 300 {
		return -1, errors.New("other errors")
	}
	return arg, nil
}

/**
 * It’s possible to use custom types as errors by implementing the Error() method on them.
 */
func selfErrFunc(arg int) (int, error) {
	if arg == 404 {
		return -1, &selfError{arg:arg, prob:"not found"}
	} else if arg == 500{
		return -1, &selfError{arg:arg, prob:"system error"}
	} else if arg >= 300 {
		return -1, &selfError{arg:arg, prob:"other errors"}
	}
	return arg, nil
}

type selfError struct {
	arg  int
	prob string
}

func (e *selfError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

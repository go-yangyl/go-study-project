package main

import "fmt"

func main() {
	err := DoSthWithBlock(func(a string) error { return do(a) })

	fmt.Println(err)
}

func DoSthWithBlock(fc func(a string) error) (err error) {
	err = fc("hello")
	if err == nil {
		fmt.Println("error is nil")
	} else {
		fmt.Println("error isn't nil")
	}
	return
}

type MyError struct {
	Code string
	Msg  string
}

type MyErrP = *MyError

func (e *MyError) Error() string {
	return e.Code + e.Msg
}

func do(a string) error {
	return nil
}

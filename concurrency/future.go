package main

import (
	"errors"
	"fmt"
	"time"
)

type OnSuccessFunc func(string)
type OnErrorFunc func(error)
type ExecuteFuncForString func() (string, error)

type FutureForString struct {
	successFunc OnSuccessFunc
	failFunc    OnErrorFunc
}

func dummySuccessFunc(_ string) {}
func dummyErrorFunc(_ error)    {}

func NewFutureForString() *FutureForString {
	return &FutureForString{
		successFunc: dummySuccessFunc,
		failFunc:    dummyErrorFunc,
	}
}

func (fut *FutureForString) Success(f OnSuccessFunc) *FutureForString {
	fut.successFunc = f
	return fut
}

func (fut *FutureForString) Fail(f OnErrorFunc) *FutureForString {
	fut.failFunc = f
	return fut
}

func (fut *FutureForString) Execute(f ExecuteFuncForString) {
	go func(fun ExecuteFuncForString) {
		str, err := fun()

		if err != nil {
			fut.failFunc(err)
		} else {
			fut.successFunc(str)
		}
	}(f)
}

func forSuccess() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "OK", nil
}

func forError() (string, error) {
	time.Sleep(time.Millisecond * 400)
	return "", errors.New("Error occured")
}

func main() {
	fmt.Println("Future design pattern")

	future := NewFutureForString()
	future.Success(func(value string) {
		fmt.Println("Finished with Success! Value = ", value)
	}).Fail(func(err error) {
		fmt.Println("Finished with Error! Error = ", err)
	})

	future.Execute(forSuccess)

	future.Execute(forError)

	time.Sleep(time.Second * 2)
}

package main

import "fmt"

// Exception is using for error data
type Exception interface{}

// FlowControl is using to make readable flow control
type FlowControl struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Call is a method using to call and start the flow
func (flow FlowControl) Call() {
	if flow.Finally != nil {
		defer flow.Finally()
	}

	if flow.Catch != nil {
		defer func() {
			if err := recover(); err != nil {
				flow.Catch(err)
			}
		}()
	}

	flow.Try()
}

// Throw is similar like throwing and error
func Throw(err Exception) {
	panic(err)
}

func main() {
	FlowControl{
		Try: func() {
			fmt.Println("Try...")
			Throw("err: oopss, something went wrong...")
		},
		Catch: func(err Exception) {
			fmt.Println("I'm catching the error, here")
			fmt.Println(err)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Call()
}

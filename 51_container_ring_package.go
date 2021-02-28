package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// circular
	cl := ring.New(5)
	for i := 0; i < cl.Len(); i++ {
		cl.Value = i
		cl = cl.Next()
	}

	cl.Do(func(value interface{}) {
		fmt.Println(value)
	})
}

/**
 * Flow control
 * - defer => finally
 * - panic => throw
 * - recover => catch
 *
 * Recover must be within the defer function
 */

package main

import "fmt"

func main() {
	isError := true

	defer func() {
		err := recover()
		fmt.Println(err)

		fmt.Println("Finally...")
	}()

	fmt.Println("Try...")
	if isError {
		panic("Error data")
	}
}

/**
Type assertion is a way to convert interface data to a specific type
Be careful to use this way, because if you assign wrong data type it will throw an error
*/

package main

import (
	"fmt"
	"reflect"
)

func returnValue(v interface{}) interface{} {
	return v
}

func main() {
	stringValue := returnValue("Hello").(string)
	intValue := returnValue(53).(int)
	boolValue := returnValue(true).(bool)
	// errorTypeAssertion := returnValue(12).(bool)

	fmt.Println(stringValue)
	fmt.Println(intValue)
	fmt.Println(boolValue)

	// get type
	stringType := reflect.TypeOf(stringValue).String()
	int16Type := reflect.TypeOf(int16(5)).String()
	fmt.Println(stringType == "string")
	fmt.Println(int16Type == "int16")

	// switch type
	// for data type conditional, you must use switch statement (no if else)
	// you better use slice / map instead of array for type conditional
	var data interface{} = 1
	switch data.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unknown")
	}
}

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Sample is sample type
type Sample struct {
	Name string `required:"true" max:"10"` // struct tage
	Age  int16  `required:"true" max:"100"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		requiredTagValue, _ := strconv.ParseBool(field.Tag.Get("required"))

		fieldValue := reflect.ValueOf(data).Field(i).Interface()
		if requiredTagValue && fieldValue == "" {
			return false
		}
	}
	return true
}

func main() {
	// getting field information
	sample := Sample{"Rico", 22}
	sampleType := reflect.TypeOf(sample)
	fieldLength := sampleType.NumField()

	fmt.Println(sampleType)
	fmt.Println(fieldLength)

	// getting field information
	nameField := sampleType.Field(0)
	ageField := sampleType.Field(1)
	fmt.Println(nameField.Name, nameField.Type, nameField.Type.String() == "string")
	fmt.Println(ageField.Name, ageField.Type, ageField.Type.String() == "int16")

	// getting struct tag information
	// tag data always returns string type, so parse it first before you get real type data
	fmt.Println(nameField.Tag.Get("required"), nameField.Tag.Get("max"))
	fmt.Println(ageField.Tag.Get("required"), ageField.Tag.Get("max"))

	// parsed data
	requiredNameField, _ := strconv.ParseBool(nameField.Tag.Get("required"))
	maxNameField, _ := strconv.ParseInt(nameField.Tag.Get("max"), 10, 16)

	fmt.Println(requiredNameField, maxNameField)
	fmt.Println(reflect.TypeOf(requiredNameField), reflect.TypeOf(maxNameField))

	// simple validation
	fmt.Println(isValid(Sample{"Rico", 22}))
	inValidSample := Sample{}
	inValidSample.Age = 22
	fmt.Println(isValid(inValidSample))
}

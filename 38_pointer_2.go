package main

import "fmt"

// Address is using to define address data
type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Sampang", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.city = "Surabaya"

	// when address2 reassign to the new address data
	// all of the references won't reference to the same data
	// so it's like creating new memory allocation for the new data
	address2 = &Address{"Jakarta", "Jawa Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	// somtimes we want to change all the references
	// so you need (*) operator
	address3 := &address1

	// so when we reassign address3 to the new value, address1 also changing
	*address3 = Address{"New City", "New Province", "New Country"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	myAddress := Address{"Sampang", "Jatim", "Indonesia"}
	yourAddress := &myAddress
	hisAddress := *yourAddress // pass by value

	yourAddress.country = "Indon"

	fmt.Println(myAddress)    // Indon
	fmt.Println(*yourAddress) // Indon
	fmt.Println(hisAddress)   // Indonesia
}

/**
& -> converting data into its memory allocation (0x8...)
* (in data type) -> you can use this as memory allocation type (*int, *Person, *interface{}, ...)
*/

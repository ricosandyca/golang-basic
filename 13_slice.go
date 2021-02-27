/**
 * Slice is created from array
 * Which Slice has same memory alloc with the based array
 *
 * Details of slice
 * - pointer => pointer to the first data of an array
 * - length => length of the slice => len(slice)
 * - capacity => capacity of the slice (count from first index to the last index of array) => cap(slice)
 *
 * Slice creation format
 * - array[low:high] - creating a slice from array starts from low index to index before high
 * - array[low:] - creating a slice from array starts from low index to the last index of the array
 * - array[:high] - creating a slice from array starts from the first index to index before high
 * - array[:] - creating a slice from array starts from the first index to the last index of the array
 * - variable_name := []int{value,value,...}
 */

package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	februaryToJuly := months[1:6]
	fmt.Println("February to June", februaryToJuly)

	augustToDecember := months[7:]
	fmt.Println("August to December", augustToDecember)

	threeMonthsAfterJanuary := months[1 : 1+3]
	fmt.Println("Three months after January", threeMonthsAfterJanuary)
}

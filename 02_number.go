/**
 * Integer data type
 * - int8: -128 - 127
 * - int16: -32768 - 32767
 * - int32: -2147483648 - 22147483647
 * - int65: -9223372036854775808 - 9223372036854775807
 *
 * Unsigned integer data type
 * - uint8: 0 - 255
 * - uint16: 0 - 65535
 * - uint32: 0 - 4294967295
 * - uint64: 0 - 18446744073709551615
 *
 * Floating point data type
 * - float32: 1.18x10^-38 - 3.4x10^38
 * - float64: 2.23x10^-308 - 1.80x10^308
 * - complex64: complex numbers with float32 real and imaginary parts
 * - complex128: complex numbers with float64 real and imaginary parts
 *
 * Alias
 * - byte = uint8
 * - rune = int32
 * - int = depends on the system architecture (x32 / x64) with minimal int32
 * - uint = depends on the system architecture (x32 / x64) with minimal uint32
 */

package main

import "fmt"

func main() {
	fmt.Println("Satu", 1)
	fmt.Println("Dua", 2)
	fmt.Println("Tiga Koma Lima", 3.5)
}

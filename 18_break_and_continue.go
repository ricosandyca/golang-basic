/**
 * Break will stop the parent looping where the break script put
 * To stop a specific looping target, you need to assign alias to the looping target
 *
 * Continue will stop the looping in the current index and get into the next index
 * Continue is like skipping current index
 * To skip a specific looping target, you need to assign alias to the looping target
 */

package main

import "fmt"

func main() {
	// break
	for i := 1; i <= 3; i++ {
		fmt.Println("i", i)
		for j := 1; j <= 3; j++ {
			fmt.Println("j", j)
			if j == 2 {
				// break the j looping
				break
			}
		}
	}

	fmt.Println("---")

	// break with looping alias
i:
	for i := 0; i <= 3; i++ {
		fmt.Println("i", i)
		for j := 1; j <= 3; j++ {
			fmt.Println("j", j)
			if j == 2 {
				// break the i looping (specific looping aliases)
				break i
			}
		}
	}

	fmt.Println("-------")

	// continue
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		for j := 0; j <= 3; j++ {
			fmt.Println("i j", i, j)
		}
	}

	fmt.Println("-------")

	// continue with looping alias
a:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue a
			}
			fmt.Println("i j", i, j)
		}
	}

}

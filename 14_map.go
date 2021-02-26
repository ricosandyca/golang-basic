/**
 * Map data type creation format
 * - var variable_name = map[key_type]value_type{[key]:value,}
 *
 * Map function
 * - len(map) => count data in map
 * - delete(map, key) => delete data by the key
 * - make(map[key_type]value_type) => make new map data
 */

package main

import "fmt"

func main() {
	var me = map[string]string{
		"name": "Rico",
		"age":  "21",
	}
	me["name"] = "Rico Sandyca Novenza"
	me["role"] = "Software Engineer"
	me["anu"] = "Anuan"

	fmt.Println("Details of me", me)

	delete(me, "anu")
	fmt.Println("Details of me", me)
}

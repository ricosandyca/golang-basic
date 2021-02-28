package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UserSlice []User

func (user UserSlice) Len() int {
	return len(user)
}

// 2,1,3

func (user UserSlice) Less(i, j int) bool {
	// true to keep position
	// false to swap the position

	// sort by age
	return user[i].age < user[j].age
}

func (user UserSlice) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

func main() {
	users := []User{
		{"Rico", 23},
		{"Sandyca", 22},
		{"Novenza", 24},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}

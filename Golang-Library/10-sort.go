package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Usia int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Usia < s[j].Usia
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	users := []User{
		{"Adrik", 30},
		{"Misbah", 25},
		{"Muhammad", 14},
		{"Misbah", 27},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}

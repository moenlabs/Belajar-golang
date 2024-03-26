package main

import (
	"fmt"
	"slices"
)

func main() {

	names := []string{"Dafi", "Misbah", "Muhammad", "Nur"}
	nilai := []int{100, 90, 95, 80}

	fmt.Println(slices.Min(nilai))
	fmt.Println(slices.Max(nilai))
	fmt.Println(slices.Contains(names, "Misbah"))
	fmt.Println(slices.Index(names, "Dafi"))
}

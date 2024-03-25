package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("misbahul munir", "munir"))
	fmt.Println(strings.Split("misbahul munir", " "))
	fmt.Println(strings.ToLower("Misbahul Munir"))
	fmt.Println(strings.ToUpper("Misbahul Munir"))
	fmt.Println(strings.Trim("      misbahul munir     ", " "))
	fmt.Println(strings.ReplaceAll("Misbahul Munir Misbahul Adrik", "Misbahul", "Muhammad"))
}

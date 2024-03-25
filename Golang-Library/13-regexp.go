package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex = regexp.MustCompile(`m([a-z])h`)
	fmt.Println(regex.MatchString("mih"))
	fmt.Println(regex.MatchString("munir"))
	fmt.Println(regex.MatchString("muh"))
}

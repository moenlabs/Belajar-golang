package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolean, err := strconv.ParseBool("false")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	result, err := strconv.Atoi("10000")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	binary := strconv.FormatInt(777, 2)

	fmt.Println(binary)

	stringInt := strconv.Itoa(777)
	fmt.Println(stringInt)

}

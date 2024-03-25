package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for a := 0; a < data.Len(); a++ {
		data.Value = "Value " + strconv.Itoa(a)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})

}

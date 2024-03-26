package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("tis is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Assalamu'alaikum\n")
	_, _ = writer.WriteString("Mari Belajar\n")

	writer.Flush()
}

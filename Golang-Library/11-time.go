package main

import (
	"fmt"
	"time"
)

func main() {
	sekarang := time.Now()
	fmt.Println(sekarang.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:052")
	fmt.Println(parse)

	// format waktu

	formatWaktu := "2006-01-02 15:04:05"

	isian := "2020-10-10 10:10:10"

	isianWaktu, err := time.Parse(formatWaktu, isian)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(isianWaktu)
	}

	// durasi

	durasiMenit := time.Minute * 10
	var durasiDetik time.Duration = 30 * time.Second

	fmt.Println(durasiMenit, durasiDetik)
}

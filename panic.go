package main

import "os"

func main() {
	panic("a panic")

	_, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}
}
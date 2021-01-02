package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\nellery\n")
	err := ioutil.WriteFile("./data.txt", d1, 0644)
	check(err)

	f, err := os.Create("./data.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes: %s\n", n2, string(d2))

	n3, err := f.WriteString("writing\n")
	fmt.Printf("worte %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("Wrote %d bytes \n", n4)

	w.Flush()
}
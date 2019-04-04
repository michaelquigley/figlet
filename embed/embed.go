package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("var data = [%d]byte{\n\t", len(bytes))

	for i, b := range bytes {
		fmt.Printf("0x%02x, ", int(b))
		if (i+1) % 24 == 0 {
			fmt.Println()
			fmt.Print("\t")
		}
	}

	fmt.Println()
	fmt.Println("}")
}
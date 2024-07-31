package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("eicar.txt")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(bytes))
	fmt.Println("HAHAHA HACKED!!!")
	fmt.Println("-Timothy")
}

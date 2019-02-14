package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\ninsert string: ")
	n, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("wrong input")
		panic(err)
	}
	splited := strings.Split(n)
	var sublen int
	var lastchar string
	for _, element := range splited {
		if element == lastchar {
			sublen++
		} else {
			lastchar = element
			sublen = 1
		}
	}
	fmt.Println("longest substring lenght: ", sublen)
}

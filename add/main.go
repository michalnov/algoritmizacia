package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var num1 num
	var num2 num

	num1.prepare()
	num2.prepare()

	result := add(num1, num2)
	result.toString()
	fmt.Println("result is: ", result.s)
}

// if n1 < 0 (if n1.signed) && n2 > 0 swap n1, n2
func add(n1 num, n2 num) num {
	var out num
	var tempInt int
	if n1.signed && !n2.signed {
		var swap num
		swap = n1
		n1 = n2
		n2 = swap
		for n1.hasNext() || n2.hasNext() {

		}
	} else if n1.signed && n2.signed {
		out.signed = true
		tempInt = tempInt + n1.element + n2.element
		if tempInt > 10 {
			out.nums = append(out.nums, tempInt%10)
			tempInt = tempInt / 10
		} else {
			out.nums = append(out.nums, tempInt)
			tempInt = 0
		}
	} else if !n1.signed && n2.signed {
		//treba dokoncit
		//tempInt = tempInt + n1.element + n2.element
		//if tempInt > 10 {
		//	out.nums = append(out.nums, tempInt%10)
		//	tempInt = tempInt / 10
		//} else {
		//	out.nums = append(out.nums, tempInt)
		//	tempInt = 0
		//}
	}
	for n1.hasNext() || n2.hasNext() {
		tempInt = tempInt + n1.element + n2.element
		if tempInt > 10 {
			out.nums = append(out.nums, tempInt%10)
			tempInt = tempInt / 10
		} else {
			out.nums = append(out.nums, tempInt)
			tempInt = 0
		}
	}
	return out
}

type num struct {
	s       string
	n       []string
	nums    []int
	signed  bool
	element int
	counter int
	len     int
}

func (n *num) hasNext() bool {
	if n.counter < n.len {
		n.element = n.nums[n.counter]
		n.counter++
		return true
	}
	n.element = 0
	return false
}

func (n *num) toString() {
	if n.signed {
		n.s = "-"
	} else {
		n.s = ""
	}
	for _, element := range n.nums {
		n.s += string(element)
	}
}

func (n *num) prepare() {
	n.read()
	n.counter = 0
	n.len = len(n.n)
}

func (n *num) read() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\ninsert first number: ")
	n.s, err = reader.ReadString('\n')
	if err != nil || !n.verify() {
		fmt.Println("wrong input")
		panic(err)
	}
}

func (n *num) verify() bool {
	if len(n.s) < 1 || len(n.s) > 80 {
		return false
	}
	n.split()
	for index, element := range n.n {
		if index == 0 && element == "-" {
			n.signed = true
		}
		i, err := strconv.Atoi(element)
		if err != nil {
			return false
		}
		n.nums = append(n.nums, i)
	}
	return true
}

func (n *num) split() {
	n.n = strings.Split(n.s, "")
	if n.n[0] == "-" {
		n.signed = true
	}
}

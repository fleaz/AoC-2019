package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func format(progString string) []int {
	var prog []int
	for _, s := range strings.Split(progString, ",") {
		cmd, _ := strconv.Atoi(s)
		prog = append(prog, cmd)
	}
	return prog
}

func execute(prog []int) []int {
	pos := 0
	for {
		cmd := prog[pos]
		switch cmd {
		case 1:
			left := prog[pos+1]
			right := prog[pos+2]
			target := prog[pos+3]
			prog[target] = prog[left] + prog[right]
			pos += 4
		case 2:
			left := prog[pos+1]
			right := prog[pos+2]
			target := prog[pos+3]
			prog[target] = prog[left] * prog[right]
			pos += 4
		case 99:
			return prog
		}
	}

}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var progString string

	for scanner.Scan() {
		progString = scanner.Text()
	}
	prog := format(progString)
	prog[1] = 12
	prog[2] = 2

	result := execute(prog)
	fmt.Printf("Result is %d\n", result[0])
}

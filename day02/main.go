package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func execute(progString string) string {
	var prog []int
	pos := 0
	for _, s := range strings.Split(progString, ",") {
		cmd, _ := strconv.Atoi(s)
		prog = append(prog, cmd)
	}

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
			return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(prog)), ","), "[]")
		}
	}

}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var programm string

	for scanner.Scan() {
		programm = scanner.Text()
	}

	fmt.Printf("Result is %q\n", execute(programm))

}

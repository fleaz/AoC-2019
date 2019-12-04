package main

import (
	"fmt"
	"strconv"
)

func isValid(password int) bool {
	pwString := strconv.Itoa(password)

	//Length criteria
	if len(pwString) != 6 {
		return false
	}

	// adjavent digits criteria
	pair := false
	for i := 0; i < len(pwString)-1; i++ {
		left := pwString[i]
		right := pwString[i+1]
		if left == right {
			pair = true
			break
		}
	}
	if pair == false {
		return false
	}

	// incresing numbers criteria
	for i := 0; i < len(pwString)-1; i++ {
		left, _ := strconv.Atoi(string(pwString[i]))
		right, _ := strconv.Atoi(string(pwString[i+1]))
		if left > right {
			return false
		}
	}

	// All check passed
	return true
}

func findValidPasswords(from int, to int) int {
	counter := 0
	for i := from; i <= to; i++ {
		if isValid(i) {
			counter++
		}
	}

	return counter
}

func main() {
	from := 359282
	to := 820401

	count := findValidPasswords(from, to)
	fmt.Printf("Result for part 1 is %d\n", count)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	fuel := math.Floor(float64(mass)/3) - 2
	return int(fuel)
}

func complexCalculateFuel(mass int) int {
	totalFuel := 0
	for {
		f := calculateFuel(mass)
		if f > 0 {
			totalFuel += f
			mass = f
		} else {
			return totalFuel
		}

	}
}

func main() {
	totalFuel := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		totalFuel += complexCalculateFuel(m)
	}

	fmt.Printf("Total fuel consumtion is %d\n", totalFuel)

}

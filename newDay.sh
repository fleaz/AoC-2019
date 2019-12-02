#! /usr/bin/env bash

DAY="$( printf '%02d' "$1" )"
FOLDER="day$DAY"

mkdir $FOLDER
touch $FOLDER/input.txt
cat <<EOT >> $FOLDER/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        fmt.Println(scanner.Text()
	}
}
EOT

cat <<EOT >> $FOLDER/main_test.go
package main

import "testing"

type testpair struct {
	value  int
	want   int
}

func Test(t *testing.T) {
	pairs := []testpair{
		{,},
		{,},
		{,},
	}

	for _, p := range pairs {
        is := functioncall()
		if is != p.want {
			t.Errorf("Wanted %d but got %d", p.want, is)
		}

	}
}
EOT

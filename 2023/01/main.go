package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(file), "\n")

	sum := 0
	for _, v := range content {
		x, err := decode(v)
		if err != nil {
			panic(err)
		}
		sum += x
	}

	fmt.Println(sum)
}

func decode(s string) (int, error) {
	first, last := -1, -1

	for _, v := range s {
		if v >= 48 && v <= 57 {
			if first < 0 {
				first = int(v) - 48
			}

			last = int(v) - 48
		}
	}

	if first < 0 || last < 0 {
		return -1, fmt.Errorf("negative value: first: %v, last: %v", first, last)
	}

	return first*10 + last, nil
}

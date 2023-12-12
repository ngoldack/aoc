package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadFile() []string {
	file, err := os.ReadFile("1.input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}

func main() {
	content := loadFile()

	sum := 0
	for i, row := range content {
		splitted := strings.Split(strings.Split(row, ":")[1], ";")
		valid := validate(splitted)

		if valid {
			sum += i + 1
		}
	}

	fmt.Println(sum)
}

var max = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func validate(raw []string) bool {
	for _, grab := range raw {
		balls := strings.Split(grab, ",")

		for _, ball := range balls {
			ball = strings.ReplaceAll(ball, " ", "")

			if index := strings.Index(ball, "red"); index != -1 {
				num, err := strconv.Atoi(ball[:index])
				if err != nil {
					panic(err)
				}

				if num > max["red"] {
					return false
				}
			}

			if index := strings.Index(ball, "green"); index != -1 {
				num, err := strconv.Atoi(ball[:index])
				if err != nil {
					panic(err)
				}

				if num > max["green"] {
					return false
				}
			}

			if index := strings.Index(ball, "blue"); index != -1 {
				num, err := strconv.Atoi(ball[:index])
				if err != nil {
					panic(err)
				}

				if num > max["blue"] {
					return false
				}
			}
		}
	}

	return true
}

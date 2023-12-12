package main

import (
	"fmt"
	"log"
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
	for i, v := range content {
		replaced := replace(v)
		x, err := decode(replaced)
		if err != nil {
			panic(err)
		}
		log.Printf("%v: %v = %v = %v", i, v, replaced, x)
		sum += x
	}

	fmt.Println(sum)
}

var numbers map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func replace(s string) string {
	ok := false
	for !ok {
		s, ok = replaceOnce(s)
	}

	return s
}

func replaceOnce(s string) (string, bool) {
	for k, v := range numbers {
		index := strings.Index(s, k)
		if index >= 0 {
			return s[:index] + v + s[index+len(v):], strings.ContainsAny(s[index+len(v):], "one two three four five six seven eight nine")
		}
	}

	return s, true
}

func decode(s string) (int, error) {
	first, last := -1, -1

	log.Printf("decode: %v", s)

	for _, v := range s {
		if v >= 48 && v <= 57 { // ASCII range for digits 0-9
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

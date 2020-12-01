package util

import (
	"bufio"
	"os"
	"strconv"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLine() string {
	file, err := os.Open("input.txt")
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	return line
}

func ReadLines() []string {
	lines := make([]string, 0)
	f, err := os.Open("input.txt")
	Check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func ReadNumbers() []int {
	numbers := make([]int, 0)
	f, err := os.Open("input.txt")
	Check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		Check(err)
		numbers = append(numbers, number)
	}
	return numbers
}

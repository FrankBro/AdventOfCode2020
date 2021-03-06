package util

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	errOpen = "failed to open file: %w"
	errRead = "failed to read file: %w"
	errAtoi = "failed to convert to int: %w"
)

func ReadLine() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", fmt.Errorf(errOpen, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	return line, nil
}

func ReadLines() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf(errOpen, err)
	}
	reader := bufio.NewReader(f)
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf(errRead, err)
	}
	split := strings.Split(string(all), "\n")
	return split, nil
}

func ReadLinesNoEmpty() ([]string, error) {
	lines := make([]string, 0)
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf(errOpen, err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func ReadNumbers() ([]int, error) {
	numbers := make([]int, 0)
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf(errOpen, err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf(errAtoi, err)
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

var ErrMissingDelimiter = errors.New("missing delimiter")

func Parse(input string, delims ...string) ([]string, error) {
	output := make([]string, len(delims)+1)
	for i, delim := range delims {
		index := strings.Index(input, delim)
		if index == -1 {
			return nil, ErrMissingDelimiter
		}
		output[i] = input[:index]
		input = input[index+len(delim):]
	}
	output[len(delims)] = input
	return output, nil
}

func Exists(value string, values ...string) bool {
	for i := range values {
		if value == values[i] {
			return true
		}
	}
	return false
}

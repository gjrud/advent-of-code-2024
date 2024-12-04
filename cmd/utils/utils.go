package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadInputToStringArray(path string, sep string) ([]string, error) {
	output, err := readInput(path, sep)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func ReadInputToIntArray(path string, sep string) ([]int, error) {
	lines, err := readInput(path, sep)
	if err != nil {
		return nil, err
	}

	output := make([]int, 0)
	for _, l := range lines {
		v, err := ParseNumber(l)
		if err != nil {
			return nil, err
		}
		output = append(output, v)
	}
	return output, nil
}

func readInput(path string, sep string) ([]string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(f), sep)
	return lines, nil
}

func ParseNumber(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/gjrud/advent-of-code-2024/cmd/utils"
)

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	list1, list2, err := readLists(input)
	if err != nil {
		log.Fatal(err)
	}

	resultP1 := solvePart1(list1, list2)
	fmt.Println(resultP1)

	locationIds := mapLocationIds(list1, list2)
	resultP2 := solvePart2(locationIds)
	fmt.Println(resultP2)
}

func solvePart1(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := 0; i < len(list1); i++ {
		sum += absDiffInt(list1[i], list2[i])
	}
	return sum
}

func readLists(lines []string) ([]int, []int, error) {
	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))
	for i, l := range lines {
		fields := strings.Fields(l)
		if len(fields) != 2 {
			return nil, nil, errors.New("invalid split")
		}
		n, err := utils.ParseNumber(fields[0])
		if err != nil {
			return nil, nil, err
		}
		list1[i] = n

		n, err = utils.ParseNumber(fields[1])
		if err != nil {
			return nil, nil, err
		}
		list2[i] = n
	}
	return list1, list2, nil
}

func absDiffInt(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func solvePart2(locationIds map[int]int) int {
	sum := 0
	for k, v := range locationIds {
		sum += k * v
	}
	return sum
}

func mapLocationIds(list1, list2 []int) map[int]int {
	locationIds := make(map[int]int)
	for _, x := range list1 {
		for _, y := range list2 {
			if x == y {
				locationIds[x]++
			}
		}
	}
	return locationIds
}

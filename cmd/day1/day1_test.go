package main

import (
	"slices"
	"testing"
)

var (
	input = []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	knownList1       = []int{3, 4, 2, 1, 3, 3}
	knownList2       = []int{4, 3, 5, 3, 9, 3}
	knownPart1Result = 11
	knownPart2Result = 31
)

func TestInputParsing(t *testing.T) {
	list1, list2, err := readLists(input)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(list1, knownList1) {
		t.Fatalf("MISMATCH\n%v\n%v", list1, knownList1)
	}
	if !slices.Equal(list2, knownList2) {
		t.Fatalf("MISMATCH\n%v\n%v", list2, knownList2)
	}
}

func TestSolvePart1(t *testing.T) {
	result := solvePart1(knownList1, knownList2)
	if result != knownPart1Result {
		t.Fatalf("MISMATCH\n%v\n%v", result, knownPart1Result)
	}
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(knownList1, knownList2)
	if result != knownPart2Result {
		t.Fatalf("MISMATCH\n%v\n%v", result, knownPart2Result)
	}
}

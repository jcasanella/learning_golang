package main

import (
	"testing"

)

func CheckValues(actual []int, expected []int, t *testing.T) {
	if len(actual) != len(expected) {
		t.Errorf("Actual len %d is different than len %d expected\n", len(actual), len(expected))
	}

	for index, value := range actual {
		if value != expected[index] {
			t.Errorf("Value %d is different than %d expected\n", value, expected[index])
		}
	}
}

func TestCase1A(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase1B(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase2A(t *testing.T) {
	numbers := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase2B(t *testing.T) {
	numbers := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase3A(t *testing.T) {
	numbers := []int{0, 8, 7, 3, 3, 4, 2}
	target := 11

	expected := []int{1, 3}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase3B(t *testing.T) {
	numbers := []int{0, 8, 7, 3, 3, 4, 2}
	target := 11

	expected := []int{1, 3}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase4A(t *testing.T) {
	numbers := []int{0, 1}
	target := 1

	expected := []int{0, 1}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase4B(t *testing.T) {
	numbers := []int{0, 1}
	target := 1

	expected := []int{0, 1}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase5A(t *testing.T) {
	numbers := []int{0, 3}
	target := 5

	expected := []int{}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase5B(t *testing.T) {
	numbers := []int{0, 3}
	target := 5

	expected := []int{}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

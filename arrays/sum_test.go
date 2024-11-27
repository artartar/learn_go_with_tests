package main

import (
	// "reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum([]int(numbers))
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Introduced in or after Go 1.21 and is part of the slices standard library. Does a simple and shallow comparison of slices.
	// First checks if they are the same size. Different sizes equals failure
	// If sizes are the same, compares each value in the arrays 1 by 1 until it encounters a failure. No failure equals success(equal)
	if !slices.Equal[[]int](got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	// Not type safe, implemented as directed in the course.
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v", got, want)
	// }
}

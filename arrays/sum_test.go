package main

import (
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

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal[[]int](got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		// checkSums := func(t testing.TB, got, want []int) {
		// 	t.Helper()
		// 	if !reflect.DeepEqual(got, want) {
		// 		t.Errorf("got %v want %v", got, want)
		// 	}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 2}, []int{0, 9, 9})
		want := []int{4, 18}

		checkSums(t, got, want)

		// if !refelct.DeepEqual(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)

		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }
	})
}

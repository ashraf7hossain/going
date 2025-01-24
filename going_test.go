package going

import (
	"testing"
)

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{1, 4, 9, 16, 25}
	squared := Map(numbers, func(x int) int {
		return x * x
	})

	for i, v := range squared {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], v)
		}
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 6}
	expected := []int{1, 3}
	odds := Filter(numbers, func(x int) bool {
		return x%2 == 1
	})

	for i, v := range odds {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], v)
		}
	}
}

func TestReduce(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	expected := 15
	sum := Reduce(ints, func(acc int, item int) int {
		return acc + item
	}, 0)

	if sum != expected {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}


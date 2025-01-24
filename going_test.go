package going

import (
	"sync"
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

func TestCMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{1, 4, 9, 16, 25}
	squared := CMap(numbers, func(x int) int {
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

// TestCMapEmptySlice tests CMap with an empty input slice
func TestCMapEmptySlice(t *testing.T) {
	list := []int{}
	f := func(x int) int {
		return x * 2
	}

	result := CMap(list, f)
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}
}

// TestCMapSingleElement tests CMap with a single-element slice
func TestCMapSingleElement(t *testing.T) {
	list := []int{42}
	f := func(x int) int {
		return x * 2
	}

	result := CMap(list, f)
	if len(result) != 1 || result[0] != 84 {
		t.Errorf("Expected [84], got %v", result)
	}
}

// TestCMapMultipleElements tests CMap with a slice of multiple elements
func TestCMapMultipleElements(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	f := func(x int) int {
		return x * 2
	}

	expected := []int{2, 4, 6, 8, 10}
	result := CMap(list, f)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
			break
		}
	}
}

// TestCMapConcurrentAccess tests CMap for race conditions
func TestCMapConcurrentAccess(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	f := func(x int) int {
		return x * 2
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = CMap(list, f)
		}()
	}
	wg.Wait()
}

// TestCMapLargeSlice tests CMap with a large slice
func TestCMapLargeSlice(t *testing.T) {
	list := make([]int, 100000)
	for i := range list {
		list[i] = i
	}

	f := func(x int) int {
		return x * 2
	}

	result := CMap(list, f)
	for i := range list {
		if result[i] != list[i]*2 {
			t.Errorf("Expected %v, got %v", list[i]*2, result[i])
			break
		}
	}
}

// // TestCMapNilFunction tests CMap with a nil function
// func TestCMapNilFunction(t *testing.T) {
// 	list := []int{1, 2, 3}
// 	defer func() {
// 		if r := recover(); r == nil {
// 			t.Errorf("Expected panic with nil function, but no panic occurred")
// 		}
// 	}()

// 	_ = CMap(list, nil) // This should panic
// }

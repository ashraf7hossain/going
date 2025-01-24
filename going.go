package going 

import (
	"fmt"
)

// Map takes a list of elements and applies the provided function `f` on each element,
// returning a new list of transformed elements. The function `f` is applied to each item
// of type T and produces a result of type U.
//
// Example:
//    numbers := []int{1, 2, 3, 4, 5}
//    squared := Map(numbers, func(x int) int {
//        return x * x
//    })
//    fmt.Println(squared)  // prints: [1 4 9 16 25]
func Map[T any, U any](list []T, f func(T) U) []U {
    result := make([]U, len(list))
    for i, v := range list {
        result[i] = f(v)
    }
    return result
}

// Filter takes a list of elements and applies the provided predicate function `f` on each element.
// It returns a new list containing only the elements for which the function `f` returns true.
//
// Example:
//    numbers := []int{1, 2, 3, 4, 5}
//    odds := Filter(numbers, func(x int) bool {
//        return x % 2 == 1
//    })
//    fmt.Println(odds)  // prints: [1 3 5]
func Filter[T any](list []T, f func(T) bool) []T {
    result := make([]T, 0, len(list))
    for _, v := range list {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

// Reduce takes a list of elements and applies the provided reducer function on each element,
// combining them into a single accumulated result. The reducer function takes an accumulator of
// type R and an item of type T, and returns a new accumulator of type R. It starts with an initial value.
//
// Example:
//    ints := []int{1, 2, 3, 4, 5}
//    sum := Reduce(ints, func(acc int, item int) int {
//        return acc + item
//    }, 0)
//    fmt.Println("Sum:", sum)  // prints: Sum: 15
func Reduce[T any, R any](list []T, reducer func(acc R, item T) R, initial R) R {
    acc := initial
    for _, item := range list {
        acc = reducer(acc, item)
    }
    return acc
}


// ForEach iterates over each element in the given list and applies the provided function `f` on each element.
// It is similar to the "forEach" function in JavaScript and other languages.
//
// Example:
//    numbers := []int{1, 2, 3, 4, 5}
//    ForEach(numbers, func(x int) {
//        fmt.Println(x * x)  // prints the square of each number
//    })
func ForEach[T any](list []T, f func(T)) {
    for _, v := range list {
        f(v)
    }
}

// Contains checks if the target element exists in the given list.
// Returns true if the element is found, otherwise false.
//
// Example:
//    numbers := []int{1, 2, 3, 4, 5}
//    found := Contains(numbers, 3)
//    fmt.Println(found)  // prints: true
//
//    found = Contains(numbers, 6)
//    fmt.Println(found)  // prints: false
func Contains[T comparable](list []T, target T) bool {
    for _, v := range list {
        if v == target {
            return true
        }
    }
    return false
}

// IndexOf searches for the target element in the list and returns its index if found.
// If the element is not found, it returns -1 and an error indicating that the element was not found.
//
// Example:
//    numbers := []int{1, 2, 3, 4, 5}
//    index, err := IndexOf(numbers, 3)
//    fmt.Println(index, err)  // prints: 2 <nil>
//
//    index, err = IndexOf(numbers, 6)
//    fmt.Println(index, err)  // prints: -1 element not found
func IndexOf[T comparable](list []T, target T) (int, error) {
    for i, v := range list {
        if v == target {
            return i, nil
        }
    }
    return -1, fmt.Errorf("element not found")
}

package govalidator

// Iterator is the function that accepts element of slice/array and its index
type Iterator[T any] func(T, int)

// ResultIterator is the function that accepts element of slice/array and its index and returns any result
type ResultIterator[T any] func(T, int) T

// ConditionIterator is the function that accepts element of slice/array and its index and returns boolean
type ConditionIterator[T any] func(T, int) bool

// ReduceIterator is the function that accepts two element of slice/array and returns result of merging those values
type ReduceIterator[T any] func(T, T) T

// Some validates that any item of array corresponds to ConditionIterator. Returns boolean.
func Some[T any](array []T, iterator ConditionIterator[T]) bool {
	res := false
	for index, data := range array {
		res = res || iterator(data, index)
	}
	return res
}

// Every validates that every item of array corresponds to ConditionIterator. Returns boolean.
func Every[T any](array []T, iterator ConditionIterator[T]) bool {
	res := true
	for index, data := range array {
		res = res && iterator(data, index)
	}
	return res
}

// Reduce boils down a list of values into a single value by ReduceIterator
func Reduce[T any](array []T, iterator ReduceIterator[T], initialValue T) T {
	for _, data := range array {
		initialValue = iterator(initialValue, data)
	}
	return initialValue
}

// Each iterates over the slice and apply Iterator to every item
func Each[T any](array []T, iterator Iterator[T]) {
	for index, data := range array {
		iterator(data, index)
	}
}

// Map iterates over the slice and apply ResultIterator to every item. Returns new slice as a result.
func Map[T any](array []T, iterator ResultIterator[T]) []T {
	var result = make([]T, len(array))
	for index, data := range array {
		result[index] = iterator(data, index)
	}
	return result
}

// Find iterates over the slice and apply ConditionIterator to every item. Returns first item that meet ConditionIterator or nil otherwise.
func Find[T any](array []T, iterator ConditionIterator[T]) (t T) {
	for index, data := range array {
		if iterator(data, index) {
			return data
		}
	}
	return
}

// Filter iterates over the slice and apply ConditionIterator to every item. Returns new slice.
func Filter[T any](array []T, iterator ConditionIterator[T]) []T {
	var result = make([]T, 0)
	for index, data := range array {
		if iterator(data, index) {
			result = append(result, data)
		}
	}
	return result
}

// Count iterates over the slice and apply ConditionIterator to every item. Returns count of items that meets ConditionIterator.
func Count[T any](array []T, iterator ConditionIterator[T]) int {
	count := 0
	for index, data := range array {
		if iterator(data, index) {
			count = count + 1
		}
	}
	return count
}

package govalidator

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Abs returns absolute value of number
func Abs[V constraints.Integer | constraints.Float](value V) V {
	return V(math.Abs(float64(value)))
}

// Sign returns signum of number: 1 in case of value > 0, -1 in case of value < 0, 0 otherwise
func Sign[V constraints.Integer | constraints.Float](value V) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

// IsNegative returns true if value < 0
func IsNegative[V constraints.Integer | constraints.Float](value V) bool {
	return value < 0
}

// IsPositive returns true if value > 0
func IsPositive[V constraints.Integer | constraints.Float](value V) bool {
	return value > 0
}

// IsNonNegative returns true if value >= 0
func IsNonNegative[V constraints.Integer | constraints.Float](value V) bool {
	return value >= 0
}

// IsNonPositive returns true if value <= 0
func IsNonPositive[V constraints.Integer | constraints.Float](value V) bool {
	return value <= 0
}

// InRange returns true if value lies between left and right border, generic type to handle int, float32, float64 and string.
// All types must the same type.
// False if value doesn't lie in range or if it incompatible or not comparable
func InRange[V constraints.Ordered](value V, left V, right V) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// IsWhole returns true if value is whole number
func IsWhole[V constraints.Integer | constraints.Float](value V) bool {
	return math.Remainder(float64(value), 1) == 0
}

// IsNatural returns true if value is natural number (positive and whole)
func IsNatural[V constraints.Integer | constraints.Float](value V) bool {
	return IsWhole(value) && IsPositive(value)
}

package govalidator

func ExampleAbs() {
	_ = Abs(-123.3e1) // 123.3e1
	_ = Abs(+0)       // 0
	_ = Abs(321)      // 321
}

func ExampleSign() {
	_ = Sign(-123) // -1
	_ = Sign(123)  // 1
	_ = Sign(0)    // 0
}

func ExampleIsNegative() {
	_ = IsNegative(-123) // true
	_ = IsNegative(0)    // false
	_ = IsNegative(123)  // false
}

func ExampleIsPositive() {
	_ = IsPositive(-123) // false
	_ = IsPositive(0)    // false
	_ = IsPositive(123)  // true
}

func ExampleIsNonNegative() {
	_ = IsNonNegative(-123) // false
	_ = IsNonNegative(0)    // true
	_ = IsNonNegative(123)  // true
}

func ExampleIsNonPositive() {
	_ = IsNonPositive(-123) // true
	_ = IsNonPositive(0)    // true
	_ = IsNonPositive(123)  // false
}

func ExampleInRange() {
	_ = InRange(10, 11, 20)             // false
	_ = InRange(10.02, -10.124, 10.234) // true
	_ = InRange("abc", "a", "cba")      // true
}

func ExampleIsWhole() {
	_ = IsWhole(1.123) // false
	_ = IsWhole(1.0)   // true
	_ = IsWhole(10)    // true
}

func ExampleIsNatural() {
	_ = IsNatural(1.123) // false
	_ = IsNatural(1.0)   // true
	_ = IsNatural(-10)   // false
}

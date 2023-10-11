package govalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

// ToString convert the input to a string.
func ToString[T any](obj T) string {
	res := fmt.Sprintf("%v", obj)
	return res
}

// ToJSON convert the input to a valid JSON string
func ToJSON[T any](obj T) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		res = []byte("")
	}
	return string(res), err
}

// ToNumber convert the input string to a number, or 0.0 if the input is not a number.
func ToNumber[T any, U constraints.Float | constraints.Integer](value T) (res U, err error) {
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = U(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = U(val.Uint())
	case reflect.Float32, reflect.Float64:
		res = U(val.Float())
	case reflect.String:
		if IsInt[string](val.String()) {
			resInt, err1 := strconv.ParseInt(val.String(), 0, 64)
			if err1 != nil {
				res = 0
			} else {
				res = U(resInt)
			}
		} else if resFloat, err1 := strconv.ParseFloat(val.String(), 64); err1 == nil {
			res = U(resFloat)
		} else {
			err = fmt.Errorf("ToInt: invalid numeric format %v", value)
			res = 0
		}

	default:
		err = fmt.Errorf("ToInt: unknown interface type %T", value)
		res = 0
	}

	return
}

// ToInt convert the input string or any int type to an integer type 64, or 0 if the input is not an integer.
func ToInt[T any](value T) (res int64, err error) {
	return ToNumber[any, int64](value)
}

// ToBoolean convert the input string to a boolean.
func ToBoolean[T ~string](str T) (bool, error) {
	return strconv.ParseBool(string(str))
}

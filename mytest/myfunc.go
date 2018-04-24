package mytest

import (
	"errors"
	"math"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return math.Trunc(a / b), nil
}

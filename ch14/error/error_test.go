package error_test

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHoundredError = errors.New("n should be not larger than 100")

func GetFibonaci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHoundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonaci(110); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number")
		}
		if err == LargerThanHoundredError {
			t.Error("Need a smaller number")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

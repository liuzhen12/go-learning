package client

import (
	"testing"

	"com.example/go_learning/src/ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonaci(5))
	t.Log(series.Square(5))
}

package ch1_test

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	fmt.Println("123")
	t.Log("123")
}

func SayHello() {
	fmt.Print("Hello")
}

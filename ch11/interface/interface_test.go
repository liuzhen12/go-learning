package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() string {
	return fmt.Sprintln("Go: Hello world")
}

func TestClient(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

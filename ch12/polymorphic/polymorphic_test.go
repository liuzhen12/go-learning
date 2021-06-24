package polymorphic_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() Code {
	return "Go: Hello world"
}

type JavaProgrammer struct {
}

func (jp *JavaProgrammer) WriteHelloWorld() Code {
	return "Java: Hello world"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}

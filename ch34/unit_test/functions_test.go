package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	input := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(input); i++ {
		ret := Square(input[i])
		// if ret != expected[i] {
		// 	t.Errorf("input is %d, the expected is %d, the actual is %d",
		// 		input[i], expected[i], ret)
		// }
		assert.Equal(t, expected[i], ret)
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}

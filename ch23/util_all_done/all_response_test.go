package utilalldone_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numberOfRunner := 10
	ch := make(chan string, numberOfRunner)
	for i := 0; i < numberOfRunner; i++ {
		go func(id int) {
			ch <- runTask(id)
		}(i)
	}
	var finalRet string
	for j := 0; j < numberOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

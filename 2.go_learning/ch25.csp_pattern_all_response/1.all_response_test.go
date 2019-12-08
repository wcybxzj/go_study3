package ch25_csp_pattern_all_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask (id int)  string{
	time.Sleep(10 *time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numberOfRunner := 10
	ch := make(chan string, numberOfRunner)
	for i := 0; i < numberOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalResult := ""
	for j:=0; j<numberOfRunner; j++ {
		finalResult += <- ch +"\n"
	}
	return finalResult
}


/*
输出:
用管道实现Waitgroup的功能
1.all_response_test.go:33: Before: 2
1.all_response_test.go:34: The result is from 1
The result is from 7
The result is from 2
The result is from 6
The result is from 9
The result is from 3
The result is from 4
The result is from 5
The result is from 8
The result is from 0
1.all_response_test.go:35: After: 2
*/
func TestAllResponse(t *testing.T)  {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Log("After:", runtime.NumGoroutine())
}
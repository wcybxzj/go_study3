package ch16_groutine

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1)
}
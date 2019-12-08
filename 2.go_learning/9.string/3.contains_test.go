package __string

import (
	"fmt"
	"strings"
	"testing"
)

//
func TestContains(t *testing.T)  {
	re := strings.Contains("abc", "a")
	fmt.Println(re)

	re = strings.Contains("a世界bc", "世")
	fmt.Println(re)
}



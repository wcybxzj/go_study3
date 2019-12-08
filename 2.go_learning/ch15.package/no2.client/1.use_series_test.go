package no2_client

import (
	no1_series "10.go_learing/ch15.package/no1.series"
	"fmt"
	"testing"
)

func Test(t *testing.T)  {
	list ,_:= no1_series.GetFib(10)
	fmt.Println(list)
}
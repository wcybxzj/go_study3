package no1_series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func Sqaure(n int) int {
	return n*n
}

var LessThanTwoError = errors.New("n less than 2")
var BiggerThanHundredError = errors.New("n bigger than 100")

func GetFib(n int) ([]int, error) {
	if n < 2 {
		return  nil,LessThanTwoError
	}

	if n >100 {
		return nil, BiggerThanHundredError
	}

	fibList := []int{1, 1}

	for i:=2; i<n; i++ {
		fibList = append(fibList,(fibList[i-1]+fibList[i-2]))
	}
	return fibList, nil
}

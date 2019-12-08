package _4_error_panic_recover

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)


var LessThanTwoError = errors.New("n less than 2")
var BiggerThanHundredError = errors.New("n bigger than 100")

//1,1,2,3,5,8,13....
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


func UseFib(num string)  {
	var(
		n int
		err error
		list []int
	)

	if n, err = strconv.Atoi(num); err != nil{
		fmt.Println("Error", err)
		return
	}

	if list, err = GetFib(n); err != nil{
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}

func TestOne(t *testing.T) {
	arr ,err := GetFib(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(arr)
}

func TestTwo(t *testing.T)  {
	arr ,err := GetFib(5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(arr)
}
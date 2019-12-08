package gc_friendly

import (
	"testing"
)

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	//	fmt.Println(&arr[2])
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	//b := *arr
	//	fmt.Println(&arr[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	//fmt.Println(&arr[2])
	withValue(arr)
	withReference(&arr)
}
//传值每次操作11622500ns
//BenchmarkPassingArrayWithValue-12    	     100	  11622500 ns/op
func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

//传地址每次操作0.51ns
//BenchmarkPassingArrayWithRef-12    	2000000000	         0.51 ns/op
func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}

package __arr_slice

import "testing"

//slice初始化测试
func TestSliceInit(t *testing.T) {

	//init method1:和数组不同这里只是声明并没有初始化里面的内容
	var s0 []int
	t.Log(len(s0), cap(s0)) // 0,0
	//t.Log(s0[0])//error index out of range
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) //1 ,1

	//init method2:
	s1 := []int{11, 22, 33, 44}
	t.Log(len(s1), cap(s1)) //4 ,4

	//init method3:只有用make生成的slice 内容才会被初始化成0
	s2 := make([]int,3, 5) //len:3 cap:5
	t.Log(len(s2), cap(s2)) //3, 5
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 333)
	t.Log(len(s2), cap(s2)) //3, 5
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

//两个slice指向同一个数组时候的测试
//两个分片共享同一个底层arr时候,修改一个slice的内容,另外一个分片也可能收到影响
func TestSliceShareMemory(t *testing.T) {
	year := [...]string{"Jan","Feb", "Mar",
					"April", "May", "June",
					"July", "Aug", "Sep",
					"Oct", "Nove", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0]= "11111"

	t.Log(Q2)
	t.Log(summer)
}

//slice的比较测试
func TestSliceComparing(t *testing.T) {
	a := []int{1,2,3,4}
	b := []int{1,2,3,4}

	//error slice只能和nil比较
	//if a == b {
	//	t.Log("equal")
	//}

	t.Log(a, b)

}
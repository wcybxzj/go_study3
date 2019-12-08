package __for_switch

import "testing"

//case有多个值
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++  {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("not 0-3")
		}
	}
}

//switch后边不加东西,实现类似 if else的效果，不实用
func TestSwitchCaseConditon(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unknow")
		}
	}
}

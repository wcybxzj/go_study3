package __string

import "testing"

func PrintLen(str string, t *testing.T)  {
	t.Log(len(str)) //len显示的是Byte数
}

func TestString(t *testing.T)  {
	var s string
	
	//exmple1:空String
	t.Log(s)//空
	PrintLen(s, t) // 0 Bytes

	//exmple2:英文String
	s="hello"
	t.Log(s) //hello
	PrintLen(s, t) // 5 Bytes

	//example3:1个二进制形式的汉字
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	t.Log(s) //严
	PrintLen(s, t) // 3 Bytes

	//example4:随便了一串无法被编码成文字的二进制数据
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s) //乱码
	PrintLen(s, t) // 4 Bytes

	//example5:string是不可变的Byte切片
	//这个和C是一样的
	//s[1]='3' //error
	//s[1]='\xAA' //error

	//example6:
	//Unicode 是一种字符集(code point) ,UTF8 是 unicode 的存储实现 (转换为字节序列列的规则)
	//  字符 			“中”
	//	Unicode			0x4E2D
	//	UTF-8			0xE4B8AD
	//	string/[]byte	[0xE4,0xB8,0xAD]

	s = "中"
	c := []rune(s)

	t.Log(len(s)) //3 bytes
	t.Log(len(c)) //1 bytes

	t.Logf("中 unicode %x", c[0]) //严 unicode 4e2d
	t.Logf("中 unicode %x", c) //严 unicode [4e2d]

	t.Logf("中 UTF8 %x",s ) //严 UTF8 e4b8ad
}

//输出:
//0, y, 79
//1, e, 65
//2, s, 73
//3, 中, 4e2d
//6, 国, 56fd


func TestStringToRune(t *testing.T)  {
	s := "yes中国"
	//range时候自动会把string中的byte转换成本rune
	for index, value := range s{
		//这里[2]意思是说这两个format都用后边第二个数据也就是value
		t.Logf("%d, %[2]c, %[2]x",index, value )
	}
}

package ch38_flexible_reflect_test

import (
	"errors"
	"reflect"
	"testing"
)

//test1:map slice深度比较
func TestDeepEqual(t *testing.T) {
	a := map[int]map[int]string{1:{100:"ybx"}, 2:{200:"wc"}}
	b := map[int]map[int]string{1:{100:"ybx"}, 2:{200:"wc200"}}
	//t.Log(a == b)

	t.Log(reflect.DeepEqual(a, b))

	s1 :=[][]int{{1, 2, 3}, {4,5,6}}
	s2 :=[][]int{{1, 2, 3}, {4,5,6}}
	s3 :=[][]int{{1, 2, 3}, {4,5,16}}
	t.Log("s1==s2?", reflect.DeepEqual(s1,s2))
	t.Log("s2==s3?", reflect.DeepEqual(s2,s3))
}

type Employee struct {
	Employee string
	Name string `format:"normal"`
	Age int
}

type Customer struct {
	CookieID string
	Name string
	Age int
}

func fillBySettings(st interface{}, settings map[string]interface{} ) error {
	var(
		ok bool
		field reflect.StructField
	)

	//1.st muse be ptr
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	//Elem返回指针指向的值
	//2.st must be struct
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	//3.循环map
	for k, v := range settings {
		vst := reflect.ValueOf(st)

		//用map中的key,从struct中取出field
		if field, ok = vst.Elem().Type().FieldByName(k); !ok {
			continue
		}

		//struct 和 map 中的字段 类型进行比较
		if field.Type == reflect.TypeOf(v) {
			vst.Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

//用1个map去填充2个类型完全不同的struct
/*
输出:
{ Dabing 30}
{ Dabing 30}
*/
func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name":"Dabing", "Age":30}

	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil{
		t.Fatal(err)
	}
	t.Log(e)

	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil{
		t.Fatal(err)
	}
	t.Log(*c)
}
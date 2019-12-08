package ch37_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

//test1:获取变量的值或者类型
/*
输出:
int64 10
int64
*/
func TestTypeAndValue(t *testing.T)  {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32,reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("unknown", t)
	}
}

//test2:
func TestBasicType(t *testing.T)  {
	var f float64 = 12
	//CheckType(f)
	CheckType(&f)
}

type Emploee struct {
	EmploeeID string
	Name      string `format:"normal"`
	Age 		int
}

func (e *Emploee) UpdateAge(newVal int) {
	e.Age = newVal
}

//test3
func TestInvokeByName(t*testing.T) {
	e := &Emploee{"1", "ybx", 100}
	t.Logf("Name: value:%[1]v, type:%[1]T", reflect.ValueOf(*e).FieldByName("Name"))

	//print tag
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' Field.")
	}else{
		t.Log("tag:format is",nameField.Tag.Get("format") )
	}

	//call method
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})

	t.Log("Update Age:", e)
}


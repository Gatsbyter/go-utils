package reflect

import (
	"reflect"
	"testing"
)

type X int
type Y X
type Z struct {
	X
}

func (X) Xv()  {}
func (*X) Xp() {}

func (Z) Zv()  {}
func (*Z) Zp() {}

type User struct {
	Name string
	Age  int
}

type Manager struct {
	User
	Title string
}

// 测试type.name 和 type.kind
func TestTypeOf(t *testing.T) {
	var a Y = 100
	printType(a, t)
}

func printType(o interface{}, t *testing.T) {
	typ := reflect.TypeOf(o)
	t.Logf("name: %s | kind: %s", typ.Name(), typ.Kind())
}

// 使用reflect构造类型
func TestMakeType(t *testing.T) {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	s := reflect.SliceOf(m)
	t.Logf("%v | %v | %v", a, m, s)
}

// 打印类型
func TestPtrType(t *testing.T) {
	x := 100

	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	t.Logf("%v, %v, %t", tx, tp, tx == tp)
	t.Logf("%v, %v", tx.Kind(), tp.Kind())
	t.Logf("%v, %v", tp.Elem(), tx == tp.Elem())
}

// 把结构体里的字段和类型都整出来
func TestStruct(t *testing.T) {
	var m Manager
	typ := reflect.TypeOf(&m)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// 打印结构体有多少个字段
	t.Logf("num field: %v", typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		t.Logf("%v, %v, %v", f.Name, f.Type, f.Offset)

		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				t.Logf("  %v, %v", af.Name, af.Type)
			}
		}
	}
}

// 找结构体里的特定字段
func TestFindField(t *testing.T) {
	var m Manager

	typ := reflect.TypeOf(m)

	name, ok := typ.FieldByName("Name")
	if !ok {
		t.Errorf("cannot find name")
		return
	}

	t.Logf("%v, %v", name.Name, name.Type)

	// 这玩意就是找多级查询，比如0, 1就是根结构的第一个结构里的第二个字段
	// {0, 1}就是age, {1}就是title, {0, 0}是name
	age := typ.FieldByIndex([]int{0, 1})
	t.Logf("%v, %v", age.Name, age.Type)
}

// 找对象的方法
func TestPrintMethod(t *testing.T) {
	var a Z

	typ := reflect.TypeOf(&a)

	// 对值和指针分别查找
	// 因为方法分为传值和传指针
	s := []reflect.Type{typ, typ.Elem()}

	for _, typ := range s {
		t.Logf("%v: %v", typ, typ.NumMethod())
		for i := 0; i < typ.NumMethod(); i++ {
			m := typ.Method(i)
			t.Logf("    %v, %v, %v, %v", m.Name, m.Type, m.Func, m.Index)
		}
	}
}

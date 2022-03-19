package generics

import (
	"fmt"
	"reflect"
)

type VectorBool Vector[bool]

type MyVectorBool struct {
	Vector[bool]
}

func ExampleVector() {
	// v0 := VectorBool{}
	// v0.Push(true) // Error: v0.Push undefined (type VectorBool has no field or method Push) https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source=gopls#MissingFieldOrMethod

	// 使用 := 实例化 Vector 为 Vector[int]
	// 此时 v1 的类型就是 Vector[int] 等价于 type Vector[int] []int ，把 Vector[int] 看成一个标识符
	v1 := Vector[int]{}
	v1.Push(1)
	v1.Push(2)
	v1.Push(3)
	fmt.Println(v1)
	_ = []int(v1) // 底层类型相同，可以这样转换
	fmt.Printf("v1 reflect: type = %s, kind = %s\n", reflect.TypeOf(v1), reflect.ValueOf(v1).Kind())

	// 使用 var 实例化 Vector 为 Vector[string]
	var v2 Vector[string]
	v2.Push("a")
	v2.Push("b")
	v2.Push("c")
	fmt.Println(v2)
	fmt.Printf("v2 reflect: type = %s, kind = %s\n", reflect.TypeOf(v2), reflect.ValueOf(v2).Kind())

	// 嵌入结构体
	v3 := MyVectorBool{}
	v3.Push(true)
	v3.Push(false)
	v3.Push(true)
	fmt.Println(v3)
	fmt.Printf("v3 reflect: type = %s, kind = %s\n", reflect.TypeOf(v3), reflect.ValueOf(v3).Kind())
	// output:
	// [1 2 3]
	// v1 reflect: type = generics.Vector[int], kind = slice
	// [a b c]
	// v2 reflect: type = generics.Vector[string], kind = slice
	// {[true false true]}
	// v3 reflect: type = generics.MyVectorBool, kind = struct
}

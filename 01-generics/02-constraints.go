package generics

import (
	"fmt"
	"reflect"
)

// Print 打印切片的元素。
// Print 有一个类型参数 T，且有一个（非类型）普通参数 s，
// s 是一个为类型为类型参数 T 的切片。
func PrintInterface[T interface{}](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Stringer 是一个类型约束。约束为 Stringer 的类型参数意味着：该类型必须有一个 String 方法。
// 在泛型函数中，使用类型为该类型参数的参数，允许在该变量上调用 String 方法。
// （这定义了与标准库的 fmt.Stringer 类型相同的接口，实际代码可能会简单地使用 fmt.Stringer。）
type Stringer interface {
	String() string
}

// Stringify 在 s 的每个元素上调用 String 方法，
// 并返回结果
func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// MyInt 定义一个类型约束，表示被约束参数必须是 int 类型
// 这约束就等价于 int，等于限定死类型参数必须是 int
type MyInt interface {
	int
}

// IntAdd2 两个 int 相加
// 这个函数声明完全等价于 func IntAdd2(x, y int) int
func IntAdd2[T MyInt](x, y T) T {
	return x + y
}

// MyIntString 定义一个需同时是 string 和 int 的类型约束
// 显然没有这种类型
type MyIntString interface {
	int
	string
}

// 这个函数不可能被调用，因为不可能存在既是 int 又是 string 的类型
func PrintMyIntString[T MyIntString](x T) {
	fmt.Println(x)
}

// MyIntWithAddOne 定义一个类型必须是 int 的类型约束，且包含一个 AddOne 方法
// 显然没有这种类型
type MyIntWithAddOne interface {
	int
	AddOne() int
}

// MyUint 定义一个类型约束，表示被约束参数的底层类型必须是 uint
type MyUint interface {
	~uint
}

// IntAdd2 底层类型为 uint 的两个变量相加
func UintAdd2[T MyUint](x, y T) T {
	return x + y
}

// MyIntWithAddAddOne 定义一个类型约束，表示被约束参数的底层类型必须是 int，且包含一个 Add 方法
type MyIntWithAdd interface {
	~int
	Add(int) int
}

func MyIntWithAddAddOne[T MyIntWithAdd](x T) int {
	return x.Add(1)
}

// MyUnion 定义一个 Union 的类型约束，表示被约束参数的类型满足如下三者之一
// a) bool
// b) 底层类型和 int 相同
// c) MyUint，即 ~uint，底层类型和 uint 相同
type MyUnion interface {
	bool | ~int | MyUint
}

// PrintMyUnionAndType 打印 MyUnion 的类型和值
func PrintMyUnionAndType[T MyUnion](x T) {
	switch v := any(x).(type) {
	case bool:
		fmt.Printf("bool = %t\n", v)
	case int:
		fmt.Printf("int = %d\n", v)
	case uint:
		fmt.Printf("uint = %d\n", v)
	default:
		// 底层类型概念：https://lingchao.xin/post/type-system-overview.html#%E6%A6%82%E5%BF%B5-%E5%BA%95%E5%B1%82%E7%B1%BB%E5%9E%8B
		// 底层类型 issue： https://github.com/golang/go/issues/39574
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int:
			fmt.Printf("~int = %d\n", v)
		case reflect.Uint:
			fmt.Printf("~uint = %d\n", v)
		default:
			fmt.Printf("dead code\n")
		}
	}
}

// 如下将报错

// type MyUnionInvalidWithInterfaceHasMethod1 interface {
// 	// cannot use xxx.Stringer in union (fmt.Stringer contains methods)
// 	// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source%3Dgopls#InvalidUnion
// 	bool | fmt.Stringer
// }

// type MyUnionInvalidWithInterfaceHasMethod2 interface {
// 	// cannot use xxx.MyIntWithAdd in union (xxx.MyIntWithAdd contains methods)
// 	// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source%3Dgopls#InvalidUnion
// 	bool | MyIntWithAdd
// }

// cannot use interface MyInt in conversion (contains specific type constraints or is comparable)
// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source%3Dgopls#MisplacedConstraintIface
// var MyInt1 = MyInt(1)

func Equals[T comparable](a, b T) bool {
	return a == b
}

type (
	A1 struct{}
	A2 struct{}
)

func (A1) String() string { return "A1" }
func (A2) String() string { return "A2" }

// // AString A1 和 A2 都拥有 String 方法，但是在 Go 1.18 中编译仍然报错
// func AString1[T A1 | A2](x T) string {
// 	return x.String() // Error: x.String undefined (type T has no field or method String) https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source=gopls#MissingFieldOrMethod
// }

// 显式的声明方法， x.String() 才不会报错
func AString1[T interface {
	A1 | A2
	String() string
	// fmt.Stringer // 这个写法也行
}](x T,
) string {
	return x.String()
}

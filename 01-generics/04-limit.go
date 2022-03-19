package generics

import "fmt"

type A struct {
	B int
}

func (A) Print() { fmt.Println("a") }

// a. Error: method must have no type parameters
// func (a A) Print[T any](a T)  {
// 	fmt.Println(a)
// }

// b. Error: complex (built-in) is not a type
// func PrintComplex[T complex](a T) {
// 	fmt.Println(a)
// }

// c. Error: a.Print undefined (type T has no field or method Print)
// func PrintA[T A](a T) {
// 	a.Print()
// }
// c. Go 1.18 处理方式为手动声明方法
func PrintA[T interface {
	A
	Print()
}](a T,
) {
	a.Print()
}

// d. Error: a.B undefined (type T has no field or method B)
// func PrintB[T A](a T) {
// 	fmt.Println(a.B)
// }
// d. Go 1.18 处理方式为通过 any 转换
func PrintB[T A](a T) {
	fmt.Println(any(a).(A).B)
}

// e. Error: embedded field type cannot be a (pointer to a) type parameter
// type EmbeddedType[T any] struct {
// 	T
// 	A int
// }

// f. Error: cannot use fmt.Stringer in union (fmt.Stringer contains methods)
// func PrintString[T string | fmt.Stringer](s T) {
// }

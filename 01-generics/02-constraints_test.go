package generics

import (
	"fmt"
	"strings"
)

func ExamplePrintInterface() {
	// 使用一个 []int 参数调用 Print。
	// print有一个类型参数T，我们要传递一个[]int，
	// 所以通过 Print[int] 来向函数 Print 的类型参数 T 传递 int，作为其参数。
	// 此时函数 Print[int] 需要一个 []int 作为参数。
	PrintInterface([]int{1, 2, 3})
	// output:
	// 1
	// 2
	// 3
}

func ExampleStringify() {
	a, b, c := &strings.Builder{}, &strings.Builder{}, &strings.Builder{}
	a.WriteString("a")
	b.WriteString("b")
	c.WriteString("c")
	fmt.Println(Stringify([]*strings.Builder{a, b, c}))
	// output:
	// [a b c]
}

func ExampleIntAdd2() {
	// 底层类型相同是不行的，如下两句会报错
	// type MyIntType int
	// fmt.Println(IntAdd2(MyIntType(1), MyIntType(2)))
	// 别名是可以的
	type MyIntAlias = int
	fmt.Println(IntAdd2(MyIntAlias(1), MyIntAlias(2)))
	fmt.Println(IntAdd2(1, 2))
	// output:
	// 3
	// 3
}

func ExampleUintAdd2() {
	type MyIntType uint
	fmt.Println(UintAdd2(MyIntType(1), MyIntType(2)))
	type MyIntAlias = uint
	fmt.Println(UintAdd2(MyIntAlias(1), MyIntAlias(2)))
	fmt.Println(UintAdd2(uint(1), uint(2)))
	// output:
	// 3
	// 3
	// 3
}

type MyIntAddType int

func (a MyIntAddType) Add(b int) int {
	return int(a) + b
}

func ExampleMyIntWithAddAddOne() {
	fmt.Println(MyIntWithAddAddOne(MyIntAddType(1)))
	// output:
	// 2
}

func ExamplePrintMyUnionAndType() {
	type MyInt int
	type MyUint uint
	PrintMyUnionAndType(true)
	PrintMyUnionAndType(1)
	PrintMyUnionAndType(MyInt(1))
	PrintMyUnionAndType(uint(2))
	PrintMyUnionAndType(MyUint(1))
	// output:
	// bool = true
	// int = 1
	// ~int = 1
	// uint = 2
	// ~uint = 1
}

func ExampleEquals() {
	fmt.Println(Equals(1, 2))
	// fmt.Println(Equals([]int{1}, []int{2})) //  切片不可比较
	fmt.Println(Equals([1]int{1}, [1]int{2})) // 长度相同数组可比较
	// output:
	// false
	// false
}

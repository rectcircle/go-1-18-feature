package generics

func ExamplePrint() {
	// 使用一个 []int 参数调用 Print。
	// print有一个类型参数T，我们要传递一个[]int，
	// 所以通过 Print[int] 来向函数 Print 的类型参数 T 传递 int，作为其参数。
	// 此时函数 Print[int] 需要一个 []int 作为参数。
	Print[int]([]int{1, 2, 3}) // 可以省略，int 类型参数。因为编译器可以从参数列表中推断
	// output:
	// 1
	// 2
	// 3
}

func ExamplePrint2() {
	Print2([]int{1, 2, 3}, []int{4, 5, 6})
	// output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExamplePrint2Same() {
	Print2Same([]int{1, 2, 3}, []int{4, 5, 6})
	// output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

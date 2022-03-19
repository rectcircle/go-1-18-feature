package generics

import "fmt"

// Print 打印切片的元素。
// Print 有一个类型参数 T，且有一个（非类型）普通参数 s，
// s 是一个元素类型为 T 的切片。
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Print2 打印两个接片的元素
// Print 有两个类型参数 T1 和 T2，且有两个（非类型）普通参数 s1 和 s2，
// s1  是一个元素类型为 T1 的切片，s2  是一个元素类型为 T1 的切片。
func Print2[T1, T2 any](s1 []T1, s2 []T2) {
	Print(s1)
	Print(s2)
}

// Print2 打印两个接片的元素
// Print 有一个类型参数 T，且有两个（非类型）普通参数 s1 和 s2，
// s1, s2 都是一个元素类型为 T 的切片
func Print2Same[T any](s1 []T, s2 []T) {
	Print(s1)
	Print(s2)
}

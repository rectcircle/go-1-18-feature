package generics

// Vector 是任何元素类型的切片。
type Vector[T any] []T

// Push 将元素添加到 Vector 的末尾。
func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}

// List 一个通用的链表类型
type List[T any] struct {
	next *List[T] // 引用自身
	val  T
}

// Adder 泛型接口
type Adder[T any] interface {
	Add(a, b T) T
}

// Object 泛型指针
type Object[T any] *T

// type T[P any] P // Error: cannot use a type parameter as RHS in type declaration https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#MisplacedTypeParam

type MyMap[K comparable, V any] map[K]V

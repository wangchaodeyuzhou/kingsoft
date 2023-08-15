package genericity

/* 泛型 key, value */
func MapKeys[K comparable, V any](m map[K]V) []K {
	l := make([]K, 0, len(m))
	for k := range m {
		l = append(l, k)
	}
	return l
}

type List[T any] struct {
	head, tail *element[T] // 链表的头与尾
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil { // 当放入的是第一个元素的时候
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else { // 赋值并移动指针
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for k := lst.head; k != nil; k = k.next {
		elems = append(elems, k.val)
	}
	return elems
}

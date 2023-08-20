package list

func NewLinkedList[T any]() List[T] {
	return newLinkedList[T]()
}

func NewLinkedListFromSlice[T any](s []T) List[T] {
	l := newLinkedList[T]()
	for _, v := range s {
		l.Add(v)
	}
	return l
}

type linkedList[T any] struct {
	List[T]

	root *node[T]
	end  *node[T]
	len  int
}

func newLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		root: nil,
		end:  nil,
		len:  0,
	}
}

func (l *linkedList[T]) Add(v T) {
	n := newNode(v)
	l.setRootIfNil(n)
	if l.root == l.end {
		l.root.next = n
	}
	l.end.next = n
	l.end = n
	l.len += 1
}

func (l *linkedList[T]) Insert(i int, v T) {
	n := newNode(v)
	l.setRootIfNil(n)
	// TODO
}

func (l *linkedList[T]) Concat(list List[T]) {
	linkedList, ok := list.(*linkedList[T])
	if ok {
		l.end.next = linkedList.root
		l.end = linkedList.end
		l.len += linkedList.len
		return
	}

	list.Each(func(v *T) {
		l.Add(*v)
	})
}

func (l *linkedList[T]) setRootIfNil(n *node[T]) {
	if l.root != nil {
		return
	}
	l.root = n
	l.end = n
	l.len = 1
}

// at n 番目の Node を取得
func (l *linkedList[T]) at(n int) *node[T] {
	if n > l.len { // Node が存在しないので、 nil を return
		return nil
	}

	ptr := l.root
	for i := 1; i < n; i++ { // n 番目の Node まで、 root から辿っていく
		ptr = ptr.next
	}
	return ptr
}

func (l linkedList[T]) index(i int) int {
	if i < 0 {
		return l.len + i
	}
	return i
}

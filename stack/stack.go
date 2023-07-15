package stack

type stack[T any] struct {
    Push func(T)
    Pop func() T
    Top func() T
    Size func() int
}

func Stack[T any]() stack[T] {
    source := make([]T, 0)
    return stack[T]{
        Push: func(i T) {
            source = append(source, i)
        },
        Pop: func() T {
            poped := source[len(source)-1]
            source = source[:len(source)-1]
            return poped
        },
        Top: func() T {
            return source[len(source)-1]
        },
        Size: func() int {
            return len(source)
        },
    }
}

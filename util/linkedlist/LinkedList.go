package linkedlist

type List[T any] interface {
    Add(data T)
    Remove(data T)
}

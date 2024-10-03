package linkedlist


// Singly Linked List Node
type SinglyNode[T any] struct {
    Data T
    Next *SinglyNode[T]
}


// Singly Linked List
type SinglyLinkedList[T any] struct {
    Head *SinglyNode[T]
}

func (l *SinglyLinkedList[T]) Add(data T) {
    newNode := &SinglyNode[T]{Data: data}
    if l.Head == nil {
        l.Head = newNode
    } else {
        current := l.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
}



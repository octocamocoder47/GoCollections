package linkedlist

import "github.com/octocamocoder47/GoCollections/util/generics"

// DoublyNode represents a node in a doubly linked list.
type DoublyNode[T any] struct {
    Data T
    Next *DoublyNode[T]
    Prev *DoublyNode[T]
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList[T any] struct {
    Head *DoublyNode[T]
    Tail *DoublyNode[T]
}

// Add appends a new node with the given data to the end of the doubly linked list.
func (l *DoublyLinkedList[T]) Add(data T) {
    newNode := &DoublyNode[T]{Data: data}
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        newNode.Prev = l.Tail
        l.Tail.Next = newNode
        l.Tail = newNode
    }
}

// Remove for comparable types
func (l *DoublyLinkedList[T]) Remove(data T) {
    if l.Head == nil {
        return // List is empty, nothing to remove
    }

    if(!generics.isComparable(data)) {
        return
    }

    current := l.Head
    // Using built-in comparison for comparable types
    for current != nil && current.Data != data {
        current = current.Next
    }

    if current == nil {
        return // Data not found in the list
    }

    l.removeNode(current)
}

// Remove with a custom comparison function for non-comparable types
func (l *DoublyLinkedList[T]) RemoveWithComparator(data T, equals func(a, b T) bool) {
    if l.Head == nil {
        return // List is empty, nothing to remove
    }

    current := l.Head
    // Use the custom comparator
    for current != nil && !equals(current.Data, data) {
        current = current.Next
    }

    if current == nil {
        return // Data not found in the list
    }

    l.removeNode(current)
}

// removeNode handles node removal common to both Remove methods.
func (l *DoublyLinkedList[T]) removeNode(node *DoublyNode[T]) {
    // If node.Prev is not nil, we are not removing the head
    if node.Prev != nil {
        node.Prev.Next = node.Next
    } else {
        // If node.Prev is nil, we're removing the head, so move head forward
        l.Head = node.Next
    }

    // If node.Next is not nil, we are not removing the tail
    if node.Next != nil {
        node.Next.Prev = node.Prev
    } else {
        // If node.Next is nil, we're removing the tail, so move tail backward
        l.Tail = node.Prev
    }
}

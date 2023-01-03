package list

import (
	"github.com/tomet/container/slices"
)

//--------------------------------------------------------------------------------
// List
//--------------------------------------------------------------------------------

type List[T comparable] struct {
	items []T
}

func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

func NewListFromSlice[T comparable](s []T) *List[T] {
	return &List[T]{
		items: s,
	}
}

func (l *List[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, e := range l.items {
			ch <- e
		}
	}()
	return ch
}

func (l *List[T]) Size() int {
	return len(l.items)
}

func (l *List[T]) Get(idx int) T {
	if l.ValidIndex(idx) {
		return l.items[idx]
	}
	var zero T
	return zero
}

func (l *List[T]) ValidIndex(idx int) bool {
	return idx >= 0 && idx < l.Size()
}

func (l *List[T]) Append(e ...T) {
	l.items = append(l.items, e...)
}

func (l *List[T]) Remove(e T) bool {
	if i := slices.Index(l.items, e); i >= 0 {
		l.items = slices.RemoveAt(l.items, i)
		return true
	}
	return false
}

func (l *List[T]) RemoveAt(idx int) bool {
	if l.ValidIndex(idx) {
		l.items = slices.RemoveAt(l.items, idx)
		return true
	}
	return false
}

func (l *List[T]) RemoveIf(f slices.FilterFunc[T]) {
	l.items = slices.Reject(l.items, f)
}

func (l *List[T]) Index(e T) int {
	return slices.Index(l.items, e)
}

//--------------------------------------------------------------------------------
// Join
//--------------------------------------------------------------------------------

func (l *List[T]) Join(delim string) {
	
}

//--------------------------------------------------------------------------------
// Stringer
//--------------------------------------------------------------------------------

func (l *List[T]) String() string {
	return "List"
}

//--------------------------------------------------------------------------------
// Filter/Reject
//--------------------------------------------------------------------------------

func (l *List[T]) Filter(f slices.FilterFunc[T]) *List[T] {
	return &List[T]{
		items: slices.Filter(l.items, f),
	}
}

func (l *List[T]) Reject(f slices.FilterFunc[T]) *List[T] {
	return NewListFromSlice(slices.Reject(l.items, f))
}

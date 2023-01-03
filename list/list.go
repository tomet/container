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

//--------------------------------------------------------------------------------
// New List
//--------------------------------------------------------------------------------


func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

func NewListFromSlice[T comparable](s []T) *List[T] {
	return &List[T]{
		items: s,
	}
}

func NewListOf[T comparable](v ...T) *List[T] {
	return &List[T]{
		items: v,
	}
}

//--------------------------------------------------------------------------------
// Iterator
//--------------------------------------------------------------------------------

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

//--------------------------------------------------------------------------------
// Size/Get/Set/ValidIdx
//--------------------------------------------------------------------------------

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

func (l *List[T]) Set(idx int, e T) {
	l.items[idx] = e
}

func (l *List[T]) ValidIndex(idx int) bool {
	return idx >= 0 && idx < l.Size()
}

func (l *List[T]) Index(e T) int {
	return slices.Index(l.items, e)
}

//--------------------------------------------------------------------------------
// Append/Insert
//--------------------------------------------------------------------------------

func (l *List[T]) Insert(i int, v ...T) {
	l.InsertSlice(i, v)
}

func (l *List[T]) InsertSlice(i int, s []T) {
	l.items = slices.InsertSlice(l.items, i, s)
}

func (l *List[T]) InsertList(i int, l2 *List[T]) {
	l.InsertSlice(i, l2.items)
}

func (l *List[T]) Append(e ...T) {
	l.items = append(l.items, e...)
}

func (l *List[T]) AppendSlice(s []T) {
	l.items = slices.AppendSlice(l.items, s)
}

func (l *List[T]) AppendList(l2 *List[T]) {
	l.AppendSlice(l2.items)
}

//--------------------------------------------------------------------------------
// Remove
//--------------------------------------------------------------------------------

func (l *List[T]) Remove(e T) bool {
	if i := slices.Index(l.items, e); i >= 0 {
		l.items = slices.RemoveAt(l.items, i)
		return true
	}
	return false
}

func (l *List[T]) RemoveAt(idx int) T {
	if l.ValidIndex(idx) {
		v := l.Get(idx)
		l.items = slices.RemoveAt(l.items, idx)
		return v
	}
	var zero T
	return zero
}

func (l *List[T]) RemoveIf(f slices.FilterFunc[T]) {
	l.items = slices.Reject(l.items, f)
}

//--------------------------------------------------------------------------------
// Join
//--------------------------------------------------------------------------------

func (l *List[T]) Join(delim string) string {
	return slices.Join(l.items, delim)
}

//--------------------------------------------------------------------------------
// Map 
//--------------------------------------------------------------------------------

// methods can't have type parameters, so i added a function
func Map[T comparable, E comparable](l *List[T], f func(T) E) *List[E] {
	return NewListFromSlice(slices.Map(l.items, f))
}

// works only if it maps to the same type!!!
func (l *List[T]) Map(f func(e T) T) *List[T] {
	return NewListFromSlice(slices.Map(l.items, f))
}

//--------------------------------------------------------------------------------
// Filter/Reject
//--------------------------------------------------------------------------------

func (l *List[T]) Filter(f slices.FilterFunc[T]) *List[T] {
	return NewListFromSlice(slices.Filter(l.items, f))
}

func (l *List[T]) Reject(f slices.FilterFunc[T]) *List[T] {
	return NewListFromSlice(slices.Reject(l.items, f))
}

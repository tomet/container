package list

import (
	"sort"

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

// Erzeugt eine neue, leere List
func New[T comparable]() *List[T] {
	return &List[T]{}
}

// Erzeugt aus einem Slice eine neue List
func FromSlice[T comparable](s []T) *List[T] {
	return &List[T]{
		items: s,
	}
}

// Erzeugt aus allen Argumenten eine neue List
func Of[T comparable](v ...T) *List[T] {
	return &List[T]{
		items: v,
	}
}

//--------------------------------------------------------------------------------
// Each
//--------------------------------------------------------------------------------

// Ruft fn für alle Einträge einmal auf
func (l *List[T]) Each(fn func(item T)) {
	for _, item := range l.items {
		fn(item)
	}
}

// Ruft fn(idx, item) für alle Einträge einmal auf.
func (l *List[T]) EachIndexed(fn func(idx int, item T)) {
	for idx, item := range l.items {
		fn(idx, item)
	}
}

//--------------------------------------------------------------------------------
// Len/Get/Set/ValidIdx
//--------------------------------------------------------------------------------

// Liefert die Länge der List (Anzahl der Elemente)
func (l *List[T]) Len() int {
	return len(l.items)
}

// Liefert den Eintrag des entsprechenden Indexes.
// Ist der Index nicht vorhanden, wird eine panic ausgelöst.
func (l *List[T]) Get(idx int) T {
	return l.items[idx]
}

// Setzt den Eintrag des entsprechenden Indexes.
func (l *List[T]) Set(idx int, e T) {
	l.items[idx] = e
}

// Prüft, ob idx ein gültiger Index ist.
func (l *List[T]) ValidIndex(idx int) bool {
	return idx >= 0 && idx < l.Len()
}

// Sucht den Eintrag e und liefert dessen Index oder -1.
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

// Entfernt den Eintrag an Index idx und liefert das entfernte Element.
func (l *List[T]) RemoveAt(idx int) T {
	if l.ValidIndex(idx) {
		v := l.Get(idx)
		l.items = slices.RemoveAt(l.items, idx)
		return v
	}
	var zero T
	return zero
}

// Löscht alle Einträge, für die fn true liefert.
func (l *List[T]) RemoveIf(fn slices.FilterFunc[T]) {
	l.items = slices.Reject(l.items, fn)
}

//--------------------------------------------------------------------------------
// Sort
//--------------------------------------------------------------------------------

// Sortiert die Liste (in-place).
// isLessFn muß true liefern, wenn a < b ist.
func (l *List[T]) Sort(isLessFn func(a, b T) bool) {
	sort.Slice(l.items, func(i, j int) bool {
		return isLessFn(l.items[i], l.items[j])
	})
}

// Sortiert die Liste (in-place) umgekehrt. (biggest first)
func (l *List[T]) SortReverse(isLessFn func(a, b T) bool) {
	sort.Slice(l.items, func(i, j int) bool {
		return !isLessFn(l.items[i], l.items[j])
	})
}

//--------------------------------------------------------------------------------
// Join
//--------------------------------------------------------------------------------

// Fügt alle Einträge zu einem String zusammen.
func (l *List[T]) Join(delim string) string {
	return slices.Join(l.items, delim)
}

//--------------------------------------------------------------------------------
// Map
//--------------------------------------------------------------------------------

// methods can't have type parameters, so i added a function
func Map[T comparable, E comparable](l *List[T], f func(T) E) *List[E] {
	return FromSlice(slices.Map(l.items, f))
}

// works only if it maps to the same type!!!
func (l *List[T]) Map(f func(e T) T) *List[T] {
	return FromSlice(slices.Map(l.items, f))
}

//--------------------------------------------------------------------------------
// Filter/Reject
//--------------------------------------------------------------------------------

// Filtert die List und liefert eine neue List.
func (l *List[T]) Filter(f slices.FilterFunc[T]) *List[T] {
	return FromSlice(slices.Filter(l.items, f))
}

// Liefert eine neue List, in der alle gefilterten Einträge entfernt sind.
func (l *List[T]) Reject(f slices.FilterFunc[T]) *List[T] {
	return FromSlice(slices.Reject(l.items, f))
}

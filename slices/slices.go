package slices

import (
	"fmt"
	"strings"
)

type FilterFunc[F comparable] func(e F) bool

func Clone[T any](s []T) []T {
	// in case we got a nil, we keep it
	if s == nil {
		return nil
	}
	s2 := make([]T, len(s), cap(s))
	copy(s2, s)
	return s2
}

//--------------------------------------------------------------------------------
// Insert
//--------------------------------------------------------------------------------

// this is the main-function, which gets called from Insert and AppendSlice
func InsertSlice[T any](s []T, i int, o []T) []T {
	l := len(s)
	var s2 []T

	for j := 0; j < i; j++ {
		s2 = append(s2, s[j])
	}

	for _, v := range o {
		s2 = append(s2, v)
	}

	for i < l {
		s2 = append(s2, s[i])
		i++
	}

	return s2
}

func AppendSlice[T any](s []T, s2 []T) []T {
	return InsertSlice(s, len(s), s2)
}

func Insert[T any](s []T, i int, v ...T) []T {
	return InsertSlice(s, i, v)
}

//--------------------------------------------------------------------------------
// Remove
//--------------------------------------------------------------------------------

func RemoveAt[T any](s []T, i int) []T {
	l := len(s) - 1
	for i < l {
		s[i] = s[i+1]
		i++
	}
	if l == 0 {
		return nil
	} else {
		return s[:l]
	}
}

func Remove[T comparable](s []T, e T) []T {
	return RemoveAt(s, Index(s, e))
}

//--------------------------------------------------------------------------------
// Index
//--------------------------------------------------------------------------------

func Index[T comparable](s []T, e T) int {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func Contains[T comparable](s []T, e T) bool {
	return Index(s, e) >= 0
}

//--------------------------------------------------------------------------------
// Map
//--------------------------------------------------------------------------------

func Map[T any, F any](s []T, f func(e T) F) []F {
	s2 := make([]F, len(s))

	for i, v := range s {
		s2[i] = f(v)
	}

	return s2
}

//--------------------------------------------------------------------------------
// Filter
//--------------------------------------------------------------------------------

func Filter[T comparable](s []T, f FilterFunc[T]) []T {
	var s2 []T
	for _, v := range s {
		if f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func Reject[T comparable](s []T, f FilterFunc[T]) []T {
	var s2 []T
	for _, v := range s {
		if !f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

//--------------------------------------------------------------------------------
// Count
//--------------------------------------------------------------------------------

func Count[T comparable](s []T, f FilterFunc[T]) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}

//--------------------------------------------------------------------------------
// Join (string)
//--------------------------------------------------------------------------------

func Join[T any](s []T, delim string) string {
	parts := Map(s, func(s T) string {
		return fmt.Sprintf("%v", s)
	})
	return strings.Join(parts, delim)
}

package slices

import (
	"fmt"
	"testing"
)

//--------------------------------------------------------------------------------
// Tests
//--------------------------------------------------------------------------------

func TestClone(t *testing.T) {
	s := []int{1, 2, 3}
	
	s2 := Clone(s)
	
	addrS := fmt.Sprintf("%p", s)
	addrS2 := fmt.Sprintf("%p", s2)
	assertNotEqual(t, addrS, addrS2)
	
	assertEqualSlice(t, s2, "[1 2 3]")
	assertEqual(t, len(s2), len(s))
	assertEqual(t, cap(s2), cap(s))
}

func TestMap(t *testing.T) {
	s := []string{"1", "22", "333"}
	
	s2 := Map(s, func(s string) int {
		return len(s)
	})
	
	assertEqualSlice(t, s2, "[1 2 3]")
}

func TestJoin(t *testing.T) {
	s := []int{1, 2, 3}
	
	s2 := Join(s, ", ")
	assertEqual(t, s2, "1, 2, 3")
}

func TestAppendSlice(t *testing.T) {
	s := []int{1, 2, 3}
	
	s = AppendSlice(s, []int{4, 5})
	
	assertEqualSlice(t, s, "[1 2 3 4 5]")
}

func TestInsert(t *testing.T) {
	s := []int{}
	
	s = Insert(s, 0, 2)
	s = Insert(s, 0, 1)
	s = Insert(s, 2, 3, 4, 5)

	assertEqualSlice(t, s, "[1 2 3 4 5]")
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}
	
	assertTrue(t, Contains(s, 2))
	assertFalse(t, Contains(s, 5))
}

func TestRemove(t *testing.T) {
	s := []int{1, 2, 3}
	
	// Remove() uses Index() and RemoveAt(), so i don't
	// have to test them
	s = Remove(s, 2)

	assertEqualSlice(t, s, "[1 3]")
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	s = Filter(s, func(v int) bool {
		return v == 2 || v == 5
	})

	assertEqualSlice(t, s, "[2 5]")
}

func TestReject(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	
	s = Reject(s, func(v int) bool {
		return v == 2 || v == 5
	})
	
	assertEqualSlice(t, s, "[1 3 4]")
}

func TestCount(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	
	n := Count(s, func(v int) bool {
		return v > 3
	})
	
	assertEqual(t, n, 2)
}

//--------------------------------------------------------------------------------
// Assertions
//--------------------------------------------------------------------------------

func assertEqualSlice(t *testing.T, got []int, want string) {
	gotStr := fmt.Sprintf("%+v", got)
	if gotStr != want {
		t.Errorf("want: %s, got: %s", want, gotStr)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func assertNotEqual[T comparable](t *testing.T, got, dont_want T) {
	if got == dont_want {
		t.Errorf("don't want: %v, got: %v", dont_want, got)
	}
}

func assertTrue(t *testing.T, got bool) {
	if got != true {
		t.Errorf("want: true, got: false")
	}
}

func assertFalse(t *testing.T, got bool) {
	if got != false {
		t.Errorf("want: false, got: true")
	}
}

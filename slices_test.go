package slices

import (
	"fmt"
	"testing"
)

//--------------------------------------------------------------------------------
// Tests
//--------------------------------------------------------------------------------

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

func assertEqual(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
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

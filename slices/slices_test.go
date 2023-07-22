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

func TestSort(t *testing.T) {
	s := []int{3, 5, 1}
	
	Sort(s, func(a, b int) bool {
		return a < b
	})
	
	assertEqualSlice(t, s, "[1 3 5]")
	
	SortReverse(s, func(a, b int) bool {
		return a < b
	})
	
	assertEqualSlice(t, s, "[5 3 1]")
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

func TestFind(t *testing.T) {
	s := []int{1, 2, 3}
	
	assertEqual(t, Find(s, func(v int) bool { return v == 2 }), 1)
	assertEqual(t, Find(s, func(v int) bool { return v == 10 }), -1)
}

func TestMatchesAny(t *testing.T) {
	s := []int{1, 2, 3}
	assertTrue(t, MatchesAny(s, func(i int) bool { return i == 1 }))
	assertFalse(t, MatchesAny(s, func(i int) bool { return i == 10 }))
}

func TestMatchesAll(t *testing.T) {
	s := []int{1, 2, 3}
	assertTrue(t, MatchesAll(s, func(i int) bool { return i < 10 }))
	assertFalse(t, MatchesAll(s, func(i int) bool { return i > 2 }))
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

func TestEach(t *testing.T) {
	s := []int{1, 2, 3}
	sum := 0
	Each(s, func(v int) { 
		sum += v
	})
	assertEqual(t, sum, 6)
	
	sum = 0
	idxs := 0
	EachIndexed(s, func(idx int, v int) {
		sum += v
		idxs += idx
	})
	assertEqual(t, sum, 6)
	assertEqual(t, idxs, 3)
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

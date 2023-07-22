package list

import (
	"fmt"
	"testing"
)

func TestIndexMethods(t *testing.T) {
	l := New[int]()
	
	l.Append(2, 30, 4, 5)
	l.Insert(0, 1)
	
	l.Set(2, 3)
	
	assertEqual(t, l.Len(), 5)
	assertEqual(t, l.Get(0), 1)
	
	assertTrue(t, l.ValidIndex(0))
	assertTrue(t, l.ValidIndex(4))
	assertFalse(t, l.ValidIndex(5))
	assertFalse(t, l.ValidIndex(-1))
	
	assertEqual(t, l.Index(2), 1)
	assertEqual(t, l.Index(10), -1)
	
	assertEqualList(t, l, "&{[1 2 3 4 5]}")
}

func TestEach(t *testing.T) {
	l := Of(1, 2, 3)
	sum := 0
	
	l.Each(func(v int) {
		sum += v
	})
	
	assertEqual(t, sum, 6)
	
	idxs := 0 
	
	l.EachIndexed(func(idx int, v int) {
		idxs += idx
	})
	
	assertEqual(t, idxs, 3)
}

func TestSort(t *testing.T) {
	l := Of(3, 5, 1)
	
	l.Sort(func(a, b int) bool {
		return a < b
	})
	assertEqualList(t, l, "&{[1 3 5]}")
	
	l.SortReverse(func(a, b int) bool {
		return a < b
	})
	assertEqualList(t, l, "&{[5 3 1]}")
}

func TestAppend(t *testing.T) {
	l := Of(1, 2, 3)
	l2 := FromSlice([]int{4, 5})
	
	l.AppendList(l2)
	
	assertEqualList(t, l, "&{[1 2 3 4 5]}")
}

func TestRemove(t *testing.T) {
	l := FromSlice([]int{1, 2, 3, 4, 5})
	
	assertTrue(t, l.Remove(2))
	assertFalse(t, l.Remove(10))
	assertEqualList(t, l, "&{[1 3 4 5]}")
	
	assertEqual(t, l.RemoveAt(0), 1)
	assertEqual(t, l.RemoveAt(10), 0)
	
	l.RemoveIf(func(v int) bool {
		return v == 5
	})
	assertEqualList(t, l, "&{[3 4]}")
}

func TestMapFunc(t *testing.T) {
	l := Of(1, 2, 3)
	
	l2 := Map(l, func(v int) string {
		return fmt.Sprintf("(%d)", v)
	})
	
	assertEqualList(t, l2, "&{[(1) (2) (3)]}")
}

func TestMapMethod(t *testing.T) {
	l := Of(1, 2, 3)
	l2 := l.Map(func(v int) int {
		return v * 10
	})
	assertEqualList(t, l2, "&{[10 20 30]}")
}

func TestJoin(t *testing.T) {
	l := Of(1, 2, 3)
	assertEqual(t, l.Join(", "), "1, 2, 3")
}

func TestFilter(t *testing.T) {
	l := Of(1, 2, 3)
	l2 := l.Filter(func(i int) bool {
		return i > 1
	})
	assertEqualList(t, l2, "&{[2 3]}")
}

func TestReject(t *testing.T) {
	l := Of(1, 2, 3)
	l2 := l.Reject(func(i int) bool {
		return i == 1
	})
	assertEqualList(t, l2, "&{[2 3]}")
}

//--------------------------------------------------------------------------------
// Assertions
//--------------------------------------------------------------------------------

func assertEqualList[T comparable](t *testing.T, got *List[T], want string) {
	gotStr := fmt.Sprintf("%v", got)
	if gotStr != want {
		t.Errorf("want: %s, got: %s", want, gotStr)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func assertNotEqual(t *testing.T, got, dont_want int) {
	if got == dont_want {
		t.Errorf("don't want: %d, got: %d", dont_want, got)
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

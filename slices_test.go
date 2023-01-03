package slices

import (
	"fmt"
	"testing"
)

//--------------------------------------------------------------------------------
// Tests
//--------------------------------------------------------------------------------

func TestInsert(t *testing.T) {
	var s []int

	s = Insert(s, 0, 2)
	s = Insert(s, 0, 1)
	s = Insert(s, 2, 3)

	assertEqual(t, s, "[1 2 3]")
}

func TestRemove(t *testing.T) {
	s := []int{1, 2, 3}
	
	s = Remove(s, 2)
	
	assertEqual(t, s, "[1 3]")
}

//--------------------------------------------------------------------------------
// Assertions
//--------------------------------------------------------------------------------

func assertEqual(t *testing.T, got []int, want string) {
	gotStr := fmt.Sprintf("%+v", got)
	if gotStr != want {
		t.Errorf("want: %s, got: %s", want, gotStr)
	}
}

//func assertTrue(t *testing.T, got bool) {
	//if got != true {
		//f.Errorf("want: true, got: false")
	//}
//}

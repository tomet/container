package maps

import (
	"testing"
	"sort"
)

func TestKeys(t *testing.T) {
	keys := Keys(map[string]bool {
		"a": true,
		"b": false,
	})
	
	sort.Strings(keys)
	assertEqual(t, keys[0], "a")
	assertEqual(t, keys[1], "b")
}

func TestValues(t *testing.T) {
	values := Values(map[string]int {
		"a": 1,
		"b": 2,
	})
	
	sort.Ints(values)
	assertEqual(t, values[0], 1)
	assertEqual(t, values[1], 2)
}

//--------------------------------------------------------------------------------
// Assertions
//--------------------------------------------------------------------------------

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

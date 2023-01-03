package slices

type FilterFunc[T comparable] func(e T) bool

func Clone[S ~[]E, E any](s S) S {
	// falls s nil ist und NICHT ein leeres Slice
	if s == nil {
		return nil
	}
	s2 := make(S, len(s), cap(s))
	copy(s2, s)
	return s2
}

func AppendSlice[S ~[]E, E any](s S, s2 S) S {
	for _, v := range s2 {
		s = append(s, v)
	}
	return s
}

func Insert[S ~[]E, E any](s S, i int, v E) S {
	l := len(s)
	var s2 S
	
	for j := 0; j < i; j++ {
		s2 = append(s2, s[j])
	}
	
	s2 = append(s2, v)
	
	for i < l {
		s2 = append(s2, s[i])
		i++
	}
	
	return s2
}

func RemoveAt[S ~[]E, E any](s S, i int) S {
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

func Remove[S ~[]E, E comparable](s S, e E) S {
	return RemoveAt(s, Index(s, e))
}

func Index[S ~[]E, E comparable](s S, e E) int {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func Contains[S ~[]E, E comparable](s S, e E) bool {
	return Index(s, e) >= 0
}

func Filter[S ~[]E, E comparable](s S, f FilterFunc[E]) S {
	var s2 S
	for _, v := range s {
		if f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func Reject[S ~[]E, E comparable](s S, f FilterFunc[E]) S {
	var s2 S
	for _, v := range s {
		if !f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func Count[S ~[]E, E comparable](s S, f FilterFunc[E]) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}
package ints

type Set []Range

func (s *Set) Add(i, n uint64) {
	*s = (*s).insert(&Range{Beg: i, Len: n})
}

func (s Set) Head() uint64 {
	if len(s) > 0 {
		r := &s[0]
		if r.Beg == 0 {
			return r.Len
		}
	}
	return 0
}

func (s Set) insert(r *Range) (dst Set) {
	i, n := s.search(r)
	if n > 0 {
		s[i].setEnd(r.End())
		j := i + 1
		j += copy(s[j:], s[i+n:])
		return s[:j]
	}
	if j := len(s) + 1; j < cap(s) {
		dst = s[:j]
	} else {
		dst = make(Set, j)
		copy(dst, s[:i])
	}
	copy(dst[i+1:], s[i:])
	dst[i] = *r
	return
}

func (s Set) search(r *Range) (i, n int) {
	// Copied and modified from sort/search.go in the standard library.
	for j := len(s); i < j; {
		mid := int(uint(i+j) >> 1) // avoid overflow
		// i ≤ mid < j

		if s[mid].End() < r.Beg {
			i = mid + 1 // false
		} else {
			j = mid // true
		}
	}
	if i == len(s) || r.Beg < s[i].Beg {
		return
	}
	n++
	end := r.End()
	if s[i].End() >= end {
		return
	}
	for {
		j := i + n
		if j >= len(s) {
			break
		}
		x := &s[j]
		if x.Beg > end {
			break
		}
		n++
		if x.Beg == end {
			r.setEnd(x.End())
			break
		}
	}
	return
}

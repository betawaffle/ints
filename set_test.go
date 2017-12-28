package ints

import (
	"testing"
)

func TestSet(t *testing.T) {
	tests := [...]struct {
		Add    []uint64
		Expect uint64
	}{
		{[]uint64{5, 0, 1, 2, 3, 4}, 6},
	}
	var s Set
	for _, test := range tests {
		for _, x := range test.Add {
			s.Add(x, 1)
			t.Logf("state: %v", s)
		}
		if v := s.Head(); v != test.Expect {
			t.Fatalf("expected Get to return %d, got %d; state: %v", test.Expect, v, s)
		}
	}
}

func BenchmarkSet(b *testing.B) {
	var s Set
	for i := 0; i < b.N; i += 6 {
		s.Add(uint64(i+5), 1)
		s.Add(uint64(i), 1)
		s.Add(uint64(i+1), 1)
		s.Add(uint64(i+2), 1)
		s.Add(uint64(i+3), 1)
		s.Add(uint64(i+4), 1)
	}
}

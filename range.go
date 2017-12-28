package ints

type Range struct {
	Beg uint64
	Len uint64
}

func (r *Range) End() uint64 {
	return saturatingAdd(r.Beg, r.Len)
}

func (r *Range) setEnd(end uint64) {
	r.Len = end - r.Beg
}

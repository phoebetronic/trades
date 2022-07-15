package floats

import (
	"sort"
)

type Floats struct {
	FL []float32
}

func (f Floats) Avg() float32 {
	return avg(f.FL)
}

func (f Floats) Med() float32 {
	l := len(f.FL)

	if l == 0 {
		return 0
	}

	c := sor(cop(f.FL))

	if l%2 == 0 {
		return avg(c[l/2-1 : l/2+1])
	}

	return c[l/2]
}

func (f Floats) Sum() float32 {
	return sum(f.FL)
}

func avg(f []float32) float32 {
	if len(f) == 0 {
		return 0
	}

	return sum(f) / float32(len(f))
}

func cop(f []float32) []float32 {
	c := make([]float32, len(f))
	copy(c, f)
	return c
}

func sor(f []float32) []float32 {
	c := cop(f)
	sort.SliceStable(c, func(i, j int) bool { return c[i] < c[j] })
	return c
}

func sum(f []float32) float32 {
	var s float32

	for _, v := range f {
		s += v
	}

	return s
}

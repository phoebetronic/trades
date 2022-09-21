package orders

import "math"

func (b *Bundle) Mid() float32 {
	var hig float32
	var low float32
	{
		hig = b.minord(b.AK)
		low = b.maxord(b.BD)
	}

	return (low + hig) / 2
}

func (b *Bundle) maxord(ord []*Order) float32 {
	var max float32
	{
		max = math.MaxFloat32 * -1
	}

	for _, x := range ord {
		if x.PR > max {
			max = x.PR
		}
	}

	return max
}

func (b *Bundle) minord(ord []*Order) float32 {
	var min float32
	{
		min = math.MaxFloat32
	}

	for _, x := range ord {
		if x.PR < min {
			min = x.PR
		}
	}

	return min
}

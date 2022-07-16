package floats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Typ_Floats_Avg(t *testing.T) {
	testCases := []struct {
		fl []float32
		re float32
	}{
		// Case 0
		{
			fl: nil,
			re: 0,
		},
		// Case 1
		{
			fl: []float32{
				33,
			},
			re: 33,
		},
		// Case 2
		{
			fl: []float32{
				33,
				37,
			},
			re: 35,
		},
		// Case 3
		{
			fl: []float32{
				33,
				37,
				36,
			},
			re: 35.333332,
		},
		// Case 4
		{
			fl: []float32{
				33,
				37,
				36,
				34,
			},
			re: 35,
		},
		// Case 5
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
			},
			re: 35.6,
		},
		// Case 6
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
				35,
			},
			re: 35.5,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f Floats
			{
				f.FL = tc.fl
			}

			var re float32
			{
				re = f.Avg()
			}

			if !reflect.DeepEqual(tc.re, re) {
				t.Fatalf("re\n\n%s\n", cmp.Diff(tc.re, re))
			}
		})
	}
}

func Test_Typ_Floats_Max(t *testing.T) {
	testCases := []struct {
		fl []float32
		re float32
	}{
		// Case 0
		{
			fl: nil,
			re: 0,
		},
		// Case 1
		{
			fl: []float32{
				33,
			},
			re: 33,
		},
		// Case 2
		{
			fl: []float32{
				-33,
			},
			re: -33,
		},
		// Case 3
		{
			fl: []float32{
				33,
				-37,
				36,
				34,
				-38,
				-35,
			},
			re: 36,
		},
		// Case 4
		{
			fl: []float32{
				-33,
				-37,
				-36,
				-34,
				-38,
				-35,
			},
			re: -33,
		},
		// Case 5
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
				35,
			},
			re: 38,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f Floats
			{
				f.FL = tc.fl
			}

			var re float32
			{
				re = f.Max()
			}

			if !reflect.DeepEqual(tc.re, re) {
				t.Fatalf("re\n\n%s\n", cmp.Diff(tc.re, re))
			}
		})
	}
}

func Test_Typ_Floats_Med(t *testing.T) {
	testCases := []struct {
		fl []float32
		re float32
	}{
		// Case 0
		{
			fl: nil,
			re: 0,
		},
		// Case 1
		{
			fl: []float32{
				33,
			},
			re: 33,
		},
		// Case 2
		{
			fl: []float32{
				33,
				37,
			},
			re: 35,
		},
		// Case 3
		{
			fl: []float32{
				33,
				37,
				36,
			},
			re: 36,
		},
		// Case 4
		{
			fl: []float32{
				33,
				37,
				36,
				34,
			},
			re: 35,
		},
		// Case 5
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
			},
			re: 36,
		},
		// Case 6
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
				35,
			},
			re: 35.5,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f Floats
			{
				f.FL = tc.fl
			}

			var re float32
			{
				re = f.Med()
			}

			if !reflect.DeepEqual(tc.re, re) {
				t.Fatalf("re\n\n%s\n", cmp.Diff(tc.re, re))
			}
		})
	}
}

func Test_Typ_Floats_Min(t *testing.T) {
	testCases := []struct {
		fl []float32
		re float32
	}{
		// Case 0
		{
			fl: nil,
			re: 0,
		},
		// Case 1
		{
			fl: []float32{
				33,
			},
			re: 33,
		},
		// Case 2
		{
			fl: []float32{
				-33,
			},
			re: -33,
		},
		// Case 3
		{
			fl: []float32{
				33,
				-37,
				36,
				34,
				-38,
				-35,
			},
			re: -38,
		},
		// Case 4
		{
			fl: []float32{
				-33,
				-37,
				-36,
				-34,
				-38,
				-35,
			},
			re: -38,
		},
		// Case 5
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
				35,
			},
			re: 33,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f Floats
			{
				f.FL = tc.fl
			}

			var re float32
			{
				re = f.Min()
			}

			if !reflect.DeepEqual(tc.re, re) {
				t.Fatalf("re\n\n%s\n", cmp.Diff(tc.re, re))
			}
		})
	}
}

func Test_Typ_Floats_Sum(t *testing.T) {
	testCases := []struct {
		fl []float32
		re float32
	}{
		// Case 0
		{
			fl: nil,
			re: 0,
		},
		// Case 1
		{
			fl: []float32{
				33,
			},
			re: 33,
		},
		// Case 2
		{
			fl: []float32{
				33,
				37,
			},
			re: 70,
		},
		// Case 3
		{
			fl: []float32{
				33,
				37,
				36,
			},
			re: 106,
		},
		// Case 4
		{
			fl: []float32{
				33,
				37,
				36,
				34,
			},
			re: 140,
		},
		// Case 5
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
			},
			re: 178,
		},
		// Case 6
		{
			fl: []float32{
				33,
				37,
				36,
				34,
				38,
				35,
			},
			re: 213,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f Floats
			{
				f.FL = tc.fl
			}

			var re float32
			{
				re = f.Sum()
			}

			if !reflect.DeepEqual(tc.re, re) {
				t.Fatalf("re\n\n%s\n", cmp.Diff(tc.re, re))
			}
		})
	}
}

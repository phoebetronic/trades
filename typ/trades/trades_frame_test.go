package trades

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/xh3b4sd/framer"
)

func Test_Typ_Trades_Frame_Next(t *testing.T) {
	testCases := []struct {
		tr *Trades
		re []*Trades
	}{
		// Case 0
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-04-02T15:55:00Z"),
				TR: []*Trade{
					{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 123, LO: 3.8, TS: newTim("2022-04-02T13:59:59Z")},
					{PR: 137, SH: 2.3, TS: newTim("2022-04-02T14:00:00Z")},
					{PR: 140, SH: 2.5, TS: newTim("2022-04-02T14:02:00Z")},
					{PR: 110, SH: 3.1, TS: newTim("2022-04-02T14:22:00Z")},
					{PR: 115, SH: 1.8, TS: newTim("2022-04-02T14:27:00Z")},
					{PR: 117, SH: 2.9, TS: newTim("2022-04-02T14:59:59Z")},
					{PR: 118, LO: 2.1, TS: newTim("2022-04-02T15:00:00Z")},
					{PR: 125, LO: 4.2, TS: newTim("2022-04-02T15:11:00Z")},
					{PR: 130, SH: 0.6, TS: newTim("2022-04-02T15:16:00Z")},
					{PR: 128, LO: 1.5, TS: newTim("2022-04-02T15:33:00Z")},
					{PR: 135, LO: 0.6, TS: newTim("2022-04-02T15:48:00Z")},
				},
			},
			re: []*Trades{
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:55:00Z"),
					EN: newTim("2022-04-02T14:00:00Z"),
					TR: []*Trade{
						{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
						{PR: 123, LO: 3.8, TS: newTim("2022-04-02T13:59:59Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T14:00:00Z"),
					EN: newTim("2022-04-02T15:00:00Z"),
					TR: []*Trade{
						{PR: 137, SH: 2.3, TS: newTim("2022-04-02T14:00:00Z")},
						{PR: 140, SH: 2.5, TS: newTim("2022-04-02T14:02:00Z")},
						{PR: 110, SH: 3.1, TS: newTim("2022-04-02T14:22:00Z")},
						{PR: 115, SH: 1.8, TS: newTim("2022-04-02T14:27:00Z")},
						{PR: 117, SH: 2.9, TS: newTim("2022-04-02T14:59:59Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T15:00:00Z"),
					EN: newTim("2022-04-02T15:55:00Z"),
					TR: []*Trade{
						{PR: 118, LO: 2.1, TS: newTim("2022-04-02T15:00:00Z")},
						{PR: 125, LO: 4.2, TS: newTim("2022-04-02T15:11:00Z")},
						{PR: 130, SH: 0.6, TS: newTim("2022-04-02T15:16:00Z")},
						{PR: 128, LO: 1.5, TS: newTim("2022-04-02T15:33:00Z")},
						{PR: 135, LO: 0.6, TS: newTim("2022-04-02T15:48:00Z")},
					},
				},
			},
		},
		// Case 1
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T11:35:00Z"),
				EN: newTim("2022-04-02T17:24:00Z"),
				TR: []*Trade{
					{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 123, LO: 3.8, TS: newTim("2022-04-02T13:59:59Z")},
					{PR: 137, SH: 2.3, TS: newTim("2022-04-02T14:00:00Z")},
					{PR: 140, SH: 2.5, TS: newTim("2022-04-02T14:02:00Z")},
					{PR: 110, SH: 3.1, TS: newTim("2022-04-02T14:22:00Z")},
					{PR: 115, SH: 1.8, TS: newTim("2022-04-02T14:27:00Z")},
					{PR: 117, SH: 2.9, TS: newTim("2022-04-02T14:59:59Z")},
					{PR: 118, LO: 2.1, TS: newTim("2022-04-02T15:00:00Z")},
					{PR: 125, LO: 4.2, TS: newTim("2022-04-02T15:11:00Z")},
					{PR: 130, SH: 0.6, TS: newTim("2022-04-02T15:16:00Z")},
					{PR: 128, LO: 1.5, TS: newTim("2022-04-02T15:33:00Z")},
					{PR: 135, LO: 0.6, TS: newTim("2022-04-02T15:48:00Z")},
				},
			},
			re: []*Trades{
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T11:35:00Z"),
					EN: newTim("2022-04-02T12:00:00Z"),
					TR: []*Trade{
						{PR: 120, LO: 3.5, TS: newTim("2022-04-02T11:35:00Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T12:00:00Z"),
					EN: newTim("2022-04-02T13:00:00Z"),
					TR: []*Trade{
						{PR: 120, LO: 3.5, TS: newTim("2022-04-02T12:00:00Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:00:00Z"),
					EN: newTim("2022-04-02T14:00:00Z"),
					TR: []*Trade{
						{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
						{PR: 123, LO: 3.8, TS: newTim("2022-04-02T13:59:59Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T14:00:00Z"),
					EN: newTim("2022-04-02T15:00:00Z"),
					TR: []*Trade{
						{PR: 137, SH: 2.3, TS: newTim("2022-04-02T14:00:00Z")},
						{PR: 140, SH: 2.5, TS: newTim("2022-04-02T14:02:00Z")},
						{PR: 110, SH: 3.1, TS: newTim("2022-04-02T14:22:00Z")},
						{PR: 115, SH: 1.8, TS: newTim("2022-04-02T14:27:00Z")},
						{PR: 117, SH: 2.9, TS: newTim("2022-04-02T14:59:59Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T15:00:00Z"),
					EN: newTim("2022-04-02T16:00:00Z"),
					TR: []*Trade{
						{PR: 118, LO: 2.1, TS: newTim("2022-04-02T15:00:00Z")},
						{PR: 125, LO: 4.2, TS: newTim("2022-04-02T15:11:00Z")},
						{PR: 130, SH: 0.6, TS: newTim("2022-04-02T15:16:00Z")},
						{PR: 128, LO: 1.5, TS: newTim("2022-04-02T15:33:00Z")},
						{PR: 135, LO: 0.6, TS: newTim("2022-04-02T15:48:00Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T16:00:00Z"),
					EN: newTim("2022-04-02T17:00:00Z"),
					TR: []*Trade{
						{PR: 135, LO: 0.6, TS: newTim("2022-04-02T16:00:00Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T17:00:00Z"),
					EN: newTim("2022-04-02T17:24:00Z"),
					TR: []*Trade{
						{PR: 135, LO: 0.6, TS: newTim("2022-04-02T17:00:00Z")},
					},
				},
			},
		},
		// Case 2
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:00:00Z"),
				EN: newTim("2022-04-02T17:00:00Z"),
				TR: []*Trade{
					{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 123, LO: 3.8, TS: newTim("2022-04-02T13:59:59Z")},
					{PR: 118, LO: 2.1, TS: newTim("2022-04-02T16:00:00Z")},
					{PR: 125, LO: 4.2, TS: newTim("2022-04-02T16:11:00Z")},
					{PR: 130, SH: 0.6, TS: newTim("2022-04-02T16:16:00Z")},
					{PR: 128, LO: 1.5, TS: newTim("2022-04-02T16:33:00Z")},
					{PR: 135, LO: 0.6, TS: newTim("2022-04-02T16:48:00Z")},
				},
			},
			re: []*Trades{
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:00:00Z"),
					EN: newTim("2022-04-02T14:00:00Z"),
					TR: []*Trade{
						{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
						{PR: 123, LO: 3.8, TS: newTim("2022-04-02T13:59:59Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T14:00:00Z"),
					EN: newTim("2022-04-02T15:00:00Z"),
					TR: []*Trade{
						{PR: 123, LO: 3.8, TS: newTim("2022-04-02T14:00:00Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T15:00:00Z"),
					EN: newTim("2022-04-02T16:00:00Z"),
					TR: []*Trade{
						{PR: 123, LO: 3.8, TS: newTim("2022-04-02T15:00:00Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T16:00:00Z"),
					EN: newTim("2022-04-02T17:00:00Z"),
					TR: []*Trade{
						{PR: 118, LO: 2.1, TS: newTim("2022-04-02T16:00:00Z")},
						{PR: 125, LO: 4.2, TS: newTim("2022-04-02T16:11:00Z")},
						{PR: 130, SH: 0.6, TS: newTim("2022-04-02T16:16:00Z")},
						{PR: 128, LO: 1.5, TS: newTim("2022-04-02T16:33:00Z")},
						{PR: 135, LO: 0.6, TS: newTim("2022-04-02T16:48:00Z")},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f *Framer
			{
				f = tc.tr.Frame(framer.Config{
					Sta: tc.tr.ST.AsTime(),
					End: tc.tr.EN.AsTime(),
					Dur: time.Hour,
				})
			}

			var re []*Trades
			for !f.Last() {
				re = append(re, f.Next())
			}

			if len(re) != len(tc.re) {
				t.Fatalf("len\n\n%s\n", cmp.Diff(len(tc.re), len(re)))
			}

			for i, r := range re {
				if !reflect.DeepEqual(tc.re[i].EX, r.EX) {
					t.Fatalf("EX index %d\n\n%s\n", i, cmp.Diff(tc.re[i].EX, r.EX))
				}
				if !reflect.DeepEqual(tc.re[i].AS, r.AS) {
					t.Fatalf("AS index %d\n\n%s\n", i, cmp.Diff(tc.re[i].AS, r.AS))
				}
				if !reflect.DeepEqual(tc.re[i].ST, r.ST) {
					t.Fatalf("ST index %d\n\n%s\n", i, cmp.Diff(tc.re[i].ST.AsTime(), r.ST.AsTime()))
				}
				if !reflect.DeepEqual(tc.re[i].EN, r.EN) {
					t.Fatalf("EN index %d\n\n%s\n", i, cmp.Diff(tc.re[i].EN.AsTime(), r.EN.AsTime()))
				}

				if len(tc.re[i].TR) != len(r.TR) {
					t.Fatalf("len index %d\n\n%s\n", i, cmp.Diff(len(tc.re[i].TR), len(r.TR)))
				}

				for j := range r.TR {
					if !reflect.DeepEqual(tc.re[i].TR[j].LI, r.TR[j].LI) {
						t.Fatalf("LI index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.re[i].TR[j].LI, r.TR[j].LI))
					}
					if !reflect.DeepEqual(tc.re[i].TR[j].PR, r.TR[j].PR) {
						t.Fatalf("PR index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.re[i].TR[j].PR, r.TR[j].PR))
					}
					if !reflect.DeepEqual(tc.re[i].TR[j].LO, r.TR[j].LO) {
						t.Fatalf("LO index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.re[i].TR[j].LO, r.TR[j].LO))
					}
					if !reflect.DeepEqual(tc.re[i].TR[j].SH, r.TR[j].SH) {
						t.Fatalf("SH index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.re[i].TR[j].SH, r.TR[j].SH))
					}
					if !reflect.DeepEqual(tc.re[i].TR[j].TS.AsTime(), r.TR[j].TS.AsTime()) {
						t.Fatalf("TS index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.re[i].TR[j].TS.AsTime(), r.TR[j].TS.AsTime()))
					}
				}
			}
		})
	}
}

func Test_Typ_Trades_Frame_Hist(t *testing.T) {
	testCases := []struct {
		tr *Trades
		hi []*Trades
		pr int
	}{
		// Case 0
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-03T03:00:00Z"),
				EN: newTim("2022-04-03T04:25:00Z"),
				TR: []*Trade{
					{PR: 122, LO: 3.5, TS: newTim("2022-04-03T04:21:01.010Z")},
					{PR: 123, LO: 3.4, TS: newTim("2022-04-03T04:21:05.012Z")},
					{PR: 118, LO: 0.1, TS: newTim("2022-04-03T04:21:13.807Z")},
					{PR: 124, LO: 3.7, TS: newTim("2022-04-03T04:21:22.002Z")},
					{PR: 121, LO: 5.2, TS: newTim("2022-04-03T04:22:08.418Z")},
					{PR: 126, LO: 2.5, TS: newTim("2022-04-03T04:22:24.663Z")},
					{PR: 125, LO: 2.4, TS: newTim("2022-04-03T04:23:33.470Z")},
					{PR: 121, LO: 6.0, TS: newTim("2022-04-03T04:23:17.297Z")},
					{PR: 118, LO: 1.7, TS: newTim("2022-04-03T04:24:19.304Z")},
					{PR: 123, LO: 3.5, TS: newTim("2022-04-03T04:24:58.779Z")},
				},
			},
			hi: []*Trades{
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-03T04:20:00.000Z"),
					EN: newTim("2022-04-03T04:20:00.500Z"),
					TR: []*Trade{
						{PR: 122, LO: 3.5, TS: newTim("2022-04-03T04:20:00.000Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-03T04:21:00.000Z"),
					EN: newTim("2022-04-03T04:21:00.500Z"),
					TR: []*Trade{
						{PR: 122, LO: 3.5, TS: newTim("2022-04-03T04:21:00.000Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-03T04:22:00.000Z"),
					EN: newTim("2022-04-03T04:22:00.500Z"),
					TR: []*Trade{
						{PR: 124, LO: 3.7, TS: newTim("2022-04-03T04:22:00.000Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-03T04:23:00.000Z"),
					EN: newTim("2022-04-03T04:23:00.500Z"),
					TR: []*Trade{
						{PR: 126, LO: 2.5, TS: newTim("2022-04-03T04:23:00.000Z")},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-03T04:24:00.000Z"),
					EN: newTim("2022-04-03T04:24:00.500Z"),
					TR: []*Trade{
						{PR: 125, LO: 2.4, TS: newTim("2022-04-03T04:24:00.000Z")},
					},
				},
			},
			pr: +1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f *Framer
			{
				f = tc.tr.Frame(framer.Config{
					Sta: tc.tr.ST.AsTime(),
					End: tc.tr.EN.AsTime(),
					Dur: 500 * time.Millisecond,
				})
			}

			for !f.Last() {
				f.Next()
			}

			var hi []*Trades
			{
				hi = f.Hist()
			}

			if len(hi) != len(tc.hi) {
				t.Fatalf("len\n\n%s\n", cmp.Diff(len(tc.hi), len(hi)))
			}

			if tc.pr != f.Pred() {
				t.Fatalf("len\n\n%s\n", cmp.Diff(tc.pr, f.Pred()))
			}

			for i, r := range hi {
				if !reflect.DeepEqual(tc.hi[i].EX, r.EX) {
					t.Fatalf("EX index %d\n\n%s\n", i, cmp.Diff(tc.hi[i].EX, r.EX))
				}
				if !reflect.DeepEqual(tc.hi[i].AS, r.AS) {
					t.Fatalf("AS index %d\n\n%s\n", i, cmp.Diff(tc.hi[i].AS, r.AS))
				}
				if !reflect.DeepEqual(tc.hi[i].ST, r.ST) {
					t.Fatalf("ST index %d\n\n%s\n", i, cmp.Diff(tc.hi[i].ST.AsTime(), r.ST.AsTime()))
				}
				if !reflect.DeepEqual(tc.hi[i].EN, r.EN) {
					t.Fatalf("EN index %d\n\n%s\n", i, cmp.Diff(tc.hi[i].EN.AsTime(), r.EN.AsTime()))
				}

				if len(tc.hi[i].TR) != len(r.TR) {
					t.Fatalf("len index %d\n\n%s\n", i, cmp.Diff(len(tc.hi[i].TR), len(r.TR)))
				}

				for j := range r.TR {
					if !reflect.DeepEqual(tc.hi[i].TR[j].LI, r.TR[j].LI) {
						t.Fatalf("LI index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.hi[i].TR[j].LI, r.TR[j].LI))
					}
					if !reflect.DeepEqual(tc.hi[i].TR[j].PR, r.TR[j].PR) {
						t.Fatalf("PR index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.hi[i].TR[j].PR, r.TR[j].PR))
					}
					if !reflect.DeepEqual(tc.hi[i].TR[j].LO, r.TR[j].LO) {
						t.Fatalf("LO index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.hi[i].TR[j].LO, r.TR[j].LO))
					}
					if !reflect.DeepEqual(tc.hi[i].TR[j].SH, r.TR[j].SH) {
						t.Fatalf("SH index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.hi[i].TR[j].SH, r.TR[j].SH))
					}
					if !reflect.DeepEqual(tc.hi[i].TR[j].TS.AsTime(), r.TR[j].TS.AsTime()) {
						t.Fatalf("TS index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.hi[i].TR[j].TS.AsTime(), r.TR[j].TS.AsTime()))
					}
				}
			}
		})
	}
}

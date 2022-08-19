package trades

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/xh3b4sd/framer"
)

func Test_Typ_Trades_Frame(t *testing.T) {
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
				EN: newTim("2022-04-02T15:48:00Z"),
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
					EN: newTim("2022-04-02T15:48:00Z"),
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
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-07-04T18:24:00Z"), // end beyond trades
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
					EN: newTim("2022-04-02T16:00:00Z"),
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
		// Case 2
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-01T00:00:00Z"), // start of the month
				EN: newTim("2022-07-04T18:24:00Z"),
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

			if len(tc.tr.TR) != 0 {
				t.Fatalf("len\n\n%s\n", cmp.Diff(0, len(tc.tr.TR)))
			}

			if len(re) != len(tc.re) {
				t.Fatalf("len\n\n%s\n", cmp.Diff(len(re), len(tc.re)))
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
				for j := range r.TR {
					if !reflect.DeepEqual(tc.re[i].TR[j], r.TR[j]) {
						t.Fatalf("TR index %d/%d\n\n%s\n", i, j, fmt.Sprintf("%#v \n\n %#v", tc.re[i].TR[j], r.TR[j]))
					}
				}
			}
		})
	}
}

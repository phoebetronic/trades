package trades

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Typ_Trades_Merge(t *testing.T) {
	testCases := []struct {
		tr []*Trades
		re *Trades
	}{
		// Case 0
		{
			tr: []*Trades{
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
			re: &Trades{
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
		},
		// Case 1
		{
			tr: []*Trades{
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
			re: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-04-02T16:00:00Z"),
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
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var tr *Trades
			{
				tr = &Trades{}
			}

			{
				tr = tr.Merge(tc.tr)
			}

			if !reflect.DeepEqual(tc.re.EX, tr.EX) {
				t.Fatalf("EX\n\n%s\n", cmp.Diff(tc.re.EX, tr.EX))
			}
			if !reflect.DeepEqual(tc.re.AS, tr.AS) {
				t.Fatalf("AS\n\n%s\n", cmp.Diff(tc.re.AS, tr.AS))
			}
			if !reflect.DeepEqual(tc.re.ST, tr.ST) {
				t.Fatalf("ST\n\n%s\n", cmp.Diff(tc.re.ST.AsTime(), tr.ST.AsTime()))
			}
			if !reflect.DeepEqual(tc.re.EN, tr.EN) {
				t.Fatalf("EN\n\n%s\n", cmp.Diff(tc.re.EN.AsTime(), tr.EN.AsTime()))
			}
			for j := range tr.TR {
				if !reflect.DeepEqual(tc.re.TR[j], tr.TR[j]) {
					t.Fatalf("TR index %d\n\n%s\n", j, fmt.Sprintf("%#v \n\n %#v", tc.re.TR[j], tr.TR[j]))
				}
			}
		})
	}
}

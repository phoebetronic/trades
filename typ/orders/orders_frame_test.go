package orders

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_Typ_Orders_Frame_Next(t *testing.T) {
	testCases := []struct {
		or *Orders
		re []*Orders
	}{
		// Case 0
		{
			or: &Orders{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:53:00Z"),
				EN: newTim("2022-04-02T13:58:00Z"),
				BU: []*Bundle{
					{
						AK: []*Order{{PR: 120, SI: 3.5}, {PR: 123, SI: 3.8}},
						BD: []*Order{{PR: 121, SI: 1.3}, {PR: 122, SI: 2.9}},
						TS: newTim("2022-04-02T13:53:05Z"),
					},
					{
						AK: []*Order{{PR: 120, SI: 3.5}, {PR: 123, SI: 3.8}},
						BD: []*Order{{PR: 121, SI: 1.3}, {PR: 122, SI: 2.9}},
						TS: newTim("2022-04-02T13:53:12Z"),
					},
					{
						AK: []*Order{{PR: 124, SI: 1.1}, {PR: 126, SI: 5.5}},
						BD: []*Order{{PR: 123, SI: 4.0}, {PR: 125, SI: 4.7}},
						TS: newTim("2022-04-02T13:53:17Z"),
					},
					{
						AK: []*Order{{PR: 125, SI: 7.1}, {PR: 118, SI: 6.7}},
						BD: []*Order{{PR: 122, SI: 2.2}, {PR: 126, SI: 4.5}},
						TS: newTim("2022-04-02T13:54:07Z"),
					},
					{
						AK: []*Order{{PR: 122, SI: 6.2}, {PR: 127, SI: 0.9}},
						BD: []*Order{{PR: 124, SI: 2.5}, {PR: 125, SI: 0.5}},
						TS: newTim("2022-04-02T13:55:28Z"),
					},
					{
						AK: []*Order{{PR: 113, SI: 8.7}, {PR: 123, SI: 6.6}},
						BD: []*Order{{PR: 123, SI: 4.4}, {PR: 122, SI: 5.7}},
						TS: newTim("2022-04-02T13:57:32Z"),
					},
				},
			},
			re: []*Orders{
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:53:00Z"),
					EN: newTim("2022-04-02T13:54:00Z"),
					BU: []*Bundle{
						{
							AK: []*Order{{PR: 120, SI: 3.5}, {PR: 123, SI: 3.8}},
							BD: []*Order{{PR: 121, SI: 1.3}, {PR: 122, SI: 2.9}},
							TS: newTim("2022-04-02T13:53:05Z"),
						},
						{
							AK: []*Order{{PR: 120, SI: 3.5}, {PR: 123, SI: 3.8}},
							BD: []*Order{{PR: 121, SI: 1.3}, {PR: 122, SI: 2.9}},
							TS: newTim("2022-04-02T13:53:12Z"),
						},
						{
							AK: []*Order{{PR: 124, SI: 1.1}, {PR: 126, SI: 5.5}},
							BD: []*Order{{PR: 123, SI: 4.0}, {PR: 125, SI: 4.7}},
							TS: newTim("2022-04-02T13:53:17Z"),
						},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:54:00Z"),
					EN: newTim("2022-04-02T13:55:00Z"),
					BU: []*Bundle{
						{
							AK: []*Order{{PR: 125, SI: 7.1}, {PR: 118, SI: 6.7}},
							BD: []*Order{{PR: 122, SI: 2.2}, {PR: 126, SI: 4.5}},
							TS: newTim("2022-04-02T13:54:07Z"),
						},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:55:00Z"),
					EN: newTim("2022-04-02T13:56:00Z"),
					BU: []*Bundle{
						{
							AK: []*Order{{PR: 122, SI: 6.2}, {PR: 127, SI: 0.9}},
							BD: []*Order{{PR: 124, SI: 2.5}, {PR: 125, SI: 0.5}},
							TS: newTim("2022-04-02T13:55:28Z"),
						},
					},
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:56:00Z"),
					EN: newTim("2022-04-02T13:57:00Z"),
					BU: nil,
				},
				{
					EX: "ftx",
					AS: "eth",
					ST: newTim("2022-04-02T13:57:00Z"),
					EN: newTim("2022-04-02T13:58:00Z"),
					BU: []*Bundle{
						{
							AK: []*Order{{PR: 113, SI: 8.7}, {PR: 123, SI: 6.6}},
							BD: []*Order{{PR: 123, SI: 4.4}, {PR: 122, SI: 5.7}},
							TS: newTim("2022-04-02T13:57:32Z"),
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f *Framer
			{
				f = tc.or.Frame(framer.Config{
					Sta: tc.or.ST.AsTime(),
					End: tc.or.EN.AsTime(),
					Dur: time.Minute,
				})
			}

			var re []*Orders
			for !f.Last() {
				re = append(re, f.Next())
			}

			if len(tc.or.BU) != 0 {
				t.Fatalf("len\n\n%s\n", cmp.Diff(len(tc.or.BU), 0))
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

				if len(tc.re[i].BU) != len(r.BU) {
					t.Fatalf("len index %d\n\n%s\n", i, cmp.Diff(len(tc.re[i].BU), len(r.BU)))
				}

				for j := range r.BU {
					if len(tc.re[i].BU[j].AK) != len(r.BU[j].AK) {
						t.Fatalf("len index %d\n\n%s\n", i, cmp.Diff(len(tc.re[i].BU[j].AK), len(r.BU[j].AK)))
					}
					if len(tc.re[i].BU[j].BD) != len(r.BU[j].BD) {
						t.Fatalf("len index %d\n\n%s\n", i, cmp.Diff(len(tc.re[i].BU[j].BD), len(r.BU[j].BD)))
					}
					if !reflect.DeepEqual(tc.re[i].BU[j].TS.AsTime(), r.BU[j].TS.AsTime()) {
						t.Fatalf("TS index %d/%d\n\n%s\n", i, j, cmp.Diff(tc.re[i].BU[j].TS.AsTime(), r.BU[j].TS.AsTime()))
					}

					for u := range r.BU[j].AK {
						if !reflect.DeepEqual(tc.re[i].BU[j].AK[u].PR, r.BU[j].AK[u].PR) {
							t.Fatalf("PR index %d/%d/%d\n\n%s\n", i, j, u, cmp.Diff(tc.re[i].BU[j].AK[u].PR, r.BU[j].AK[u].PR))
						}
						if !reflect.DeepEqual(tc.re[i].BU[j].AK[u].SI, r.BU[j].AK[u].SI) {
							t.Fatalf("SI index %d/%d/%d\n\n%s\n", i, j, u, cmp.Diff(tc.re[i].BU[j].AK[u].SI, r.BU[j].AK[u].SI))
						}
					}
					for u := range r.BU[j].BD {
						if !reflect.DeepEqual(tc.re[i].BU[j].BD[u].PR, r.BU[j].BD[u].PR) {
							t.Fatalf("PR index %d/%d/%d\n\n%s\n", i, j, u, cmp.Diff(tc.re[i].BU[j].BD[u].PR, r.BU[j].BD[u].PR))
						}
						if !reflect.DeepEqual(tc.re[i].BU[j].BD[u].SI, r.BU[j].BD[u].SI) {
							t.Fatalf("SI index %d/%d/%d\n\n%s\n", i, j, u, cmp.Diff(tc.re[i].BU[j].BD[u].SI, r.BU[j].BD[u].SI))
						}
					}
				}
			}
		})
	}
}

func newTim(str string) *timestamppb.Timestamp {
	var err error

	var tim time.Time
	{
		tim, err = time.Parse(time.RFC3339, str)
		if err != nil {
			panic(err)
		}
	}

	return timestamppb.New(tim)
}

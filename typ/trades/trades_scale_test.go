package trades

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_Typ_Trades_Scale(t *testing.T) {
	testCases := []struct {
		tr *Trades
		sc float32
		re *Trades
	}{
		// Case 0
		{
			tr: &Trades{},
			sc: 0,
			re: &Trades{},
		},
		// Case 1
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-04-05T13:55:00Z"),
				TR: []*Trade{
					{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 140, SH: 2.5, TS: newTim("2022-04-02T14:02:00Z")},
					{PR: 110, SH: 3.1, TS: newTim("2022-04-02T14:22:00Z")},
					{PR: 115, SH: 1.8, TS: newTim("2022-04-02T14:27:00Z")},
					{PR: 125, LO: 4.2, TS: newTim("2022-04-02T15:11:00Z")},
					{PR: 130, SH: 0.6, TS: newTim("2022-04-02T15:16:00Z")},
					{PR: 128, LO: 1.5, TS: newTim("2022-04-02T15:33:00Z")},
					{PR: 135, LO: 0.6, TS: newTim("2022-04-02T15:48:00Z")},
				},
			},
			sc: +0.1,
			re: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-04-05T13:55:00Z"),
				TR: []*Trade{
					{PR: 132.0, LO: 3.8500000, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 154.0, SH: 2.7500000, TS: newTim("2022-04-02T14:02:00Z")},
					{PR: 121.0, SH: 3.4099998, TS: newTim("2022-04-02T14:22:00Z")},
					{PR: 126.5, SH: 1.9799999, TS: newTim("2022-04-02T14:27:00Z")},
					{PR: 137.5, LO: 4.6200000, TS: newTim("2022-04-02T15:11:00Z")},
					{PR: 143.0, SH: 0.6600000, TS: newTim("2022-04-02T15:16:00Z")},
					{PR: 140.8, LO: 1.6500000, TS: newTim("2022-04-02T15:33:00Z")},
					{PR: 148.5, LO: 0.6600000, TS: newTim("2022-04-02T15:48:00Z")},
				},
			},
		},
		// Case 2
		{
			tr: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-04-05T13:55:00Z"),
				TR: []*Trade{
					{PR: 120, LO: 3.5, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 140, SH: 2.5, TS: newTim("2022-04-02T14:02:00Z")},
					{PR: 110, SH: 3.1, TS: newTim("2022-04-02T14:22:00Z")},
					{PR: 115, SH: 1.8, TS: newTim("2022-04-02T14:27:00Z")},
					{PR: 125, LO: 4.2, TS: newTim("2022-04-02T15:11:00Z")},
					{PR: 130, SH: 0.6, TS: newTim("2022-04-02T15:16:00Z")},
					{PR: 128, LO: 1.5, TS: newTim("2022-04-02T15:33:00Z")},
					{PR: 135, LO: 0.6, TS: newTim("2022-04-02T15:48:00Z")},
				},
			},
			sc: -0.15,
			re: &Trades{
				EX: "ftx",
				AS: "eth",
				ST: newTim("2022-04-02T13:55:00Z"),
				EN: newTim("2022-04-05T13:55:00Z"),
				TR: []*Trade{
					{PR: 102.00, LO: 2.9750000, TS: newTim("2022-04-02T13:55:00Z")},
					{PR: 119.00, SH: 2.1250000, TS: newTim("2022-04-02T14:02:00Z")},
					{PR: 93.500, SH: 2.6350000, TS: newTim("2022-04-02T14:22:00Z")},
					{PR: 97.750, SH: 1.5300000, TS: newTim("2022-04-02T14:27:00Z")},
					{PR: 106.25, LO: 3.5699997, TS: newTim("2022-04-02T15:11:00Z")},
					{PR: 110.50, SH: 0.5100000, TS: newTim("2022-04-02T15:16:00Z")},
					{PR: 108.80, LO: 1.2750000, TS: newTim("2022-04-02T15:33:00Z")},
					{PR: 114.75, LO: 0.5100000, TS: newTim("2022-04-02T15:48:00Z")},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var re *Trades
			{
				re = tc.tr.Scale(tc.sc)
			}

			if !reflect.DeepEqual(tc.re.EX, re.EX) {
				t.Fatalf("EX\n\n%s\n", cmp.Diff(tc.re.EX, re.EX))
			}
			if !reflect.DeepEqual(tc.re.AS, re.AS) {
				t.Fatalf("AS\n\n%s\n", cmp.Diff(tc.re.AS, re.AS))
			}
			if !reflect.DeepEqual(tc.re.ST, re.ST) {
				t.Fatalf("ST\n\n%s\n", cmp.Diff(tc.re.ST, re.ST))
			}
			if !reflect.DeepEqual(tc.re.EN, re.EN) {
				t.Fatalf("EN\n\n%s\n", cmp.Diff(tc.re.EN, re.EN))
			}
			for i := range tc.re.TR {
				if !reflect.DeepEqual(tc.re.TR[i], re.TR[i]) {
					t.Fatalf("TR index %d\n\n%s\n", i, fmt.Sprintf("%#v \n\n %#v", tc.re.TR[i], re.TR[i]))
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

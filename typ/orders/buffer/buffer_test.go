package buffer

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/trades/fix"
	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
)

func Test_Typ_Orders_Buffer_Finish(t *testing.T) {
	testCases := []struct {
		dur time.Duration
		set func(c Buffer)
		tim []time.Time
		cou int
		ord int
	}{
		// case 0
		{
			dur: 15 * time.Second,
			set: func(c Buffer) {
				finish(c, "2022-05-01T23:59:59.500Z")
				finish(c, "2022-05-02T00:00:00.500Z")
				finish(c, "2022-05-02T00:00:02.000Z")
				buffer(c, "2022-05-02T00:00:00.000Z")
				finish(c, "2022-05-02T00:00:02.300Z")
				finish(c, "2022-05-02T00:00:04.500Z")
				buffer(c, "2022-05-02T00:00:05.000Z")
				finish(c, "2022-05-02T00:00:07.300Z")
				finish(c, "2022-05-02T00:00:08.500Z")
				buffer(c, "2022-05-02T00:00:10.000Z")
				finish(c, "2022-05-02T00:00:12.300Z")
				finish(c, "2022-05-02T00:00:14.500Z")
				buffer(c, "2022-05-02T00:00:15.000Z")
				finish(c, "2022-05-02T00:00:15.500Z")
			},
			tim: []time.Time{
				newTim("2022-05-02T00:00:00.000Z"),
				newTim("2022-05-02T00:00:05.000Z"),
				newTim("2022-05-02T00:00:10.000Z"),
			},
			cou: 1,
			ord: 14,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var buf Buffer
			{
				buf = New(Config{
					Mar: newMar(tc.dur),
				})
			}

			var cou int
			var abc int
			var wei sync.WaitGroup
			var ord []*orders.Orders

			wei.Add(1)
			go func() {
				{
					defer wei.Done()
				}

				for c := range buf.Orders() {
					{
						cou += len(c.BU)
					}

					{
						ord = append(ord, c)
					}

					for _, x := range c.BU {
						abc += len(x.AK)
						abc += len(x.BD)
					}
				}
			}()

			{
				tc.set(buf)
				close(buf.Orders())
				wei.Wait()
			}

			var exp int
			{
				exp = len(tc.tim)
			}

			if cou != exp {
				t.Fatalf("cou\n\n%s\n", cmp.Diff(exp, cou))
			}
			if abc != tc.ord {
				t.Fatalf("cou\n\n%s\n", cmp.Diff(tc.ord, abc))
			}

			if len(ord) != tc.cou {
				t.Fatalf("len\n\n%s\n", cmp.Diff(tc.cou, len(ord)))
			}

			var bu int
			{
				bu = buf.Metric()
			}

			if bu > 2 {
				t.Fatalf("buf\n\n%s\n", cmp.Diff(2, bu))
			}

			for i := range ord {
				{
					if ord[i].EX != "ftx" {
						t.Fatalf("EX\n\n%s\n", cmp.Diff("ftx", ord[i].EX))
					}
					if ord[i].AS != "eth" {
						t.Fatalf("AS\n\n%s\n", cmp.Diff("eth", ord[i].AS))
					}
					if !ord[i].ST.AsTime().Equal(tc.tim[i]) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff(tc.tim[i].String(), ord[i].ST.AsTime().String()))
					}
					if !ord[i].EN.AsTime().Equal(tc.tim[i].Add(tc.dur)) {
						t.Fatalf("EN\n\n%s\n", cmp.Diff(tc.tim[i].String(), ord[i].EN.AsTime().String()))
					}
					if len(ord[i].BU) != len(tc.tim) {
						t.Fatal("amount of bundles must match")
					}
				}

				for j := range ord[i].BU {
					if ord[i].BU[j].TS.AsTime().Before(ord[i].ST.AsTime()) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff(ord[i].BU[j].TS.AsTime().String(), ord[i].ST.AsTime().String()))
					}
					if ord[i].BU[j].TS.AsTime().Equal(ord[i].EN.AsTime()) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff("", ord[i].EN.AsTime().String()))
					}
					if ord[i].BU[j].TS.AsTime().After(ord[i].EN.AsTime()) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff(ord[i].BU[j].TS.AsTime().String(), ord[i].EN.AsTime().String()))
					}
				}
			}
		})
	}
}

func Test_Orders_Buffer_Empty_Two(t *testing.T) {
	var buf Buffer
	{
		buf = New(Config{
			Mar: newMar(5 * time.Second),
		})
	}

	{
		finish(buf, "1994-12-23T18:42:22Z")
		finish(buf, "1994-12-23T18:42:23Z")
	}

	{
		select {
		case <-buf.Orders():
			t.Fatal("expected no trades results at all")
		default:
		}
	}
}

func Test_Orders_Buffer_Empty_Thr(t *testing.T) {
	var buf Buffer
	{
		buf = New(Config{
			Mar: newMar(5 * time.Second),
		})
	}

	{
		finish(buf, "1994-12-23T18:42:22Z")
		finish(buf, "1994-12-23T18:42:23Z")
		finish(buf, "1994-12-23T18:42:24Z")
		finish(buf, "1994-12-23T18:42:24Z")
		finish(buf, "1994-12-23T18:42:24Z")
		finish(buf, "1994-12-23T18:42:26Z")
	}

	{
		select {
		case <-buf.Orders():
			t.Fatal("expected no trades results at all")
		default:
		}
	}
}

func buffer(buf Buffer, str string) {
	var bun *orders.Bundle
	{
		bun = newBun(newTim(str))
	}

	{
		buf.Buffer(bun)
	}
}

func finish(buf Buffer, str string) {
	var err error

	var tim time.Time
	{
		tim, err = time.Parse(time.RFC3339, str)
		if err != nil {
			panic(err)
		}
	}

	{
		buf.Finish(tim.UTC())
	}
}

func newMar(dur time.Duration) market.Market {
	return market.New(market.Config{
		Exc: "ftx",
		Ass: "eth",
		Dur: dur,
	})
}

func newTim(str string) time.Time {
	var err error

	var tim time.Time
	{
		tim, err = time.Parse(time.RFC3339, str)
		if err != nil {
			panic(err)
		}
	}

	return tim
}

func newBun(tim time.Time) *orders.Bundle {
	return fix.Ordmap()[tim]
}

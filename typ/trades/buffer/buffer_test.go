package buffer

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetronic/trades/fix"
	"github.com/phoebetronic/trades/typ/market"
	"github.com/phoebetronic/trades/typ/trades"
)

func Test_Typ_Trades_Buffer_Finish(t *testing.T) {
	testCases := []struct {
		dur time.Duration
		set func(c Buffer)
		tim []time.Time
	}{
		// case 0
		{
			dur: 3 * time.Second,
			set: func(c Buffer) {
				finish(c, "2022-05-01T23:59:59.500Z")
				finish(c, "2022-05-02T00:00:00.600Z")
				finish(c, "2022-05-02T00:00:01.000Z")
				buffer(c, "2022-05-02T00:00:00.000Z")
				finish(c, "2022-05-02T00:00:01.300Z")
				finish(c, "2022-05-02T00:00:03.500Z")
				buffer(c, "2022-05-02T00:00:03.000Z")
			},
			tim: []time.Time{
				newTim("2022-05-02T00:00:00.000Z"),
			},
		},
		// case 1
		{
			dur: 3 * time.Second,
			set: func(c Buffer) {
				finish(c, "2022-05-01T23:59:59.500Z")
				finish(c, "2022-05-02T00:00:00.500Z")
				finish(c, "2022-05-02T00:00:00.600Z")
				finish(c, "2022-05-02T00:00:01.000Z")
				buffer(c, "2022-05-02T00:00:00.000Z")
				finish(c, "2022-05-02T00:00:01.300Z")
				finish(c, "2022-05-02T00:00:03.500Z")
				buffer(c, "2022-05-02T00:00:03.000Z")
				finish(c, "2022-05-02T00:00:05.600Z")
			},
			tim: []time.Time{
				newTim("2022-05-02T00:00:00.000Z"),
			},
		},
		// case 2
		{
			dur: 3 * time.Second,
			set: func(c Buffer) {
				finish(c, "2022-05-01T23:59:59.500Z")
				finish(c, "2022-05-02T00:00:00.500Z")
				finish(c, "2022-05-02T00:00:00.600Z")
				finish(c, "2022-05-02T00:00:01.000Z")
				buffer(c, "2022-05-02T00:00:00.000Z")
				finish(c, "2022-05-02T00:00:01.300Z")
				finish(c, "2022-05-02T00:00:03.500Z")
				buffer(c, "2022-05-02T00:00:03.000Z")
				finish(c, "2022-05-02T00:00:05.600Z")
				finish(c, "2022-05-02T00:00:06.300Z")
			},
			tim: []time.Time{
				newTim("2022-05-02T00:00:00.000Z"),
				newTim("2022-05-02T00:00:03.000Z"),
			},
		},
		// case 3
		{
			dur: 3 * time.Second,
			set: func(c Buffer) {
				finish(c, "2022-05-02T00:00:16.500Z")
				finish(c, "2022-05-02T00:00:16.600Z")
				finish(c, "2022-05-02T00:00:17.000Z")
				buffer(c, "2022-05-02T00:00:18.000Z")
				finish(c, "2022-05-02T00:00:18.500Z")
				finish(c, "2022-05-02T00:00:20.500Z")
				buffer(c, "2022-05-02T00:00:21.000Z")
				finish(c, "2022-05-02T00:00:21.600Z")
				finish(c, "2022-05-02T00:00:22.300Z")
				finish(c, "2022-05-02T00:00:22.300Z")
				finish(c, "2022-05-02T00:00:23.500Z")
				buffer(c, "2022-05-02T00:00:24.000Z")
				finish(c, "2022-05-02T00:00:24.100Z")
				finish(c, "2022-05-02T00:00:27.100Z")
			},
			tim: []time.Time{
				newTim("2022-05-02T00:00:18.000Z"),
				newTim("2022-05-02T00:00:21.000Z"),
				newTim("2022-05-02T00:00:24.000Z"),
			},
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
			var wei sync.WaitGroup
			var tra []*trades.Trades

			wei.Add(1)
			go func() {
				{
					defer wei.Done()
				}

				for c := range buf.Trades() {
					{
						cou += len(c.TR)
					}

					{
						tra = append(tra, c)
					}
				}
			}()

			{
				tc.set(buf)
				close(buf.Trades())
				wei.Wait()
			}

			var exp int
			for _, t := range tc.tim {
				l := len(newTra(t))
				if l != 0 {
					exp += l
				} else {
					exp++ // account for the filled trade of buffer 2022-05-02T00:00:24Z
				}
			}

			if cou != exp {
				t.Fatalf("cou\n\n%s\n", cmp.Diff(exp, cou))
			}

			if len(tra) != len(tc.tim) {
				t.Fatalf("len\n\n%s\n", cmp.Diff(len(tc.tim), len(tra)))
			}

			var bu int
			{
				bu = buf.Metric()
			}

			if bu > 2 {
				t.Fatalf("buf\n\n%s\n", cmp.Diff(2, bu))
			}

			for i := range tra {
				{
					if tra[i].EX != "ftx" {
						t.Fatalf("EX\n\n%s\n", cmp.Diff("ftx", tra[i].EX))
					}
					if tra[i].AS != "eth" {
						t.Fatalf("AS\n\n%s\n", cmp.Diff("eth", tra[i].AS))
					}
					if !tra[i].ST.AsTime().Equal(tc.tim[i]) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff(tc.tim[i].String(), tra[i].ST.AsTime().String()))
					}
					if !tra[i].EN.AsTime().Equal(tc.tim[i].Add(tc.dur)) {
						t.Fatalf("EN\n\n%s\n", cmp.Diff(tc.tim[i].String(), tra[i].EN.AsTime().String()))
					}
					if len(tra[i].TR) == 0 {
						t.Fatal("trades must not be empty")
					}
				}

				// Case 3 represents a special use case where the trades of our
				// fixtures do not define any trades for the buffer between
				// seconds 24 and 27. The buffer implementation is designed to
				// fill missing trades by carrying over the last known trade in
				// order to always provide the last known price and volume
				// information. The empty buffer starts at 2022-05-02T00:00:24Z
				// and the last trade before that buffer is defined as follows.
				//
				//     {
				//       "PR": 38487,
				//       "LO": 0.1838,
				//       "TS": "2022-05-02T00:00:23.579116Z"
				//     }
				//
				if len(tra[i].TR) == 1 {
					if tra[i].TR[0].PR != 38487 {
						t.Fatalf("PR\n\n%s\n", cmp.Diff(38487, tra[i].PR))
					}
					if tra[i].TR[0].LO != 0.1838 {
						t.Fatalf("LO\n\n%s\n", cmp.Diff(0.1838, tra[i].LO))
					}
					if tra[i].TR[0].SH != 0 {
						t.Fatalf("SH\n\n%s\n", cmp.Diff(0, tra[i].SH))
					}
				}

				for j := range tra[i].TR {
					if tra[i].TR[j].TS.AsTime().Before(tra[i].ST.AsTime()) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff(tra[i].TR[j].TS.AsTime().String(), tra[i].ST.AsTime().String()))
					}
					if tra[i].TR[j].TS.AsTime().Equal(tra[i].EN.AsTime()) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff("", tra[i].EN.AsTime().String()))
					}
					if tra[i].TR[j].TS.AsTime().After(tra[i].EN.AsTime()) {
						t.Fatalf("ST\n\n%s\n", cmp.Diff(tra[i].TR[j].TS.AsTime().String(), tra[i].EN.AsTime().String()))
					}
				}
			}
		})
	}
}

func Test_Buffer_Empty_One(t *testing.T) {
	var buf Buffer
	{
		buf = New(Config{
			Mar: newMar(3 * time.Second),
		})
	}

	{
		finish(buf, "1994-12-23T18:42:22Z") // first observed, already started
		finish(buf, "1995-06-11T04:53:13Z") // fully observed from beginning
		finish(buf, "1996-02-01T13:13:44Z") // wrapping up first complete bucket
	}

	{
		select {
		case tra := <-buf.Trades():
			if len(tra.TR) != 1 {
				t.Fatal("expected one empty trade")
			}
		default:
			t.Fatal("expected empty trades result")
		}
	}

	{
		select {
		case <-buf.Trades():
			t.Fatal("expected no more trades results")
		default:
		}
	}
}

func Test_Buffer_Empty_Two(t *testing.T) {
	var buf Buffer
	{
		buf = New(Config{
			Mar: newMar(3 * time.Second),
		})
	}

	{
		finish(buf, "1994-12-23T18:42:22Z")
		finish(buf, "1994-12-23T18:42:23Z")
	}

	{
		select {
		case <-buf.Trades():
			t.Fatal("expected no trades results at all")
		default:
		}
	}
}

func Test_Buffer_Empty_Thr(t *testing.T) {
	var buf Buffer
	{
		buf = New(Config{
			Mar: newMar(3 * time.Second),
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
		case <-buf.Trades():
			t.Fatal("expected no trades results at all")
		default:
		}
	}
}

func buffer(buf Buffer, str string, ind ...int) {
	var tra []*trades.Trade
	{
		tra = newTra(newTim(str))
	}

	{
		if len(ind) == 0 {
			for _, t := range tra {
				buf.Buffer(t)
			}
		}
	}

	{
		if len(ind) != 0 {
			for _, i := range ind {
				buf.Buffer(tra[i])
			}
		}
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

func newTra(tim time.Time) []*trades.Trade {
	var tra []*trades.Trade
	for _, v := range fix.Tramap()[tim] {
		tra = append(tra, &trades.Trade{
			LI: v.LI,
			PR: v.PR,
			LO: v.LO,
			SH: v.SH,
			TS: v.TS,
		})
	}

	return tra
}

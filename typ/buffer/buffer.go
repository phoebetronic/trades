package buffer

import (
	"sync"
	"time"

	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/trades"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Buffer struct {
	buf map[time.Time][]*trades.Trade
	cha chan *trades.Trades
	ini bool
	mar *market.Market
	mut sync.Mutex
	tim time.Time
	tra *trades.Trade
}

func New(con Config) *Buffer {
	{
		con.Verify()
	}

	return &Buffer{
		buf: map[time.Time][]*trades.Trade{},
		cha: make(chan *trades.Trades, 1),
		mar: con.Mar,
		tra: &trades.Trade{},
	}
}

func (b *Buffer) Buffer(tra *trades.Trade) {
	var tim time.Time
	{
		tim = tra.TS.AsTime().Truncate(b.mar.Dur())
	}

	{
		b.mut.Lock()
		b.buf[tim] = append(b.buf[tim], tra)
		b.mut.Unlock()
	}
}

func (b *Buffer) Finish(tim time.Time) {
	{
		b.mut.Lock()
		defer b.mut.Unlock()
	}

	var buc time.Time
	{
		buc = tim.Truncate(b.mar.Dur())
	}

	{
		// The first time Collector.Finish is called we are within the first
		// time bucked that has already started. For this very bucket we
		// initialize our internal bucket time.
		if b.tim.IsZero() {
			b.tim = buc
		}
	}

	{
		// Whenever the currently perceived bucket time is equal to the
		// currently tracked bucket time, we simply return because our job of
		// completing the current bucket is not yet done.
		if b.tim.Equal(buc) {
			return
		}
	}

	var tra []*trades.Trade
	{
		tra = b.buf[b.tim]
	}

	if len(tra) == 0 {
		tra = append(tra, &trades.Trade{
			LI: b.tra.LI,
			PR: b.tra.PR,
			LO: b.tra.LO,
			SH: b.tra.SH,
			TS: timestamppb.New(b.tim),
		})
	} else {
		b.tra = tra[len(tra)-1]
	}

	{
		if b.ini {
			b.cha <- &trades.Trades{
				EX: b.mar.Exc(),
				AS: b.mar.Ass(),
				ST: timestamppb.New(b.tim),
				EN: timestamppb.New(b.tim.Add(b.mar.Dur())),
				TR: tra,
			}
		}
	}

	{
		// The very first time the currently perceived bucket time and the
		// currently tracked bucket time is not equal anymore, we collected
		// trades from a time bucket that we did not observe from the start.
		// Therefore we MUST not emit a candle, but only cleanup, track the new
		// bucket time and continue collecting.
		{
			b.ini = true
		}
	}

	{
		delete(b.buf, b.tim)
		b.tim = buc
	}
}

func (b *Buffer) Metric() int {
	return len(b.buf)
}

func (b *Buffer) Trades() chan *trades.Trades {
	return b.cha
}

package buffer

import (
	"sync"
	"time"

	"github.com/phoebetron/trades/typ/market"
	"github.com/phoebetron/trades/typ/orders"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type B struct {
	buf map[time.Time][]*orders.Bundle
	cha chan *orders.Orders
	ini bool
	mar market.Market
	mut sync.Mutex
	tim time.Time
}

func New(con Config) *B {
	{
		con.Verify()
	}

	return &B{
		buf: map[time.Time][]*orders.Bundle{},
		cha: make(chan *orders.Orders, 1),
		mar: con.Mar,
	}
}

func (b *B) Buffer(bun *orders.Bundle) {
	var tim time.Time
	{
		tim = bun.TS.AsTime().Truncate(b.mar.Dur())
	}

	{
		b.mut.Lock()
		b.buf[tim] = append(b.buf[tim], bun)
		b.mut.Unlock()
	}
}

func (b *B) Finish(tim time.Time) {
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

	var bun []*orders.Bundle
	{
		bun = b.buf[b.tim]
	}

	{
		if b.ini {
			b.cha <- &orders.Orders{
				EX: b.mar.Exc(),
				AS: b.mar.Ass(),
				ST: timestamppb.New(b.tim),
				EN: timestamppb.New(b.tim.Add(b.mar.Dur())),
				BU: bun,
			}
		}
	}

	{
		// The very first time the currently perceived bucket time and the
		// currently tracked bucket time is not equal anymore, we collected
		// orders from a time bucket that we did not observe from the start.
		// Therefore we MUST not emit a send, but only cleanup, track the new
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

func (b *B) Metric() int {
	return len(b.buf)
}

func (b *B) Orders() chan *orders.Orders {
	return b.cha
}

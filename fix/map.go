package fix

import (
	"encoding/json"
	"time"

	"github.com/phoebetronic/trades/typ/orders"
	"github.com/phoebetronic/trades/typ/trades"
)

func Ordmap() map[time.Time]*orders.Bundle {
	var m map[time.Time]*orders.Bundle
	{
		err := json.Unmarshal([]byte(Orders), &m)
		if err != nil {
			panic(err)
		}
	}

	return m
}

func Tramap() map[time.Time][]*trades.Trade {
	var m map[time.Time][]*trades.Trade
	{
		err := json.Unmarshal([]byte(Trades), &m)
		if err != nil {
			panic(err)
		}
	}

	return m
}

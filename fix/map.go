package fix

import (
	"encoding/json"
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func Map() map[time.Time][]*trades.Trade {
	var m map[time.Time][]*trades.Trade
	{
		err := json.Unmarshal([]byte(Trades), &m)
		if err != nil {
			panic(err)
		}
	}

	return m
}

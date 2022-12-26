package main

import (
	"fmt"

	"github.com/uncle-gua/go-binance/v2/futures"
)

const (
	ApiKey    = ""
	ApiSecret = ""
)

func main() {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		fmt.Printf("%d %d %d %s %s %d %d %.2f %.2f %.2f %.2f %.2f %d %v %.2f %.2f %.2f\n",
			event.Time,
			event.Kline.StartTime,
			event.Kline.EndTime,
			event.Kline.Symbol,
			event.Kline.Interval,
			event.Kline.FirstTradeID,
			event.Kline.LastTradeID,
			event.Kline.Open,
			event.Kline.High,
			event.Kline.Low,
			event.Kline.Close,
			event.Kline.Volume,
			event.Kline.TradeNum,
			event.Kline.IsFinal,
			event.Kline.QuoteVolume,
			event.Kline.ActiveBuyVolume,
			event.Kline.ActiveBuyQuoteVolume,
		)
	}

	errHandler := func(err error) {
		fmt.Println(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	doneC, _, err := futures.WsKlineServe("ETHUSDT", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

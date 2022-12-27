package main

import (
	"context"
	"ticker/models"

	"github.com/uncle-gua/go-binance/v2/futures"
	"github.com/uncle-gua/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ApiKey    = ""
	ApiSecret = ""
)

func main() {
	wsKlineHandler := func(event *futures.WsKlineEvent) {
		ticker := models.Ticker{
			ID:                   primitive.NewObjectID(),
			EventTime:            event.Time,
			StartTime:            event.Kline.StartTime,
			EndTime:              event.Kline.EndTime,
			Symbol:               event.Kline.Symbol,
			Interval:             event.Kline.Interval,
			FirstTradeID:         event.Kline.FirstTradeID,
			LastTradeID:          event.Kline.LastTradeID,
			Open:                 event.Kline.Open,
			High:                 event.Kline.High,
			Low:                  event.Kline.Low,
			Close:                event.Kline.Close,
			Volume:               event.Kline.Volume,
			TradeNum:             event.Kline.TradeNum,
			IsFinal:              event.Kline.IsFinal,
			QuoteVolume:          event.Kline.QuoteVolume,
			ActiveBuyVolume:      event.Kline.ActiveBuyVolume,
			ActiveBuyQuoteVolume: event.Kline.ActiveBuyQuoteVolume,
		}
		if _, err := models.TickerCollection.InsertOne(context.TODO(), &ticker, options.InsertOne()); err != nil {
			log.Error(err)
		}
	}

	errHandler := func(err error) {
		log.Error(err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	doneC, _, err := futures.WsKlineServe("ETHUSDT", "1m", wsKlineHandler, errHandler)
	if err != nil {
		log.Error(err)
		return
	}
	<-doneC
}

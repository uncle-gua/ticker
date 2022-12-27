package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticker struct {
	ID                   primitive.ObjectID `bson:"_id"`
	EventTime            int64              `bson:"eventTime"`
	StartTime            int64              `bson:"startTime"`
	EndTime              int64              `bson:"endTime"`
	Symbol               string             `bson:"symbol"`
	Interval             string             `bson:"interval"`
	FirstTradeID         int64              `bson:"firstTradeID"`
	LastTradeID          int64              `bson:"lastTradeID"`
	Open                 float64            `bson:"open"`
	High                 float64            `bson:"high"`
	Low                  float64            `bson:"low"`
	Close                float64            `bson:"close"`
	Volume               float64            `bson:"volume"`
	TradeNum             int64              `bson:"tradeNum"`
	IsFinal              bool               `bson:"isFinal"`
	QuoteVolume          float64            `bson:"quoteVolume"`
	ActiveBuyVolume      float64            `bson:"activeBuyVolume"`
	ActiveBuyQuoteVolume float64            `bson:"activeBuyQuoteVolume"`
}

package marketdata

import (
	"testing"

	"github.com/billfort/binance-usdmfuture/pub"
	"github.com/stretchr/testify/require"
)

// go test -v -run TestTestConnectivity
func TestConnectivity(t *testing.T) {
	err := Connectivity()
	require.Nil(t, err)
}

// go test -v -run TestCheckServerTime
func TestCheckServerTime(t *testing.T) {
	serverTime, err := CheckServerTime()
	require.Nil(t, err)
	require.Greater(t, serverTime, int64(0))
}

// go test -v -run TestExchangeInfo
func TestExchangeInfo(t *testing.T) {
	ei, err := ExchangeInfo()
	require.Nil(t, err)
	require.NotNil(t, ei)
}

// go test -v -run TestOrderBook
func TestOrderBook(t *testing.T) {
	symbol := "BTCUSDT"
	limit := 5
	ob, err := OrderBook(symbol, limit)
	require.Nil(t, err)
	require.NotNil(t, ob)
	require.Equal(t, limit, len(ob.Asks))
	require.Equal(t, limit, len(ob.Bids))
}

// go test -v -run TestRecentMarketTrade
func TestRecentMarketTrade(t *testing.T) {
	symbol := "BTCUSDT"
	limit := 5
	trades, err := RecentMarketTrades(symbol, limit)
	require.Nil(t, err)
	require.NotNil(t, trades)
	require.Equal(t, len(trades), limit)
}

// go test -v -run TestHistoricalTrades
func TestHistoricalTrades(t *testing.T) {
	symbol := "BTCUSDT"
	limit := 5
	fromTradeId := 0
	trades, err := HistoricalTrades(symbol, fromTradeId, limit)
	require.Nil(t, err)
	require.NotNil(t, trades)
	require.Equal(t, len(trades), limit)
}

// go test -v -run TestAggregatedTrades
func TestAggregatedTrades(t *testing.T) {
	symbol := "BTCUSDT"
	fromId := 0
	limit := 5
	startTime := 0
	endTime := 0
	trades, err := AggregatedTrades(symbol, fromId, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, trades)
	require.Equal(t, len(trades), limit)
}

// go test -v -run TestKline
func TestKline(t *testing.T) {
	symbol := "BTCUSDT"
	interval := pub.KI_Hour1
	limit := 5
	startTime := 0
	endTime := 0
	klines, err := Klines(symbol, interval, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, klines)
	require.Equal(t, len(klines), limit)
}

// go test -v -run TestMarkPrice
func TestMarkPrice(t *testing.T) {
	symbol := "BTCUSDT"
	markPrice, err := MarkPrice(symbol)
	require.Nil(t, err)
	require.NotNil(t, markPrice)
	require.Equal(t, len(markPrice), 1)
}

// go test -v -run TestFundingRateHistory
func TestFundingRateHistory(t *testing.T) {
	symbol := "BTCUSDT"
	limit := 5
	startTime := 0
	endTime := 0
	rates, err := FundingRateHistory(symbol, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, rates)
	require.Equal(t, len(rates), limit)
}

// go test -v -run TestFundingInfo
func TestFundingInfo(t *testing.T) {
	info, err := FundingInfo()
	require.Nil(t, err)
	require.NotNil(t, info)
	require.Greater(t, len(info), 0)
}

// go test -v -run TestTickerPrice
func TestTickerPriceStatistics24hr(t *testing.T) {
	symbol := "BTCUSDT"
	ticker, err := TickerPriceStatistics24hr(symbol)
	require.Nil(t, err)
	require.NotNil(t, ticker)
}

// go test -v -run TestTickerPrice
func TestTickerPrice(t *testing.T) {
	ticker, err := TickerPrice("", "v2")
	require.Nil(t, err)
	require.NotNil(t, ticker)
	require.Greater(t, len(ticker), 0)
}

// go test -v -run TestBookTicker
func TestBookTicker(t *testing.T) {
	symbol := "BTCUSDT"
	ticker, err := BookTicker(symbol)
	require.Nil(t, err)
	require.NotNil(t, ticker)
	require.Greater(t, len(ticker), 0)
}

// go test -v -run TestDeliveryPrice
func TestDeliveryPrice(t *testing.T) {
	symbol := "BTCUSDT"
	price, err := DeliveryPrice(symbol)
	require.Nil(t, err)
	require.NotNil(t, price)
}

// go test -v -run TestOpenInterest
func TestOpenInterest(t *testing.T) {
	symbol := "BTCUSDT"
	oi, err := OpenInterest(symbol)
	require.Nil(t, err)
	require.NotNil(t, oi)
}

// go test -v -run TestOpenInterestHist
func TestOpenInterestHist(t *testing.T) {
	symbol := "BTCUSDT"
	period := pub.KI_Hour1
	limit := 5
	startTime := 0
	endTime := 0
	oi, err := OpenInterestHist(symbol, period, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, oi)
	require.Greater(t, len(oi), 0)
}

// go test -v -run TestTopLongShortPositionRatio
func TestTopLongShortPositionRatio(t *testing.T) {
	symbol := "BTCUSDT"
	period := pub.KI_Hour1
	limit := 5
	startTime := 0
	endTime := 0
	ratios, err := TopLongShortPositionRatio(symbol, period, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, ratios)
	require.Greater(t, len(ratios), 0)
}

// go test -v -run TestTopLongShortAccountRatio
func TestTopLongShortAccountRatio(t *testing.T) {
	symbol := "BTCUSDT"
	period := pub.KI_Hour1
	limit := 5
	startTime := 0
	endTime := 0
	ratios, err := TopLongShortAccountRatio(symbol, period, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, ratios)
	require.Greater(t, len(ratios), 0)
}

// go test -v -run TestGlobalLongShortAccountRatio
func TestGlobalLongShortAccountRatio(t *testing.T) {
	symbol := "BTCUSDT"
	period := pub.KI_Hour1
	limit := 5
	startTime := 0
	endTime := 0
	ratios, err := GlobalLongShortAccountRatio(symbol, period, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, ratios)
	require.Greater(t, len(ratios), 0)
}

// go test -v -run TestTakerLongShortRatio
func TestTakerLongShortRatio(t *testing.T) {
	symbol := "BTCUSDT"
	period := pub.KI_Hour1
	limit := 5
	startTime := 0
	endTime := 0
	ratios, err := TakerLongShortRatio(symbol, period, startTime, endTime, limit)
	require.Nil(t, err)
	require.NotNil(t, ratios)
	require.Greater(t, len(ratios), 0)
}

// // go test -v -run TestBlvtKlines
// func TestBlvtKlines(t *testing.T) {
// 	symbol := "BTCUSDT"
// 	interval := pub.Hour1
// 	limit := 5
// 	startTime := 0
// 	endTime := 0
// 	klines, err := BlvtKlines(symbol, interval, startTime, endTime, limit)
// 	require.Nil(t, err)
// 	require.NotNil(t, klines)
// 	require.Equal(t, len(klines), limit)
// }

// go test -v -run TestCompositeIndexInfo
func TestCompositeIndexInfo(t *testing.T) {
	symbol := ""
	indexInfo, err := CompositeIndexInfo(symbol)
	require.Nil(t, err)
	require.NotNil(t, indexInfo)
	require.Greater(t, len(indexInfo), 0)
}

// go test -v -run TestAssetIndex
func TestAssetIndex(t *testing.T) {
	symbol := "BTCUSD"
	indexInfo, err := AssetIndex(symbol)
	require.Nil(t, err)
	require.NotNil(t, indexInfo)
	require.Greater(t, len(indexInfo), 0)
}

// go test -v -run TestIndexConstituents
func TestIndexConstituents(t *testing.T) {
	symbol := "BTCUSDT"
	cons, err := IndexConstituents(symbol)
	require.Nil(t, err)
	require.NotNil(t, cons)
}

package marketdata

import (
	"encoding/json"
	"log"

	"github.com/billfort/binance-usdmfuture/pub"
)

// Test connectivity to the Rest API.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api
func Connectivity() error {
	resBody, err := pub.GetNoSign("/fapi/v1/ping", nil)
	if err != nil {
		return err
	}
	log.Println("TestConnectivity: ", string(resBody)) // "{}"
	return nil
}

// Test connectivity to the Rest API and get the current server time.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Check-Server-Time
func CheckServerTime() (int64, error) {
	resBody, err := pub.GetNoSign("/fapi/v1/time", nil)
	if err != nil {
		return 0, err
	}

	type st struct {
		ServerTime int64 `json:"serverTime"`
	}
	var s st
	err = json.Unmarshal(resBody, &s)
	if err != nil {
		return 0, err
	}

	return s.ServerTime, nil
}

// Current exchange trading rules and symbol information
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Exchange-Information
func ExchangeInfo() (*ExchInfo, error) {
	resBody, err := pub.GetNoSign("/fapi/v1/exchangeInfo", nil)
	if err != nil {
		return nil, err
	}
	var ei ExchInfo
	err = json.Unmarshal(resBody, &ei)
	if err != nil {
		return nil, err
	}
	log.Printf("market.GetExchInfo, rateLimits len %v, Assets len %v, Symbols len %v\n", len(ei.RateLimits), len(ei.Assets), len(ei.Symbols))

	return &ei, nil
}

// Query symbol orderbook
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Order-Book
func OrderBook(symbol string, limit int) (*orderBook, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/depth", params)
	if err != nil {
		return nil, err
	}
	var d orderBook
	err = json.Unmarshal(resBody, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// Get recent market trades filled in the order book. Only market trades will be returned,
// which means the insurance fund trades and ADL trades won't be returned.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Recent-Trades-List
func RecentMarketTrades(symbol string, limit int) ([]marketTrade, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/trades", params)
	if err != nil {
		return nil, err
	}
	var trades []marketTrade
	err = json.Unmarshal(resBody, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}

// Get older market historical trades.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Old-Trades-Lookup
func HistoricalTrades(symbol string, fromId, limit int) ([]marketTrade, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if fromId > 0 {
		params["fromId"] = fromId
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/historicalTrades", params)
	if err != nil {
		return nil, err
	}
	var trades []marketTrade
	err = json.Unmarshal(resBody, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}

// Get compressed, aggregate market trades. Market trades that fill in 100ms with the same price
// and the same taking side will have the quantity aggregated.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Compressed-Aggregate-Trades-List
func AggregatedTrades(symbol string, fromId, startTime, endTime, limit int) ([]aggTrade, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if fromId > 0 {
		params["fromId"] = fromId
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/aggTrades", params)
	if err != nil {
		return nil, err
	}
	var trades []aggTrade
	err = json.Unmarshal(resBody, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}

// convert slice of interface to slice of kData
func sliceToKdata(arr [][]interface{}) []kData {
	kDataArr := make([]kData, len(arr))
	for i := 0; i < len(arr); i++ {
		kDataArr[i].OpenTime = int64(arr[i][0].(float64))
		kDataArr[i].Open = arr[i][1].(string)
		kDataArr[i].High = arr[i][2].(string)
		kDataArr[i].Low = arr[i][3].(string)
		kDataArr[i].Close = arr[i][4].(string)
		kDataArr[i].Volume = arr[i][5].(string)
		kDataArr[i].CloseTime = int64(arr[i][6].(float64))
		kDataArr[i].QuoteAssetVolume = arr[i][7].(string)
		kDataArr[i].NumberOfTrades = int(arr[i][8].(float64))
		kDataArr[i].TakerBuyBaseAssetVolume = arr[i][9].(string)
		kDataArr[i].TakerBuyQuoteAssetVolume = arr[i][10].(string)
		kDataArr[i].Ignore = arr[i][11].(string)
	}
	return kDataArr
}

// Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Kline-Candlestick-Data
func Klines(symbol string, interval pub.KlineInterval, startTime, endTime, limit int64) ([]kData, error) {
	params := map[string]interface{}{
		"symbol":   symbol,
		"interval": interval,
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/klines", params)
	if err != nil {
		return nil, err
	}

	var arr [][]interface{}
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	kDataArr := sliceToKdata(arr)
	return kDataArr, nil
}

// Kline/candlestick bars for a specific contract type. Klines are uniquely identified by their open time.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Continuous-Contract-Kline-Candlestick-Data
func ContinuousKlines(pair string, contractType pub.ContractType, interval pub.KlineInterval,
	startTime, endTime, limit int) ([]kData, error) {
	params := map[string]interface{}{
		"pair":         pair,
		"contractType": contractType,
		"interval":     interval,
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/continuousKlines", params)
	if err != nil {
		return nil, err
	}

	var arr [][]interface{}
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	kDataArr := sliceToKdata(arr)
	return kDataArr, nil
}

// Kline/candlestick bars for the index price of a pair. Klines are uniquely identified by their open time.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Index-Price-Kline-Candlestick-Data
func IndexPriceKlines(pair string, contractType pub.ContractType, interval pub.KlineInterval,
	startTime, endTime, limit int) ([]kData, error) {
	params := map[string]interface{}{
		"pair":         pair,
		"contractType": contractType,
		"interval":     interval,
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/indexPriceKlines", params)
	if err != nil {
		return nil, err
	}

	var arr [][]interface{}
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	kDataArr := sliceToKdata(arr)
	return kDataArr, nil
}

// Kline/candlestick bars for the mark price of a symbol. Klines are uniquely identified by their open time.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Mark-Price-Kline-Candlestick-Data
func MarkPriceKlines(symbol string, interval pub.KlineInterval, startTime, endTime, limit int) ([]kData, error) {
	params := map[string]interface{}{
		"symbol":   symbol,
		"interval": interval,
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/markPriceKlines", params)
	if err != nil {
		return nil, err
	}

	var arr [][]interface{}
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	kDataArr := sliceToKdata(arr)
	return kDataArr, nil
}

// Premium index kline bars of a symbol. Klines are uniquely identified by their open time.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Premium-Index-Kline-Data
func PremiumIndexKlines(symbol string, interval pub.KlineInterval, startTime, endTime, limit int) ([]kData, error) {
	params := map[string]interface{}{
		"symbol":   symbol,
		"interval": interval,
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/premiumIndexKlines", params)
	if err != nil {
		return nil, err
	}

	var arr [][]interface{}
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	kDataArr := sliceToKdata(arr)
	return kDataArr, nil
}

// Mark Price and Funding Rate
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Mark-Price
func MarkPrice(symbol string) ([]markPrice, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	resBody, err := pub.GetNoSign("/fapi/v1/premiumIndex", params)
	if err != nil {
		return nil, err
	}

	if symbol != "" {
		var mp markPrice
		err = json.Unmarshal(resBody, &mp)
		if err != nil {
			return nil, err
		}
		return []markPrice{mp}, nil
	} else {
		var mps []markPrice
		err = json.Unmarshal(resBody, &mps)
		if err != nil {
			return nil, err
		}
		return mps, nil
	}
}

// Get Funding Rate History
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Get-Funding-Rate-History
func FundingRateHistory(symbol string, startTime, endTime, limit int) ([]fundingRate, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/fapi/v1/fundingRate", params)
	if err != nil {
		return nil, err
	}

	var frs []fundingRate
	err = json.Unmarshal(resBody, &frs)
	if err != nil {
		return nil, err
	}
	return frs, nil
}

// Query funding rate info for symbols that had FundingRateCap/ FundingRateFloor / fundingIntervalHours adjustment
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Get-Funding-Rate-Info
func FundingInfo() ([]fundingInfo, error) {
	resBody, err := pub.GetNoSign("/fapi/v1/fundingInfo", nil)
	if err != nil {
		return nil, err
	}

	var fis []fundingInfo
	err = json.Unmarshal(resBody, &fis)
	if err != nil {
		return nil, err
	}
	return fis, nil
}

// 24 hour rolling window price change statistics.
// Careful when accessing this with no symbol.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/24hr-Ticker-Price-Change-Statistics
func TickerPriceStatistics24hr(symbol string) ([]ticker24hr, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	resBody, err := pub.GetNoSign("/fapi/v1/ticker/24hr", params)
	if err != nil {
		return nil, err
	}

	if symbol != "" {
		var t ticker24hr
		err = json.Unmarshal(resBody, &t)
		if err != nil {
			return nil, err
		}
		return []ticker24hr{t}, nil
	} else {
		var ts []ticker24hr
		err = json.Unmarshal(resBody, &ts)
		if err != nil {
			return nil, err
		}
		return ts, nil
	}
}

// Latest price for a symbol or symbols.
// version: v1, v2
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Symbol-Price-Ticker
func TickerPrice(symbol string, version string) ([]tickerPrice, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	resBody, err := pub.GetNoSign("/fapi/"+version+"/ticker/price", params)
	if err != nil {
		return nil, err
	}

	if symbol != "" {
		var t tickerPrice
		err = json.Unmarshal(resBody, &t)
		if err != nil {
			return nil, err
		}
		return []tickerPrice{t}, nil
	} else {
		var ts []tickerPrice
		err = json.Unmarshal(resBody, &ts)
		if err != nil {
			return nil, err
		}
		return ts, nil
	}
}

// Best price/qty on the order book for a symbol or symbols.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Symbol-Order-Book-Ticker
func BookTicker(symbol string) ([]bookTicker, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	resBody, err := pub.GetNoSign("/fapi/v1/ticker/bookTicker", params)
	if err != nil {
		return nil, err
	}

	if symbol != "" {
		var t bookTicker
		err = json.Unmarshal(resBody, &t)
		if err != nil {
			return nil, err
		}
		return []bookTicker{t}, nil
	} else {
		var ts []bookTicker
		err = json.Unmarshal(resBody, &ts)
		if err != nil {
			return nil, err
		}
		return ts, nil
	}
}

func DeliveryPrice(pair string) ([]deliveryPrice, error) {
	params := map[string]interface{}{
		"pair": pair,
	}
	resBody, err := pub.GetNoSign("/futures/data/delivery-price", params)
	if err != nil {
		return nil, err
	}

	var arr []deliveryPrice
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// Get present open interest of a specific symbol.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Open-Interest
func OpenInterest(symbol string) (*openInterest, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	resBody, err := pub.GetNoSign("/fapi/v1/openInterest", params)
	if err != nil {
		return nil, err
	}

	var oi openInterest
	err = json.Unmarshal(resBody, &oi)
	if err != nil {
		return nil, err
	}
	return &oi, nil
}

// Open Interest Statistics
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Open-Interest-Statistics
func OpenInterestHist(symbol string, period pub.KlineInterval, startTime, endTime, limit int) ([]openInterestHist, error) {
	params := map[string]interface{}{
		"symbol": symbol,
		"period": string(period),
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/futures/data/openInterestHist", params)
	if err != nil {
		return nil, err
	}

	var arr []openInterestHist
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// The proportion of net long and net short positions to total open positions of the top 20% users with the highest margin balance.
// Long Position % = Long positions of top traders / Total open positions of top traders
// Short Position % = Short positions of top traders / Total open positions of top traders
// Long/Short Ratio (Positions) = Long Position % / Short Position %
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Top-Trader-Long-Short-Ratio
func TopLongShortPositionRatio(symbol string, period pub.KlineInterval, startTime, endTime, limit int) ([]longShortRatio, error) {
	params := map[string]interface{}{
		"symbol": symbol,
		"period": string(period),
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/futures/data/topLongShortPositionRatio", params)
	if err != nil {
		return nil, err
	}

	var arr []longShortRatio
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// The proportion of net long and net short accounts to total accounts of the top 20% users with the highest margin balance. Each account is counted once only.
// Long Account % = Accounts of top traders with net long positions / Total accounts of top traders with open positions
// Short Account % = Accounts of top traders with net short positions / Total accounts of top traders with open positions
// Long/Short Ratio (Accounts) = Long Account % / Short Account %
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Top-Long-Short-Account-Ratio
func TopLongShortAccountRatio(symbol string, period pub.KlineInterval, startTime, endTime, limit int) ([]longShortRatio, error) {
	params := map[string]interface{}{
		"symbol": symbol,
		"period": string(period),
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/futures/data/topLongShortAccountRatio", params)
	if err != nil {
		return nil, err
	}

	var arr []longShortRatio
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// Query symbol Long/Short Ratio
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Long-Short-Ratio
func GlobalLongShortAccountRatio(symbol string, period pub.KlineInterval, startTime, endTime, limit int) ([]longShortRatio, error) {
	params := map[string]interface{}{
		"symbol": symbol,
		"period": string(period),
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/futures/data/globalLongShortAccountRatio", params)
	if err != nil {
		return nil, err
	}

	var arr []longShortRatio
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// Taker Buy/Sell Volume
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Taker-BuySell-Volume
func TakerLongShortRatio(symbol string, period pub.KlineInterval, startTime, endTime, limit int) ([]takerLongShortRatio, error) {
	params := map[string]interface{}{
		"symbol": symbol,
		"period": string(period),
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}
	resBody, err := pub.GetNoSign("/futures/data/takerlongshortRatio", params)
	if err != nil {
		return nil, err
	}

	var arr []takerLongShortRatio
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// the response of /fapi/v1/lvtKlines is `page not found`.
// TODO: find the correct endpoint
// func BlvtKlines(symbol string, interval pub.KlineInterval, startTime, endTime, limit int) ([]kData, error) {
// 	params := map[string]interface{}{
// 		"symbol":   symbol,
// 		"interval": interval,
// 	}
// 	if startTime > 0 {
// 		params["startTime"] = startTime
// 	}
// 	if endTime > 0 {
// 		params["endTime"] = endTime
// 	}
// 	if limit > 0 {
// 		params["limit"] = limit
// 	}
// 	resBody, err := pub.GetNoSign("/fapi/v1/lvtKlines", params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("resBody: %v\n", string(resBody))

// 	var arr [][]interface{}
// 	err = json.Unmarshal(resBody, &arr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(arr) == 0 {
// 		return nil, nil
// 	}
// 	kDataArr := sliceToKdata(arr)
// 	return kDataArr, nil
// }

// Query composite index symbol information
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Composite-Index-Symbol-Information
func CompositeIndexInfo(symbol string) ([]indexInfo, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	resBody, err := pub.GetNoSign("/fapi/v1/indexInfo", params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var arr []indexInfo
	err = json.Unmarshal(resBody, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// asset index for Multi-Assets mode
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Multi-Assets-Mode-Asset-Index
func AssetIndex(symbol string) ([]assetIndex, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}
	resBody, err := pub.GetNoSign("/fapi/v1/assetIndex", params)
	if err != nil {
		return nil, err
	}
	if symbol != "" {
		var ai assetIndex
		err = json.Unmarshal(resBody, &ai)
		if err != nil {
			return nil, err
		}
		return []assetIndex{ai}, nil
	} else {
		var arr []assetIndex
		err = json.Unmarshal(resBody, &arr)
		if err != nil {
			return nil, err
		}
		return arr, nil
	}
}

// Query index price constituents
// https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Index-Constituents
func IndexConstituents(symbol string) (*indexConstituents, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	resBody, err := pub.GetNoSign("/fapi/v1/constituents", params)
	if err != nil {
		return nil, err
	}
	var ic indexConstituents
	err = json.Unmarshal(resBody, &ic)
	if err != nil {
		return nil, err
	}
	return &ic, nil
}

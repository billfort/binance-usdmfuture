package convert

import (
	"encoding/json"

	"github.com/billfort/binance-usdmfuture/pub"
)

// Query for all convertible token pairs and the tokens’ respective upper/lower limits
// https://developers.binance.com/docs/derivatives/usds-margined-futures/convert
func ListPairs(fromAsset, toAsset string) ([]convertPair, error) {
	params := map[string]interface{}{
		"fromAsset": fromAsset,
		"toAsset":   toAsset,
	}
	resBody, err := pub.GetNoSign("/fapi/v1/convert/exchangeInfo", params)
	if err != nil {
		return nil, err
	}
	var resp []convertPair
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Request a quote for the requested token pairs
// validTime: 10s, 30s, 1m, default 10s
// https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Send-quote-request
func RequestQuote(key *pub.Key, fromAsset, toAsset, fromAssetAmount, toAmount, validTime string) (*quote, error) {
	params := map[string]interface{}{
		"fromAsset":       fromAsset,
		"toAsset":         toAsset,
		"fromAssetAmount": fromAssetAmount,
		"toAmount":        toAmount,
		"validTime":       validTime,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/convert/getQuote", params)
	if err != nil {
		return nil, err
	}
	var resp quote
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Accept the offered quote by quote ID.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Accept-Quote
func AcceptQuote(key *pub.Key, quoteID string) (*acceptQuote, error) {
	params := map[string]interface{}{
		"quoteId": quoteID,
	}
	resBody, _, err := pub.PostWithSign(key, "/fapi/v1/convert/confirm", params)
	if err != nil {
		return nil, err
	}
	var resp acceptQuote
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Query order status by order ID.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Order-Status
func QueryOrderStatus(key *pub.Key, orderID string) (*orderStatus, error) {
	params := map[string]interface{}{
		"orderId": orderID,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/convert/orderStatus", params)
	if err != nil {
		return nil, err
	}
	var resp orderStatus
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

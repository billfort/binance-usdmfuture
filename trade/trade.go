package trade

import (
	"encoding/json"
	"fmt"

	"github.com/billfort/binance-usdmfuture/pub"
)

func NewOrder(key *pub.Key, op *OrderParam) (*orderResponse, error) {
	params := pub.StructToMap(op)

	resBody, errMsg, err := pub.PostWithSign(key, "/fapi/v1/order", params)
	if err != nil {
		return nil, err
	}
	if errMsg.Code != 0 {
		return nil, fmt.Errorf("%+v", errMsg)
	}

	var resp orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func TestOrder(key *pub.Key, op *OrderParam) (*orderResponse, error) {
	params := pub.StructToMap(op)

	resBody, errMsg, err := pub.PostWithSign(key, "/fapi/v1/order/test", params)
	if err != nil {
		return nil, err
	}
	if errMsg.Code != 0 {
		return nil, fmt.Errorf("%+v", errMsg)
	}

	var resp orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func BatchOrders(key *pub.Key, ops []OrderParam) ([]orderResponse, error) {
	b, err := json.Marshal(ops)
	if err != nil {
		return nil, err
	}
	batchParams := struct {
		BatchOrders string `json:"batchOrders"`
	}{
		BatchOrders: string(b),
	}
	params := pub.StructToMap(&batchParams)

	resBody, errMsg, err := pub.PostWithSign(key, "/fapi/v1/batchOrders", params)
	if err != nil {
		return nil, err
	}
	if errMsg.Code != 0 {
		return nil, fmt.Errorf("%+v", errMsg)
	}

	var resp []orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ModifyOrder(key *pub.Key, mp *modifyParam) (*orderResponse, error) {
	params := pub.StructToMap(mp)

	resBody, err := pub.PutWithSign(key, "/fapi/v1/order", params)
	if err != nil {
		return nil, err
	}

	var resp orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func ModifyBatchOrders(key *pub.Key, mps []modifyParam) ([]orderResponse, error) {
	b, err := json.Marshal(mps)
	if err != nil {
		return nil, err
	}
	batchParams := struct {
		BatchOrders string `json:"batchOrders"`
	}{
		BatchOrders: string(b),
	}
	params := pub.StructToMap(&batchParams)

	resBody, err := pub.PutWithSign(key, "/fapi/v1/batchOrders", params)
	if err != nil {
		return nil, err
	}

	var resp []orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CancelOrder(key *pub.Key, symbol string, orderId int64, origClientOrderId string) (*orderResponse, error) {
	params := map[string]interface{}{
		"symbol":            symbol,
		"orderId":           orderId,
		"origClientOrderId": origClientOrderId,
	}

	resBody, err := pub.DeleteWithSign(key, "/fapi/v1/order", params)
	if err != nil {
		return nil, err
	}

	var resp orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func CancelBatchOrders(key *pub.Key, symbol string, orderIdList []int64, origClientOrderIdList []string) ([]orderResponse, error) {
	params := map[string]interface{}{
		"symbol":                symbol,
		"orderIdList":           orderIdList,
		"origClientOrderIdList": origClientOrderIdList,
	}

	resBody, err := pub.DeleteWithSign(key, "/fapi/v1/batchOrders", params)
	if err != nil {
		return nil, err
	}

	var resp []orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CancelAllOpenOrders(key *pub.Key, symbol string) error {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resBody, err := pub.DeleteWithSign(key, "/fapi/v1/allOpenOrders", params)
	if err != nil {
		return err
	}

	var resp pub.ErrMsg
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return err
	}
	if resp.Code == 200 { // "msg": "The operation of cancel all open order is done."
		return nil
	}

	return fmt.Errorf("%+v", resp)
}

func CountdownCancleAll(key *pub.Key, symbol string, countdownTime int64) error {
	params := map[string]interface{}{
		"symbol":        symbol,
		"countdownTime": countdownTime,
	}

	resBody, err := pub.DeleteWithSign(key, "/fapi/v1/countdownCancelAll", params)
	if err != nil {
		return err
	}

	var resp = struct {
		Symbol        string `json:"symbol"`
		CountdownTime string `json:"countdownTime"`
	}{}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return err
	}

	return nil
}

func QueryOrder(key *pub.Key, symbol string, orderId int64, origClientOrderId string) (*orderResponse, error) {
	params := map[string]interface{}{
		"symbol":            symbol,
		"orderId":           orderId,
		"origClientOrderId": origClientOrderId,
	}

	resBody, err := pub.GetWithSign(key, "/fapi/v1/order", params)
	if err != nil {
		return nil, err
	}

	var resp orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func QueryAllOrders(key *pub.Key, symbol string, orderId int64, startTime int64, endTime int64, limit int) ([]orderResponse, error) {
	params := map[string]interface{}{
		"symbol":    symbol,
		"orderId":   orderId,
		"startTime": startTime,
		"endTime":   endTime,
		"limit":     limit,
	}

	resBody, err := pub.GetWithSign(key, "/fapi/v1/allOrders", params)
	if err != nil {
		return nil, err
	}

	var resp []orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func QueryOpenOrders(key *pub.Key, symbol string) ([]orderResponse, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resBody, err := pub.GetWithSign(key, "/fapi/v1/openOrders", params)
	if err != nil {
		return nil, err
	}

	var resp []orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func QueryOpenOrder(key *pub.Key, symbol string, orderId int64, origClientOrderId string) (*orderResponse, error) {
	params := map[string]interface{}{
		"symbol":            symbol,
		"orderId":           orderId,
		"origClientOrderId": origClientOrderId,
	}

	resBody, err := pub.GetWithSign(key, "/fapi/v1/openOrder", params)
	if err != nil {
		return nil, err
	}

	var resp orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func QueryForceOrders(key *pub.Key, symbol string, autoCloseType string, startTime int64, endTime int64, limit int) ([]orderResponse, error) {
	params := map[string]interface{}{
		"symbol":        symbol,
		"autoCloseType": autoCloseType,
		"startTime":     startTime,
		"endTime":       endTime,
		"limit":         limit,
	}

	resBody, err := pub.GetWithSign(key, "/fapi/v1/forceOrders", params)
	if err != nil {
		return nil, err
	}

	var resp []orderResponse
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func QueryUserTrades(key *pub.Key, symbol string, orderId int64, startTime, endTime, fromId int64, limit int) ([]tradeInfo, error) {
	params := map[string]interface{}{
		"symbol":    symbol,
		"orderId":   orderId,
		"startTime": startTime,
		"endTime":   endTime,
		"fromId":    fromId,
		"limit":     limit,
	}

	resBody, err := pub.GetWithSign(key, "/fapi/v1/userTrades", params)
	if err != nil {
		return nil, err
	}

	var resp []tradeInfo
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SetMarginType(key *pub.Key, symbol string, marginType pub.MarginType) error {
	params := map[string]interface{}{
		"symbol":     symbol,
		"marginType": marginType,
	}

	_, errMsg, err := pub.PostWithSign(key, "/fapi/v1/marginType", params)
	if err != nil {
		return err
	}
	if errMsg.Code < 0 {
		return fmt.Errorf("%+v", errMsg)
	}

	return nil
}

// dualSidePosition: "true": Enable Hedge Mode, "false": one-way mode
func SetPositionMode(key *pub.Key, dualSidePosition bool) error {
	params := map[string]interface{}{
		"dualSidePosition": dualSidePosition,
	}

	_, errMsg, err := pub.PostWithSign(key, "/fapi/v1/positionSide/dual", params)
	if err != nil {
		return err
	}
	if errMsg.Code < 0 {
		return fmt.Errorf("%+v", errMsg)
	}

	return nil
}

func SetLeverage(key *pub.Key, symbol string, leverage int) (*leverageInfo, error) {
	params := map[string]interface{}{
		"symbol":   symbol,
		"leverage": leverage,
	}

	resBody, errMsg, err := pub.PostWithSign(key, "/fapi/v1/leverage", params)
	if err != nil {
		return nil, err
	}
	if errMsg.Code < 0 {
		return nil, fmt.Errorf("%+v", errMsg)
	}

	var resp leverageInfo
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// multiAssetMargin: true: Multi-asset mode, false: Single-asset mode
func SetMarginAssetMode(key *pub.Key, symbol string, multiAssetMargin bool) error {
	params := map[string]interface{}{
		"symbol":           symbol,
		"multiAssetMargin": multiAssetMargin,
	}

	_, errMsg, err := pub.PostWithSign(key, "/fapi/v1/multiAssetsMargin", params)
	if err != nil {
		return err
	}
	if errMsg.Code < 0 {
		return fmt.Errorf("%+v", errMsg)
	}

	return nil
}

// type_: 1: Add position margin, 2: Reduce position margin
func ModifyPositionMargin(key *pub.Key, symbol string, positionSide pub.PositionSide, amount string,
	type_ int) error {
	params := map[string]interface{}{
		"symbol":       symbol,
		"positionSide": positionSide,
		"amount":       amount,
		"type":         type_,
	}

	_, errMsg, err := pub.PostWithSign(key, "/fapi/v1/positionMargin", params)
	if err != nil {
		return err
	}
	if errMsg.Code < 0 {
		return fmt.Errorf("%+v", errMsg)
	}

	var resp = struct {
		Code   int     `json:"code"`
		Msg    string  `json:"msg"`
		Amount float64 `json:"amount"`
		Type   int     `json:"type"`
	}{}
	err = json.Unmarshal([]byte(errMsg.Msg), &resp)
	if err != nil {
		return err
	}

	return nil
}

func GetPositionInfoV2(symbol string) ([]positionInfo, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resBody, err := pub.GetWithSign(nil, "/fapi/v2/positionRisk", params)
	if err != nil {
		return nil, err
	}

	var resp []positionInfo
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetPositionInfoV3(symbol string) ([]positionInfo, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resBody, err := pub.GetWithSign(nil, "/fapi/v2/positionRisk", params)
	if err != nil {
		return nil, err
	}

	var resp []positionInfo
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func AdlQuantile(symbol string) ([]adlQuantile, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resBody, err := pub.GetWithSign(nil, "/fapi/v1/adlQuantile", params)
	if err != nil {
		return nil, err
	}

	var resp []adlQuantile
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// type,	1: Add position margin，2: Reduce position margin
func GetPositionMarginHistory(symbol string, type_ int, startTime, endTime int64, limit int) ([]positionMarginHist, error) {
	params := map[string]interface{}{
		"symbol":    symbol,
		"type":      type_,
		"startTime": startTime,
		"endTime":   endTime,
		"limit":     limit,
	}

	resBody, err := pub.GetWithSign(nil, "/fapi/v1/positionMargin/history", params)
	if err != nil {
		return nil, err
	}

	var resp []positionMarginHist
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

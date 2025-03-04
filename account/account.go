package account

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/billfort/binance-usdmfuture/pub"
)

// Get account balance (USER_DATA)
// version: v2, v3
// v3: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V3
// v2: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V2
func AccountBalance(key *pub.Key, version string) ([]accountBalance, error) {
	resBody, err := pub.GetWithSign(key, fmt.Sprintf("/fapi/%v/balance", version), nil)
	if err != nil {
		return nil, err
	}

	var resp []accountBalance
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get current account information. User in single-asset/ multi-assets mode will see different value
// version: v2, v3
// v3: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Information-V3
// v2: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Information-V2
func AccountInfo(key *pub.Key, version string) (*accountInfo, error) {
	resBody, err := pub.GetWithSign(key, fmt.Sprintf("/fapi/%v/account", version), nil)
	if err != nil {
		return nil, err
	}

	var resp accountInfo
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Get User Commission Rate
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/User-Commission-Rate
func CommissionRate(key *pub.Key) (*commissionRate, error) {
	resBody, err := pub.GetWithSign(key, "/fapi/v1/commissionRate", nil)
	if err != nil {
		return nil, err
	}

	var resp commissionRate
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Query account configuration
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Config
func AccountConfiguration(key *pub.Key) (*accountConfiguration, error) {
	resBody, err := pub.GetWithSign(key, "/fapi/v1/accountConfiguration", nil)
	if err != nil {
		return nil, err
	}

	var resp accountConfiguration
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Get current account symbol configuration.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Symbol-Config
func SymbolConfiguration(key *pub.Key, symbol string) (*symbolConfiguration, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/positionRisk", params)
	if err != nil {
		return nil, err
	}

	var resp symbolConfiguration
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Query User Rate Limit
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Query-Rate-Limit
func UserRateLimit(key *pub.Key) ([]rateLimitInfo, error) {
	resBody, err := pub.GetWithSign(key, "/fapi/v1/rateLimit/order", nil)
	if err != nil {
		return nil, err
	}

	var resp []rateLimitInfo
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Query user notional and leverage bracket on speicfic symbol
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Notional-and-Leverage-Brackets
func LeverageBracket(key *pub.Key, symbol string) ([]leverageBracket, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/leverageBracket", params)
	if err != nil {
		return nil, err
	}

	if symbol != "" {
		var resp leverageBracket
		err = json.Unmarshal(resBody, &resp)
		if err != nil {
			return nil, err
		}
		return []leverageBracket{resp}, nil
	} else {
		var resp []leverageBracket
		err = json.Unmarshal(resBody, &resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

// Get user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Current-Multi-Assets-Mode
func MultiAssetsMargin(key *pub.Key) (bool, error) {
	resBody, err := pub.GetWithSign(key, "/fapi/v1/multiAssetsMargin", nil)
	if err != nil {
		return false, err
	}

	var resp = struct {
		MultiAssetsMargin bool `json:"multiAssetsMargin"`
	}{}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return false, err
	}

	return resp.MultiAssetsMargin, nil
}

// Get user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Current-Position-Mode
func DualSidePosition(key *pub.Key) (bool, error) {
	resBody, err := pub.GetWithSign(key, "/fapi/v1/positionSide/dual", nil)
	if err != nil {
		return false, err
	}

	var resp = struct {
		DualSidePosition bool `json:"dualSidePosition"`
	}{}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return false, err
	}

	return resp.DualSidePosition, nil
}

// Query income history
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Income-History
func IncomeHistory(key *pub.Key, symbol string, incomeType pub.IncomeType, startTime, endTime string, page, limit int) ([]income, error) {
	params := map[string]interface{}{
		"symbol":     symbol,
		"incomeType": incomeType,
		"startTime":  startTime,
		"endTime":    endTime,
		"page":       page,
		"limit":      limit,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/income", params)
	if err != nil {
		return nil, err
	}

	var resp []income
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Futures trading quantitative rules indicators
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Trading-Quantitative-Rules-Indicators
func QuantitativeIndicator(key *pub.Key, symbol string) (*indicators, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/apiTradingStatus", params)
	if err != nil {
		return nil, err
	}

	var resp indicators
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Get download id for futures transaction history
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Transaction-History
// downloadType: "income", "order", "trade"
func GetDownloadId(key *pub.Key, downloadType string, startTime, endTime int64) (*downloadId, error) {
	params := map[string]interface{}{
		"startTime": startTime,
		"endTime":   endTime,
	}
	resBody, err := pub.GetWithSign(key, fmt.Sprintf("/fapi/v1/%v/asyn", downloadType), params)
	if err != nil {
		return nil, err
	}

	var resp downloadId
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Get futures transaction history download link by Id
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Futures-Transaction-History-Download-Link-by-Id
// downloadType: "income", "order", "trade"
func GetDownloadUrl(key *pub.Key, downloadType, downloadId string) (*downloadUrl, error) {
	params := map[string]interface{}{
		"downloadId": downloadId,
	}
	resBody, err := pub.GetWithSign(key, fmt.Sprintf("/fapi/v1/%v/asyn/id", downloadType), params)
	if err != nil {
		return nil, err
	}

	var resp downloadUrl
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Change user's BNB Fee Discount (Fee Discount On or Fee Discount Off ) on EVERY symbol
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Toggle-BNB-Burn-On-Futures-Trade
func ToggleBNBFee(key *pub.Key, feeBurn bool) error {
	params := map[string]interface{}{
		"feeBurn": feeBurn,
	}
	_, errMsg, err := pub.PostWithSign(key, "/fapi/v1/feeBurn", params)
	if err != nil {
		return err
	}
	if errMsg.Code == 200 {
		return nil
	}
	return fmt.Errorf("%+v", errMsg)
}

// Get user's BNB Fee Discount (Fee Discount On or Fee Discount Off )
// https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-BNB-Burn-Status
func GetBNBFeeStatus(key *pub.Key) (bool, error) {
	resBody, err := pub.GetWithSign(key, "/fapi/v1/feeBurn", nil)
	if err != nil {
		return false, err
	}

	var resp = struct {
		FeeBurn bool `json:"feeBurn"`
	}{}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return false, err
	}

	return resp.FeeBurn, nil
}

// Get future asset transfer history (USER_DATA)
// https://developers.binance.com/docs/wallet/asset/query-user-universal-transfer
// startTime: milliSecond, max 6 months
// asset: "usdt", "btc", "eth", "bnb", "busd", "usdc", etc.
func GetInternalTransferHist(key *pub.Key, asset string, startTime int64) (list []TfrRow, err error) {
	// params: {"startTime": milliSecond, (optional)"asset": "usdt", (optional)"endTime": milliSecond}
	if startTime == 0 {
		startTime = time.Now().UnixMilli() - 6*720*3600000 // 6个月以来
	}

	params := map[string]interface{}{
		"startTime": fmt.Sprintf("%v", startTime), // max 6 months, default 7 days
		"asset":     asset,
		"size":      100, // max 100, default 10
	}

	resBody, err := pub.SpotGetWithSign(key, "/sapi/v1/futures/transfer", params)
	if err != nil {
		return nil, err
	}
	var r tfrHist
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return nil, err
	}
	fmt.Printf("transfer total: %v\n", r.Total)
	for i := 0; i < len(r.Rows); i++ {
		switch r.Rows[i].Type {
		case 1:
			r.Rows[i].CType = "Spot2USDM"
		case 2:
			r.Rows[i].CType = "USDM2Spot"
		case 3:
			r.Rows[i].CType = "Spot2CoinM"
		case 4:
			r.Rows[i].CType = "CoinM2Spot"
		default:
			r.Rows[i].CType = "ToBeDefined"
		}
		r.Rows[i].CTime = time.Unix(r.Rows[i].Timestamp/1000, 0).Format("2006-01-02 15:04:05")
	}

	return r.Rows, nil
}

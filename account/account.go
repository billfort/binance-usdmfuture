package account

import (
	"encoding/json"
	"fmt"

	"github.com/billfort/binance-usdmfuture/pub"
)

// version: v2, v3
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

// version: v2, v3
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

package portfoliomargin

import (
	"encoding/json"

	"github.com/billfort/binance-usdmfuture/pub"
)

// Get Classic Portfolio Margin current account information.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/portfolio-margin-endpoints
func GetPmAccountInfo(key *pub.Key, asset string) (*accountInfo, error) {
	params := map[string]interface{}{
		"asset": asset,
	}
	resBody, err := pub.GetWithSign(key, "/fapi/v1/pmAccountInfo", params)
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

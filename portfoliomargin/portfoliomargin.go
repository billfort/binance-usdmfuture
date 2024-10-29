package portfoliomargin

import (
	"encoding/json"

	"github.com/billfort/binance-usdmfuture/pub"
)

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

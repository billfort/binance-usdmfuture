package portfoliomargin

type accountInfo struct {
	MaxWithdrawAmountUSD string `json:"maxWithdrawAmountUSD"`
	Asset                string `json:"asset"`
	MaxWithdrawAmount    string `json:"maxWithdrawAmount"`
}

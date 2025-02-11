package account

type accountBalance struct {
	AccountAlias       string `json:"accountAlias"`       // unique account code
	Asset              string `json:"asset"`              // asset name
	Balance            string `json:"balance"`            // wallet balance
	CrossWalletBalance string `json:"crossWalletBalance"` // crossed wallet balance
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"` // whether the asset can be used as margin in Multi-Assets mode
	UpdateTime         int64  `json:"updateTime"`
}

type asset struct {
	Asset                  string `json:"asset"`                  // asset name
	WalletBalance          string `json:"walletBalance"`          // wallet balance
	UnrealizedProfit       string `json:"unrealizedProfit"`       // unrealized profit
	MarginBalance          string `json:"marginBalance"`          // margin balance
	MaintMargin            string `json:"maintMargin"`            // maintenance margin required
	InitialMargin          string `json:"initialMargin"`          // total initial margin required with current mark price
	PositionInitialMargin  string `json:"positionInitialMargin"`  // initial margin required for positions with current mark price
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // initial margin required for open orders with current mark price
	CrossWalletBalance     string `json:"crossWalletBalance"`     // crossed wallet balance
	CrossUnPnl             string `json:"crossUnPnl"`             // unrealized profit of crossed positions
	AvailableBalance       string `json:"availableBalance"`       // available balance
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`      // maximum amount for transfer out
	UpdateTime             int64  `json:"updateTime"`             // last update time
}

type position struct {
	Symbol           string `json:"symbol"`           // symbol
	PositionSide     string `json:"positionSide"`     // position side
	PositionAmt      string `json:"positionAmt"`      // position amount
	UnrealizedProfit string `json:"unrealizedProfit"` // unrealized profit
	IsolatedMargin   string `json:"isolatedMargin"`   // isolated margin
	Notional         string `json:"notional"`         // notional value
	IsolatedWallet   string `json:"isolatedWallet"`   // isolated wallet
	InitialMargin    string `json:"initialMargin"`    // initial margin required with current mark price
	MaintMargin      string `json:"maintMargin"`      // maintenance margin required
	UpdateTime       int64  `json:"updateTime"`       // last update time
	EntryPrice       string `json:"entryPrice"`       // entry price
	Leverage         string `json:"leverage"`         // leverage
}

type accountInfo struct {
	TotalInitialMargin          string     `json:"totalInitialMargin"`          // total initial margin required with current mark price (useless with isolated positions), only for USDT asset
	TotalMaintMargin            string     `json:"totalMaintMargin"`            // total maintenance margin required, only for USDT asset
	TotalWalletBalance          string     `json:"totalWalletBalance"`          // total wallet balance, only for USDT asset
	TotalUnrealizedProfit       string     `json:"totalUnrealizedProfit"`       // total unrealized profit, only for USDT asset
	TotalMarginBalance          string     `json:"totalMarginBalance"`          // total margin balance, only for USDT asset
	TotalPositionInitialMargin  string     `json:"totalPositionInitialMargin"`  // initial margin required for positions with current mark price, only for USDT asset
	TotalOpenOrderInitialMargin string     `json:"totalOpenOrderInitialMargin"` // initial margin required for open orders with current mark price, only for USDT asset
	TotalCrossWalletBalance     string     `json:"totalCrossWalletBalance"`     // crossed wallet balance, only for USDT asset
	TotalCrossUnPnl             string     `json:"totalCrossUnPnl"`             // unrealized profit of crossed positions, only for USDT asset
	AvailableBalance            string     `json:"availableBalance"`            // available balance, only for USDT asset
	MaxWithdrawAmount           string     `json:"maxWithdrawAmount"`           // maximum amount for transfer out, only for USDT asset
	Assets                      []asset    `json:"assets"`
	Positions                   []position `json:"positions"`
}

type commissionRate struct {
	Symbol              string `json:"symbol"`              // symbol
	MakerCommissionRate string `json:"makerCommissionRate"` // maker commission rate (0.02%)
	TakerCommissionRate string `json:"takerCommissionRate"` // taker commission rate (0.04%)
}

type accountConfiguration struct {
	FeeTier           int  `json:"feeTier"`           // account commission tier
	CanTrade          bool `json:"canTrade"`          // if can trade
	CanDeposit        bool `json:"canDeposit"`        // if can transfer in asset
	CanWithdraw       bool `json:"canWithdraw"`       // if can transfer out asset
	DualSidePosition  bool `json:"dualSidePosition"`  // if can enable dual side position
	UpdateTime        int  `json:"updateTime"`        // reserved property, please ignore
	MultiAssetsMargin bool `json:"multiAssetsMargin"` // if can enable multi-assets margin
	TradeGroupId      int  `json:"tradeGroupId"`      // trade group id
}

type symbolConfiguration struct {
	Symbol           string `json:"symbol"`           // symbol
	MarginType       string `json:"marginType"`       // margin type
	IsAutoAddMargin  string `json:"isAutoAddMargin"`  // is auto add margin
	Leverage         int    `json:"leverage"`         // leverage
	MaxNotionalValue string `json:"maxNotionalValue"` // max notional value
}

type rateLimitInfo struct {
	RateLimitType string `json:"rateLimitType"` // rate limit type
	Interval      string `json:"interval"`      // interval
	IntervalNum   int    `json:"intervalNum"`   // interval num
	Limit         int    `json:"limit"`         // limit
}

type leverageBracket struct {
	Symbol       string `json:"symbol"`       // symbol
	NotionalCoef string `json:"notionalCoef"` // user symbol bracket multiplier, only appears when user's symbol bracket is adjusted
	Brackets     []struct {
		Bracket          int     `json:"bracket"`          // Notional bracket
		InitialLeverage  int     `json:"initialLeverage"`  // Max initial leverage for this bracket
		NotionalCap      float64 `json:"notionalCap"`      // Cap notional of this bracket
		NotionalFloor    float64 `json:"notionalFloor"`    // Notional threshold of this bracket
		MaintMarginRatio float64 `json:"maintMarginRatio"` // Maintenance ratio for this bracket
		Cum              int     `json:"cum"`              // Auxiliary number for quick calculation
	} `json:"brackets"`
}

type income struct {
	Symbol     string `json:"symbol"`     // trade symbol, if existing
	IncomeType string `json:"incomeType"` // income type
	Income     string `json:"income"`     // income amount
	Asset      string `json:"asset"`      // income asset
	Info       string `json:"info"`       // extra information
	Time       int64  `json:"time"`
	TranID     string `json:"tranId"`  // transaction id
	TradeID    string `json:"tradeId"` // trade id, if existing
}

type indicator struct {
	IsLocked           bool   `json:"isLocked"`           // whether the indicator is locked
	PlannedRecoverTime int64  `json:"plannedRecoverTime"` // planned recover time
	Indicator          string `json:"indicator"`          // indicator
	Value              string `json:"value"`              // current value
	TriggerValue       string `json:"triggerValue"`       // trigger value
}

type indicators struct {
	Indicators map[string][]indicator `json:"indicators"`
	UpdateTime int64                  `json:"updateTime"`
}

type downloadId struct {
	AvgCostTimestampOfLast30d int64  `json:"avgCostTimestampOfLast30d"` // Average time taken for data download in the past 30 days
	DownloadID                string `json:"downloadId"`                // download id
}

type downloadUrl struct {
	DownloadID          string `json:"downloadId"`          // download id
	Status              string `json:"status"`              // download status: completed, processing
	URL                 string `json:"url"`                 // download url
	Notified            bool   `json:"notified"`            // ignore
	ExpirationTimestamp int64  `json:"expirationTimestamp"` // download url expiration timestamp
	IsExpired           bool   `json:"isExpired"`           // ignore
}

type TfrRow struct {
	Asset     string `json:"asset"`     // "USDT", 资产
	TransId   string `json:"transId"`   // 划转ID
	Amount    string `json:"amount"`    // 数量
	Type      int    `json:"type"`      // "1", 划转方向: 1( 现货向USDT本位合约), 2( USDT本位合约向现货), 3( 现货向币本位合约), and 4( 币本位合约向现货)
	CType     string `json:"cType"`     //
	Timestamp int64  `json:"timestamp"` // 时间戳
	CTime     string `json:"ctime"`     // 时间
	Status    string `json:"status"`    // PENDING (等待执行), CONFIRMED (成功划转), FAILED (执行失败);
}
type tfrHist struct {
	Rows  []TfrRow `json:"rows"`
	Total int      `json:"total"`
}

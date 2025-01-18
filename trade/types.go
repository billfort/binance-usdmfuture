package trade

import "github.com/billfort/binance-usdmfuture/pub"

type OrderParam struct {
	Symbol              string
	Side                pub.OrderSide
	PositionSide        pub.PositionSide
	Type                pub.OrderType
	TimeInForce         pub.TimeInForce
	Quantity            string
	ReduceOnly          string
	Price               string
	NewClientOrderId    string
	StopPrice           string
	ClosePosition       string // true, false. Close-All, used with STOP_MARKET and TAKE_PROFIT_MARKET
	ActivationPrice     string
	CallbackRate        string
	WorkingType         pub.WorkingType
	PriceProtect        string           // "TRUE", "FALSE", default "FALSE",
	NewOrderRespType    pub.ResponseType // "ACK", "RESULT", default "ACK"
	PriceMatch          pub.PriceMatch
	SelfTradePrevention pub.StpMode
	GoodTillDate        int64
	RecvWindow          int64
	Timestamp           int64
}

type orderResponse struct {
	ClientOrderId       string           `json:"clientOrderId"`
	CumQty              string           `json:"cumQty"`
	CumQuote            string           `json:"cumQuote"`
	ExecutedQty         string           `json:"executedQty"`
	OrderId             int64            `json:"orderId"`
	AvgPrice            string           `json:"avgPrice"`
	OrigQty             string           `json:"origQty"`
	Price               string           `json:"price"`
	ReduceOnly          bool             `json:"reduceOnly"`
	Side                pub.OrderSide    `json:"side"`
	PositionSide        pub.PositionSide `json:"positionSide"`
	Status              pub.OrderStatus  `json:"status"`
	StopPrice           string           `json:"stopPrice"`
	ClosePosition       bool             `json:"closePosition"`
	Symbol              string           `json:"symbol"`
	TimeInForce         pub.TimeInForce  `json:"timeInForce"`
	Type                pub.OrderType    `json:"type"`
	OrigType            pub.OrderType    `json:"origType"`
	ActivatePrice       string           `json:"activatePrice"`
	PriceRate           string           `json:"priceRate"`
	UpdateTime          int64            `json:"updateTime"`
	WorkingType         pub.WorkingType  `json:"workingType"`
	PriceProtect        bool             `json:"priceProtect"`
	PriceMatch          pub.PriceMatch   `json:"priceMatch"`
	SelfTradePrevention pub.StpMode      `json:"selfTradePreventionMode"`
	GoodTillDate        int64            `json:"goodTillDate"`
	// if error
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type modifyParam struct {
	Symbol            string         `json:"symbol"`
	Side              pub.OrderSide  `json:"side"`
	Quantity          string         `json:"quantity"`
	Price             string         `json:"price"`
	OrderId           int64          `json:"orderId"`
	OrigClientOrderId string         `json:"origClientOrderId"`
	PriceMatch        pub.PriceMatch `json:"priceMatch"`
	RecvWindow        int64          `json:"recvWindow"`
	Timestamp         int64          `json:"timestamp"`
}

type tradeInfo struct {
	Buyer           bool             `json:"buyer"`
	Commission      string           `json:"commission"`
	CommissionAsset string           `json:"commissionAsset"`
	Id              int64            `json:"id"`
	Maker           bool             `json:"maker"`
	OrderId         int64            `json:"orderId"`
	Price           string           `json:"price"`
	Qty             string           `json:"qty"`
	QuoteQty        string           `json:"quoteQty"`
	RealizedPnl     string           `json:"realizedPnl"`
	Side            pub.OrderSide    `json:"side"`
	PositionSide    pub.PositionSide `json:"positionSide"`
	Symbol          string           `json:"symbol"`
	Time            int64            `json:"time"`
}

type leverageInfo struct {
	Symbol           string `json:"symbol"`
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
}

type positionInfo struct {
	Symbol                 string `json:"symbol"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	EntryPrice             string `json:"entryPrice"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	MarkPrice              string `json:"markPrice"`
	UnRealizedProfit       string `json:"unRealizedProfit"`
	LiquidationPrice       string `json:"liquidationPrice"`
	IsolatedMargin         string `json:"isolatedMargin"`
	Notional               string `json:"notional"`
	MarginAsset            string `json:"marginAsset"`
	IsolatedWallet         string `json:"isolatedWallet"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Adl                    int    `json:"adl"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	Leverage               string `json:"leverage"`
	MaxNotionalValue       string `json:"maxNotionalValue"`
	MarginType             string `json:"marginType"`
	IsAutoAddMargin        string `json:"isAutoAddMargin"`
	UpdateTime             int64  `json:"updateTime"`
}

type adlQuantile struct {
	Symbol      string `json:"symbol"`
	AdlQuantile struct {
		Long  int `json:"LONG"`  // adl quantile for "LONG" position in hedge mode
		Short int `json:"SHORT"` // adl qauntile for "SHORT" position in hedge mode
		Both  int `json:"BOTH"`  // adl qunatile for position in one-way mode
	} `json:"adlQuantile"`
}

type positionMarginHist struct {
	Symbol       string `json:"symbol"`
	Type         int    `json:"type"`
	DeltaType    string `json:"deltaType"`
	Amount       string `json:"amount"`
	Asset        string `json:"asset"`
	Time         int64  `json:"time"`
	PositionSide string `json:"positionSide"`
}

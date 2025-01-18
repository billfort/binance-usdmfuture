package streamuserdata

type listenKey struct {
	ListenKey string `json:"listenKey"`
}

type streamHeader struct {
	EventType string `json:"e"` // listenKeyExpired
	EventTime int64  `json:"E"`
}

type AccountUpdate struct {
	UserId          int64
	EventType       string `json:"e"` // ACCOUNT_UPDATE
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Data            struct {
		EventReasonType string `json:"m"`
		Balances        []struct {
			Asset              string `json:"a"`
			WalletBalance      string `json:"wb"`
			CrossWalletBalance string `json:"cw"`
			BalanceChange      string `json:"bc"`
		} `json:"B"`
		Positions []struct {
			Symbol              string `json:"s"`
			PositionAmount      string `json:"pa"`
			EntryPrice          string `json:"ep"`
			BreakEvenPrice      string `json:"bep"`
			AccumulatedRealized string `json:"cr"`
			UnrealizedPnL       string `json:"up"`
			MarginType          string `json:"mt"`
			IsolatedWallet      string `json:"iw"`
			PositionSide        string `json:"ps"`
		} `json:"P"`
	} `json:"a"`
}

type MarginCall struct {
	UserId      int64
	EventType   string `json:"e"` // MARGIN_CALL
	EventTime   int64  `json:"E"`
	CrossWallet string `json:"cw"`
	Positions   []struct {
		Symbol            string `json:"s"`
		PositionSide      string `json:"ps"`
		PositionAmount    string `json:"pa"`
		MarginType        string `json:"mt"`
		IsolatedWallet    string `json:"iw"` // if isolated position
		MarkPrice         string `json:"mp"`
		UnrealizedPnL     string `json:"up"`
		MaintenanceMargin string `json:"mm"`
	} `json:"p"`
}

type Order struct {
	Symbol            string `json:"s"`
	ClientOrderID     string `json:"c"`
	Side              string `json:"S"`
	OrderType         string `json:"o"`
	TimeInForce       string `json:"f"`
	OriginalQuantity  string `json:"q"`
	OriginalPrice     string `json:"p"`
	AveragePrice      string `json:"ap"`
	StopPrice         string `json:"sp"`
	ExecutionType     string `json:"x"`
	OrderStatus       string `json:"X"`
	OrderID           int64  `json:"i"`
	OrderLastFilled   string `json:"l"`
	OrderFilled       string `json:"z"`
	LastFilledPrice   string `json:"L"`
	CommissionAsset   string `json:"N"`
	Commission        string `json:"n"`
	OrderTradeTime    int64  `json:"T"`
	TradeID           int64  `json:"t"`
	BidsNotional      string `json:"b"`
	AskNotional       string `json:"a"`
	IsMaker           bool   `json:"m"`
	IsReduceOnly      bool   `json:"R"`
	StopPriceWorking  string `json:"wt"`
	OriginalOrderType string `json:"ot"`
	PositionSide      string `json:"ps"`
	CloseAll          bool   `json:"cp"`
	ActivationPrice   string `json:"AP"`
	CallbackRate      string `json:"cr"`
	PriceProtection   bool   `json:"pP"`
	RealizedProfit    string `json:"rp"`
	STPMode           string `json:"V"`
	PriceMatchMode    string `json:"pm"`
	GTD               int64  `json:"gtd"`
}

type OrderTradeUpdate struct {
	UserId          int64
	EventType       string `json:"e"` // ORDER_TRADE_UPDATE
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Order           Order  `json:"o"`
}

type TradeLite struct {
	UserId          int64
	EventType       string `json:"e"` // TRADE_LITE
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Symbol          string `json:"s"`
	Quantity        string `json:"q"`
	Price           string `json:"p"`
	IsMaker         bool   `json:"m"`
	ClientOrderID   string `json:"c"`
	Side            string `json:"S"`
	LastFilledPrice string `json:"L"`
	OrderLastFilled string `json:"l"`
	TradeID         int64  `json:"t"`
	OrderID         int64  `json:"i"`
}

type AccountConfigUpdate struct {
	UserId          int64
	EventType       string `json:"e"` // ACCOUNT_CONFIG_UPDATE
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	AccountLeverage struct {
		Symbol   string `json:"s"`
		Leverage int    `json:"l"`
	} `json:"ac"`
	AccountAssetMode struct {
		MultiAssets bool `json:"j"`
	} `json:"ai"`
}

type StrategyUpdate struct {
	UserId          int64
	EventType       string `json:"e"` // STRATEGY_UPDATE
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	StrategyUpdate  struct {
		StrategyID     int    `json:"si"`
		StrategyType   string `json:"st"`
		StrategyStatus string `json:"ss"`
		Symbol         string `json:"s"`
		UpdateTime     int64  `json:"ut"`
		OpCode         int    `json:"c"`
	} `json:"su"`
}

type GridUpdate struct {
	UserId          int64
	EventType       string `json:"e"` // GRID_UPDATE
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	GridUpdate      struct {
		StrategyID        int    `json:"si"`
		StrategyType      string `json:"st"`
		StrategyStatus    string `json:"ss"`
		Symbol            string `json:"s"`
		RealizedPNL       string `json:"r"`
		UnmatchedAvgPrice string `json:"up"`
		UnmatchedQty      string `json:"uq"`
		UnmatchedFee      string `json:"uf"`
		MatchedPNL        string `json:"mp"`
		UpdateTime        int64  `json:"ut"`
	} `json:"gu"`
}

type ConditionalOrderTriggerReject struct {
	UserId          int64
	EventType       string `json:"e"` // "CONDITIONAL_ORDER_TRIGGER_REJECT"
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Order           struct {
		Symbol       string `json:"s"`
		OrderID      int64  `json:"i"`
		RejectReason string `json:"r"`
	} `json:"or"`
}

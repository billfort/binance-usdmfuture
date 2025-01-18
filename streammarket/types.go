package streammarket

type SubUnsub struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int64    `json:"id"`
}

type StreamData struct {
	Stream string                 `json:"stream"`
	Data   map[string]interface{} `json:"data"`
}

type AggTrade struct {
	EventType string `json:"e"` // "aggTrade",
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	TradeID   int64  `json:"a"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
	FirstID   int64  `json:"f"`
	LastID    int64  `json:"l"`
	Time      int64  `json:"T"`
	IsBuyer   bool   `json:"m"`
}

type MarkPriceUpdate struct {
	EventType            string `json:"e"` // "markPriceUpdate",
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	MarkPrice            string `json:"p"`
	IndexPrice           string `json:"i"`
	EstimatedSettlePrice string `json:"P"`
	FundingRate          string `json:"r"`
	NextFundingTime      int64  `json:"T"`
}

type Kline struct {
	EventType    string `json:"e"` // "kline", "continuous_kline"
	EventTime    int64  `json:"E"`
	Symbol       string `json:"s"`
	Pair         string `json:"ps"` // :"BTCUSDT",
	ContractType string `json:"ct"` // :"PERPETUAL"
	K            struct {
		StartTime           int64  `json:"t"`
		CloseTime           int64  `json:"T"`
		Symbol              string `json:"s"`
		Interval            string `json:"i"`
		FirstTradeID        int64  `json:"f"`
		LastTradeID         int64  `json:"L"`
		OpenPrice           string `json:"o"`
		ClosePrice          string `json:"c"`
		HighPrice           string `json:"h"`
		LowPrice            string `json:"l"`
		BaseAssetVolume     string `json:"v"`
		NumberOfTrades      int    `json:"n"`
		IsKlineClosed       bool   `json:"x"`
		QuoteAssetVolume    string `json:"q"`
		TakerBuyBaseVolume  string `json:"V"`
		TakerBuyQuoteVolume string `json:"Q"`
		Ignore              string `json:"B"`
	} `json:"k"`
}

type MiniTicker struct {
	EventType        string `json:"e"` // "24hrMiniTicker",
	EventTime        int64  `json:"E"`
	Symbol           string `json:"s"`
	ClosePrice       string `json:"c"`
	OpenPrice        string `json:"o"`
	HighPrice        string `json:"h"`
	LowPrice         string `json:"l"`
	BaseAssetVolume  string `json:"v"` // Total traded base asset volume
	QuoteAssetVolume string `json:"q"` // Total traded quote asset volume
}

type Ticker struct {
	EventType        string `json:"e"` // "24hrTicker",
	EventTime        int64  `json:"E"`
	Symbol           string `json:"s"`
	PriceChange      string `json:"p"`
	PriceChangePct   string `json:"P"`
	WeightedAvgPrice string `json:"w"`
	LastPrice        string `json:"c"`
	LastQty          string `json:"Q"`
	OpenPrice        string `json:"o"`
	HighPrice        string `json:"h"`
	LowPrice         string `json:"l"`
	BaseAssetVolume  string `json:"v"`
	QuoteAssetVolume string `json:"q"`
	OpenTime         int64  `json:"O"`
	CloseTime        int64  `json:"C"`
	FirstTradeID     int64  `json:"F"`
	LastTradeID      int64  `json:"L"`
	NumberOfTrades   int    `json:"n"`
}

type BookTicker struct {
	EventType string `json:"e"` // "bookTicker",
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	BidPrice  string `json:"b"` // best bid price
	BidQty    string `json:"B"` // best bid qty
	AskPrice  string `json:"a"` // best ask price
	AskQty    string `json:"A"` // best ask qty
}

type ForceOrder struct {
	EventType string `json:"e"` // "forceOrder",
	EventTime int64  `json:"E"`
	Order     struct {
		Symbol             string `json:"s"`
		Side               string `json:"S"`
		OrderType          string `json:"o"`
		TimeInForce        string `json:"f"`
		OriginalQuantity   string `json:"q"`
		Price              string `json:"p"`
		AveragePrice       string `json:"ap"`
		OrderStatus        string `json:"X"`
		OrderLastFilledQty string `json:"l"`
		OrderFilledQty     string `json:"z"`
		OrderTradeTime     int64  `json:"T"`
	} `json:"o"`
}

type DepthUpdate struct {
	EventType string     `json:"e"` // "depthUpdate",
	EventTime int64      `json:"E"`
	Symbol    string     `json:"s"`
	FirstID   int64      `json:"U"`
	LastID    int64      `json:"u"`
	PrevLast  int64      `json:"pu"` // Final update Id in last stream(ie `u` in last stream)
	Bids      [][]string `json:"b"`  // []string{Price level to be updated, qty}
	Asks      [][]string `json:"a"`  // []string{Price level to be updated, qty}
}

type CompositeIndex struct {
	EventType   string `json:"e"` // "compositeIndex",
	EventTime   int64  `json:"E"`
	Symbol      string `json:"s"`
	Price       string `json:"p"`
	Composition []struct {
		BaseAsset        string `json:"b"`
		QuoteAsset       string `json:"q"`
		Weight           string `json:"w"`
		WeightPercentage string `json:"W"`
		IndexPrice       string `json:"i"`
	} `json:"c"`
}

type ContractInfo struct {
	EventType    string `json:"e"` // "contractInfo",
	EventTime    int64  `json:"E"`
	Symbol       string `json:"s"`
	Pair         string `json:"ps"` // :"BTCUSDT",
	ContractType string `json:"ct"` // :"PERPETUAL"
	DeliveryTime int64  `json:"dt"`
	OnboardTime  int64  `json:"ot"`
	Status       string `json:"cs"`
	Brackets     []struct {
		Bracket          int     `json:"bs"`
		BracketFloor     int     `json:"bnf"`
		BracketCeiling   int     `json:"bnc"`
		MaintenanceRatio float64 `json:"mmr"`
		CalculationField int     `json:"cf"`
		MinLeverage      int     `json:"mi"`
		MaxLeverage      int     `json:"ma"`
	} `json:"bks"`
}

type AssetIndexUpdate struct {
	EventType     string `json:"e"` // "assetIndexUpdate",
	EventTime     int64  `json:"E"`
	Symbol        string `json:"s"`
	Index         string `json:"i"`
	BidBuffer     string `json:"b"`
	AskBuffer     string `json:"a"`
	BidRate       string `json:"B"`
	AskRate       string `json:"A"`
	AutoBidBuffer string `json:"q"`
	AutoAskBuffer string `json:"g"`
	AutoBidRate   string `json:"Q"`
	AutoAskRate   string `json:"G"`
}

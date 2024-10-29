package marketdata

type rateLimit struct { // API访问的限制
	Interval      string `json:"interval"`      // "MINUTE" 按照分钟计算
	IntervalNum   int    `json:"intervalNum"`   // 1 按照1分钟计算
	Limit         int    `json:"limit"`         // 2400 上限次数
	RateLimitType string `json:"rateLimitType"` // "REQUEST_WEIGHT"按照访问权重来计算 "ORDERS" 按照订单数量来计算
}

type asset struct { // 资产信息
	Asset             string `json:"asset"`             // "BUSD",
	MarginAvailabel   bool   `json:"marginAvailable"`   // true 是否可用作保证金
	AutoAssetExchange string `json:"autoAssetExchange"` // "0" 保证金资产自动兑换阈值
}

type filter struct {
	FilterType        string `json:"filterType"`        // "PRICE_FILTER" 价格限制
	MaxPrice          string `json:"maxPrice"`          // "300" 价格上限, 最大价格
	MinPrice          string `json:"minPrice"`          // "0.0001" 价格下限, 最小价格
	TickSize          string `json:"tickSize"`          // "0.0001" // 订单最小价格间隔
	MaxQty            string `json:"maxQty"`            // "10000000" 数量上限, 最大数量
	MinQty            string `json:"minQty"`            // "1" 数量下限, 最小数量
	StepSize          string `json:"stepSize"`          // "1" // 订单最小数量间隔
	Limit             int    `json:"limit"`             // 200
	Notional          string `json:"notional"`          // "1",
	MultiplierUp      string `json:"multiplierUp"`      // "1.1500" 价格上限百分比
	MultiplierDown    string `json:"multiplierDown"`    // "0.8500" 价格下限百分比
	MultiplierDecimal string `json:"multiplierDecimal"` // "4"
}

type symbol struct { // 交易对信息
	Symbol                string   `json:"symbol"`                // "BLZUSDT" 交易对
	Pair                  string   `json:"pair"`                  // "BLZUSDT" 标的交易对
	ContractType          string   `json:"contractType"`          // "PERPETUAL" 合约类型
	DeliveryDate          int64    `json:"deliveryDate"`          // 4133404800000 交割日期
	OnboardDate           int64    `json:"onboardDate"`           // 1598252400000,     // 上线日期
	Status                string   `json:"status"`                // "TRADING" 交易对状态
	MaintMarginPercent    string   `json:"maintMarginPercent"`    // "2.5000" 请忽略
	RequiredMarginPercent string   `json:"requiredMarginPercent"` // "5.0000" 请忽略
	BaseAsset             string   `json:"baseAsset"`             // "BLZ" 标的资产
	QuoteAsset            string   `json:"quoteAsset"`            // "USDT" 报价资产
	MarginAsset           string   `json:"marginAsset"`           // "USDT" 保证金资产
	PricePrecision        int      `json:"pricePrecision"`        // 5 价格小数点位数(仅作为系统精度使用，注意同tickSize 区分）
	VolPrecision          int      `json:"quantityPrecision"`     // 0 数量小数点位数(仅作为系统精度使用，注意同stepSize 区分）
	BaseAssetPrecision    int      `json:"baseAssetPrecision"`    // 8 标的资产精度
	QuotePrecision        int      `json:"quotePrecision"`        // 8 报价资产精度
	UnderlyingType        string   `json:"underlyingType"`        // "COIN",
	UnderlyingSubType     []string `json:"underlyingSubType"`     // ["STORAGE"],
	SettlePlan            int      `json:"settlePlan"`            // 0,
	TriggerProtect        string   `json:"triggerProtect"`        // "0.15" 开启"priceProtect"的条件订单的触发阈值
	Filters               []filter `json:"filters"`
	OrderType             []string `json:"OrderType"`       // 订单类型 "LIMIT",  "MARKET", "STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET", "TRAILING_STOP_MARKET" // 跟踪止损市价单
	TimeInForce           []string `json:"timeInForce"`     // 有效方式 "GTC" 成交为止, 一直有效 "IOC" 无法立即成交(吃单)的部分就撤销 "FOK" 无法全部立即成交就撤销 "GTX" 无法成为挂单方就撤销
	LiquidationFee        string   `json:"liquidationFee"`  // "0.010000", 强平费率
	MarketTakeBound       string   `json:"marketTakeBound"` // "0.30", 市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
}

type ExchInfo struct {
	// "exchangeFilters": [],
	RateLimits []rateLimit `json:"rateLimits"` // API访问的限制
	ServerTime int64       `json:"serverTime"` // 1565613908500 请忽略。如果需要获取当前系统时间，请查询接口 “GET /fapi/v1/time”
	Assets     []asset     `json:"assets"`
	Symbols    []symbol    `json:"symbols"`  // 交易对信息
	TimeZone   string      `json:"timezone"` // "UTC" // 服务器所用的时间区域
}

type orderBook struct {
	LastUpdateId int64      `json:"lastUpdateId"` // 1027024,
	E            int64      `json:"E"`            // Message output time
	T            int64      `json:"T"`            // Transaction time
	Bids         [][]string `json:"bids"`         // [][price, qty]
	Asks         [][]string `json:"asks"`         // [][price, qty]
}

type marketTrade struct {
	ID           int64  `json:"id"`           // 28457
	Price        string `json:"price"`        // "4.00000100"
	Qty          string `json:"qty"`          // "12.00000000"
	QuoteQty     string `json:"quoteQty"`     // "48.00"
	Time         int64  `json:"time"`         // 1499865549590
	IsBuyerMaker bool   `json:"isBuyerMaker"` // true / false
}

type aggTrade struct {
	AggTradeId   int64  `json:"a"` // 26129,         // Aggregate tradeId
	Price        string `json:"p"` // "0.01633102",  // Price
	Qty          string `json:"q"` // "4.70443515",  // Quantity
	FirstTradeId int64  `json:"f"` // 27781,         // First tradeId
	LastTradeId  int64  `json:"l"` // 27781,         // Last tradeId
	Timestamp    int64  `json:"T"` // 1498793709153, // Timestamp
	IsBuyerMaker bool   `json:"m"` // true,          // Was the buyer the maker?
}

type kData struct {
	OpenTime                 int64  `json:"t"` // 1499040000000,      // Open time
	Open                     string `json:"o"` // "0.01634790",       // Open
	High                     string `json:"h"` // "0.80000000",       // High
	Low                      string `json:"l"` // "0.01575800",       // Low
	Close                    string `json:"c"` // "0.01577100",       // Close
	Volume                   string `json:"v"` // "148976.11427815",  // Volume
	CloseTime                int64  `json:"T"` // 1499644799999,      // Close time
	QuoteAssetVolume         string `json:"q"` // "2434.19055334",    // Quote asset volume
	NumberOfTrades           int    `json:"n"` // 308,                // Number of trades
	TakerBuyBaseAssetVolume  string `json:"V"` // "1756.87402397",    // Taker buy base asset volume
	TakerBuyQuoteAssetVolume string `json:"Q"` // "28.46694368",      // Taker buy quote asset volume
	Ignore                   string `json:"I"` // "17928899.62484339" // Ignore.
}

type markPrice struct {
	Symbol               string `json:"symbol"`               // "BTCUSDT",
	MarkPrice            string `json:"markPrice"`            // "11793.63104562",	// mark price
	IndexPrice           string `json:"indexPrice"`           // "11781.80495970",	// index price
	EstimatedSettlePrice string `json:"estimatedSettlePrice"` // "11781.16138815", // Estimated Settle Price, only useful in the last hour before the settlement starts.
	LastFundingRate      string `json:"lastFundingRate"`      // "0.00038246",  // This is the Latest funding rate
	NextFundingTime      int64  `json:"nextFundingTime"`      // 1597392000000,
	InterestRate         string `json:"interestRate"`         // "0.00010000",
	Time                 int64  `json:"time"`                 // 1597370495002
}

type fundingRate struct {
	Symbol      string `json:"symbol"`      // "BTCUSDT",
	FundingRate string `json:"fundingRate"` // "-0.03750000",
	FundingTime int64  `json:"fundingTime"` // 1570608000000,
	MarkPrice   string `json:"markPrice"`   // "34287.54619963"   // mark price associated with a particular funding fee charge
}

type fundingInfo struct {
	Symbol                   string `json:"symbol"`                   // "BTCUSDT",
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`   // "0.02500000",
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"` // "-0.02500000",
	FundingIntervalHours     int    `json:"fundingIntervalHours"`     // 8,
	Disclaimer               bool   `json:"disclaimer"`               // false, ignore
}

type ticker24hr struct {
	Symbol             string `json:"symbol"`             // "BTCUSDT",
	PriceChange        string `json:"priceChange"`        // "-94.99999800",
	PriceChangePercent string `json:"priceChangePercent"` // "-95.960",
	WeightedAvgPrice   string `json:"weightedAvgPrice"`   // "0.29628482",
	LastPrice          string `json:"lastPrice"`          // "4.00000200",
	LastQty            string `json:"lastQty"`            // "200.00000000",
	OpenPrice          string `json:"openPrice"`          // "99.00000000",
	HighPrice          string `json:"highPrice"`          // "100.00000000",
	LowPrice           string `json:"lowPrice"`           // "0.10000000",
	Volume             string `json:"volume"`             // "8913.30000000",
	QuoteVolume        string `json:"quoteVolume"`        // "15.30000000",
	OpenTime           int64  `json:"openTime"`           // 1499783499040,
	CloseTime          int64  `json:"closeTime"`          // 1499869899040,
	FirstId            int64  `json:"firstId"`            // 28385,   // First tradeId
	LastId             int64  `json:"lastId"`             // 28460,    // Last tradeId
	Count              int    `json:"count"`              // 76         // Trade count
}

type tickerPrice struct {
	Symbol string `json:"symbol"` // "BTCUSDT",
	Price  string `json:"price"`  // "6000.01",
	Time   int64  `json:"time"`   // 1589437530011   // Transaction time
}

type bookTicker struct {
	Symbol   string `json:"symbol"`   // "BTCUSDT",
	BidPrice string `json:"bidPrice"` // "4.00000000",
	BidQty   string `json:"bidQty"`   // "431.00000000",
	AskPrice string `json:"askPrice"` // "4.00000200",
	AskQty   string `json:"askQty"`   // "9.00000000",
	Time     int64  `json:"time"`     // 1589437530011   // Transaction time
}

type deliveryPrice struct {
	DeliTime  int64   `json:"deliveryTime"`  // 1695945600000,
	DeliPrice float64 `json:"deliveryPrice"` // 27103.00000000
}

type openInterest struct {
	Symbol       string `json:"symbol"`       // "BTCUSDT",
	OpenInterest string `json:"openInterest"` // "10659.509",
	Time         int64  `json:"time"`         // 1589437530011   // Transaction time
}

type openInterestHist struct {
	Symbol               string `json:"symbol"`               // "BTCUSDT",
	SumOpenInterest      string `json:"sumOpenInterest"`      // "10659.509",
	SumOpenInterestValue string `json:"sumOpenInterestValue"` // "10659.509",
	Timestamp            int64  `json:"timestamp"`            // 1589437530011
}

type longShortRatio struct {
	Symbol         string `json:"symbol"`         // "BTCUSDT",
	LongShortRatio string `json:"longShortRatio"` // "1.4342",// long/short position ratio of top traders
	LongAccount    string `json:"longAccount"`    // "0.5891", // long positions ratio of top traders
	ShortAccount   string `json:"shortAccount"`   // "0.4108", // short positions ratio of top traders
	Timestamp      int64  `json:"timestamp"`      // 1583139600000
}

type takerLongShortRatio struct {
	BuySellRatio string `json:"buySellRatio"` // "1.5586",// buy/sell volume ratio
	BuyVol       string `json:"buyVol"`       // "387.3300", // buy volume
	SellVol      string `json:"sellVol"`      // "248.5030", // sell volume
	Timestamp    int64  `json:"timestamp"`    // 1585614900000
}

type indexMember struct {
	BaseAsset          string `json:"baseAsset"`          // "BTC",
	QuoteAsset         string `json:"quoteAsset"`         // "USDT",
	WeightInQuantity   string `json:"weightInQuantity"`   // "1.04406228",
	WeightInPercentage string `json:"weightInPercentage"` // "0.02783900"
}

type indexInfo struct {
	Symbol        string        `json:"symbol"`        // "DEFIUSDT",
	Time          int64         `json:"time"`          // 1589437530011,    // Current time
	Component     string        `json:"component"`     // "baseAsset", // Component asset
	BaseAssetList []indexMember `json:"baseAssetList"` //
}

type assetIndex struct {
	Symbol                string `json:"symbol"`                // "ADAUSD",
	Time                  int64  `json:"time"`                  // 1635740268004,
	Index                 string `json:"index"`                 // "1.92957370",
	BidBuffer             string `json:"bidBuffer"`             // "0.10000000",
	AskBuffer             string `json:"askBuffer"`             // "0.10000000",
	BidRate               string `json:"bidRate"`               // "1.73661633",
	AskRate               string `json:"askRate"`               // "2.12253107",
	AutoExchangeBidBuffer string `json:"autoExchangeBidBuffer"` // "0.05000000",
	AutoExchangeAskBuffer string `json:"autoExchangeAskBuffer"` // "0.05000000",
	AutoExchangeBidRate   string `json:"autoExchangeBidRate"`   // "1.83309501",
	AutoExchangeAskRate   string `json:"autoExchangeAskRate"`   // "2.02605238"
}

type constituent struct {
	Symbol   string `json:"symbol"`   // "BTCUSDT",
	Exchange string `json:"exchange"` // "Binance",
}

type indexConstituents struct {
	Symbol      string        `json:"symbol"`       // "BTCUSDT",
	Time        int64         `json:"time"`         // 1697421272043,
	Constiuents []constituent `json:"constituents"` //
}

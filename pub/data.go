package pub

type Key struct {
	UserId    int64 // any id to identify this user key paire
	ApiKey    string
	SecretKey string
}

type ErrMsg struct {
	Code int
	Msg  string
}

/*
// webstocket account update data
type BalanceUpdate struct {
	Symbol        string
	WalletBalance float64
	Change        float64
}
type PosiUpdate struct {
	Symbol           string
	PositionSide     string // BOTH, LONG, SHORT
	Amouont          float64
	OpenPrice        float64
	RelizedProfit    float64
	UnrealizedProfit float64
	MarginType       string
	IsolatedMargin   float64
}
type AccData struct {
	UpdateReason    string // 事件推出原因
	BalanceUpdates  []BalanceUpdate
	PositionUpdates []PosiUpdate
}
type AccountUpdateData struct {
	UserId    int64
	Event     string // ACCOUNT_UPDATE
	EventTime int64  // event time
	MatchTime int64  // 撮合 time
	AccUpdate AccData
}

// websocket order update data
type OrderData struct {
	Symbol            string
	ClientOrderId     string  // 客户端自定订单ID, 特殊的自定义订单ID: "autoclose-"开头的字符串: 系统强平订单, "adl_autoclose": ADL自动减仓订单, "settlement_autoclose-": 下架或交割的结算订单
	Dir               string  // SELL, BUY
	OrderType         string  // 订单类型, Market, TRAILING_STOP_MARKET
	ValidType         string  // 有效方式, GTC, ...
	OrderVol          float64 // 订单原始数量
	OrderPrice        float64 // 订单原始价格
	SuccPrice         float64 // 订单平均价格
	TriggerPrice      float64 // 条件订单触发价格，对追踪止损单无效
	ExecType          string  // 本次事件的具体执行类型
	OrderStatus       string  // 订单的当前状态  NEW, FILLED, PARTIALLY_FILLED, CANCELED, EXPIRED, NEW_INSURANCE 风险保障基金(强平), NEW_ADL 自动减仓序列(强平)
	OrderId           int64   // 订单ID
	LastSuccVol       float64 // 订单末次成交量
	SuccVol           float64 // 订单累计已成交量
	LastSuccPrice     float64 // 订单末次成交价格
	FeeToken          string  // 手续费资产类型, USDT
	Fee               float64 // 手续费数量
	SuccTime          int64   // unixt time, milli-second
	SuccId            int64   // 成api
	BuyNetVal         float64 // 买单净值
	SellNetVal        float64 // 卖单净值
	Maker             bool    // 该成交是作为挂单成交吗？
	ReduceOnly        bool    // 是否是只减仓单
	TriggerType       string  // 触发价类型
	OriginOrderType   string  // 原始订单类型
	Position          string  // 持仓方向 LONG, SHORT
	TriggerOrder      bool    // 是否为触发平仓单; 仅在条件订单情况下会推送此字段
	TrailLossPrice    float64 // 追踪止损激活价格, 仅在追踪止损单时会推送此字段
	TrailLossRetrieve float64 // 追踪止损回调比例, 仅在追踪止损单时会推送此字段
	Profit            float64 // 该交易实现盈亏
}
type OrderUpdateData struct {
	UserId    int64
	Event     string // ACCOUNT_UPDATE
	EventTime int64  // event time
	MatchTime int64  // 撮合 time
	Data      struct {
		Symbol            string
		ClientOrderId     string  // 客户端自定订单ID, 特殊的自定义订单ID: "autoclose-"开头的字符串: 系统强平订单, "adl_autoclose": ADL自动减仓订单, "settlement_autoclose-": 下架或交割的结算订单
		Dir               string  // SELL, BUY
		OrderType         string  // 订单类型, Market, TRAILING_STOP_MARKET
		ValidType         string  // 有效方式, GTC, ...
		OrderVol          float64 // 订单原始数量
		OrderPrice        float64 // 订单原始价格
		SuccPrice         float64 // 订单平均价格
		TriggerPrice      float64 // 条件订单触发价格，对追踪止损单无效
		ExecType          string  // 本次事件的具体执行类型
		OrderStatus       string  // 订单的当前状态  NEW, FILLED, PARTIALLY_FILLED, CANCELED, EXPIRED, NEW_INSURANCE 风险保障基金(强平), NEW_ADL 自动减仓序列(强平)
		OrderId           int64   // 订单ID
		LastSuccVol       float64 // 订单末次成交量
		SuccVol           float64 // 订单累计已成交量
		LastSuccPrice     float64 // 订单末次成交价格
		FeeToken          string  // 手续费资产类型, USDT
		Fee               float64 // 手续费数量
		SuccTime          int64   // unixt time, milli-second
		SuccId            int64   // 成交ID
		BuyNetVal         float64 // 买单净值
		SellNetVal        float64 // 卖单净值
		Maker             bool    // 该成交是作为挂单成交吗？
		ReduceOnly        bool    // 是否是只减仓单
		TriggerType       string  // 触发价类型
		OriginOrderType   string  // 原始订单类型
		Position          string  // 持仓方向 LONG, SHORT
		TriggerOrder      bool    // 是否为触发平仓单; 仅在条件订单情况下会推送此字段
		TrailLossPrice    float64 // 追踪止损激活价格, 仅在追踪止损单时会推送此字段
		TrailLossRetrieve float64 // 追踪止损回调比例, 仅在追踪止损单时会推送此字段
		Profit            float64 // 该交易实现盈亏
	}
}

*/

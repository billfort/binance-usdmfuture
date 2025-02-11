package pub

const (
	futureBaseUrl = "https://fapi.binance.com"
	spotBaseUrl   = "https://api.binance.com" // api1...., api2...., api3..., api4....
	futureWssUrl  = "wss://fstream.binance.com"

	recvWindow = "5000"
	WsChanLen  = 128 // chan lengh for websocket message
)

var TestKey = &Key{
	UserId:    123456,
	ApiKey:    "Y3H9yFQTFh55vnlowVrQHBPMRALEfmAfa8Q3yj62HLNolUu9O7TVtvhDio62tfgD", // for testing
	SecretKey: "FiAAghn8vzGYZ6xYGDpE3Dw8uknzrIrRXJMLjiEiYKBJdHMt6xmiJCJvVwuUXuYD", // for testing
}

type ContractType string

const (
	CT_Perpetual           ContractType = "PERPETUAL"
	CT_CurrentMonth        ContractType = "CURRENT_MONTH"
	CT_NextMonth           ContractType = "NEXT_MONTH"
	CT_CurrentQuarter      ContractType = "CURRENT_QUARTER"
	CT_NextQuarter         ContractType = "NEXT_QUARTER"
	CT_PerpetualDelivering ContractType = "PERPETUAL_DELIVERING"
)

type ContractStatus string

const (
	CS_PendingTrading ContractStatus = "PENDING_TRADING"
	CS_Trading        ContractStatus = "TRADING"
	CS_PreDelivering  ContractStatus = "PRE_DELIVERING"
	CS_Delivering     ContractStatus = "DELIVERING"
	CS_Delivered      ContractStatus = "DELIVERED"
	CS_PreSettle      ContractStatus = "PRE_SETTLE"
	CS_Settling       ContractStatus = "SETTLING"
	CS_Close          ContractStatus = "CLOSE"
)

type OrderStatus string

const (
	OS_New             OrderStatus = "NEW"
	OS_PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OS_Filled          OrderStatus = "FILLED"
	OS_Canceled        OrderStatus = "CANCELED"
	OS_Rejected        OrderStatus = "REJECTED"
	OS_Expired         OrderStatus = "EXPIRED"
)

type OrderType string

const (
	OT_Limit              OrderType = "LIMIT"
	OT_Market             OrderType = "MARKET"
	OT_Stop               OrderType = "STOP"
	OT_StopMarket         OrderType = "STOP_MARKET"
	OT_TakeProfit         OrderType = "TAKE_PROFIT"
	OT_TakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OT_TrailingStopMarket OrderType = "TRAILING_STOP_MARKET"
)

type OrderSide string

const (
	OS_Buy  OrderSide = "BUY"
	OS_Sell OrderSide = "SELL"
)

type PositionSide string

const (
	PS_Long  PositionSide = "LONG"
	PS_Short PositionSide = "SHORT"
	PS_Both  PositionSide = "BOTH"
)

type TimeInForce string

const (
	TIF_GTC TimeInForce = "GTC" // Good Till Cancel(GTC order valitidy is 1 year from placement)
	TIF_IOC TimeInForce = "IOC" // Immediate or Cancel
	TIF_FOK TimeInForce = "FOK" // Fill or Kill
	TIF_GTX TimeInForce = "GTX" // Good Till Crossing (Post Only)
	TIF_GTD TimeInForce = "GTD" // Good Till Date
)

type WorkingType string

const (
	WT_MarkPrice     WorkingType = "MARK_PRICE"
	WT_ContractPrice WorkingType = "CONTRACT_PRICE"
)

type ResponseType string

const (
	RT_Result ResponseType = "RESULT"
	RT_Ack    ResponseType = "ACK"
)

type KlineInterval string

const (
	KI_Minute1  KlineInterval = "1m"
	KI_Minute3  KlineInterval = "3m"
	KI_Minute5  KlineInterval = "5m"
	KI_Minute15 KlineInterval = "15m"
	KI_Minute30 KlineInterval = "30m"
	KI_Hour1    KlineInterval = "1h"
	KI_Hour2    KlineInterval = "2h"
	KI_Hour4    KlineInterval = "4h"
	KI_Hour6    KlineInterval = "6h"
	KI_Hour8    KlineInterval = "8h"
	KI_Hour12   KlineInterval = "12h"
	KI_Day1     KlineInterval = "1d"
	KI_Day3     KlineInterval = "3d"
	KI_Week1    KlineInterval = "1w"
	KI_Month1   KlineInterval = "1M"
)

type StpMode string // self trade prevention mode
const (
	STP_None        StpMode = "NONE"
	STP_ExpireTaker StpMode = "EXPIRE_TAKER"
	STP_ExpireBoth  StpMode = "EXPIRE_BOTH"
	STP_ExpireMaker StpMode = "EXPIRE_MAKER"
)

type PriceMatch string

const (
	PM_None        PriceMatch = "NONE"
	PM_Opponent    PriceMatch = "OPPONENT"    // counterparty best price
	PM_Opponent_5  PriceMatch = "OPPONENT_5"  // the 5th best price from the counterparty
	PM_Opponent_10 PriceMatch = "OPPONENT_10" // the 10th best price from the counterparty
	PM_Opponent_20 PriceMatch = "OPPONENT_20" // the 20th best price from the counterparty
	PM_Queue       PriceMatch = "QUEUE"       // the best price on the same side of the order book
	PM_Queue_5     PriceMatch = "QUEUE_5"     // the 5th best price on the same side of the order book
	PM_Queue_10    PriceMatch = "QUEUE_10"    // the 10th best price on the same side of the order book
	PM_Queue_20    PriceMatch = "QUEUE_20"    // the 20th best price on the same side of the order book
)

type MarginType string

const (
	MT_Isolated MarginType = "ISOLATED"
	MT_Cross    MarginType = "CROSSED"
)

type IncomeType string

const (
	IT_Transfer                 IncomeType = "TRANSFER"
	IT_WelcomeBonus             IncomeType = "WELCOME_BONUS"
	IT_RealizedPnl              IncomeType = "REALIZED_PNL"
	IT_FundingFee               IncomeType = "FUNDING_FEE"
	IT_Commission               IncomeType = "COMMISSION"
	IT_InsuranceClear           IncomeType = "INSURANCE_CLEAR"
	IT_ReferralKickback         IncomeType = "REFERRAL_KICKBACK"
	IT_CommissionRebate         IncomeType = "COMMISSION_REBATE"
	IT_ApiRebate                IncomeType = "API_REBATE"
	IT_ContestReward            IncomeType = "CONTEST_REWARD"
	IT_CrossCollateralTransfer  IncomeType = "CROSS_COLLATERAL_TRANSFER"
	IT_OptionsPremiumFee        IncomeType = "OPTIONS_PREMIUM_FEE"
	IT_OptionsSettleProfit      IncomeType = "OPTIONS_SETTLE_PROFIT"
	IT_InternalTransfer         IncomeType = "INTERNAL_TRANSFER"
	IT_AutoExchange             IncomeType = "AUTO_EXCHANGE"
	IT_DeliveredSettelment      IncomeType = "DELIVERED_SETTELMENT"
	IT_CoinSwapDeposit          IncomeType = "COIN_SWAP_DEPOSIT"
	IT_CoinSwapWithdraw         IncomeType = "COIN_SWAP_WITHDRAW"
	IT_PositionLimitIncreaseFee IncomeType = "POSITION_LIMIT_INCREASE_FEE"
)

package okxv5

type Category string

const (
	Spot    Category = "SPOT"
	Margin  Category = "MARGIN"
	Swap    Category = "SWAP"
	Futures Category = "FUTURES"
	Option  Category = "OPTION"
)

type MarginMode string

const (
	Cross    MarginMode = "cross"
	Isolated MarginMode = "isolated"
)

type TradeMode string

const (
	CrossMode    TradeMode = "cross"
	IsolatedMode TradeMode = "isolated"
	CashMode     TradeMode = "cash"
)

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type PosSide string

const (
	LongSide  PosSide = "long"
	ShortSide PosSide = "short"
	NetSide   PosSide = "net"
)

type OrderType string

const (
	LimitType       OrderType = "limit"
	MarketType      OrderType = "market"
	PostOnly        OrderType = "post_only"
	Fok             OrderType = "fok"
	Ioc             OrderType = "ioc"
	OptimalLimitIoc OrderType = "optimal_limit_ioc"
	Mmp             OrderType = "mmp"
	MmpAndPostOnly  OrderType = "mmp_and_post_only"
)

type OrderState string

const (
	Canceled        OrderState = "canceled"
	Live            OrderState = "live"
	PartiallyFilled OrderState = "partially_filled"
	Filled          OrderState = "filled"
	MmpCanceled     OrderState = "mmp_canceled"
)

type TriggerPriceType string

const (
	Last  TriggerPriceType = "last"
	Index TriggerPriceType = "index"
	Mark  TriggerPriceType = "mark"
)

type OrderCategory string

const (
	Normal             OrderCategory = "normal"
	Twap               OrderCategory = "twap"
	Adl                OrderCategory = "adl"
	FullLiquidation    OrderCategory = "full_liquidation"
	PartialLiquidation OrderCategory = "partial_liquidation"
	Delivery           OrderCategory = "delivery"
	Ddh                OrderCategory = "ddh" //Delta dynamic hedge
)

type MarginType string

const (
	Manual     MarginType = "manual"
	AutoBorrow MarginType = "auto_borrow"
	AutoRepay  MarginType = "auto_repay"
)

// https://www.okx.com/docs-v5/en/#order-book-trading-market-data-ws-order-book-channel

type OrderbookType string

const (
	Books  OrderbookType = "books"
	Books5 OrderbookType = "books5"
	Bbotbt OrderbookType = "bbo-tbt"
	//Log in required:
	//Booksl2tbt   OrderbookType = "books-l2-tbt"
	//Books50l2tbt OrderbookType = "books50-l2-tbt"
)

type ContractIsolatedMarginTradingSettings string

const (
	Auto    ContractIsolatedMarginTradingSettings = "automatic"
	Manualy ContractIsolatedMarginTradingSettings = "autonomy"
)

type GreekType string

const (
	Pa GreekType = "PA"
	Bs GreekType = "BS"
)

type MgnIsoMode string

const (
	Automatic   MgnIsoMode = "automatic"
	QuickMargin MgnIsoMode = "quick_margin"
)

type PositionMode string

const (
	LongShortMode PositionMode = "long_short_mode"
	NetMode       PositionMode = "net_mode"
)

type CandleStickType string

const (
	Candle3M     CandleStickType = "candle3M"
	Candle1M     CandleStickType = "candle1M"
	Candle1W     CandleStickType = "candle1W"
	Candle1D     CandleStickType = "candle1D"
	Candle2D     CandleStickType = "candle2D"
	Candle3D     CandleStickType = "candle3D"
	Candle5D     CandleStickType = "candle5D"
	Candle12H    CandleStickType = "candle12H"
	Candle6H     CandleStickType = "candle6H"
	Candle4H     CandleStickType = "candle4H"
	Candle2h     CandleStickType = "candle2H"
	Candle1H     CandleStickType = "candle1H"
	Candle30m    CandleStickType = "candle30m"
	Candle15m    CandleStickType = "candle15m"
	Candle5m     CandleStickType = "candle5m"
	Candle3m     CandleStickType = "candle3m"
	Candle1m     CandleStickType = "candle1m"
	Candle1s     CandleStickType = "candle1s"
	Candle3Mutc  CandleStickType = "candle3Mutc"
	Candle1Mutc  CandleStickType = "candle1Mutc"
	Candle1Wutc  CandleStickType = "candle1Wutc"
	Candle1Dutc  CandleStickType = "candle1Dutc"
	Candle2Dutc  CandleStickType = "candle2Dutc"
	Candle3Dutc  CandleStickType = "candle3Dutc"
	Candle5Dutc  CandleStickType = "candle5Dutc"
	Candle12Hutc CandleStickType = "candle12Hutc"
	Candle6Hutc  CandleStickType = "candle6Hutc"
)

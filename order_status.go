package binance

// OrderStatus represents order status enum.
type OrderStatus string

// OrderType represents order type enum.
type OrderType string

// OrderSide represents order side enum.
type OrderSide string

// OrderResponse represents order response type enum.
type OrderResponse string

var (
	StatusNew             = OrderStatus("NEW")
	StatusPartiallyFilled = OrderStatus("PARTIALLY_FILLED")
	StatusFilled          = OrderStatus("FILLED")
	StatusCancelled       = OrderStatus("CANCELED")
	StatusPendingCancel   = OrderStatus("PENDING_CANCEL")
	StatusRejected        = OrderStatus("REJECTED")
	StatusExpired         = OrderStatus("EXPIRED")

	TypeLimit  = OrderType("LIMIT")
	TypeMarket = OrderType("MARKET")

	SideBuy  = OrderSide("BUY")
	SideSell = OrderSide("SELL")

	RespTypeACK    = OrderResponse("ACK")
	RespTypeResult = OrderResponse("RESULT")
	RespTypeFull   = OrderResponse("FULL")
)

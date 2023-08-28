package utils

type CreateOrderRequest struct {
	UserId     string `json:"user_id"`
	Item       string `json:"item"`
	Quantity   int    `json:"quantity"`
	TotalPrice int64  `json:"total_price"`
}

type CreateOrderResponse struct {
	UserId      string `json:"user_id"`
	OrderID     string `json:"order_id"`
	TotalPrice  int64  `json:"total_price"`
	CreateOrder bool   `json:"create_order"`
}

type CreateOrderEvent struct {
	OrderID    string `json:"order_id"`
	TotalPrice int64  `json:"total_price"`
}

type ProcessPaymentData struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
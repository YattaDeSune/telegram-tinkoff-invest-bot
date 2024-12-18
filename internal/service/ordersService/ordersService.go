package ordersservice

import (
	service "telegram_invest_bot/internal/service/client"

	"github.com/tinkoff/invest-api-go-sdk/investgo"
)

type OrdersService *investgo.OrdersServiceClient

func NewOrdersServiceClient(client service.Client) OrdersService {
	return client.Client.NewOrdersServiceClient()
}

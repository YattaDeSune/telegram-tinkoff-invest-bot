package operationsservice

import (
	service "telegram_invest_bot/internal/service/client"

	"github.com/tinkoff/invest-api-go-sdk/investgo"
)

type OperationsService *investgo.OperationsServiceClient

func NewOperationsServiceClient(client service.Client) OperationsService {
	return client.Client.NewOperationsServiceClient()
}

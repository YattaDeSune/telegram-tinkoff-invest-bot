package usersservice

import (
	service "telegram_invest_bot/internal/service/client"

	"github.com/tinkoff/invest-api-go-sdk/investgo"
)

type UsersService *investgo.UsersServiceClient

func NewUsersServiceClient(client service.Client) UsersService {
	return client.Client.NewUsersServiceClient()
}

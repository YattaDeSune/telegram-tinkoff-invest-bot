package service

import (
	"context"

	"github.com/tinkoff/invest-api-go-sdk/investgo"
)

type Client struct {
	Client *investgo.Client
}

func NewClient(ctx context.Context, cfg investgo.Config, logger investgo.Logger) *Client {
	// создаем клиента для investAPI, он позволяет создавать нужные сервисы и уже через них вызывать нужные методы
	c, err := investgo.NewClient(ctx, cfg, logger)
	if err != nil {
		logger.Fatalf("client creating error %v", err.Error())
	}
	defer func() {
		logger.Infof("closing client connection")
		err := c.Stop()
		if err != nil {
			logger.Errorf("client shutdown error %v", err.Error())
		}
	}()

	client := &Client{Client: c}

	return client
}

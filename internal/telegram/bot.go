package telegram

import (
	"context"
	"os/signal"
	"syscall"
	"telegram_invest_bot/internal/config"
	"telegram_invest_bot/internal/logger"
	service "telegram_invest_bot/internal/service/client"
	operationsservice "telegram_invest_bot/internal/service/operationsService"
	ordersservice "telegram_invest_bot/internal/service/ordersService"
	usersservice "telegram_invest_bot/internal/service/usersService"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot    *tgbotapi.BotAPI
	Logger logger.Logger
	Cfg    config.Config
}

func NewBot() (*Bot, error) {
	logger, err := logger.NewLogger()
	if err != nil {
		logger.Errorf("logger create error %v:", err.Error())
		return &Bot{}, ErrLoggerError
	}

	cfg, err := config.New()
	if err != nil {
		logger.Errorf("config loading error %v:", err.Error())
		return &Bot{}, ErrCfgError
	}

	b, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		logger.Errorf("bot create error %v:", err.Error())
		return &Bot{}, ErrBotError
	}

	bot := &Bot{
		Bot:    b,
		Logger: logger,
		Cfg:    *cfg,
	}

	return bot, nil
}

// логика

func (b *Bot) RunBot() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	client := service.NewClient(ctx, b.Cfg.Tinkoff, b.Logger)

	usersService := usersservice.NewUsersServiceClient(*client)
	operationsService := operationsservice.NewOperationsServiceClient(*client)
	ordersService := ordersservice.NewOrdersServiceClient(*client)

	return nil
}

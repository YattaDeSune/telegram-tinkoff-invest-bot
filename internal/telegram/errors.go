package telegram

import "errors"

var (
	// create errors
	ErrCfgError    = errors.New("config loading error")
	ErrLoggerError = errors.New("logger create error")
	ErrBotError    = errors.New("bot create error")

	// run errors
)

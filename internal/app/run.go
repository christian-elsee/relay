package app

import (
	"context"
	"log/slog"
)

type App struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *App {
	return &App{
		logger: logger,
	}
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info("Enter", "ref", "app/internal/run.go")
	defer a.logger.Info("Exit", "ref", "app/internal/run.go")

	// Start services, run work, or wait for shutdown here.

	return nil
}

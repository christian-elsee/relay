package main

import (
	"os"
	"context"
	"log/slog"

	"github.com/scratch/golang-hello-world/internal/app"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(
		os.Stderr, nil,
	))
	application := app.New(logger)

	if err := application.Run(context.Background()); err != nil {
		logger.Error("Application run failure")
		os.Exit(1)
	}

}


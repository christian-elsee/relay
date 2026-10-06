package main

import "os"
import "log/slog"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	slog.Info("Hello World", "author", "christian")

}

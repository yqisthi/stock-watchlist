package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger() {
	Log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
}

func Error(err error, msg ...string) {
	message := err.Error()

	if len(msg) > 0 {
		message = msg[0]
	}
	Log.Error(message, "error", err)
}

func Info(msg string) {
	Log.Info(msg)
}
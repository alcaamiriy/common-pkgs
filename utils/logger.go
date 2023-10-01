package utils

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"runtime/debug"
	"time"
)

var Logger zerolog.Logger

func InitLogger() {
	buildInfo, _ := debug.ReadBuildInfo()

	var writers []io.Writer
	writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	if EnvConfig.Environment == "production" {

		file, err := os.OpenFile(
			"app.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			panic(err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		writers = append(writers, file)
	}
	mw := io.MultiWriter(writers...)

	Logger = zerolog.New(mw).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		//Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}

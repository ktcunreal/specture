package main

import (
	"errors"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
	"specture/api"
	"specture/internal/config"
)

func main() {
	config.LoadConfig()
	InitLogger(config.GetLoglevel())
	router := api.InitializeMainRouter()

	log.Info("Server started")
	log.Debug(config.GetGlobalConfig())
	log.Infof("You should be able to get your QR code at %s/qr/%s\n", config.GetBaseUrl(),config.GetPresharedKey())
	http.ListenAndServe(config.GetListenAddress(), router)
}

func InitLogger(conf string) {
	var loglevel log.Level
	switch conf {
	case "debug":
		loglevel = log.DebugLevel
	case "info":
		loglevel = log.InfoLevel
	case "warn":
		loglevel = log.WarnLevel
	case "error":
		loglevel = log.ErrorLevel
	case "fatal":
		loglevel = log.FatalLevel
	default:
		panic(errors.New("Invalid log level!"))
	}
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level:           loglevel,
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	log.SetDefault(logger)
}

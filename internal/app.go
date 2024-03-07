package internal

import (
	"context"
	"github.com/ttdung/du/config"
	"github.com/ttdung/du/logger"

	"os"
	"os/signal"
	"syscall"
)

// App is the main application of the project.
type App struct {
	log logger.Logger
	//listener   *listener.Listener
	//whm        *webhook.WebHookManager
	//httpServer *api.HTTPServer
}

// NewApp creates a new main application.
func NewApp(cfg *config.Config) (*App, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, err
	}

	// Setup logger
	var log logger.Logger
	if cfg.Logger.Color {
		log = logger.NewZeroLoggerWithColor(cfg.Logger.LogPath, "APP")
	} else {
		log = logger.NewZeroLogger(cfg.Logger.LogPath, "APP")
	}
	log.SetLogLevel(logger.LogLevel(cfg.Logger.Level))

	return &App{
		log: log.WithPrefix("App"),
	}, nil
}

func (app *App) Start(ctx context.Context) {

	sysErr := make(chan os.Signal, 1)
	signal.Notify(sysErr,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case _ = <-ctx.Done():
		app.log.Infof("terminating due to ctx.Done")
		//app.whm.Stop()
		return
	case sig := <-sysErr:
		app.log.Infof("terminating got `[%v]` signal", sig)
		return
	}
}

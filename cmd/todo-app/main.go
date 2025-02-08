package main

import (
	"context"
	logBase "log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mnsavag/fast-todo-app.git/internal/lib/log"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		logBase.Fatalln(err)
	}

	logger, err := log.NewLogger(cfg.Logger)
	if err != nil {
		logBase.Fatalln(err)
	}

	app := todoapp.NewApp(cfg, logger)
	if err := app.Init(ctx); err != nil {
		logBase.Fatalln(err)
	}

	if err := app.Start(ctx); err != nil {
		logBase.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	if err := app.Stop(); err != nil {
		logBase.Fatalln(err)
	}
}

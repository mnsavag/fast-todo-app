package todoapp

import (
	"context"
	"fmt"

	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/api"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/config"
	appgrpc "github.com/mnsavag/fast-todo-app.git/internal/todoapp/grpc"
	apphttp "github.com/mnsavag/fast-todo-app.git/internal/todoapp/http"
	"github.com/mnsavag/fast-todo-app.git/internal/todoapp/service"
	"golang.org/x/sync/errgroup"

	"github.com/sirupsen/logrus"
)

// App struct, NewApp, Start(), Stop(), dbconn, repo, controller, service)

type App struct {
	httpServer *apphttp.Server
	grpcServer *appgrpc.Server
	cfg        config.Config
	logger     *logrus.Logger
}

func NewApp(cfg config.Config, logger *logrus.Logger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Init(ctx context.Context) error {

	todoAppService := service.NewService(a.logger)

	// http
	a.httpServer = apphttp.NewServer(a.cfg, a.logger)
	err := a.httpServer.Register(ctx)
	if err != nil {
		return fmt.Errorf("cant register http server: %w", err)
	}

	// grpc
	a.grpcServer = appgrpc.NewServer(a.cfg, a.logger)
	apiServer := api.NewServer(a.logger, todoAppService)
	a.grpcServer.RegisterTodoAppServiceV1(apiServer)

	return nil
}

func (a *App) Start(ctx context.Context) error {
	errG, _ := errgroup.WithContext(ctx)
	errG.Go(func() error {
		return a.grpcServer.Start()
	})
	errG.Go(func() error {
		return a.httpServer.Start()
	})
	return nil
}

func (a *App) Stop() error {
	a.grpcServer.Stop()
	if err := a.httpServer.Stop(); err != nil {
		return fmt.Errorf("cannot stop http server, :%w", err)
	}
	a.logger.Info("gracefull stop")
	return nil
}

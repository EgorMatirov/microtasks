package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/EgorMatirov/microtasks/internal/domain"
	ginrouter "github.com/EgorMatirov/microtasks/internal/implemention/gin"
	"github.com/EgorMatirov/microtasks/internal/infrastructure"
	"github.com/EgorMatirov/microtasks/internal/usecase"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//nolint // it's main func
func main() {
	config, err := infrastructure.NewConfigFromEnv()
	if err != nil {
		panic(fmt.Sprintf("can't build config: %v", err))
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info(
		"starting service",
		zap.String("name", infrastructure.AppName),
		zap.String("version", infrastructure.AppTag),
	)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	dbPool, err := infrastructure.InitPGPool(
		ctx,
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
		config.Database.MaxConn,
	)

	defer dbPool.Close()

	crudRepo := domain.NewRepo(logger)
	handlerConstructor := usecase.HandlerConstructor{
		Crud: crudRepo,
	}

	ucHandler := handlerConstructor.New()

	r := infrastructure.InitGinRouter(logger)
	routerHandler := ginrouter.NewRouter(ucHandler, logger)
	routerHandler.SetRoutes(r)

	srv := http.Server{
		Addr:    config.HTTPServer.Addr,
		Handler: r,
	}

	httSrvWaitCh := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			httSrvWaitCh <- err
			close(httSrvWaitCh)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		logger.Info("stabilized shutdown...")
	case err := <-httSrvWaitCh:
		logger.Error("emergency shutdown on http error", zap.Error(err))
	}

}

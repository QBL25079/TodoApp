package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_middleware "github.com/QBL25079/TodoApp/internal/core/transport/http/middleware"
	core_http_server "github.com/QBL25079/TodoApp/internal/core/transport/http/server"
	user_transport_http "github.com/QBL25079/TodoApp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init app logger: %w", err)
		os.Exit(1)
	}

	defer logger.Close()

	logger.Debug("Starting to do app")

	usersTransportHTTP := user_transport_http.NewUsersHTTPHandler(nil)

	userRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)

	apiVersionRouter.RegisterRoutes(userRoutes...)

	httpServer := core_http_server.NewHTTPServer(core_http_server.NewConfigMust(), logger, core_http_middleware.RequestID(), core_http_middleware.Logger(logger), core_http_middleware.Panic(), core_http_middleware.Trace())

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

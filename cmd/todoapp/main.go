package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_pgx_pool "github.com/QBL25079/TodoApp/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/QBL25079/TodoApp/internal/core/transport/http/middleware"
	core_http_server "github.com/QBL25079/TodoApp/internal/core/transport/http/server"
	task_postgres_repo "github.com/QBL25079/TodoApp/internal/features/tasks/repository/postgres"
	task_service "github.com/QBL25079/TodoApp/internal/features/tasks/service"
	task_transport_http "github.com/QBL25079/TodoApp/internal/features/tasks/transport/http"
	users_postgres_repo "github.com/QBL25079/TodoApp/internal/features/users/repository/postgres"
	user_service "github.com/QBL25079/TodoApp/internal/features/users/service"
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

	logger.Debug("initializing connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("Initializing feature ", zap.String("feature", "users"))
	usersRepository := users_postgres_repo.NewUsersRepository(pool)
	userService := user_service.NewUserService(usersRepository)
	usersTransportHTTP := user_transport_http.NewUsersHTTPHandler(userService)

	logger.Debug("Initializing feature ", zap.String("feature", "tasks"))

	tasksRepository := task_postgres_repo.NewTaskRepository(pool)
	taskService := task_service.NewTaskService(tasksRepository)
	tasksTransportHTTP := task_transport_http.NewTaskHTTPHandler(taskService)

	logger.Debug("initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(core_http_server.NewConfigMust(), logger, core_http_middleware.RequestID(), core_http_middleware.Logger(logger), core_http_middleware.Trace(), core_http_middleware.Panic())

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)

	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

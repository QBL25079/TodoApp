package core_http_server

import (
	"context"
	"fmt"
	"errors"
	"net/http"
	"go.uber.org/zap"
	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
)

type HTTPSever struct {
	mux *http.ServeMux
	config Config
	log core_logger.Logger
}


func NewHTTPServer(config Config, log core_logger.Logger) *HTTPSever {
	return &HTTPSever{mux: http.NewServeMux(), config: config, log: log}
}

func (h *HTTPSever) Run(ctx context.Context) error {
	server := &http.Server{Addr: h.config.Addr, Handler: h.mux}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		h.log.Warn("Start HTTP server", zap.String("addr: ", h.config.Addr))

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("ListenAndServe HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown http server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
		h.log.Warn("HTTP server stopped")

	}
	return nil
}
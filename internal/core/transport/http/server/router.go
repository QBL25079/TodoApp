package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/QBL25079/TodoApp/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("/v1")
	ApiVersion2 = ApiVersion("/v2")
	ApiVersion3 = ApiVersion("/v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []core_http_middleware.Middleware
}

func NewApiVersionRouter(apiVersion ApiVersion, middleware ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{ServeMux: http.NewServeMux(), apiVersion: apiVersion, middleware: middleware}
}

func (a *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		a.Handle(pattern, route.WithMiddleware())
	}
}

func (a *APIVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(a, a.middleware...)
}

package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/QBL25079/TodoApp/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no keys=%w", core_errors.ErrInvalidArgument)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("path value='%s' by key='%s' not a valid integer: %v: %w", pathValue, key, err, core_errors.ErrInvalidArgument)
	}

	if val < 0 {
		return 0, fmt.Errorf("ID always positiv: %w", core_errors.ErrInvalidArgument)
	}

	return val, nil
} 
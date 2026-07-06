package task_transport_http

import (
	"fmt"
	"net/http"

	"github.com/QBL25079/TodoApp/internal/core/domain"
	core_logger "github.com/QBL25079/TodoApp/internal/core/logger"
	core_http_request "github.com/QBL25079/TodoApp/internal/core/transport/http/request"
	core_http_response "github.com/QBL25079/TodoApp/internal/core/transport/http/response"
	core_http_types "github.com/QBL25079/TodoApp/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed core_http_types.Nullable[bool] `json:"completed"`
}

type PatchTaskResponse TaskDTOResponse

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Title can`t bu NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("Title length must be between 1 fnd 100 chars")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			desLen := len([]rune(*r.Description.Value))

			if desLen < 1 || desLen > 1000 {
				return fmt.Errorf("Description length must be between 1 and 1000 chars")
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("Completed can`t be NULL")
		}
	}

	return nil
}

func (h TaskHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get id query param ")
		return 
	}

	var req PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	taskPatch := taskPatchFromRequest(req)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return 
	}

	response := PatchTaskResponse(taskDTOFromDomain(taskDomain)) 

	responseHandler.JSONResponse(response, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(request.Title.ToDomain(), request.Description.ToDomain(), request.Completed.ToDomain())
}
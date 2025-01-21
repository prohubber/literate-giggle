// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint   `json:"id,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
	Task   *string `json:"task,omitempty"`
	UserId *uint   `json:"user_id,omitempty"`
}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// PatchTasksIdJSONRequestBody defines body for PatchTasksId for application/json ContentType.
type PatchTasksIdJSONRequestBody = Task

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
	// Get tasks by user ID
	// (GET /tasks/user/{user_id})
	GetTasksUserUserId(ctx echo.Context, userId uint) error
	// Delete a task
	// (DELETE /tasks/{id})
	DeleteTasksId(ctx echo.Context, id uint) error
	// Update a task partially
	// (PATCH /tasks/{id})
	PatchTasksId(ctx echo.Context, id uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// GetTasksUserUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasksUserUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "user_id" -------------
	var userId uint

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", ctx.Param("user_id"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasksUserUserId(ctx, userId)
	return err
}

// DeleteTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTasksId(ctx, id)
	return err
}

// PatchTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTasksId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)
	router.GET(baseURL+"/tasks/user/:user_id", wrapper.GetTasksUserUserId)
	router.DELETE(baseURL+"/tasks/:id", wrapper.DeleteTasksId)
	router.PATCH(baseURL+"/tasks/:id", wrapper.PatchTasksId)

}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks201JSONResponse Task

func (response PostTasks201JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type GetTasksUserUserIdRequestObject struct {
	UserId uint `json:"user_id"`
}

type GetTasksUserUserIdResponseObject interface {
	VisitGetTasksUserUserIdResponse(w http.ResponseWriter) error
}

type GetTasksUserUserId200JSONResponse []Task

func (response GetTasksUserUserId200JSONResponse) VisitGetTasksUserUserIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTasksIdRequestObject struct {
	Id uint `json:"id"`
}

type DeleteTasksIdResponseObject interface {
	VisitDeleteTasksIdResponse(w http.ResponseWriter) error
}

type DeleteTasksId204Response struct {
}

func (response DeleteTasksId204Response) VisitDeleteTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PatchTasksIdRequestObject struct {
	Id   uint `json:"id"`
	Body *PatchTasksIdJSONRequestBody
}

type PatchTasksIdResponseObject interface {
	VisitPatchTasksIdResponse(w http.ResponseWriter) error
}

type PatchTasksId200JSONResponse Task

func (response PatchTasksId200JSONResponse) VisitPatchTasksIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
	// Get tasks by user ID
	// (GET /tasks/user/{user_id})
	GetTasksUserUserId(ctx context.Context, request GetTasksUserUserIdRequestObject) (GetTasksUserUserIdResponseObject, error)
	// Delete a task
	// (DELETE /tasks/{id})
	DeleteTasksId(ctx context.Context, request DeleteTasksIdRequestObject) (DeleteTasksIdResponseObject, error)
	// Update a task partially
	// (PATCH /tasks/{id})
	PatchTasksId(ctx context.Context, request PatchTasksIdRequestObject) (PatchTasksIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context) error {
	var request GetTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(ctx echo.Context) error {
	var request PostTasksRequestObject

	var body PostTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx.Request().Context(), request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		return validResponse.VisitPostTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTasksUserUserId operation middleware
func (sh *strictHandler) GetTasksUserUserId(ctx echo.Context, userId uint) error {
	var request GetTasksUserUserIdRequestObject

	request.UserId = userId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasksUserUserId(ctx.Request().Context(), request.(GetTasksUserUserIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasksUserUserId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksUserUserIdResponseObject); ok {
		return validResponse.VisitGetTasksUserUserIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTasksId operation middleware
func (sh *strictHandler) DeleteTasksId(ctx echo.Context, id uint) error {
	var request DeleteTasksIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTasksId(ctx.Request().Context(), request.(DeleteTasksIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTasksId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTasksIdResponseObject); ok {
		return validResponse.VisitDeleteTasksIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchTasksId operation middleware
func (sh *strictHandler) PatchTasksId(ctx echo.Context, id uint) error {
	var request PatchTasksIdRequestObject

	request.Id = id

	var body PatchTasksIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTasksId(ctx.Request().Context(), request.(PatchTasksIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTasksId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTasksIdResponseObject); ok {
		return validResponse.VisitPatchTasksIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

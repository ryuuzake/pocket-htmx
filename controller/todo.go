package controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"

	"github.com/ryuuzake/pocket-htmx/query"
	"github.com/ryuuzake/pocket-htmx/utils"
	"github.com/ryuuzake/pocket-htmx/view/todos"
)

const (
	TODOS = "todos"
)

type TodoHandler struct{ Handler }

func NewTodoHandler(app *pocketbase.PocketBase) TodoHandler {
	return TodoHandler{NewHandler(app)}
}

func (h TodoHandler) GetTodos(c echo.Context) error {
	todosList, err := query.GetListTodos(h.app.Dao())
	if err != nil {
		return apis.NewBadRequestError("", err)
	}
	return utils.Render(todos.Todos(todosList), c)
}

func (h TodoHandler) PostToggleTodo(c echo.Context) error {
	record, err := h.app.Dao().FindRecordById(TODOS, c.PathParam("id"))
	if err != nil {
		return apis.NewNotFoundError("", err)
	}
	canAccess, err := h.app.Dao().CanAccessRecord(record, apis.RequestInfo(c), record.Collection().UpdateRule)
	if !canAccess {
		return apis.NewForbiddenError("", err)
	}
	var isComplete bool = record.Get("is_complete").(bool)
	record.Set("is_complete", !isComplete)
	if err := h.app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/todos")
}

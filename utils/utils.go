package utils

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func Render(cmp templ.Component, c echo.Context) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func SetRequestContextValue(c echo.Context, key string, val any) {
	c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), key, val)))
}

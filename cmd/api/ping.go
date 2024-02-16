package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) ping(c echo.Context) error {

	env := map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	return c.JSON(http.StatusOK, env)

}

package mw

import (
	"github.com/labstack/echo/v4"
	"log"
)

func RoleCheckMW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header := ctx.Request().Header.Get("User-Role")

		if header == "admin" {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}

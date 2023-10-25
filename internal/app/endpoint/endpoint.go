package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int
}
type Endpoint struct {
	service Service
}

func New(service Service) *Endpoint {
	return &Endpoint{
		service: service,
	}
}

func (endpoint *Endpoint) Status(ctx echo.Context) error {
	numberOfDays := endpoint.service.DaysLeft()
	dur := fmt.Sprintf("Number of days: %d", numberOfDays)

	err := ctx.String(http.StatusOK, dur)
	if err != nil {
		return err
	}

	return nil
}

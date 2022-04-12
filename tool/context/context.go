package context

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type DataRes struct {
	Status string      `json:"Status"`
	Data   interface{} `json:"Data"`
}

type ErrorRes struct {
	Status string      `json:"Status"`
	ErrMsg interface{} `json:"Data"`
}

func RetData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, DataRes{
		Status: "200",
		Data:   data,
	})
}

func RetError(c echo.Context, code int, status, errMsg string) error {
	return c.JSON(code, ErrorRes{
		Status: status,
		ErrMsg: errMsg,
	})
}
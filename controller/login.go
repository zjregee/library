package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func Login(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["token"] = "daada"
	return context.RetData(c, resData)
}

func Register(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	return context.RetData(c, resData)
}

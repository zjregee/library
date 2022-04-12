package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func BookIsFavorite(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["is_favorite"] = "1"
	return context.RetData(c, resData)
}

func BookChangeFavorite(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	return context.RetData(c, resData)
}
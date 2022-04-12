package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func BorrowBook(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	return context.RetData(c, resData)
}

func ReturnBook(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	return context.RetData(c, resData)
}

func BorrowBookTime(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["borrow_days"] = "10"
	resData["remain_days"] = "8"
	return context.RetData(c, resData)
}

func BorrowBookNum(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["borrow_nums"] = "10"
	resData["remain_nums"] = "8"
	return context.RetData(c, resData)
}
package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func GetHotBook(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "3"
	temp := []interface{}{}
	t := map[string]interface{}{}
	t["book_id"] = "adad"
	t["title"] = "三体"
	t["author"] = "路遥"
	t["picture"] = "12121"
	temp = append(temp, t)
	temp = append(temp, t)
	temp = append(temp, t)
	resData["books"] = temp
	return context.RetData(c, resData)
}

func GetNewBook(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "3"
	temp := []interface{}{}
	t := map[string]interface{}{}
	t["book_id"] = "adad"
	t["title"] = "三体"
	t["author"] = "路遥"
	t["picture"] = "12121"
	temp = append(temp, t)
	temp = append(temp, t)
	temp = append(temp, t)
	resData["books"] = temp
	return context.RetData(c, resData)
}
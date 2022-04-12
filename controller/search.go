package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func SearchBooks(c echo.Context) error {
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

func SearchLost(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "3"
	temp := []interface{}{}
	t := map[string]interface{}{}
	t["title"] = "蓝色的水杯"
	t["time"] = "2022-03-01"
	t["position"] = "图书馆一楼"
	t["picture"] = "12344"
	t["id"] = "1231"
	temp = append(temp, t)
	temp = append(temp, t)
	temp = append(temp, t)
	resData["goods"] = temp
	return context.RetData(c, resData)
}

func SearchActivity(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "3"
	temp := []interface{}{}
	t := map[string]interface{}{}
	t["title"] = "您的新学期图书馆讲座"
	t["time"] = "2022-03-18 00:00"
	t["picture"] = "12121"
	t["site"] = "线上"
	t["state"] = "回顾"
	temp = append(temp, t)
	temp = append(temp, t)
	temp = append(temp, t)
	resData["activity"] = temp
	return context.RetData(c, resData)
}
package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func GerHotActivity(c echo.Context) error {
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

func GetNewActivity(c echo.Context) error {
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

func GetActivityInformation(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	temp := []interface{}{}
	temp = append(temp, "阅读打开")
	temp = append(temp, "迎新")
	resData["labels"] = temp
	resData["labels_count"] = "2"
	resData["view_count"] = "2204"
	resData["people_num"] = "200"
	resData["max_num"] = "200"
	resData["introduce"] = "dasfsdfsfsfsafs"
	resData["is_apply"] = "1"
	return context.RetData(c, resData)
}

func ApplyActivity(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	return context.RetData(c, resData)
}
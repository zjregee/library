package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func GetAllLost(c echo.Context) error {
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

func GetLostInformation(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["title"] = "13321322222"
	resData["picture"] = "13123"
	resData["notes"] = "如图"
	resData["how_text"] = "qq联系"
	resData["place"] = "1"
	resData["time"] = "1"
	return context.RetData(c, resData)
}
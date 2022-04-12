package controller

import (
	"library/tool/context"
	"github.com/labstack/echo/v4"
)

func GetPersonalInformation(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["name"] = "123"
	resData["student_id"] = "U201917288"
	resData["college"] = "软件学院"
	resData["head_portrait"] = "21"
	return context.RetData(c, resData)
}

func SetPersonalInformation(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	return context.RetData(c, resData)
}

func GetPersonalBorrow(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "3"
	resData["max_count"] = "20"
	temp := []interface{}{}
	t := map[string]interface{}{}
	t["title"] = "三体"
	t["picture"] = "12121"
	t["borrow_time"] = "2022-03-01"
	t["return_time"] = "2022-04-30"
	t["borrow_days"] = "10"
	t["return_days"] = "4"
	temp = append(temp, t)
	temp = append(temp, t)
	temp = append(temp, t)
	resData["books"] = temp
	return context.RetData(c, resData)
}

func GetPersonalFavoriteNum(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "10"
	return context.RetData(c, resData)
}

func GetPersonalFavorite(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["count"] = "3"
	temp := []interface{}{}
	t := map[string]interface{}{}
	t["title"] = "三体"
	t["picture"] = "12121"
	t["author"] = "我"
	t["state"] = "1"
	temp = append(temp, t)
	temp = append(temp, t)
	temp = append(temp, t)
	resData["books"] = temp
	return context.RetData(c, resData)
}



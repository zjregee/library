package controller

import (
	"github.com/labstack/echo/v4"
	"library/tool/context"
)

func GetBookBasicInformation(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["title"] = "1"
	resData["picture"] = "1"
	resData["author"] = "1"
	resData["publish"] = "1"
	resData["publish_time"] = "1"
	resData["introduce_first"] = "1"
	resData["introduce_second"] = "1"
	return context.RetData(c, resData)
}

func GetBookSpecificInformation(c echo.Context) error {
	resData := map[string]interface{}{}
	resData["flag"] = "1"
	resData["guancangdidian"] = "1"
	resData["suoshuhao"] = "1"
	resData["guancangzhuangtai"] = "1"
	resData["zaitixingtai"] = "1"
	resData["ISBN"] = "1"
	resData["zhuti"] = "1"
	resData["jianjie"] = "1"
	return context.RetData(c, resData)
}
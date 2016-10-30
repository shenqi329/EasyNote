package controller

import (
	"easynote/service"
	"github.com/labstack/echo"
	"net/http"
)

func GetRegister(c echo.Context) error {

	service.IsUserExist()
	test := User{Name: "liujunshi", Password: "123456"}
	return c.JSON(http.StatusOK, test)

}

package controller

import (
	"fmt"
	"easynote/bean"
	"easynote/service"
	"github.com/labstack/echo"
	"net/http"
)

func GetRegister(c echo.Context) error {
	fmt.Println("GetRegister");
	service.IsUserExist()
	test := bean.User{Name: "liujunshi", Password: "123456"}
	return c.JSON(http.StatusOK, test)

}

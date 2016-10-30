package controller

import (
	//"fmt"
	//"easynote/bean"
	//"easynote/service"
	"github.com/labstack/echo"
	"net/http"
)

var (
	users = []string{"Joxxe", "Veerxx", "Zioxxn"}
)

func GetRegister(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}


// func GetRegister(c echo.Context) error {
// 	fmt.Println("GetRegister");
// 	service.IsUserExist()
// 	test := bean.User{Name: "liujunshi", Password: "123456"}
// 	return c.JSON(http.StatusOK, test)

// }

package controller

import (
	"easynote/bean"
	"easynote/service"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	_ "os"
	"time"
)

func IsUserExist(c echo.Context) error {

	userName := c.Param("name")

	isUserExist := service.IsUserExist(userName)

	userExistBean := bean.UserExistBean{Exist: isUserExist, Name: userName}

	data, _ := json.Marshal(userExistBean)

	responseBean := bean.ResponseBean{Code: "000001", Desc: "success", Data: string(data)}

	cookie := &standard.Cookie{&http.Cookie{
		Name:  "session",
		Value: "securetoken",
		//Path:    "/",
		Domain:  "192.168.99.100",
		Expires: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		//Expires: time.Date(2016, time.December, 1, 0, 0, 0, 0, time.UTC),
		//Secure:  true,
		//HttpOnly: true,
	}}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, responseBean)

}

func UserRegister(c echo.Context) error {

	user := new(bean.UserBean)
	if error := c.Bind(user); error != nil {
		return error
	}
	fmt.Println("username:" + user.Name + "password" + user.Password)

	if !service.UserRegister(user) {
		responseBean := bean.ResponseBean{Code: "000002", Desc: "resource is exist"}
		return c.JSON(http.StatusOK, responseBean)
	}

	responseBean := bean.ResponseBean{Code: "000001", Desc: "success"}
	return c.JSON(http.StatusOK, responseBean)
}

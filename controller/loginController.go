package controller

import (
	"easynote/bean"
	"easynote/service"
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	_ "os"
	"fmt"
)


func IsUserExist(c echo.Context) error {

	userName := c.Param("name")

	isUserExist := service.IsUserExist(userName)

	userExistBean := bean.UserExistBean{Exist:isUserExist,Name:userName}

	data,_ := json.Marshal(userExistBean);

	responseBean := bean.ResponseBean{Code:"000001",Desc:"success",Data:string(data)}
	
	return c.JSON(http.StatusOK, responseBean)

}

func UserRegister(c echo.Context) error{

	user := new(bean.UserBean)
	if error := c.Bind(user) ; error != nil {
		return error
	}
	fmt.Println("username:"+user.Name+"password"+user.Password)

	if(!service.UserRegister(user)){
		responseBean := bean.ResponseBean{Code:"000002",Desc:"resource is exist"}
		return c.JSON(http.StatusOK,responseBean);
	}
	
	responseBean := bean.ResponseBean{Code:"000001",Desc:"success"}
	return c.JSON(http.StatusOK,responseBean);
}
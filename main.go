package main

import (
	"easynote/controller"
	//"errors"
	//easynotemiddleware "easynote/middleware"
	//"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	glog "github.com/labstack/gommon/log"
	"net/http"
)

func defaultServer(c echo.Context) error {
	//return errors.New("api did not exist!")
	return c.JSON(http.StatusOK, "a message from server")
}

func main() {

	e := echo.New()

	e.Use(echomiddleware.Logger())
	//e.Use(easynotemiddleware.ReqRespLogger())

	e.GET("/user/:name", controller.IsUserExist)
	e.POST("/user/register", controller.UserRegister)
	e.GET("/", defaultServer)

	e.SetDebug(true)
	e.SetLogLevel(glog.DEBUG)

	e.Run(standard.WithTLS(":80", "./study/https/home/ca.crt", "./study/https/home/ca_decrypted.key"))
	e.Run(standard.New(":80"))
}

// var (
// 	users = []string{"Joe", "Veer", "Zion"}
// )

// func getUsers(c echo.Context) error {
// 	return c.JSON(http.StatusOK, users)
// }

// type (
// 	User struct {
// 		Name     string
// 		Password string
// 	}
// )

// func getRegister(c echo.Context) error {
// 	test := User{Name: "liujunshi", Password: "123456"}
// 	return c.JSON(http.StatusOK, test)
// }

//package main

// import (
// 	"fmt"
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// 	"log"
// )

// type Person struct {
// 	Name  string
// 	Phone string
// }

// func main() {
// 	session, err := mgo.Dial("mongodb://localhost:27017")

// 	if err != nil {
// 		fmt.Println("连接失败")
// 		panic(err)
// 		return
// 	}
// 	defer session.Close()

// 	session.SetMode(mgo.Monotonic, true)
// 	c := session.DB("test").C("person")

// 	if c == nil {
// 		return
// 	}

// 	err = c.Insert(&Person{"Ale", "+86 13488888888"},
// 		&Person{"Cla", "+86 13788888888"})

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	result := Person{}

// 	err = c.Find(bson.M{"name": "Ale"}).One(&result)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Phone", result.Phone)
// }

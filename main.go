package main

import (
	"easynote/controller"
	"fmt"
	//easynotemiddleware "easynote/middleware"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	glog "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"net/http"
)

func defaultServer(c echo.Context) error {
	//return c.Redirect(http.StatusTemporaryRedirect, "http://www.baidu.com/")
	return c.JSON(http.StatusOK, `{"name":"Alice","body":"Hello","time":1294706395881547000}`)
}

func main() {

	e := echo.New()

	e.Use(echomiddleware.Logger())
	//e.Use(easynotemiddleware.ReqRespLogger())

	viper.SetConfigType("json")
	viper.SetConfigFile("./viperconf.json")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file change:", e.Name)
	})

	fmt.Println("name:", viper.GetString("name"))
	fmt.Println("password:", viper.GetString("password"))

	e.GET("/user/:name", controller.IsUserExist)
	e.POST("/user/register", controller.UserRegister)

	e.GET("/", defaultServer)

	e.SetDebug(true)
	e.SetLogLevel(glog.DEBUG)

	//fmt.Println("server run on port:80")
	e.Run(standard.New(":80"))
	e.Run(standard.WithTLS(":80", "./study/https/mac/server.crt", "./study/https/mac/server.key"))
}

//https://192.168.99.100/user/1
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

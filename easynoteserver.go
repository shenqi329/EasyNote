package main

import (
	"easynote/controller"
	easynoteGrpc "easynote/grpc"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpcPb "im/logicserver/grpc/pb"
	"log"
	"net"
	"net/http"
)

func defaultServer(c echo.Context) error {
	return c.JSON(http.StatusOK, "a message from server")
}

func grpcConnServer(tcpAddr string) *grpc.ClientConn {
	//Set up a connection to the server.
	conn, err := grpc.Dial(tcpAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func grpcServerRegister(tcpAddr string) {

	accessServerConn := grpcConnServer("localhost:6004")
	imServerConn := grpcConnServer("localhost:6005")

	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//log.Println("设置环境变量")
		ctx = netContext.WithValue(ctx, "AccessServerConn", accessServerConn)
		ctx = netContext.WithValue(ctx, "LogicServerConn", imServerConn)
		return handler(ctx, req)
	}))

	//grpcPb.RegisterSessionServer(s, &easynoteGrpc.Sesion{})

	forward := &easynoteGrpc.Forward{}
	forward.AddHandleFunc(grpcPb.MessageTypeCreateSessionRequest, grpcPb.MessageTypeCreateSessionResponse, easynoteGrpc.CreateSession)
	forward.AddHandleFunc(grpcPb.MessageTypeAddSessionUsersRequest, grpcPb.MessageTypeAddSessionUsersResponse, easynoteGrpc.AddSessionUsers)
	forward.AddHandleFunc(grpcPb.MessageTypeDeleteSessionUsersRequest, grpcPb.MessageTypeDeleteSessionUsersResponse, easynoteGrpc.DeleteSessionUsers)
	grpcPb.RegisterForwardToLogicServer(s, forward)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go grpcServerRegister("localhost:6006")

	e := echo.New()

	e.Use(echomiddleware.Logger())
	//e.Use(easynotemiddleware.ReqRespLogger())

	e.GET("/user/:name", controller.IsUserExist)
	e.POST("/user/register", controller.UserRegister)

	e.GET("/", defaultServer)

	fmt.Println("server run on port:8082")
	e.Run(standard.New(":8082"))
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

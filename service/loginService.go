package service

import (
	"easynote/bean"
	"easynote/mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func IsUserExist(userName *string) bool {

	session := mongodb.GetSession()

	c := session.DB("test").C("person")
	if c == nil {
		return false
	}

	result := bean.User{}
	err := c.Find(bson.M{"Name": userName}).One(&result)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Name", result.Name)

	return true
}

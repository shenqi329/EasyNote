package service

import (
	"easynote/bean"
	"easynote/mongodb"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	//"strings"
	//"log"
)

func IsUserExist(name string) bool {

	session := mongodb.GetSession()
	fmt.Println("name = %s",name)
	c := session.DB("test").C("t_user")
	if c == nil {
		fmt.Println("get t_user fail");
		return false
	}
	fmt.Println("get t_user success");

	result := bean.UserBean{}
	err := c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return false
	}
	fmt.Println("query name = %s,password = %s",result.Name,result.Password);
	return true
}

func UserRegister(user *bean.UserBean) bool{
	session := mongodb.GetSession()

	c := session.DB("test").C("t_user")
	if c == nil {
		fmt.Println("get t_user fail");
		return false
	}

	if IsUserExist(user.Name) {
		fmt.Println("Insert fail,user already exist:" + user.Name);
		return false;
	}

	err := c.Insert(user)

	if err != nil {
		fmt.Println("Insert fail:" + user.Name);
		return false;
	}
	return true;
}
package bean

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	UserGenderMale   = "m"
	UserGenderFemale = "f"
)

type (
	UserBean struct {
		ID       bson.ObjectId `bson:"_id,omitempty"`
		Name     string        `json:"name" xml:"name" form:"name"`
		Password string        `json:"password" xml:"password" form:"password"`
		NickName string        `json:"nickName"`
		Gender   string        `json:"gender"`
	}
)

type (
	UserExistBean struct {
		Name  string `json:"name"`
		Exist bool   `json:"exist"`
	}
)

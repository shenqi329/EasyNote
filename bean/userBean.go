package bean

type (
	UserBean struct {
		Name     string `json:"name" xml:"name" form:"name"`
		Password string `json:"password" xml:"password" form:"password"`
	}
)

type (
	UserExistBean struct {
		Name  string
		Exist bool
	}
)

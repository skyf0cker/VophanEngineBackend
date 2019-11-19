package test

import (
	"VophanEngineBackend/model"
	"fmt"
	"testing"
)

//success
func TestUser_CreateUser(t *testing.T) {
	u := &model.User{
		Username:"testing789",
		Password:"testing555",
		Email:"809866729@qq.com",
		UserChoice: []model.UserChoice{
			model.UserChoice{Content:"test content"},
			model.UserChoice{Content:"test content2"},
			model.UserChoice{Content:"test content3"},
		},
	}

	if err := model.CreateUser(u);err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("success!")
	}
}

// success
func TestUser_UpdateByName(t *testing.T) {
	err := model.UpdateUser("testing567", map[string]interface{}{"password":"chnaged"})
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("success!")
	}
}

//success
func TestUser_DeleteUser(y *testing.T) {
	err := model.DeleteUserByName("testing789")
	if err != nil{
		fmt.Print(err)
	} else {
		fmt.Println("success!")
	}
}

//success
func TestUser_GetUserByName(t *testing.T) {
	u, err := model.GetUserByName("testing789")
	if err == nil {
		fmt.Println(u.Password)
	} else {
		fmt.Println(err)
	}
}

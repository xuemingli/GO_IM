package controller

import (
	"GO_IM/model"
	"GO_IM/service"
	"GO_IM/util"
	"fmt"
	"math/rand"
	"net/http"
)

var userService *service.UserService

func UserLogin(w http.ResponseWriter, r *http.Request) {
	//1.获取前端传递的参数
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "mobile or passwd error!")
	}
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	//1.获取前端传递的参数
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	plainpwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "success")

	}
}

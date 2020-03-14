package main

import (
	"GO_IM/controller"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		//
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	//提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir(".")))

	//提供指定的静态资源目录支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	//登录和注册页面模板渲染
	RegisterView()

	//绑定请求和处理函数
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}

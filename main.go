package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"html/template"
	"log"
	"net/http"
)

var DbEngin *xorm.Engine

type H struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func init(){
	driverName := "mysql"
	DsName := "root:root@(127.0.0.1:3306)/chat?charset=utf-8"
	DbEngin,err := xorm.NewEngine(driverName,DsName)
	if err != nil{
		log.Fatal(err.Error())
	}
	//显示sql语句
	DbEngin.ShowSQL(true)
	//设置数据库最大连接数
	DbEngin.SetMaxOpenConns(2)
	//自动user
	//DbEngin.Sync2(new(User))
	fmt.Println("Init data ok.")
}

func resp(w http.ResponseWriter, code int, data interface{},msg string){
	//设置header为JSON，需要将默认的text/html改为application/json
	w.Header().Set("Content-Type","application/json")
	//设置返回200状态码
	w.WriteHeader(http.StatusOK)

	h := H{
		Code:code,
		Msg:msg,
		Data:data,
	}

	ret,err := json.Marshal(h)
	if err != nil{
		log.Println(err.Error())
	}
	//输出信息
	w.Write(ret)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	//1.获取前端传递的参数
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")
	loginok := false
	if mobile == "110" && passwd == "123456" {
		loginok = true
	}
	if loginok{
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		resp(w,0,data,"success")
	}else{
		resp(w,-1,nil,"mobile or passwd failed!")
	}
}

func RegisterView(){
	//一次解析出全部模板
	tpl,err := template.ParseGlob("view/**/*")
	if nil!=err{
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _,v := range tpl.Templates(){
		//
		tplname := v.Name();
		fmt.Println("HandleFunc     "+v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     "+v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w,tplname,nil)
			if err!=nil{
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	//提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir(".")))

	//提供指定的静态资源目录支持
	http.Handle("/asset/",http.FileServer(http.Dir(".")))

	RegisterView()

	//绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)

	//启动web服务器
	http.ListenAndServe(":8080",nil)
}

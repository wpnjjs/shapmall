package main

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	_ "./mongooper"
	"./priv_utils"
	"./tryit"
	// "schema"
)


func main(){
	fmt.Println("Test")
	http.HandleFunc("/test",tryit.Testhandler)
	http.HandleFunc("/login",loginhandler)
	http.HandleFunc("/pages/",getres)
	err:= http.ListenAndServe("0.0.0.0:8080",nil)
	if err != nil {
		fmt.Println(err)
	}
}

// 处理登录和首页展示
func loginhandler(w http.ResponseWriter, r *http.Request) {
	// 根据是否上传post数据区分操作行为
	// 一 、解析处理ajax上传的数据
	fmt.Println(r)
	errf := r.ParseForm()
	if errf != nil {
		fmt.Println("解析表单数据失败！")
	}
	// var decoder = schema.NewDecoder()
	fmt.Println(r.FormValue("usern"))
	fmt.Println(r.PostFormValue("passw"))
	// 二 、表单直接提交，不通过ajax提交
	// if r.PostFormValue("username") != "" && r.PostFormValue("password") != "" {
	// 	name := r.PostFormValue("username")
	// 	pwd := r.PostFormValue("password")
	// 	// 获取username对应的用户数据信息
	// 	fmt.Println(name, pwd)
	// 	accountfind := mongooper.AccountManagement{}
	// 	oper := &accountfind
	// 	bak := oper.Retrieve(name)
	// 	fmt.Println("after Retrieve", bak)
	// }
	pathstr, errp := os.Getwd()
	if errp != nil {
		panic(errp)
	}
	fmt.Println("OK", pathstr)
	var htmlstr []byte
	filename := "./pages/login.html"
	if !priv_utils.CheckFileIsExist(filename) {
		filename = "./pages/error.html"
	}
	log.Println("read file is ", filename)
	// 读取资源文件数据
	htmlstr, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// log.Println("file content is ", htmlstr)
	// 返回客户端
	w.Write([]byte(htmlstr))
	// var oper mongooper.Operator
	// accountcreate := mongooper.AccountManagement{
	// 	"admin",
	// 	"admin123",
	// 	"~~~~",
	// 	"",
	// 	"",
	// 	""}
	// oper = &accountcreate
	// oper.Create()
}

func getres(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getres", r.URL)
	filename := "." + r.URL.Path
	fmt.Println("filename", filename)
	res, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(res))
}


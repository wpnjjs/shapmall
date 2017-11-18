package logicalcontrol

import (
	"time"
	"fmt"
	"net/http"
	"../mongooper"
	"../priv_utils"
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
)

// 处理登录和首页展示
func Loginhandler(w http.ResponseWriter, r *http.Request) {
	// 根据是否上传post数据区分操作行为
	// 一 、解析处理ajax上传的数据
	// fmt.Println(r)
	var htmlstr []byte
	errf := r.ParseForm()
	if errf != nil {
		fmt.Println("解析表单数据失败！")
	}
	// var decoder = schema.NewDecoder()
	// fmt.Println(r.PostForm)
	// fmt.Println(r.FormValue("username"))
	// fmt.Println(r.PostFormValue("username"))
	// 二 、表单直接提交，不通过ajax提交
	if r.PostFormValue("username") != "" && r.PostFormValue("password") != "" {
		name := r.PostFormValue("username")
		pwd := r.PostFormValue("password")
		// 获取username对应的用户数据信息
		fmt.Println(name, pwd)
		accountfind := mongooper.AccountManagement{}
		oper := &accountfind
		bak, obj := oper.Retrieve(name)
		if bak != -1 {
			fmt.Println("id:", obj.Id_)
			fmt.Println("id:", obj.Id_.Hex())
			if obj.AccountName == name && obj.Password == pwd {
				htmlstr = []byte("1|./pages/error.html|" + obj.Id_.Hex())
			}else{
				htmlstr = []byte("0")
			}
		}else{
			if bak == -1 {
				count := oper.RetrieveCount()
				fmt.Println("first access!", count)
				if count == 0 {
					// var oper mongooper.Operator
					id := bson.NewObjectId()
					fmt.Println("createid is >>", id)
					accountcreate := mongooper.AccountManagement{
						id,
						name,
						pwd,
						"0000",
						time.Now().String(),
						"",
						"0"}
					oper := &accountcreate
					createflag := oper.Create()
					if createflag == 1 {
						htmlstr = []byte("1|./pages/error.html|"+id.Hex())
					}
				}
			}else{
				htmlstr = []byte("0")
			}
		}
	}else{
		pathstr, errp := os.Getwd()
		if errp != nil {
			panic(errp)
		}
		fmt.Println("OK", pathstr)
		filename := "./pages/login.html"
		if !priv_utils.CheckFileIsExist(filename) {
			filename = "./pages/error.html"
		}
		log.Println("read file is ", filename)
		// 读取资源文件数据
		html, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		htmlstr = html
	}
	// log.Println("file content is ", htmlstr)
	// 返回客户端
	w.Write([]byte(htmlstr))
}
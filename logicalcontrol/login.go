package logicalcontrol

import (
	"time"
	"fmt"
	"net/http"
	"../mongooper"
	"../priv_utils"
	// "os"
	"log"
	// "io/ioutil"
	"gopkg.in/mgo.v2/bson"
	// "strings"
)

// 处理登录和首页展示
func Loginhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Loginhandler")
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
		// fmt.Println(name, pwd)
		accountfind := mongooper.AccountManagement{}
		oper := &accountfind
		bak, obj := oper.Retrieve(name)
		if bak != -1 {
			// fmt.Println("id:", obj.Id_)
			// fmt.Println("id:", obj.Id_.Hex())
			decryptstr, _ := priv_utils.DesDecrypt(obj.Password,[]byte(priv_utils.KeyGen(name)))
			// fmt.Println("decryptstr", decryptstr)
			if obj.AccountName == name && string(decryptstr) == pwd {
				// enidstr, _ := priv_utils.DesEncrypt([]byte(obj.Id_.Hex()),[]byte(priv_utils.KeyGen(obj.Id_.Hex())))
				// fmt.Println("enidstr",enidstr)
				// fmt.Println("enidstr",string(enidstr))
				htmlstr = []byte("1|./managecenter.html|" + obj.Id_.Hex())
				lstatus:= mongooper.LoginStatue{
					bson.NewObjectId(),
					name,
					time.Now()}
				oper := &lstatus
				oper.DeleteByName(name)
				oper.Create()
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
					
					encryptstr, _ := priv_utils.DesEncrypt([]byte(pwd), []byte(priv_utils.KeyGen(name)))
					// fmt.Println("encryptstr", encryptstr)
					accountcreate := mongooper.AccountManagement{
						id,
						name,
						encryptstr,
						"0000",
						time.Now(),
						"",
						time.Now(),
						"0"}
					oper := &accountcreate
					createflag := oper.Create()
					if createflag == 1 {
						// enidstr, _ := priv_utils.DesEncrypt([]byte(id.Hex()),[]byte(priv_utils.KeyGen(id.Hex())))
						// fmt.Println("first enidstr",enidstr)
						// fmt.Println("first enidstr",string(enidstr))
						htmlstr = []byte("1|./managecenter.html|"+id.Hex())
					}
				}
			}else{
				htmlstr = []byte("0")
			}
		}
	}
	// else{
	// 	pathstr, errp := os.Getwd()
	// 	if errp != nil {
	// 		panic(errp)
	// 	}
	// 	fmt.Println("OK", pathstr)
	// 	filename := "./pages/login.html"
	// 	if !priv_utils.CheckFileIsExist(filename) {
	// 		filename = "./pages/error.html"
	// 	}
	// 	log.Println("read file is ", filename)
	// 	// 读取资源文件数据
	// 	html, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	htmlstr = html
	// }
	// log.Println("file content is ", htmlstr)
	// 返回客户端
	w.Write([]byte(htmlstr))
}

func Logouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logouthandler",r.URL, r.URL.Path)
	errf := r.ParseForm()
	if errf != nil {
		fmt.Println("解析表单数据失败！")
	}
	name := r.PostFormValue("username")
	fmt.Println("name>>>",name)
	lstatus:= mongooper.LoginStatue{}
	oper := &lstatus
	oper.DeleteByName(name)
	w.Write([]byte("./login.html"))
}
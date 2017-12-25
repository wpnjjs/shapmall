package logicalcontrol

import (
	"time"
	"strings"
	"encoding/json"
	"fmt"
	"net/http"
	"../mongooper"
	"../priv_utils"
	"gopkg.in/mgo.v2/bson"
)

func AbutmentConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AbutmentConfigHandler",r.URL)
	pathstrarr := strings.Split(r.URL.Path,"/")
	operstr := pathstrarr[2]
	fmt.Println("get operator string is >>", operstr)

	var htmlstr []byte

	errf := r.ParseForm()
	if errf != nil {
		fmt.Println("解析表单数据失败！")
	}

	switch operstr {
		case "r" :
			fmt.Println("retrive")
			anabutmentconfig := mongooper.AbutmentConfig{}
			oper := &anabutmentconfig
			allrecord := oper.RetrieveAll("")
			jstr, err := json.Marshal(allrecord)
			if err != nil { panic(err)}
			htmlstr = []byte(jstr)
		case "c" :
			fmt.Println("create")
			name := r.PostFormValue("abutmentsystemname")
			code := r.PostFormValue("abutmentsystemcode")
			uid := r.PostFormValue("uid")
			fmt.Println("name>>>",name,code,uid)
			anabutmentconfig := mongooper.AbutmentConfig{
				bson.NewObjectId(),
				name,
				code,
				uid,
				time.Now()}
			oper := &anabutmentconfig
			oper.Create()
		case "u" :
			fmt.Println("update")
			id := r.PostFormValue("id")
			name := r.PostFormValue("abutmentsystemname")
			code := r.PostFormValue("abutmentsystemcode")
			uid := r.PostFormValue("uid")
			fmt.Println("name>>>",name,code,uid)
			anabutmentconfig := mongooper.AbutmentConfig{
				bson.ObjectIdHex(id),
				name,
				code,
				uid,
				time.Now()}
			oper := &anabutmentconfig
			oper.Update()
		case "d" :
			fmt.Println("delete")
			anabutmentconfig := mongooper.AbutmentConfig{}
			oper := &anabutmentconfig
			id := r.PostFormValue("id")
			fmt.Println("delete id is >", id)
			bak := oper.Delete(id)
			fmt.Println("delete", bak)
			htmlstr = []byte(string(bak))
		case "bd" :
			fmt.Println("batch delete")
			anabutmentconfig := mongooper.AbutmentConfig{}
			oper := &anabutmentconfig
			idarr := r.PostFormValue("bid")
			idarrf := priv_utils.ExStringToArray(idarr)
			// fmt.Println("ex arr is >",idarrf)
			for i:=0;i<len(idarrf);i++ {
				bak := oper.Delete(idarrf[i])
				fmt.Println("delete",bak)
			}
			htmlstr = []byte(string(1))
	}

	w.Write([]byte(htmlstr))
}
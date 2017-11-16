package tryit

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"../priv_utils"
)

func Testhandler(w http.ResponseWriter, r *http.Request) {
	pathstr, errp := os.Getwd()
	if errp != nil {
		panic(errp)
	}
	fmt.Println("OK", pathstr)
	var htmlstr []byte
	filename := "./pages/testvue.html"
	if !priv_utils.CheckFileIsExist(filename) {
		filename = "./pages/error.html"
	}
	log.Println("read file is ", filename)
	// 读取资源文件数据
	htmlstr, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// 返回客户端
	w.Write([]byte(htmlstr))
}


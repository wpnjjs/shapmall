package main

import (
	"time"
	"io/ioutil"
	"fmt"
	"net/http"
	"./tryit"
	"./logicalcontrol"
	// "schema"
)


func main(){
	fmt.Println("["+time.Now().String()+"]:Go Service is running...")
	http.HandleFunc("/test",tryit.Testhandler)
	http.HandleFunc("/login",logicalcontrol.Loginhandler)
	http.HandleFunc("/pages/",getres)
	err:= http.ListenAndServe("0.0.0.0:8080",nil)
	if err != nil {
		fmt.Println(err)
	}
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


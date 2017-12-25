// package main
package tryit

import (
	"reflect"
	"fmt"
	"time"
	"../priv_utils"
	"encoding/json"
	_ "gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string `bson:"name"`
	Age int `bson:"age"`
}

func main() {
	// fmt.Println("main begin")
	// defer func() {
	// 	if err:=recover(); err != nil {
	// 		fmt.Println("get panic value is ", err)
	// 	}
	// }()
	// tryone()
	// fmt.Println("main end")
	people := &Person{"lucy",15}
	s := reflect.TypeOf(people).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Name)
	}

	jstr, err := json.Marshal(people)
	if err != nil { panic(err)}
	fmt.Println(jstr)
	fmt.Println(string(jstr))
	peoplea := &Person{}
	json.Unmarshal(jstr, peoplea)
	fmt.Println(peoplea.Name,peoplea.Age)
	fmt.Println(time.Now().String())
	// fmt.Println(time.Location.Local())
	DeEnCryptTest()
}

func tryone() {
	fmt.Println("panic,defer and recover")
	fmt.Println("begin panic")
	panic(55)
}

// 加解密测试
func DeEnCryptTest(){
	fmt.Println(len([]byte("abcd1234")))
	fmt.Println([]byte("for Encrypt string"))
	encryptstr, _ := priv_utils.DesEncrypt([]byte("for Encrypt string"), []byte("abcd1234"))
	fmt.Println(encryptstr)
	decryptstr, _ := priv_utils.DesDecrypt(encryptstr,[]byte("abcd1234"))
	fmt.Println(decryptstr)
	fmt.Println(string(decryptstr))
}
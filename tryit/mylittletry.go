package tryit

import (
	"fmt"
	"time"
	"../priv_utils"
)

func main() {
	// fmt.Println("main begin")
	// defer func() {
	// 	if err:=recover(); err != nil {
	// 		fmt.Println("get panic value is ", err)
	// 	}
	// }()
	// tryone()
	// fmt.Println("main end")
	fmt.Println(time.Now().String())
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
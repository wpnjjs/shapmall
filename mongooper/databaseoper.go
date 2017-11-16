package mongooper

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "fmt"

func Connect() (s *mgo.Session, db *mgo.Database){
	// "连接到mongo服务器，获取数据库对象。返回session和database"
	// 非本地mongo服务连接 mongodb://username:password@host:port[,host:port]/db
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	// defer session.Close(),此处不要关闭，不然会提示Session Already Closed panic
	session.SetMode(mgo.Monotonic,true)
	database := session.DB("test")
	fmt.Println("返回数据库操作对象")
	return session, database
}

type Operator interface {
	// 定义数据操作统一接口CRUD操作
	Create() int
	Delete() int
	Update() int
	Retrieve(idx string) interface{}
}

// 后台用户管理（Account Management）
type AccountManagement struct {
	AccountName string //账号名
	Password string //账号密码
	CreatorCode string  //创建者
	CreateDate string //创建时间
	ModifiedDate string //修改时间
	RollCode string //角色码
}

func (account *AccountManagement) Create() int {
	session, database := Connect()
	defer session.Close()
	coll := database.C("accountmanagement")
	fmt.Println(account.AccountName)
	err := coll.Insert(account)
	fmt.Println("err", err)
	var rst int
	if err != nil {
		rst = 0
		panic(err)
	}else{
		rst = 1
	}
	return rst
}

func (account *AccountManagement) Delete() int {
	return 0
}

func (account *AccountManagement) Update() int {
	return 0
}

func (account *AccountManagement) Retrieve(idx string) interface{} {
	fmt.Println("Retrieve", idx)
	session, database := Connect()
	defer session.Close()
	fmt.Println("before get collection")
	coll := database.C("accountmanagement")
	var accountone *AccountManagement
	fmt.Println("before find")
	err := coll.Find(bson.M{"accountname":idx}).One(&accountone)
	fmt.Println("after find")
	if err != nil { 
		fmt.Println("find error")
		return -1
		// panic(err) 
	}
	fmt.Println("Retrieve", accountone)
	var oper Operator = accountone
	fmt.Println()
	return oper
}
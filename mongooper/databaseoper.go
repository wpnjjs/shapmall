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
	database := session.DB("shopmall")
	fmt.Println("返回数据库操作对象")
	return session, database
}

type Operator interface {
	// 定义数据操作统一接口CRUD操作
	Create() int
	Delete() int
	Update() int
	Retrieve(idx string) interface{}
	RetrieveCount() int
	// Get_ID(name string) interface{}
}

// 后台用户管理（Account Management）
type AccountManagement struct {
	Id_ bson.ObjectId `bson:"_id"`
	AccountName string `bson:"accountname"` //账号名
	Password string `bson:"password"` //账号密码
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建时间
	ModifiedDate string `bson:"modifieddate"` //修改时间
	RollCode string `bson:"rollcode"` //角色码
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

func (account *AccountManagement) Retrieve(idx string) (int, *AccountManagement) {
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
		return -1, nil
		// panic(err) 
	}
	fmt.Println("Retrieve", accountone)
	fmt.Println("Retrieve", accountone.Id_)
	// var oper Operator = accountone
	return 1, accountone
}

func (account *AccountManagement) RetrieveCount() int {
	fmt.Println("RetrieveCount")
	session, database := Connect()
	defer session.Close()
	fmt.Println("before get collection")
	coll := database.C("accountmanagement")
	// var accountone *AccountManagement
	fmt.Println("before find")
	count, err := coll.Count()
	fmt.Println("after find")
	if err != nil { 
		fmt.Println("find error")
		return -1
		// panic(err) 
	}
	return count
}

// func (account *AccountManagement) Get_ID(name string) interface{} {
// 	fmt.Println("Get_ID")
// 	session, database := Connect()
// 	defer session.Close()
// 	fmt.Println("before get collection")
// 	coll := database.C("accountmanagement")
// 	// var accountone *AccountManagement
// 	fmt.Println("before find")
// 	var result *Result
// 	coll.Find(bson.M{"accountname":name}).One(&result) //.Select(bson.M{"_id":1})
// 	fmt.Println("after find", result.AccountId)
// 	return 0
// }
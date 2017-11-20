package mongooper

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "fmt"
import "time"

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
	Password []byte `bson:"password"` //账号密码
	CreatorCode string  `bson:"creatorcode"` //创建者
	// CreateDate string `bson:"createdate"` //创建时间
	CreateDate time.Time `bson:"createdate"` //创建时间
	Modifier string `bson:"modifier"` //修改者
	ModifiedDate string `bson:"modifieddate"` //修改时间
	RollCode string `bson:"rollcode"` //角色码
}

func (account *AccountManagement) Create() int {
	session, database := Connect()
	defer session.Close()
	coll := database.C("accountmanagement")
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

func (account *AccountManagement) Delete(id string) int {
	fmt.Println("Delete", id)
	session, database := Connect()
	defer session.Close()
	coll := database.C("accountmanagement")
	err := coll.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		if err.Error() == "not found" {
			return 0
		}else{
			panic(err)
		}
	}
	return 1
}

func (account *AccountManagement) Update() int {
	return 0
}

func (account *AccountManagement) Retrieve(idx string) (int, *AccountManagement) {
	fmt.Println("Retrieve", idx)
	session, database := Connect()
	defer session.Close()
	coll := database.C("accountmanagement")
	var accountone *AccountManagement
	err := coll.Find(bson.M{"accountname":idx}).One(&accountone)
	if err != nil { 
		return -1, nil
		// panic(err) 
	}
	// var oper Operator = accountone
	return 1, accountone
}

func (account *AccountManagement) RetrieveCount() int {
	fmt.Println("RetrieveCount")
	session, database := Connect()
	defer session.Close()
	coll := database.C("accountmanagement")
	count, err := coll.Count()
	if err != nil { 
		return -1
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

// 货币管理Currency

// 货币兑率（Currency Exchange Rate），管理虚拟货币之间的兑率。兑率管理
type CurrencyExchangeRate struct {
	Id_ bson.ObjectId `bson:"_id"`
	BaseCurrencyName string `bson:"basecurrencyname"` //基础货币名称
	ExTo string  `bson:"exto"` //兑率JSON串
	AbutmentSystemCode string `bson:"abutmentsystemcode"` //游戏ID
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate time.Time `bson:"createdate"` //创建日期
}

// 货币类型管理（Currency Type Management），管理虚拟货币类型。
type CurrencyTypeManagement struct {
	Id_ bson.ObjectId `bson:"_id"`
	CurrencyName string `bson:"currencyname"` //货币名称
	AbutmentSystemCode string `bson:"abutmentsystemcode"`//对接系统码
	CreatorCode string `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建日期
}

// vip管理
// VIP等级管理（VIP Level Managment），管理VIP登记控制。对接系统码区分不同系统的VIP差异。
type VIPLevelManager struct {
	Id_ bson.ObjectId `bson:"_id"`
	LevelCode string `bson:"levelcode"` //VIP代码
	LevelName string `bson:"levelname"`//VIP级别名称
	LevelDesc string `bson:"leveldesc"`//VIP级别描述
	LevelRight string `bson:"levelright"`//VIP权益（JSON）串
	LevelQualification string `bson:"levelqualification"`//VIP资质
	ValidTime int32 `bson:"validtime"`//VIP有效时间（单位s）
	AbutmentSystemCode string `bson:"abutmentsystemcode"`//对接系统码
	CreatorCode string  `bson:"creatorcode"`//创建者
	CreateDate string `bson:"createdate"`//创建日期
}

// 会员管理（Relationship of VIPs and Users）
type RVIPAndUsers struct {
	Id_ bson.ObjectId `bson:"_id"`
	Uid string `bson:"uid"`//用户ID
	LevelCode string `bson:"levelcode"`//VIP代码
	IsValid int `bson:"isvalid"`//是否有效VIP
	InValidDate string `bson:"invaliddate"`//失效日期
	CreatorCode string  `bson:"creatorcode"`//创建者
	CreateDate string `bson:"createdate"`//创建日期
	ModifiedDate string `bson:"modifieddate"`//修改日期
	AbutmentSystemCode string  `bson:"abutmentsystemcode"`//对接系统码
}

// 商品管理
// 商品类型管理 （Goods Types）
type GoodsTypes struct{
	Id_ bson.ObjectId `bson:"_id"`
	GoodsTypeCode string `bson:"goodstypecode"` //商品类型码
	GoodsDesc string `bson:"goodsdesc"` //商品类型描述
	AbutmentSystemCode string `bson:"abutmentsystemcode"` //对接系统码
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建日期
}

// 商品品质管理（Goods Quality）
type GoodsQuality struct {
	Id_ bson.ObjectId `bson:"_id"`
	GoodsQualityCode string `bson:"goodsqualitycode"` //商品级别码
	GoodsQualityDesc string `bson:"goodsqualitydesc"` //级别描述
	AbutmentSystemCode string `bson:"abutmentsystemcode"`  //对接系统码
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建日期
}

// 商品标识管理（Goods Flags）
type GoodsFlags struct {
	Id_ bson.ObjectId `bson:"_id"`
	GoodsFlagCode string `bson:"goodsflagcode"` //商品标识码
	GoodsFlagDesc string `bson:"goodsflagdesc"` //商品标识描述
	AbutmentSystemCode string  `bson:"abutmentsystemcode"` //对接系统码
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建日期
}

// 商品购买限制管理（Goods Restrict Management），定制商品购买制度。如，达到多少级限制，VIP达到限制等。
type GoodsRestrict struct {
	Id_ bson.ObjectId `bson:"_id"`
	GoodsRestrictCode string  `bson:"goodsrestrictcode"` //商品限制码
	GoodsRestrictDesc string  `bson:"goodsrestrictdesc"` //商品限制描述
	RestrictElementCode string `bson:"restrictelementcode"` //限制因素
	RestrictCondition string `bson:"restrictcondition"` //限制条件
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建日期
	AbutmentSystemCode string  `bson:"abutmentsystemcode"` //对接系统码
}

// 限制因子管理 （Restrict Element）
type RestrictElement struct {
	Id_ bson.ObjectId `bson:"_id"`
	RestrictElementCode string `bson:"restrictelementcode"` //限制因素
	RestrictElementDesc string `bson:"restrictelementdesc"` //因子描述
	AbutmentSystemCode string  `bson:"abutmentsystemcode"` //对接系统码
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建时间
}

// 活动管理（Activity Management）
type ActivityManagement struct {
	Id_ bson.ObjectId `bson:"_id"`
	ActivityCode string `bson:"activitycode"` //活动码
	ActivityName string `bson:"activityname"` //活动名
	ActivityTypeCode string `bson:"activitytypecode"` //活动类型码
	IsValid int `bson:"isvalid"` //是否失效
	StartDate string `bson:"startdate"` //活动开始时间
	EndDate string `bson:"enddate"` //活动结束时间
	RestrictElementCode string `bson:"restrictelementcode"` //限制因素
	JoinGoodsCodelist string `bson:"joingoodscodelist"` //活动参与的商品编码数组，JSON串，k:v。key商品码，value权值likes
	CreatorCode string  `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建时间
	AbutmentSystemCode string `bson:"abutmentsystemcode"` //对接系统码
}

// 活动类型管理（Activity Type Management）
type ActivityTypeManagement struct {
	Id_ bson.ObjectId `bson:"_id"`
	ActivityTypeCode string `bson:"activitytypecode"` //活动类型码
	ActivityTypeDesc string `bson:"activitytypedesc"` //活动类型描述
	CreatorCode string `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建时间
	AbutmentSystemCode string `bson:"abutmentsystemcode"` //对接系统码
}

// 玩家背包（Player Backpack）
type Backpack struct {
	Id_ bson.ObjectId `bson:"_id"`
	Uid string `bson:"uid"` //前端用户的ID
	BackpackSize int `bson:"backpacksize"` //背包大小
	GoodsList string `bson:"goodslist"` //背包内容
	AbutmentSystemCode string `bson:"abutmentsystemcode"` //对接系统码
}


// 对接配置（Abutment Managment），对对接的系统后台进行管理对接码（此处为有游戏的唯一码。），以对接码区分各游戏接入数据。
type AbutmentConfig struct {
	Id_ bson.ObjectId `bson:"_id"`
	AbutmentSystemName string `bson:"abutmentsystemName"` //对接系统名称
	AbutmentSystemCode string `bson:"abutmentsystemcode"` //对接系统码
	CreatorCode string `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建日期 
}

func (abutment *AbutmentConfig) Create() int {
	fmt.Println("Create")
	session, database := Connect()
	defer session.Close()
	coll := database.C("abutmentconfig")
	err := coll.Insert(abutment)
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

func (abutment *AbutmentConfig) Delete(id string) int {
	fmt.Println("Delete", id)
	session, database := Connect()
	defer session.Close()
	coll := database.C("abutmentconfig")
	err := coll.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		if err.Error() == "not found" {
			return 0
		}else{
			panic(err)
		}
	}
	return 1
}

func (abutment *AbutmentConfig) Update() int {
	return 0
}

func (abutment *AbutmentConfig) Retrieve(idx string) (int, *AbutmentConfig) {
	fmt.Println("Retrieve", idx)
	session, database := Connect()
	defer session.Close()
	coll := database.C("abutmentconfig")
	var anabutment *AbutmentConfig
	err := coll.Find(bson.M{"accountname":idx}).One(&anabutment)
	if err != nil { 
		return -1, nil
	}
	return 1, anabutment
}

func (abutment *AbutmentConfig) RetrieveCount() int {
	fmt.Println("RetrieveCount")
	session, database := Connect()
	defer session.Close()
	coll := database.C("abutmentconfig")
	count, err := coll.Count()
	if err != nil { 
		return -1
	}
	return count
}

// 角色管理（Roll Management），进入此系统的角色。创建用户分配角色，控制访问权限。 todo
type RollManagement struct {
	Id_ bson.ObjectId `bson:"_id"`
	RollName string `bson:"rollname"` //角色名
	RollCode string `bson:"rollcode"` //角色码
	RollRight string `bson:"rollright"` //访问页面权限JSON串
	CreatorCode string `bson:"creatorcode"` //创建者
	CreateDate string `bson:"createdate"` //创建时间
	ModifiedDate string `bson:"modifieddate"` //修改时间
}
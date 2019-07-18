package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	driverName string
	UserList map[string]*User
)

func init() {
	//init Mysql connect


}

type User struct {
	Id int64
	Username string
	Password string
	Cteatetime int64
	Ip_address string
	Nickname string
	Updatetime int64
	Lastlogintime int64
	Icon string
	Description string
	Token string
}


//func AddUser(u User) string {
//	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
//	UserList[u.Id] = &u
//	return u.Id
//}
//
func UserLogin(username string, password string) (u *User, err error) {
	o := orm.NewOrm()
	user := User{Username:username,Password:password}
	err = o.Read(&user)
	return &user, err
}

func createToken(uid int64){

}

func GetAllUsers() map[string]*User {
	aliasName  := beego.AppConfig.String("dbAliasName")
	driverName := beego.AppConfig.String("dbDriverName")
	userName   := beego.AppConfig.String("dbUserName")
	password   := beego.AppConfig.String("dbPassword")
	host       := beego.AppConfig.String("dbHost")
	port 	   := beego.AppConfig.String("dbPort")
	dbName 	   := beego.AppConfig.String("dbName")

	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbConn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	err := orm.RegisterDataBase(aliasName,driverName,dbConn)
	beego.Debug(err)
	if err != nil {
		beego.Debug("connect mysql error")
		panic(false)
	}
	beego.Debug("connet mysql success")
	qb,_ := orm.NewQueryBuilder(driverName)
	qb.Select("*").From("blog_user")
	sql := qb.String()
	beego.Debug(sql)
	o := orm.NewOrm()
	userList := []User{}
	_,err = o.Raw(sql).QueryRows(&userList)
	beego.Debug(err)
	return UserList
}

//func UpdateUser(uid string, uu *User) (a *User, err error) {
//	if u, ok := UserList[uid]; ok {
//		if uu.Username != "" {
//			u.Username = uu.Username
//		}
//		if uu.Password != "" {
//			u.Password = uu.Password
//		}
//		if uu.Profile.Age != 0 {
//			u.Profile.Age = uu.Profile.Age
//		}
//		if uu.Profile.Address != "" {
//			u.Profile.Address = uu.Profile.Address
//		}
//		if uu.Profile.Gender != "" {
//			u.Profile.Gender = uu.Profile.Gender
//		}
//		if uu.Profile.Email != "" {
//			u.Profile.Email = uu.Profile.Email
//		}
//		return u, nil
//	}
//	return nil, errors.New("User Not Exist")
//}
//
//func Login(username, password string) bool {
//	for _, u := range UserList {
//		if u.Username == username && u.Password == password {
//			return true
//		}
//	}
//	return false
//}
//
//func DeleteUser(uid string) {
//	delete(UserList, uid)
//}

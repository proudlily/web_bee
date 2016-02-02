package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//	"log"
	"os"
	"path"
	"time"
)

const (
	//数据库名字
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

//博客
type Topic struct {
	Id          int64
	Uid         int64
	Title       string
	Content     string `orm:"size(5000)"`
	Attachment  string
	Created     time.Time `orm:"index"`
	Updated     time.Time `orm:"index"`
	views       int64     `orm:"index"`
	Auther      string
	Replytime   time.Time `orm:"index"`
	ReplyCount  int64
	ReplyUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	//注册模型
	orm.RegisterModel(new(Category), new(Topic))
	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}

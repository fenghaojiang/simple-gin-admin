package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"simple-gin-admin/conf"
)

var GlobalDao *Dao

type Dao struct {
	db *sqlx.DB
}

func NewDao() *Dao {
	db := NewDB()
	dao := &Dao{db: db}
	return dao
}

func NewDB() *sqlx.DB {
	c := mysql.NewConfig()
	c.Addr = conf.GlobalConfig.MySQLConf.Addr
	c.Net = conf.GlobalConfig.MySQLConf.Network
	c.DBName = conf.GlobalConfig.MySQLConf.DBName
	c.User = conf.GlobalConfig.MySQLConf.User
	c.Passwd = conf.GlobalConfig.MySQLConf.Password
	c.Collation = "utf8mb4_bin"
	c.ParseTime = true
	db, err := sqlx.Connect("mysql", c.FormatDSN())
	if err != nil {
		panic("Connect Mysql fail!")
	}
	return db
}

func (d *Dao) Close() {
	_ = d.db.Close()
}

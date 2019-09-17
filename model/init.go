package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Database struct {
	adminDB *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)

	if err != nil {
		fmt.Println("Database connection failed. Database name: ", name)
		fmt.Println("error", err)
	}

	setup(db)
	return db
}

func setup(db *gorm.DB) {
	db.DB().SetMaxIdleConns(0)
	db.LogMode(viper.GetBool("db.LogMode"))
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"))
}

func getAdminDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{
		adminDB: getAdminDB(),
	}
}

func (db *Database) Close() {
	DB.adminDB.Close()
}

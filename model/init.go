package model

import (
	"fmt"
	"gin-demo/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	AdminDB *gorm.DB
	MongoDB *mongo.Database
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

func getMongoDB() *mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/test")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.HandlerLogger().Error(err)
	}

	return client.Database("test")
}

func (db *Database) Init() {
	DB = &Database{
		AdminDB: getAdminDB(),
		MongoDB: getMongoDB(),
	}
}

func (db *Database) Close() {
	_ = DB.adminDB.Close()
}

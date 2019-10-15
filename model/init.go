package model

import (
	"context"
	"fmt"
	"gin-demo/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		logger.HandlerLogger().
			WithFields(logrus.Fields{"error": err}).
			Error("Database connection failed. Database name: ", name)
	}
	setup(db)
	return db
}

func setup(db *gorm.DB) {
	db.DB().SetMaxIdleConns(0)
	db.LogMode(viper.GetBool("mysql.logMode"))
}

func getAdminDB() *gorm.DB {
	return openDB(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"))
}

func getMongoUrl(dbType string) (url string) {
	username := viper.GetString("mongoDB." + dbType + ".username")
	password := viper.GetString("mongoDB." + dbType + ".password")
	addr := viper.GetString("mongoDB." + dbType + ".addr")
	name := viper.GetString("mongoDB." + dbType + ".dataname")
	url = fmt.Sprintf("mongodb://%s:%s@%s/%s",
		username,
		password,
		addr,
		name)
	return
}

func getMongoDB() *mongo.Database {
	// Set client options
	url := getMongoUrl("default")
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.HandlerLogger().Error(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.HandlerLogger().WithFields(logrus.Fields{"error": err}).Error("mongo connection failed.")
		err = client.Disconnect(context.TODO())
		if err != nil {
			logger.HandlerLogger().Error(err)
		}
	}
	return client.Database("dipoletest")
}

func (db *Database) Init() {
	DB = &Database{
		AdminDB: getAdminDB(),
		MongoDB: getMongoDB(),
	}
}

func (db *Database) Close() {
	DB.AdminDB.Close()
	//DB.MongoDB.
}

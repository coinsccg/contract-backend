package database

import (
	"context"
	"fmt"
	"time"

	"auction/config"
	"auction/logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*gorm.DB
}

var (
	DB          *gorm.DB
	mongoClient *mongo.Database
)

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	dbSource := config.GetConfig().Database.DbUsername + ":" + config.GetConfig().Database.DbPwd + "@tcp(localhost:" + config.GetConfig().Database.DbPort + ")/" + config.GetConfig().Database.DbSchemaName + "?" + config.GetConfig().Database.DbArgs
	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		logs.GetLogger().Fatal("db err: ", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	db.LogMode(config.GetConfig().Dev)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func SaveOne(data interface{}) error {
	db := GetDB()
	err := db.Save(data).Error
	return err
}

func SaveOneWithTransaction(data interface{}) error {
	tx := GetDB().Begin()
	err := tx.Set("gorm:query_option", "FOR UPDATE").Save(data).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	err = tx.Commit().Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}

func InitMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	oc := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", config.GetConfig().Database.DbUsername, config.GetConfig().Database.DbPwd, config.GetConfig().Database.DbHost, config.GetConfig().Database.DbPort))
	oc.SetMaxPoolSize(100)
	client, err := mongo.Connect(ctx, oc)
	if err != nil {
		panic(err)
	}
	mongoClient = client.Database(config.GetConfig().Database.DbSchemaName)
	return client
}

func GetMongoClient() *mongo.Database {
	return mongoClient
}

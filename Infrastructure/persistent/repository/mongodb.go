package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//连接mongodb数据库

var (
	mongodb *mongo.Client
)

//InitMongoDB 初始化
func InitMongoDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	mongodb, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = mongodb.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return mongodb
}

//Close 关闭数据库
// func Close() {
// 	mongodb.Disconnect()
// }

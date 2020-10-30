package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	var err error
	// mongodb+srv://free-mongo:<password>@cluster0.8pa0j.mongodb.net/<dbname>?retryWrites=true&w=majority
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb+srv://free-mongo:Wukong233.@cluster0.8pa0j.mongodb.net/users?retryWrites=true&w=majority")

	// 连接到MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

type Student struct {
	Id string `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
	CreatedAt string
	UpdatedAt string
}

func main() {
	var err error
	// 指定获取要操作的数据集
	collection := client.Database("test").Collection("users")

	//// 创建一个Student变量用来接收查询的结果
	//var result = Student{
	//	Name:  "test",
	//	Email: "test@qq.com",
	//}
	//res,err := collection.InsertOne(context.TODO(), &result)
	//if err!=nil {
	//	panic(err.Error())
	//}
	//println(res.InsertedID)



	var result2 Student
	filter := bson.M{"name": "test"}
	err = collection.FindOne(context.TODO(), filter).Decode(&result2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result2)
	oid,_ := primitive.ObjectIDFromHex(result2.Id)
	fmt.Println(oid.Timestamp().String())
}

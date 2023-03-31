package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var opt *options.ClientOptions
var collection *mongo.Collection

func init() {
	// clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 检测连接
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected to MongoDB!")

	// client.Database(global.COLLECTION_NAME).CreateCollection(context.TODO(), global.QUESTION_INFO)

	// CheckError(err)
	// collection = client.Database("cui").Collection("question_info") //选择对应集合

	// 建立连接
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 选择数据库my_db
	database := client.Database(global.COLLECTION_NAME)
	// 选择表my_collection
	collection = database.Collection(global.QUESTION_INFO)
	fmt.Println("connet to mongodb successfully")

}

func Insert(ctx context.Context) {

	// 插入记录，_id默认生成
	// record := map[string]interface{}{
	// 	"QuestionDesc": "25",
	// 	"AnswerDesc":   "sssssss",
	// }

	record := model.QuestionEntity{
		QuestionDesc: "how to use mongodb",
		Language:     "go",
		AnswerDesc:   "search it in internet",
		CreateDate:   time.Now().String(),
		CreateBy:     "zhangsan",
		UpdateDate:   time.Now().String(),
		UpdateBy:     "zhangsan",
	}

	insertResult, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("插入记录的ID：", insertResult.InsertedID)
}

func Update(ctx context.Context) {
	filter := bson.D{{"city", "北京"}}
	update := bson.D{{"$inc", bson.D{{"score", 5}}}} //inc为increase
	res, err := collection.UpdateMany(ctx, filter, update)
	CheckError(err)
	fmt.Printf("update %d doc \n", res.ModifiedCount)
}

func Query(ctx context.Context) {
	sort := bson.D{{"name", 1}} //1为升序
	filter := bson.D{{"score", bson.D{{"$gt", 3}}}}
	findOption := options.Find()
	findOption.SetSort(sort)
	findOption.SetLimit(10)
	findOption.SetSkip(1)
	cursor, err := collection.Find(ctx, filter, findOption)
	CheckError(err)
	for cursor.Next(ctx) {
		var doc model.QuestionEntity
		err := cursor.Decode(&doc)
		CheckError(err)
		fmt.Printf("%s %d %d\n", doc.Id, doc.QuestionDesc, doc.AnswerDesc)
	}
}

func CheckError(err error) {
	panic("Get Error " + err.Error())
}

func delete(ctx context.Context) {
	filter := bson.D{{"_id", ""}}
	res, err := collection.DeleteMany(ctx, filter)
	CheckError(err)
	fmt.Printf("delete %d doc \n", res.DeletedCount)
}

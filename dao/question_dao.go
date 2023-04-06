package dao

import (
	"context"
	"fmt"

	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var opt *options.ClientOptions
var collection *mongo.Collection

func Insert(ctx context.Context, questions []model.QuestionEntity) {

	// 插入记录，_id默认生成
	insertResult, err := DB.Collection(global.COLLECTION_QUESTION_INFO).InsertMany(context.TODO(), questions)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("插入记录的ID：", insertResult)
}
func InsertOne(ctx context.Context, question model.QuestionEntity) {

	// 插入记录，_id默认生成
	insertResult, err := DB.Collection(global.COLLECTION_QUESTION_INFO).InsertOne(context.TODO(), question)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("插入记录的ID：", insertResult)
}
func Update(ctx context.Context) {
	filter := bson.D{{"city", "北京"}}
	update := bson.D{{"$inc", bson.D{{"score", 5}}}} //inc为increase
	res, err := collection.UpdateMany(ctx, filter, update)
	CheckError(err)
	fmt.Printf("update %d doc \n", res.ModifiedCount)
}

func QueryAll(ctx context.Context) {
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
func Query(ctx context.Context) {
	sort := bson.D{{"QuestionDesc", 1}} //1为升序
	// filter := bson.D{{"score", bson.D{{"$gt", 3}}}}
	filter := bson.D{}
	findOption := options.Find()
	findOption.SetSort(sort)
	//findOption.SetLimit(10)
	//findOption.SetSkip(1)
	if collection == nil {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 选择数据库my_db
		database := client.Database(global.DB_NAME)
		// 选择表my_collection
		collection = database.Collection(global.COLLECTION_QUESTION_INFO)
	}
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

func Delete(ctx context.Context, id string) {
	filter := bson.D{{"_id", id}}
	res, err := collection.DeleteMany(ctx, filter)
	CheckError(err)
	fmt.Printf("delete %d doc \n", res.DeletedCount)
}
func DeleteAll(ctx context.Context) {
	coll := DB.Collection(global.COLLECTION_QUESTION_INFO)
	coll.RemoveAll(ctx, bson.D{{}})
}

func QueryQuestions(ctx context.Context) []model.QuestionEntity {

	coll := DB.Collection(global.COLLECTION_QUESTION_INFO)
	batch := []model.QuestionEntity{}

	err := coll.Find(ctx, bson.D{{}}).Sort("language").Limit(100).All(&batch)

	if err != nil {
		fmt.Println(err)
	}
	return batch
}

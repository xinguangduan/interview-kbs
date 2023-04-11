package dao

import (
	"context"
	"fmt"

	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/model"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type QuestionDao struct {
}

// var opt *options.ClientOptions
// var collection *mongo.Collection

func (m *QuestionDao) Insert(ctx context.Context, questions []model.QuestionEntity) {

	// 插入记录，_id默认生成
	insertResult, err := DB.Collection(global.COLLECTION_QUESTION_INFO).InsertMany(context.TODO(), questions)
	if err != nil {
		fmt.Println(err)
		return
	}
	global.Logger.Info("插入记录的ID：", insertResult.InsertedIDs)
}
func (m *QuestionDao) InsertOne(ctx context.Context, question model.QuestionEntity) error {

	// 插入记录，_id默认生成
	insertResult, err := DB.Collection(global.COLLECTION_QUESTION_INFO).InsertOne(context.TODO(), question)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	global.Logger.Info("插入记录的ID：", insertResult.InsertedID)
	return err
}
func (m *QuestionDao) UpdateQuestion(question model.QuestionEntity) {
	// filter := bson.D{{"city", "北京"}}
	// update := bson.D{{"$inc", bson.D{{"score", 5}}}} //inc为increase
	//res, err := collection.UpdateMany(ctx, filter, update)
	//CheckError(err)
	// fmt.Printf("update %d doc \n", res.ModifiedCount)
}

func (m *QuestionDao) QueryAll(ctx context.Context) {
	// sort := bson.D{{"name", 0}} //1为升序
	// filter := bson.D{{"score", bson.D{{"$gt", 3}}}}
	// findOption := options.Find()
	// findOption.SetSort(sort)
	// findOption.SetLimit(10)
	// findOption.SetSkip(1)
	// cursor, err := collection.Find(ctx, filter, findOption)
	// CheckError(err)
	// for cursor.Next(ctx) {
	// 	var doc model.QuestionEntity
	// 	err := cursor.Decode(&doc)
	// 	CheckError(err)
	// 	fmt.Printf("%s %d %d\n", doc.Uid, doc.QuestionDesc, doc.AnswerDesc)
	// }
}
func (m *QuestionDao) Query(ctx context.Context) {
	//sort := bson.D{{"QuestionDesc", 1}} //1为升序
	// filter := bson.D{{"score", bson.D{{"$gt", 3}}}}
	// filter := bson.D{}
	// findOption := options.Find()
	//findOption.SetSort(sort)
	//findOption.SetLimit(10)
	//findOption.SetSkip(1)
	// if collection == nil {
	// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// 选择数据库my_db
	// 	database := client.Database(global.DB_NAME)
	// 	// 选择表my_collection
	// 	collection = database.Collection(global.COLLECTION_QUESTION_INFO)
	// }
	// cursor, err := collection.Find(ctx, filter, findOption)
	// CheckError(err)
	// for cursor.Next(ctx) {
	// 	var doc model.QuestionEntity
	// 	err := cursor.Decode(&doc)
	// 	CheckError(err)
	// 	fmt.Printf("%s %d %d\n", doc.Uid, doc.QuestionDesc, doc.AnswerDesc)
	// }
}

func (m *QuestionDao) GetQuestionByUid(uid string) model.QuestionEntity {
	var question model.QuestionEntity
	filter := bson.D{{Key: "uid", Value: uid}}
	DB.Collection(global.COLLECTION_QUESTION_INFO).Find(context.TODO(), filter).One(&question)
	return question
}

func CheckError(err error) {
	if err != nil {
		global.Logger.Fatal("Got Error ", err.Error())
	}

}

func (m *QuestionDao) DeleteQuestionByUid(uid string) error {
	filter := bson.D{{Key: "uid", Value: uid}}
	collection := DB.Collection(global.COLLECTION_QUESTION_INFO)
	err := collection.Remove(context.TODO(), filter)
	CheckError(err)
	return err
}
func (m *QuestionDao) DeleteAll(ctx context.Context) (result *qmgo.DeleteResult, err error) {
	coll := DB.Collection(global.COLLECTION_QUESTION_INFO)
	res, err := coll.RemoveAll(ctx, bson.D{{}})
	CheckError(err)
	return res, err
}

func (m *QuestionDao) QueryQuestions(ctx context.Context) []model.QuestionEntity {

	collection := DB.Collection(global.COLLECTION_QUESTION_INFO)
	batch := []model.QuestionEntity{}
	filter := bson.D{{}}
	//findOption.SetSkip(1)
	//cursor, err := collection.Find(ctx, filter, findOption)
	// err := coll.Find(ctx, filter, findOption).All(&batch)

	err := collection.Find(ctx, filter).Sort("-createDate").All(&batch)
	// err := collection.Find(ctx, filter).Sort("-createDate", "-updateDate").All(&batch)

	if err != nil {
		fmt.Println(err)
	}
	return batch
}

func (m *QuestionDao) GetQuestionList() ([]model.QuestionEntity, int, error) {
	collection := DB.Collection(global.COLLECTION_QUESTION_INFO)
	batch := []model.QuestionEntity{}
	filter := bson.D{{}}
	err := collection.Find(context.TODO(), filter).Sort("-createDate").All(&batch)
	return batch, len(batch), err
}

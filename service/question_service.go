package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/lightsoft/interview-knowledge-base/dao"
	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/model"
	"github.com/lightsoft/interview-knowledge-base/service/dto"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var questionService *QuestionService

type QuestionService struct {
	BaseService
	Dao *dao.QuestionDao
}

func NewQuestionService() *QuestionService {
	if questionService == nil {
		questionService = &QuestionService{
			Dao: &dao.QuestionDao{},
		}
	}

	return questionService
}

func SetTokenToRedis(uid uint, token string) error {
	return global.RedisClient.Set(strings.Replace(global.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(uid)), -1), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
}
func (m *QuestionService) BatchAddQuestion(iQuestionDTOs []*dto.QuestionDTO) error {
	var newDataArray []model.QuestionEntity
	for _, v := range iQuestionDTOs {
		var q model.QuestionEntity
		v.Uid = primitive.NewObjectID().Hex()
		// v.CreateDate = utils.GetNowDate()
		// v.CreateDate = utils.GetNowDate()
		v.CreateDate = time.Now()
		v.UpdateDate = time.Now()

		v.ConvertToModel(&q)

		// question.Language = m.Language
		// question.AnswerDesc = m.AnswerDesc
		// question.KeyWord = m.KeyWord
		// question.QuestionDesc = m.QuestionDesc
		// question.Priority = m.Priority
		// question.Uid = m.Uid

		newDataArray = append(newDataArray, q)
	}
	m.Dao.Insert(context.TODO(), newDataArray)
	return nil
}
func (m *QuestionService) AddQuestion(iQuestionDTO *dto.QuestionDTO) error {
	var questionEntity model.QuestionEntity
	iQuestionDTO.CreateBy = "张三"
	iQuestionDTO.CreateDate = time.Now()
	iQuestionDTO.UpdateBy = "张三"
	iQuestionDTO.UpdateDate = time.Now()
	mongoTimestamp := bson.MongoTimestamp(time.Now().Unix())
	iQuestionDTO.LastUpdateTime = int64(mongoTimestamp)
	iQuestionDTO.ConvertToModel(&questionEntity)
	return m.Dao.InsertOne(context.TODO(), questionEntity)
}

func (m *QuestionService) GetQuestionByUid(commonDTO *dto.CommonDTO) (model.QuestionEntity, error) {
	return m.Dao.GetQuestionByUid(commonDTO.Uid), nil
}

func (m *QuestionService) GetQuestionList(questionListDTO *dto.QuestionListDTO) ([]model.QuestionEntity, int, error) {
	return m.Dao.GetQuestionList()
}

func (m *QuestionService) UpdateQuestion(questionUpdateDTO *dto.QuestionUpdateDTO) error {
	if questionUpdateDTO.Uid == "" {
		return errors.New("Invalid Question Uid")
	}
	var questionUpdate model.QuestionEntity
	questionUpdateDTO.ConvertToModel(&questionUpdate)
	m.Dao.UpdateQuestion(questionUpdate)
	return nil
}

func (m *QuestionService) DeleteQuestionByUid(commonDTO *dto.CommonDTO) error {
	m.Dao.DeleteQuestionByUid(commonDTO.Uid)
	return nil
}
func (m *QuestionService) DeleteAllQuestion() error {
	m.Dao.DeleteAll(context.TODO())
	return nil
}

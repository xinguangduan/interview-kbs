package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/service"
	"github.com/lightsoft/interview-knowledge-base/service/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ERR_CODE_ADD_QUESTION       = 10011
	ERR_CODE_GET_QUESTION_BY_ID = 10012
	ERR_CODE_GET_QUESTION_LIST  = 10013
	ERR_CODE_UPDATE_QUESTION    = 10014
	ERR_CODE_DELETE_QUESTION    = 10015
)

type QuestionApi struct {
	BaseApi
	Service *service.QuestionService
}

func NewQuestionApi() QuestionApi {
	return QuestionApi{
		BaseApi: NewBaseApi(),
		Service: service.NewQuestionService(),
	}
}

// func (m QuestionApi) CreateOneQuestion(c *gin.Context) {

// 	var json dto.QuestionDTO
// 	err := c.BindJSON(&json)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	global.Logger.Info(json)
// 	json.Uid = primitive.NewObjectID().Hex()
// 	json.CreateDate = time.Now()
// 	json.UpdateDate = time.Now()

// 	//dao.InsertOne(context.TODO(), json)

// 	m.Service.AddQuestion(&json)
// 	c.AbortWithStatusJSON(http.StatusOK, "ok")
// }

func (m QuestionApi) AddQuestion(c *gin.Context) {
	var questionDTO dto.QuestionDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &questionDTO}).GetError(); err != nil {
		return
	}
	questionDTO.Uid = primitive.NewObjectID().Hex()

	err := m.Service.AddQuestion(&questionDTO)

	if err != nil {
		m.ServerFail(ResponseMessage{
			Code: ERR_CODE_ADD_QUESTION,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseMessage{
		Data: questionDTO,
		Msg:  "add successfully",
	})
}

func (m QuestionApi) BatchAddQuestion(c *gin.Context) {
	var postData []*dto.QuestionDTO

	if err := c.ShouldBind(&postData); err != nil {
		global.Logger.Error(err)
		return
	}
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: nil}).GetError(); err != nil {
		return
	}

	var newDataArray []*dto.QuestionDTO
	for _, v := range postData {
		v.Uid = primitive.NewObjectID().Hex()
		// v.CreateDate = utils.GetNowDate()
		// v.CreateDate = utils.GetNowDate()
		newDataArray = append(newDataArray, v)
	}

	// dao.InsertUser(context.TODO(), newDataArray)
	err := m.Service.BatchAddQuestion(newDataArray)

	global.Logger.Info(err)

	m.OK(ResponseMessage{
		Msg: "Batch added  successfully",
	})
}

func (m QuestionApi) GetQuestionByUid(c *gin.Context) {
	var commonDTO dto.CommonDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &commonDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	questionDTO, err := m.Service.GetQuestionByUid(&commonDTO)
	if err != nil {
		m.ServerFail(ResponseMessage{
			Code: ERR_CODE_GET_QUESTION_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseMessage{
		Data: questionDTO,
		Msg:  "Get successfully",
	})
}

func (m QuestionApi) GetQuestionList(c *gin.Context) {
	var questionListDTO dto.QuestionListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &questionListDTO}).GetError(); err != nil {
		return
	}

	questionList, nTotal, err := m.Service.GetQuestionList(&questionListDTO)
	global.Logger.Info("question list", questionList, nTotal, err)

	if err != nil {
		m.ServerFail(ResponseMessage{
			Code: ERR_CODE_GET_QUESTION_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseMessage{
		Data:  questionList,
		Total: nTotal,
	})
}

func (m QuestionApi) UpdateQuestion(c *gin.Context) {
	var questionUpdateDTO dto.QuestionUpdateDTO
	//strId := c.Param("id")
	//fmt.Println("strId:" + strId)
	//
	//id, _ := strconv.Atoi(strId)
	//uid := uint(id)
	//questionUpdateDTO.ID = uid

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &questionUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateQuestion(&questionUpdateDTO)

	if err != nil {
		m.ServerFail(ResponseMessage{
			Code: ERR_CODE_UPDATE_QUESTION,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseMessage{
		Msg: "updated successfully",
	})
}

func (m QuestionApi) DeleteQuestionByUid(c *gin.Context) {
	var commonDTO dto.CommonDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &commonDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteQuestionByUid(&commonDTO)
	if err != nil {
		m.ServerFail(ResponseMessage{
			Code: ERR_CODE_DELETE_QUESTION,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseMessage{
		Msg: "deleted successfully",
	})
}

func (m QuestionApi) DeleteAll(c *gin.Context) {
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: nil}).GetError(); err != nil {
		return
	}
	err := m.Service.DeleteAllQuestion()
	if err != nil {
		m.ServerFail(ResponseMessage{
			Code: ERR_CODE_DELETE_QUESTION,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseMessage{
		Msg: "Deleted successfully",
	})
}

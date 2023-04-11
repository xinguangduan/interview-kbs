package global

import (
	"github.com/lightsoft/interview-knowledge-base/configuration"
	"go.uber.org/zap"
)

const DB_NAME = "interview_knowledage_base"
const COLLECTION_QUESTION_INFO = "question_info"
const COLLECTION_USER_INFO = "user_info"

const LOGIN_USER_TOKEN_REDIS_KEY = "LOGIN_USER_TOKEN_REDIS_KEY"

// configuration
const SERVER_PORT = "server.port"
const MONGO_DB_PATH = "db.path"

var (
	Logger      *zap.SugaredLogger
	RedisClient *configuration.RedisClient
)

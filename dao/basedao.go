package dao

import (
	"context"

	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
)

var DB *qmgo.Database

func InitDatabase() {

	ctx := context.Background()
	url := viper.GetString(global.MONGO_DB_PATH)
	//连接数据库
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: url})
	if err != nil {
		return
	}
	//选择数据库
	DB = client.Database(global.DB_NAME)

}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfoEntity struct {
	Id           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Age          uint16             `bson:"age"`
	Weight       uint32             `bson:"weight"`
	CreateAt     time.Time          `bson:"createAt"`
	UpdateAt     int64              `bson:"updateAt"`
	CreateTimeAt time.Time          `bson:"createTimeAt"`
	UpdateTimeAt int64              `bson:"updateTimeAt"`
}

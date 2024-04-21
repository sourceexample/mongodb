package modMongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mdbLog struct {
	ID      primitive.ObjectID `bson:"_id"`
	AddTime time.Time          `bson:"add_time"`
	Text    string             `bson:"text"`
	Appid   string             `bson:"appid"`
}

type CLog struct {
}

var g_singleLog *CLog = &CLog{}

func GetSingleLog() *CLog {
	return g_singleLog
}

func (pInst *CLog) AddLog(text string, appid string) error {
	log1 := &mdbLog{
		ID:      primitive.NewObjectID(),
		AddTime: time.Now(),
		Text:    text,
		Appid:   appid,
	}
	err := getSingleMongoDB().addLog(log1)
	return err
}

func (pInst *CLog) QueryAppid(appid string, limit int) ([]string, error) {

	condition := bson.D{{"appid", appid}} //bson.D{Key: "appid", Value: appid}

	logs, err := getSingleMongoDB().queryAppid(condition, limit)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(logs))
	for _, v := range logs {
		ret = append(ret, v.Text)
	}
	return ret, nil
}

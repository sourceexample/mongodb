package modMongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CMongoDB struct {
	logCollection *mongo.Collection
	ctx           context.Context
}

var g_singleMongoDB *CMongoDB = &CMongoDB{}

func getSingleMongoDB() *CMongoDB {
	return g_singleMongoDB
}
func (pInst *CMongoDB) getLogCollection() *mongo.Collection {
	return pInst.logCollection
}
func (pInst *CMongoDB) addLog(logdata *mdbLog) error {
	_, err := pInst.logCollection.InsertOne(pInst.ctx, logdata)
	return err
}
func (pInst *CMongoDB) queryAppid(qry primitive.D, limit int) ([]mdbLog, error) {

	cur, err := pInst.logCollection.Find(pInst.ctx, qry)
	if err != nil {
		return nil, err
	}
	defer cur.Close(pInst.ctx)
	var result []mdbLog

	for cur.Next(pInst.ctx) {

		// create a value into which the single document can be decoded
		var elem mdbLog
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		result = append(result, elem)

	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (pInst *CMongoDB) Initialize(url string) error {
	pInst.ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(pInst.ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = client.Ping(pInst.ctx, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pInst.logCollection = client.Database("gitlog").Collection("log1")

	return nil
}

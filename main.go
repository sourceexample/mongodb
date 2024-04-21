package main

import (
	"testmongodb/modMongodb"
)

const mongodbcon = "mongodb+srv://gitworkone:gwq31Irr4Ym1RQxh@gitlog.3cuns5c.mongodb.net/?retryWrites=true&w=majority&appName=gitlog"

func main() {
	err := modMongodb.MDB_Initialize(mongodbcon)
	if err != nil {
		return
	}

	inputhandler()
}

// func addData() error {
// 	err = modMongodb.GetSingleLog().AddLog("web log 1", "web2")
// 	err = modMongodb.GetSingleLog().AddLog("web log 2", "web2")

// }

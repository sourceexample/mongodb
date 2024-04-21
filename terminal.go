package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testmongodb/modMongodb"
)

func inputhandler() {
	reader := bufio.NewReader(os.Stdin)
	bQuit := false
	for !bQuit {
		fmt.Print(">: ")
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)
		text = strings.Trim(text, "\r")
		text = strings.Trim(text, "\n")
		command := strings.Split(text, " ")
		commandlength := len(command)
		if commandlength < 1 {
			continue
		}

		switch command[0] {
		case "quit":
			bQuit = true
			break
		case "exit":
			bQuit = true
			break
		case "addapp":
			if commandlength < 2 {
				fmt.Println("add app log error, params not enough")
				continue
			}
			err := modMongodb.GetSingleLog().AddLog(command[1], "app1")
			if err != nil {
				fmt.Println(err)
			}
		case "addweb":
			if commandlength < 2 {
				fmt.Println("add web log error, params not enough")
				continue
			}
			err := modMongodb.GetSingleLog().AddLog(command[1], "web2")
			if err != nil {
				fmt.Println(err)
			}
		case "queryappid":
			if commandlength < 2 {
				fmt.Println("queyr by appid error, params not enough")
				continue
			}
			data, err := modMongodb.GetSingleLog().QueryAppid(command[1], 1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(data)
			}
		}

	}
}

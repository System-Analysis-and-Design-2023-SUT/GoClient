package main

import (
	"fmt"
	"github.com/System-Analysis-and-Design-2023-SUT/GoClient/sadqueue"
	"log"
)

func main() {
	var opt string
	for {
		fmt.Println("1. manual push\n2. manual pull\n3. manual subscribe")

		if scanOpt(&opt) {
			return
		}

		switch opt {
		case "1":
			var key, value string

			fmt.Print("key=")
			if scanOpt(&key) {
				return
			}

			fmt.Print("value=")
			if scanOpt(&value) {
				return
			}

			err := sadqueue.Push(key, value)
			if err != nil {
				fmt.Println(err)
			}
		case "2":
			key, message, err := sadqueue.Pull()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("key:" + key + "\nmessage:" + message)
			}
		case "3":
			err := sadqueue.Subscribe(logInConsole)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func scanOpt(opt any) bool {
	_, err := fmt.Scanln(opt)
	if err != nil {
		return true
	}
	return false
}

func logInConsole(key string, value string) {
	log.Printf("key: %s\tvalue: %s", key, value)
}

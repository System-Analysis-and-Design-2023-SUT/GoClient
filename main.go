package main

import (
	"fmt"
	"github.com/System-Analysis-and-Design-2023-SUT/GoClient/sadqueue"
)

func main() {
	resp, err := sadqueue.Pull()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resp)
	}
}

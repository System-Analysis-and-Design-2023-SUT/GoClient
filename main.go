package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/System-Analysis-and-Design-2023-SUT/GoClient/client"
	"github.com/System-Analysis-and-Design-2023-SUT/GoClient/websocketclient"
)

func main() {
	host := client.FindLivingHost([]string{"http://host1", "http://host2"})
	if host == "" {
		fmt.Println("No available hosts found.")
		return
	}

	fmt.Printf("Connected to %s.\n", host)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Choose an action (push, pull, subscribe or exit):")
		scanner.Scan()
		action := scanner.Text()

		switch action {
		case "push":
			fmt.Println("Enter a message to push:")
			scanner.Scan()
			message := scanner.Text()
			client.PushMessage(host, message)
		case "pull":
			client.PullMessage(host)
		case "subscribe":
			websocketclient.SubscribeToMessages(host, func(message string) {
				fmt.Printf("New message received: %s\n", message)
			})
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}

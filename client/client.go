package client

import (
	// "fmt"
	"net/http"
	// "net/url"
)

func CheckHost(host string) bool {
	resp, err := http.Get(host + "/check")
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

func FindLivingHost(hosts []string) string {
	for _, host := range hosts {
		if CheckHost(host) {
			return host
		}
	}
	return ""
}

func PushMessage(host, message string) {
	// Implementation of push message with query params
}

func PullMessage(host string) {
	// Implementation of pull message
}

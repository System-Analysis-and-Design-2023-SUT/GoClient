package sadqueue

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	hosts    = []string{"http://65.109.208.41:8080", "http://91.107.240.167:8080", "http://91.107.169.86:8080"}
	liveHost string
	mu       sync.Mutex
)

func init() {
	randIndex := rand.Intn(len(hosts))
	liveHost = hosts[randIndex] // Initially, set the first host as live
}

func checkHost(host string) error {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(host + "/-/ready")
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return ErrHostNotAvailable
	}
	return nil
}

func findLiveHost() (string, error) {
	for _, host := range hosts {
		if checkHost(host) == nil {
			return host, nil
		}
	}
	return "", ErrLiveHostNotFound
}

func getLiveHost() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if checkHost(liveHost) == nil {
		return liveHost, nil
	}
	return findLiveHost()
}

func Push(message string) error {
	host, err := getLiveHost()
	if err != nil {
		return err
	}

	key := uuid.New().String()
	resp, err := http.PostForm(fmt.Sprintf("%s/push?key=%s&values%s", host, key, message), nil)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Print(resp)
		return ErrPushFailed
	}
	return nil
}

func Pull() (string, error) {
	host, err := getLiveHost()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(host + "/pull")
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", ErrPullFailed
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	return result["value"], nil
}

func Subscribe() error {
	host, err := getLiveHost()
	if err != nil {
		return err
	}

	wsUrl := strings.Replace(host, "http", "ws", 1) + "/subscribe"
	ws, resp, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return ErrSubscribeFailed
	}

	go func() {
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				fmt.Print(err)
			}
		}(ws)

		err := ws.WriteMessage(websocket.TextMessage, []byte("subscribe\n"))
		if err != nil {
			return
		}
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				break
			}
			fmt.Print(string(message))
		}
	}()
	return nil
}

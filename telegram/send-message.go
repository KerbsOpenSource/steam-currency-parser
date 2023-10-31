// The telegram package is needed to send a message
package telegram

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	chatID   = 0
	apiToken = ""
	baseURL  = "https://api.telegram.org/bot" + apiToken + "/sendMessage"
)

type BotSendMessageID struct {
	Result struct {
		Message_id int
	}
}

func SendMessage(text string) {
	postData := url.Values{}
	postData.Set("chat_id", strconv.Itoa(chatID))
	postData.Set("text", text)
	_, err := http.PostForm(baseURL, postData)
	log.Fatal(err)
}

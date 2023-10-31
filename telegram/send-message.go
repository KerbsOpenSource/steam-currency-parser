// The telegram package is needed to send a message
package telegram

import (
	"log"
	"net/http"
	"net/url"
)

func SendMessage(text string, apiToken string, chatID string) {
	baseURL := "https://api.telegram.org/bot" + apiToken + "/sendMessage"
	postData := url.Values{}
	postData.Set("chat_id", chatID)
	postData.Set("text", text)
	_, err := http.PostForm(baseURL, postData)
	if err != nil {
		log.Fatal(err)
	}
}

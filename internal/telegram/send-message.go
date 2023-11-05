// The telegram package is needed to send a message
package telegram

import (
	"log"
	"net/http"
	"net/url"
)

func SendMessage(text string, telegramBotToken string, chatID string) {
	baseURL := "https://api.telegram.org/bot" + telegramBotToken + "/sendMessage"
	postData := url.Values{}
	postData.Set("chat_id", chatID)
	postData.Set("text", text)
	_, err := http.PostForm(baseURL, postData)
	if err != nil {
		log.Fatal(err)
	}
}

package clients

import (
	"io"
	"log"
	"net/http"
	"os"
)

var telegramApiKey string = os.Getenv("TELEGRAM_API_KEY")

func SendMessage(channelId string, text string) string {
	request, error := http.NewRequest("GET", "https://api.telegram.org/bot"+telegramApiKey+"/sendMessage", nil)
	if error != nil {
		log.Print(error)
		os.Exit(1)
	}

	q := request.URL.Query()

	q.Add("chat_id", channelId)
	q.Add("text", text)

	request.URL.RawQuery = q.Encode()

	var client = &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()
	bodyBytes, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}
	return string(bodyBytes)
}

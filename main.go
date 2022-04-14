package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Input values
	// keyword := "Book tickets"
	eventId := "ET00110845"
	// regionSlug := "manipal"
	// regionCode := "MANI"

	// Get env variables
	// telegramApiKey := os.Getenv("TELEGRAM_API_KEY")
	// telegramChannelId := os.Getenv("TELEGRAM_CHANNEL_ID")
	bookingUrl := os.Getenv("BOOKING_URL")

	fmt.Println(bookingUrl)

	client := &http.Client{}

	// Hit booking url
	request, error := http.NewRequest("GET", bookingUrl, nil)
	if error != nil {
		log.Print(error)
		os.Exit(1)
	}

	q := request.URL.Query()
	q.Add("eventcode", eventId)
	q.Add("isdesktop", "true")
	q.Add("channel", "web")

	request.URL.RawQuery = q.Encode()

	fmt.Println(request.URL.String())

	response, error := client.Do(request)
	fmt.Println(response.Body.Read(make([]byte, 1024)))
	defer response.Body.Close()
}

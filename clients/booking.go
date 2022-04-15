package clients

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"undercover-bookie-go/models"
)

var client = &http.Client{}
var bookingUrl = os.Getenv("BOOKING_URL")

func GetMovieSynopsis(eventId string, regionCode string, regionSlug string) models.MovieSynopsis {
	request, error := http.NewRequest("GET", bookingUrl, nil)
	if error != nil {
		log.Print(error)
		os.Exit(1)
	}

	q := request.URL.Query()
	q.Add("eventcode", eventId)
	q.Add("isdesktop", "true")
	q.Add("channel", "web")

	request.Header.Set("authority", request.Host)
	request.Header.Set("x-app-code", "WEB")
	request.Header.Set("x-region-code", regionCode)
	request.Header.Set("x-region-slug", regionSlug)
	request.Header.Set("accept", "application/json, text/plain, */*")

	request.URL.RawQuery = q.Encode()

	response, error := client.Do(request)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	bodyBytes, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}
	var synopsis models.MovieSynopsis
	json.Unmarshal(bodyBytes, &synopsis)
	return synopsis
}

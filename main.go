package main

import (
	"fmt"
	"os"
	"strings"
	"undercover-bookie-go/clients"
)

func main() {
	// Input values
	keyword := os.Args[1]
	eventId := os.Args[2]
	regionSlug := os.Args[3]
	regionCode := os.Args[4]
	telegramChannelId := os.Args[5]
	telegramDebugChannelId := "@nandeeshkamathdev"

	response, error := clients.GetMovieSynopsis(eventId, regionCode, regionSlug)
	if error != nil {
		clients.SendMessage(telegramDebugChannelId, fmt.Sprint(error))
	}

	if len(response.BannerWidget.PageCta) == 0 {
		return
	}
	bookingOpen := strings.Contains(response.BannerWidget.PageCta[0].Text, keyword)

	if bookingOpen {
		var bookingUrl string
		if len(response.Seo.MetaProperties) > 0 {
			bookingUrl = response.Seo.MetaProperties[7].Value
		} else {
			bookingUrl = "[ No booking url found ]"
		}
		eventName := response.Meta.Event.EventName
		message := fmt.Sprintf("%s is now ready to be booked at %s.\n%s", eventName, regionSlug, bookingUrl)
		_, error := clients.SendMessage(telegramChannelId, message)
		if error != nil {
			clients.SendMessage(telegramDebugChannelId, fmt.Sprint(error))
		}
	}
}

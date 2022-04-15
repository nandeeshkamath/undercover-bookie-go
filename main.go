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

	// if len(os.Args) > 5 {
	// 	telegramDebugChannelId = os.Args[6]
	// }

	response, error := clients.GetMovieSynopsis(eventId, regionCode, regionSlug)
	if error != nil {
		clients.SendMessage(telegramDebugChannelId, fmt.Sprint(error))
	}

	eventName := response.Meta.Event.EventName
	if len(response.BannerWidget.PageCta) == 0 {
		clients.SendMessage(telegramDebugChannelId, fmt.Sprintf("No booking found yet for %s.", eventName))
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
		message := fmt.Sprintf("%s is now ready to be booked at %s.\n%s", eventName, regionSlug, bookingUrl)
		clients.SendMessage(telegramChannelId, message)
	}
}

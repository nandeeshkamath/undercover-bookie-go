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

	response := clients.IsBookingOpen(eventId, regionCode, regionSlug)

	if strings.Contains(response.BannerWidget.PageCta[0].Text, keyword) {
		eventName := response.Meta.Event.EventName
		bookingUrl := response.Seo.MetaProperties[7].Value
		message := fmt.Sprintf("%s is now ready to be booked at %s.\n%s", eventName, regionSlug, bookingUrl)
		clients.SendMessage(telegramChannelId, message)
	}
}

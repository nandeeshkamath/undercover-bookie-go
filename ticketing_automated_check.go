package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type AutomatedCheck struct {
	Urls      map[string]string `json:"urls"`
	EventName string            `json:"event_name"`
}

func (a *AutomatedCheck) Load() {
	configEncoded := os.Getenv("AUTOMATED_CHECK_CONFIG")

	configBytes, err := base64.StdEncoding.DecodeString(configEncoded)
	if err != nil {
		log.Fatal("Failed to decode base64 config:", err)
	}
	configJson := string(configBytes)

	if configJson == "" {
		log.Fatal("No configuration found for AutomatedCheck")
	}

	if err := json.Unmarshal([]byte(configJson), &a); err != nil {
		log.Fatalf("Failed to parse configuration: %v", err)
	}
}

func (a *AutomatedCheck) Check() []string {
	eventUrls := make([]string, 0)
	for slug, url := range a.Urls {
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
		    chromedp.Headless,
		    chromedp.DisableGPU,
		    chromedp.NoSandbox,
		)
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()
		
		chromedp.WithPollingTimeout(30 * time.Second)

		var eventUrl string
		err := chromedp.Run(ctx,
			// Navigate to console
			chromedp.Navigate(url),
			chromedp.WaitVisible(fmt.Sprintf(`//span[contains(normalize-space(text()), "%s")]`, slug), chromedp.BySearch),
			chromedp.Evaluate(fmt.Sprintf(`(function() {
				var node = document.evaluate('//a[contains(normalize-space(text()), "%s")]', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				return node ? node.href : "";
			})()`, a.EventName), &eventUrl),
		)

		if err != nil {
			log.Fatal(err)
		}

		eventUrls = append(eventUrls, eventUrl)
		time.Sleep(1 * time.Second)
	}

	return eventUrls
}

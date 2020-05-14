package main

import (
	"fmt"
	"log"

	"github.com/wedojava/crawler_corps/pkg/corps"
	corpsDw "github.com/wedojava/crawler_corps/pkg/corps/dwnews"
)

func main() {
	fmt.Println("Crawler Corps fighting against Dwnews!")
	dw := corpsDw.DwnewsSite{}
	dw.Domain = "www.dwnews.com"
	dw.StartUrl = "https://" + dw.Domain
	dw.GetUrls() // get and set dw.Urls if method success.
	for _, url := range dw.Urls {
		dwCrawler := corpsDw.DwnewsCrawler{}
		dwCrawler.Url = url
		raw, err := corps.HttpGetBody(url, 10)
		if err != nil {
			log.Println("\n[-] main() -> HttpGetBody() err occur! \n[-] ", err)
		}
		dwCrawler.Raw = raw
		dwCrawler.GetTitle()
		dwCrawler.GetContent()
		dwCrawler.GetDatetime()
	}
}

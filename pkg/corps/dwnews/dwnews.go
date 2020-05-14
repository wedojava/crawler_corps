package corps

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/wedojava/crawler_corps/pkg/corps"
	"github.com/wedojava/gears"
)

type DwnewsCLeader struct {
	corps.CrawlerLeader
}

type DwnewsCrawler struct {
	corps.Crawler
}

type CrawlerDwnews struct {
	Dwnews   DwnewsCLeader
	Crawlers []DwnewsCrawler
}

type Paragraph struct {
	Type    string
	Content string
}

func (web *DwnewsCLeader) GetUrls() error {
	var ret_lst []string
	var reLink = regexp.MustCompile(`(?m)<a\shref\s?=\s?"(?P<href>/.{2}/\d{8}/.+?)".*?>`)
	raw, err := gears.HttpGetBody(web.StartUrl, 10)
	if err != nil {
		return errors.Wrap(err, "\n[-] (web *DwNews)GetUrls()>HttpGetBody() error!\n[-] ")
	}
	for _, v := range reLink.FindAllStringSubmatch(raw, -1) {
		ret_lst = append(ret_lst, web.StartUrl+v[1])
	}
	web.Urls = gears.StrSliceDeDupl(ret_lst)
	return nil
}

func (dwCrawler *DwnewsCrawler) GetContent() error {
	var jsTxtBody = "["
	var body string // splice contents
	var reContent = regexp.MustCompile(`"htmlTokens":\[\[(?P<contents>.*?)\]\]`)
	for _, v := range reContent.FindAllStringSubmatch(dwCrawler.Raw, -1) {
		jsTxtBody += v[1] + ","
	}
	if jsTxtBody == "[" { // this means jsTxtBody got northing, so it may be pic news.
		reContent = regexp.MustCompile(`"\d{7}":{"caption":"(?P<title>.*?)"`)
		for _, v := range reContent.FindAllStringSubmatch(dwCrawler.Raw, -1) {
			body += v[1] + "  \n"
		}
	} else {
		jsTxtBody = strings.ReplaceAll(jsTxtBody, "],[", ",")
		jsTxtBody = jsTxtBody[:len(jsTxtBody)-1] + "]" // now body json data prepared done.
		// Unmarshal the json data
		var paragraph []Paragraph
		err := json.Unmarshal([]byte(jsTxtBody), &paragraph)
		if err != nil {
			return fmt.Errorf("[-] fetcher.FmtBodyDwnews()>Unmarshal() Error: %q", err)
		}
		for _, p := range paragraph {
			if p.Type == "boldText" {
				body += "**" + p.Content + "**  \n"
			} else {
				body += p.Content + "  \n"
			}

		}
	}
	dwCrawler.Content = body

	return nil
}

func (dwCrawler *DwnewsCrawler) GetDatetime() error {
	var a = regexp.MustCompile(`(?m)<meta name="parsely-pub-date" content="(?P<date>.*?)".*?/>`)
	rt := a.FindStringSubmatch(dwCrawler.Raw)
	if rt != nil {
		dwCrawler.Datetime = rt[1]
		return nil
	}
	return errors.New("[-] (dwCrawler *DwnewsCrawler) GetDatetime(): regexp matched nothing.")
}

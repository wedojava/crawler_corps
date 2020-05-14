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

type DwnewsSite struct {
	corps.WebSite
}

type DwnewsPost struct {
	corps.Post
}

type CrawlerDwnews struct {
	Dwnews DwnewsSite
	Posts  []DwnewsPost
}

type Paragraph struct {
	Type    string
	Content string
}

func (web *DwnewsSite) GetUrls() error {
	var ret_lst []string
	var reLink = regexp.MustCompile(`(?m)<a\shref\s?=\s?"(?P<href>/.{2}/\d{8}/.+?)".*?>`)
	raw, err := corps.HttpGetBody(web.StartUrl, 10)
	if err != nil {
		return errors.Wrap(err, "\n[-] (web *DwNews)GetUrls()>HttpGetBody() error!\n[-] ")
	}
	for _, v := range reLink.FindAllStringSubmatch(raw, -1) {
		ret_lst = append(ret_lst, web.StartUrl+v[1])
	}
	web.Urls = gears.StrSliceDeDupl(ret_lst)
	return nil
}

func (dwPost *DwnewsPost) GetContent() error {
	var jsTxtBody = "["
	var body string // splice contents
	var reContent = regexp.MustCompile(`"htmlTokens":\[\[(?P<contents>.*?)\]\]`)
	for _, v := range reContent.FindAllStringSubmatch(dwPost.Raw, -1) {
		jsTxtBody += v[1] + ","
	}
	if jsTxtBody == "[" { // this means jsTxtBody got northing, so it may be pic news.
		reContent = regexp.MustCompile(`"\d{7}":{"caption":"(?P<title>.*?)"`)
		for _, v := range reContent.FindAllStringSubmatch(dwPost.Raw, -1) {
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
	dwPost.Content = body

	return nil
}

func (dwPost *DwnewsPost) GetDatetime() error {
	var a = regexp.MustCompile(`(?m)<meta name="parsely-pub-date" content="(?P<date>.*?)".*?/>`)
	rt := a.FindStringSubmatch(dwPost.Raw)
	if rt != nil {
		dwPost.Datetime = rt[1]
		return nil
	}
	return errors.New("[-] (dwPost *DwnewsPost) GetDatetime(): regexp matched nothing.")
}
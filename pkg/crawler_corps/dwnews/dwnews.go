package crawler_corps

import (
	"regexp"

	"github.com/pkg/errors"
	cc "github.com/wedojava/crawler_corps/crawler_corps"
	"github.com/wedojava/gears"
)

type CLeaderDwnews struct {
	cc.CrawlerLeader
}

func (cl *CLeaderDwnews) GetUrls() error {
	var ret_lst []string
	var reLink = regexp.MustCompile(`(?m)<a\shref\s?=\s?"(?P<href>/.{2}/\d{8}/.+?)".*?>`)
	raw, err := gears.HttpGetBody(cl.CrawlerLeader.StartUrl, 10)

	if err != nil {
		return errors.Wrap(err, "\n[-] (web *DwNews)GetUrls()>HttpGetBody() error!\n[-] ")
	}
	for _, v := range reLink.FindAllStringSubmatch(raw, -1) {
		ret_lst = append(ret_lst, cl.CrawlerLeader.StartUrl+v[1])
	}
	cl.CrawlerLeader.Urls = gears.StrSliceDeDupl(ret_lst)
	return nil

}

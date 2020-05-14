package corps

import (
	"errors"
	"regexp"
)

func (c *Crawler) GetTitle() error {
	var a = regexp.MustCompile(`(?m)<meta name="twitter:title" content="(?P<title>.*?)"`)
	rt := a.FindStringSubmatch(c.Raw)
	if rt != nil {
		c.Title = rt[1]
		return nil
	} else {
		return errors.New("\n[-] GetTitle()>FindStringSubmatch() Error!\n")
	}
}

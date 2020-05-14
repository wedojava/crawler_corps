package corps

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

type Post struct {
	Url      string
	Raw      string
	Title    string
	Content  string
	Datetime string
}

type WebSite struct {
	Domain   string
	StartUrl string
	Urls     []string
}

type IPost interface {
	GetContent()
	GetDatetime()
}

type IWebSite interface {
	GetUrls()
}

func (c *Post) GetTitle() error {
	var a = regexp.MustCompile(`(?m)<meta name="twitter:title" content="(?P<title>.*?)"`)
	rt := a.FindStringSubmatch(c.Raw)
	if rt != nil {
		c.Title = rt[1]
		return nil
	} else {
		return errors.New("\n[-] GetTitle()>FindStringSubmatch() Error!\n")
	}

}

func HttpGetBody(url string, n int) (string, error) {
	raw, err := http.Get(url)
	for err != nil && n > 0 {
		raw, err = http.Get(url)
		time.Sleep(time.Minute * 1)
		n--
	}
	if err != nil {
		return "", errors.Wrapf(err, "\n[-] HttpGetBody()>Get() try times, but error occur still!\n[-] ")
	}
	rawBody, err := ioutil.ReadAll(raw.Body)
	defer raw.Body.Close()
	if err != nil {
		return "", errors.Wrap(err, "\n[-] HttpGetBody()>ReadAll() Error!\n[-] ")
	}
	if raw.StatusCode != 200 {
		return "", errors.Wrap(err, "\n[-] HttpGetBody()>Get() Error! Message: Cannot open the url.\n[-] ")
	}
	return string(rawBody), nil
}

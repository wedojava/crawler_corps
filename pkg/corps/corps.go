package corps

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/wedojava/gears"
)

type Corps struct {
	Targets []CrawlerLeader
}

type Crawler struct {
	Url      string
	Raw      string
	Title    string
	Content  string
	Datetime string
}

type CrawlerLeader struct {
	Domain   string
	StartUrl string
	Urls     []string
	Crawlers []ICrawler
}

type ICrawler interface {
	// Setting done the Url and Raw
	GetRaw() error
	// Get then Set Content
	GetContent() error
	// Get then Set Datetime
	GetDatetime() error
	// Save to file
	Save()
}

type ICrawlerLeader interface {
	// Get urls from StartUrl
	GetUrls() error
	NewCrawlers()
	// Assign crawlers to fetching
	Assign()
	// Delete files saved days ago
	DelRoutine(daysOffset int) error
}

func (c *Corps) Start(icls []ICrawlerLeader) {
	for _, icl := range icls {
		icl.GetUrls()
		// icl.NewCrawlers()
		icl.Assign()
	}
}

// func (c *CrawlerLeader) NewCrawlers() {
//         for _, url := range c.Urls {
//                 crawler := &ICrawler{}
//                 crawler.Url = url
//                 c.Crawlers = append(c.Crawlers, *crawler)
//         }
// }
func (c *Crawler) GetContent() error {
	return nil
}

func (c *Crawler) GetRaw() error {
	raw, err := gears.HttpGetBody(c.Url, 10)
	if err != nil {
		return errors.Wrap(err, "\n[-] CrawlerInit()>gears.HttpGetBody() error.\n[-] ")
	}
	c.Raw = raw
	return nil
}

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

func (c *Crawler) Save() {
	t, err := time.Parse(time.RFC3339, c.Datetime)
	if err != nil {
		fmt.Printf("\n[-] SaveOneDwnew()>time.Parse(%s) error.\n%v\n", c.Datetime, err)
	}
	filename := fmt.Sprintf("[%02d.%02d][%02d%02dH]%s%s", t.Month(), t.Day(), t.Hour(), t.Minute(), c.Title, ".txt")
	// Save Body to file named title in folder twitter site content
	domainStr := gears.HttpGetDomain(c.Url)
	gears.MakeDirAll(domainStr)
	savePath := filepath.Join(domainStr, filename)
	if !gears.Exists(savePath) {
		err = ioutil.WriteFile(filepath.Join(domainStr, filename), []byte(c.Content), 0644)
		if err != nil {
			fmt.Printf("\n[-] SaveOneDwnew()>WriteFile(%s) error.\n%v\n", filepath.Join(domainStr, filename), err)
		}
	}
}

// DelRoutine remove files in folder days ago
func (cl *CrawlerLeader) DelRoutine(daysOffset int) error {
	folder := cl.Domain
	if !gears.Exists(folder) {
		fmt.Printf("\n[-] DelRoutine() err: Folder(%s) does not exist.\n", folder)
		return nil
	}
	for i := 0; i < 3; i++ { // deal with file n+i days ago
		a := time.Now().AddDate(0, 0, -(daysOffset + i))
		b := fmt.Sprintf("[%02d.%02d]", a.Month(), a.Day())
		c, _ := gears.GetPrefixedFiles(folder, b)
		for _, f := range c {
			// fmt.Println("DelRoutine will delete: ", f)
			os.Remove(f)
		}

	}
	return nil
}

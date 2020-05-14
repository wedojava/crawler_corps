package crawler_corps

type Corps struct {
	CLeaders []CrawlerLeader
}

type CrawlerLeader struct {
	Domain   string
	StartUrl string
	Urls     []string
}

type Crawler struct {
	Url      string
	Raw      string
	Title    string
	Content  string
	Datetime string
}

type ICorps interface {
	Fire()
}

type ICrawlerLeader interface {
	GetUrls()
	Assign(ICrawler)
}

type ICrawler interface {
	GetTitle()
	GetContent()
	GetDatetime()
}

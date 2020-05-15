package corps

import (
	"strings"
	"testing"

	"github.com/wedojava/crawler_corps/pkg/corps"
	"github.com/wedojava/gears"
)

var dwtest DwnewsSite = DwnewsSite{
	CrawlerLeader: corps.CrawlerLeader{
		Domain:   "www.dwnews.com",
		StartUrl: "https://www.dwnews.com",
		Urls:     nil,
	},
}

func TestGetUrls(t *testing.T) {
	dwtest.GetUrls()
	got := len(dwtest.Urls)
	want := 42
	if got <= want {
		t.Errorf("\n[GOT]:\n%v\n[WANT]\n%v\n", got, want)
	}
}

func TestGetContent(t *testing.T) {
	dwpost := &DwnewsCrawler{
		Crawler: corps.Crawler{
			Url:      "https://www.dwnews.com/全球/60178371",
			Raw:      "",
			Title:    "",
			Content:  "",
			Datetime: "",
		},
	}
	dwpost.Raw, _ = gears.HttpGetBody(dwpost.Url, 10)

	dwpost.GetContent()
	dwpost.GetTitle()
	wantTitle := "【新冠肺炎】北京给澳大利亚的经济教训 大国的强硬无差别"
	gotTitle := dwpost.Title
	if wantTitle != gotTitle {
		t.Errorf("\n[WANT]\n%v\n[GOT]\n%v\n", wantTitle, gotTitle)
	}

	wantContentHas := "对此，中国外交部先在4月23日批评了澳方的“政治操弄”行为，中国驻澳大使亦在4月28日接受采访时指出，澳方行为伤害中国感情，亦称中国消费者可能拒买澳红酒、牛肉等。至此，澳大利亚的遭遇业已成为北京与堪培拉之间心照不宣的结果。"
	gotContent := dwpost.Content
	if !strings.Contains(gotContent, wantContentHas) {
		t.Errorf("Can not find out the string the content must has.")
	}
}

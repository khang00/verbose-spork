package search

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	defaultCookies           = "1P_JAR=2024-02-11-13; AEC=Ae3NU9P7GiXs8p-whOlnW7jsXdTYMRZ72V8hCSofbb2D2zAKmrSOov45ow; DV=k7GnnwzkVFcV0J-HU_sc0eeC8jmG2Rg; NID=511=SHf6RR5xYryKuyYYvM6OysCj7fTuwyKDXpYXk5k-Jz5ZbJjP3LEsIGAMyjnpdoZ7fZLvOXwseoSX1FnKd3m35zgyFk8qXWGNYptjC3V78V5VqJEyYmHskDkOl8LKZZ4AiYrSjPJtkt8-tBfEqytu1_Go8gQDjTX9MUFGWxDScOwiBjzlqBNybDf333KiNButMk6MMU6fbzhUiWQI6PudN5OxAf8c5iPdDbY"
	defaultUserAgent         = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	defaultSearchURLTemplate = "https://www.google.com/search?q=%s"
	defaultStatSelector      = "#result-stats"
	defaultAdsSelector       = "#tads > div"
	defaultLinkSelector      = "a[href]"
	defaultLinkRegex         = "^https"
	defaultStatRegex         = "(\\d+\\.)+\\d+"
)

type GoogleQuerier struct {
	searchURLTemplate string
	statSelector      string
	adsSelector       string
	linkSelector      string
	linkRegex         *regexp.Regexp
	statRegex         *regexp.Regexp
	collector         *colly.Collector
}

func NewGoogleSearchQuerier() *GoogleQuerier {
	linkRegex, _ := regexp.Compile(defaultLinkRegex)
	statRegex, _ := regexp.Compile(defaultStatRegex)

	cookies := getCookiesFromString(defaultCookies)
	collector := colly.NewCollector()
	collector.SetCookies(defaultSearchURLTemplate, cookies)
	collector.UserAgent = defaultUserAgent

	return &GoogleQuerier{
		searchURLTemplate: defaultSearchURLTemplate,
		statSelector:      defaultStatSelector,
		adsSelector:       defaultAdsSelector,
		linkSelector:      defaultLinkSelector,
		linkRegex:         linkRegex,
		statRegex:         statRegex,
		collector:         collector,
	}
}

func getCookiesFromString(cookie string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookie)
	request := http.Request{Header: header}

	return request.Cookies()
}

func (r *GoogleQuerier) Search(keyword string) (*Result, error) {
	result := &Result{}
	var err error
	r.collector.OnHTML("html", func(e *colly.HTMLElement) {
		result.HTMLPage = r.parseHTMLPage(e)
		stats, errParse := r.parseStats(e)
		if errParse != nil {
			err = errParse
		}

		adsNum, errParse := r.parseAdvertises(e)
		if errParse != nil {
			err = errParse
		}

		links, errParse := r.parseLinks(e)
		if errParse != nil {
			err = errParse
		}

		result.ResultStats = stats
		result.NumberOfLinks = len(links)
		result.NumberOfAds = adsNum
	})

	url := r.buildUrl(keyword)
	err = r.collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GoogleQuerier) buildUrl(keyword string) string {
	return fmt.Sprintf(r.searchURLTemplate, keyword)
}

func (r *GoogleQuerier) parseHTMLPage(elem *colly.HTMLElement) string {
	return string(elem.Response.Body)
}

func (r *GoogleQuerier) parseStats(elem *colly.HTMLElement) (int, error) {
	statsText := elem.DOM.Find(r.statSelector).Text()
	statsNum := r.statRegex.Find([]byte(statsText))
	return strconv.Atoi(strings.Replace(string(statsNum), ".", "", -1))
}

func (r *GoogleQuerier) parseAdvertises(elem *colly.HTMLElement) (int, error) {
	return len(elem.DOM.Find(r.adsSelector).Nodes), nil
}

func (r *GoogleQuerier) parseLinks(elem *colly.HTMLElement) ([]Link, error) {
	linkAttr := "href"

	links := make([]string, 0)
	elem.DOM.Find(r.linkSelector).Each(func(i int, selection *goquery.Selection) {
		link, exists := selection.Attr(linkAttr)
		if exists && r.linkRegex.Match([]byte(link)) {
			links = append(links, link)
		}
	})

	return links, nil
}

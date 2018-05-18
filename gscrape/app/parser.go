package app

import (
	"net/url"

	"github.com/gocolly/colly"
)

type Link struct {
	Anchor string
	URL    string
}

func Parse(uri *url.URL) chan *Link {
	out := make(chan *Link)
	go func(out chan *Link) {
		defer close(out)
		c := colly.NewCollector(
			colly.AllowedDomains(uri.Host),
			colly.MaxDepth(1),
		)
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			if e.Attr("href") != "" {
				l := &Link{
					Anchor: e.Text,
					URL:    e.Request.AbsoluteURL(e.Attr("href")),
				}
				out <- l
				c.Visit(e.Request.AbsoluteURL(l.URL))
			}
		})
		c.Visit(uri.String())
	}(out)
	return out
}

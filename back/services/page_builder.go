package services

import (
	"back/repositories"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type PageBuilder struct {
	document           *goquery.Document
	url                string
	title              string
	headingsCount      int
	internalLinksCount int
	externalLinksCount int
}

func NewPageBuilder(url string) (*PageBuilder, error) {
	pageContent := getPageContent(url)

	document, err := prepareQueryableDocument(pageContent)

	if err != nil {
		return nil, err
	}

	return &PageBuilder{
		document: document,
		url:      url,
	}, nil
}

func (b *PageBuilder) GetPage() *repositories.Page {
	return &repositories.Page{
		URL:                b.url,
		Title:              b.title,
		HeadingsCount:      b.headingsCount,
		InternalLinksCount: b.internalLinksCount,
		ExternalLinksCount: b.externalLinksCount,
	}
}

func (b *PageBuilder) SetTitle() *PageBuilder {
	titleNode := b.document.Find("title").First().Nodes[0]

	if titleNode == nil {
		b.title = ""
		return b
	}

	b.title = titleNode.FirstChild.Data
	return b
}

func (b *PageBuilder) SetHeadingCount() *PageBuilder {
	count := 0

	b.document.Find("h1").Each(func(i int, s *goquery.Selection) {
		count++
	})

	b.document.Find("h2").Each(func(i int, s *goquery.Selection) {
		count++
	})

	b.document.Find("h3").Each(func(i int, s *goquery.Selection) {
		count++
	})

	b.document.Find("h4").Each(func(i int, s *goquery.Selection) {
		count++
	})

	b.document.Find("h5").Each(func(i int, s *goquery.Selection) {
		count++
	})

	b.document.Find("h6").Each(func(i int, s *goquery.Selection) {
		count++
	})

	b.headingsCount = count

	return b
}

func (b *PageBuilder) SetLinksCount() *PageBuilder {
	internal := 0
	external := 0

	b.document.Find("a").Each(func(i int, s *goquery.Selection) {
		href := s.Nodes[0]

		if href == nil {
			return
		}

		for _, attr := range href.Attr {
			if attr.Key != "href" {
				break
			}

			if !strings.HasPrefix(attr.Val, "http") {
				internal++
				break
			}

			if strings.HasPrefix(attr.Val, b.url) {
				internal++
				break
			}

			external++
		}
	})

	b.internalLinksCount = internal
	b.externalLinksCount = external

	return b
}

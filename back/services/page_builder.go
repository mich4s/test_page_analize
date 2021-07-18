package services

import (
	"back/repositories"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type PageBuilder struct {
	document           *goquery.Document
	documentContent    string
	url                string
	title              string
	htmlVersion        string
	headingsCount      int
	internalLinksCount int
	externalLinksCount int
	hasLoginForm       bool
}

func NewPageBuilder(url string) (*PageBuilder, error) {
	pageContent := getPageContent(url)

	document, err := prepareQueryableDocument(pageContent)

	if err != nil {
		return nil, err
	}

	documentContent, err := document.Html()

	if err != nil {
		return nil, err
	}

	return &PageBuilder{
		document:        document,
		url:             url,
		documentContent: documentContent,
	}, nil
}

func (b *PageBuilder) GetPage() *repositories.Page {
	return &repositories.Page{
		URL:                b.url,
		Title:              b.title,
		HTMLVersion:        b.htmlVersion,
		HeadingsCount:      b.headingsCount,
		InternalLinksCount: b.internalLinksCount,
		ExternalLinksCount: b.externalLinksCount,
		HasLoginForm:       b.hasLoginForm,
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

func (b *PageBuilder) SetHtmlVersion() *PageBuilder {

	if strings.HasPrefix(b.documentContent, "<!DOCTYPE html>") {
		b.htmlVersion = "HTML 5"
		return b
	}

	if strings.HasPrefix(b.documentContent, "# HTML 4.0") {
		b.htmlVersion = "HTML 4"
		return b
	}

	if strings.HasPrefix(b.documentContent, "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4") {
		b.htmlVersion = "HTML 4"
		return b
	}

	return b
}

func (b *PageBuilder) SetHeadingCount() *PageBuilder {
	count := 0

	b.document.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
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

func (b *PageBuilder) SetLoginForm() *PageBuilder {
	hasLoginForm := false

	b.document.Find("form").Each(func(i int, s *goquery.Selection) {
		s.Find("input").Each(func(k int, input *goquery.Selection) {
			for _, node := range input.Nodes {
				for _, attr := range node.Attr {
					if strings.Contains(attr.Val, "login") {
						hasLoginForm = true
						break
					}
				}
			}
		})
	})

	b.hasLoginForm = hasLoginForm

	return b
}

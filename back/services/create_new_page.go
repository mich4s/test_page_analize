package services

import (
	"back/repositories"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

func CreateNewPage(url string) (*repositories.Page, error) {
	pageBuilder, err := NewPageBuilder(url)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	pageBuilder.SetTitle().SetHeadingCount().SetLinksCount()

	page := pageBuilder.GetPage()

	return (&repositories.PagesRepository{}).Save(page)
}

func getPageContent(url string) io.Reader {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	return res.Body
}

func prepareQueryableDocument(documentContent io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(documentContent)
}

package repositories

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	URL                string
	Title              string
	HTMLVersion        string
	HeadingsCount      int
	InternalLinksCount int
	ExternalLinksCount int
	HasLoginForm       bool
}

type PagesRepository struct{}

func (r *PagesRepository) Save(page *Page) (*Page, error) {
	err := connection.Save(page).Error

	return page, err
}

func (r *PagesRepository) FindAll() ([]Page, error) {
	pages := make([]Page, 0)

	err := connection.Find(&pages).Error

	return pages, err
}

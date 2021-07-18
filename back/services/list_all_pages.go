package services

import "back/repositories"

func ListAllPages() ([]repositories.Page, error) {
	return (&repositories.PagesRepository{}).FindAll()
}

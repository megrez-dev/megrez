package dao

import "github.com/megrez/pkg/entity/po"

// ListAllLinks return all links
func (dao *DAO) ListAllLinks() ([]po.Link, error) {
	links := []po.Link{}
	result := dao.db.Find(&links)
	return links, result.Error
}

// CreateLink handle create link
func (dao *DAO) CreateLink(link *po.Link) error {
	result := dao.db.Create(&link)
	return result.Error
}

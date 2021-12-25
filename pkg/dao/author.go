package dao

import (
	"github.com/megrez/pkg/entity/po"
)

// GetAuthorByID return author by id
func (dao *DAO) GetAuthorByID(id uint) (po.Author, error) {
	author := po.Author{}
	result := dao.db.First(&author, id)
	return author, result.Error
}

// GetAuthorByMail return author by mail
func (dao *DAO) GetAuthorMyMail(mail string) (po.Author, error) {
	author := po.Author{}
	result := dao.db.First(&author, "`mail` = ?", mail)
	return author, result.Error
}

// CreateAuthor handle create author
func (dao *DAO) CreateAuthor(author *po.Author) error {
	result := dao.db.Create(author)
	return result.Error
}

// UpdateAuthor handle update author
func (dao *DAO) UpdateAuthor(author *po.Author) error {
	result := dao.db.Save(author)
	return result.Error
}

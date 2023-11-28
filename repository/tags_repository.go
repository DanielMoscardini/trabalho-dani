package repository

import "golang-crud-gin/model"

// define métodos para operações básicas de persistência relacionadas à entidade "tags"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}

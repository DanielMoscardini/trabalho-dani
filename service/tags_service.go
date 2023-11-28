package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

//  define métodos para manipulação de dados relacionados à entidade "tags"
type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}

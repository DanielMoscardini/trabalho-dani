package request

// // representa os dados de uma requisicao JSON na atualizacao de uma tag
type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

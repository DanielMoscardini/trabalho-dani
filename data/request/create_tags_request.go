package request

// representa os dados de uma requisicao JSON na criacao de uma tag
type CreateTagsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

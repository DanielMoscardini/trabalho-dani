package response

// representa os dados de resposta em formato JSON para a entidade tag
type TagsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

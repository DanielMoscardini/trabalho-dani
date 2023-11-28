package model

// define o modelo de dados para a entidade "tag", especificando campos como "Id" e "Name"
type Tags struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

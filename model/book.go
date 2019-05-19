package model

type Book struct {
	TableName []byte `json:"table_name" sql:"book.books"`
	Id int32 `json:"id" sql:",pk"`
	Name string `json:"name"`
	Author string `json:"author"`
	Category string `json:"category"`
}
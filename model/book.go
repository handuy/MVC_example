package model

import (
	"log"

	"github.com/go-pg/pg"
)

type Book struct {
	TableName []byte `json:"table_name" sql:"book.books"`
	Id        int32  `json:"id" sql:",pk"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Category  string `json:"category"`
	
}

type LessonContent struct {
	// Tạo bảng
	TableName []byte `json:"table_name" sql:"book.lesson_content"`
	// Id của item
	LessonId string `json:"lesson_id" sql:",pk"`
	// Id của item
	ContentId string `json:"content_id" sql:",pk"`
	// Loại content: text, slide, video, file, problem, quiz_list, exercise
	ContentType string `json:"content_type"`
	// Text
	Text string `json:"text"`
}

type BookService struct {
	DB *pg.DB
}

func NewBookService(db *pg.DB) (*BookService, error) {
	return &BookService{
		DB: db,
	}, nil
}

func (b *BookService) GetBooks() []Book {
	var books []Book
	_, err := b.DB.Query(&books, `
		SELECT *
		FROM book.books
	`)
	if err != nil {
		log.Println(err)
	}
	return books
}

func (b *BookService) GetLastBookId() int32 {
	var lastBookId int32
	_, err := b.DB.Query(&lastBookId, `
		SELECT max(id) AS last_book_id
		FROM book.books
	`)
	if err != nil {
		log.Println(err)
		return 0
	}
	return lastBookId
}

func (b *BookService) CreateBook(book *Book) error {
	err := b.DB.Insert(book)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Update test update các trường hợp
func (b *BookService) Update(book *Book, rawQueryWhere string, whereParams []interface{}, columnUpdate []string) error {

	_, err := b.DB.Model(book).Column(columnUpdate...).Where(rawQueryWhere, whereParams...).Update()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Update test update các trường hợp
func (b *BookService) DeleteLesson(lesson *LessonContent) error {

	_, err := b.DB.Model(lesson).WherePK().Update()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

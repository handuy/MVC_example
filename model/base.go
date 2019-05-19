package model

import (
	"fmt"
	"unicode"

	"git.hocngay.com/techmaster-example/config"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func ConnectDb(user string, password string, database string, address string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     address,
	})

	return db
}

//
func MigrationDb(db *pg.DB, config config.Config) error {
	// Tạo schema
	var schemas = []string{"book"}
	for _, schema := range schemas {
		_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + ";")
		if err != nil {
			return err
		}
	}

	var book Book
	err := createTable(&book, "book", "books", db, config)
	if err != nil {
		return err
	}

	return nil
}

// TableIsExists kiểm tra bảng đã tồn tại
func TableIsExists(schema, tableName string, db *pg.DB) (bool, error) {
	var exist bool
	_, err := db.Query(&exist, `
		SELECT EXISTS (
			SELECT 1
			FROM   information_schema.tables 
			WHERE  table_schema = ?
			AND    table_name = ?
			)`, schema, tableName)
	if err != nil {
		return exist, err
	}
	return exist, err
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

// LogQueryToConsole Log câu lệnh query
func LogQueryToConsole(db *pg.DB) {
	db.AddQueryHook(dbLogger{})
}

func createTable(model interface{}, schema, tableName string, db *pg.DB, config config.Config) error {
	exist, err := TableIsExists(schema, tableName, db)
	if err != nil {
		return err
	}
	if !exist {
		err = db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})

		if err != nil {
			return err
		}
	}

	return err
}

// ToSnake Change word to Snake Case
func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

// ResultResponse Response với mã trạng thái
type ResultResponse struct {
	// Kết quả trả về
	Status bool `json:"status"`
	// Thông báo lỗi
	Messages []string `json:"messages"`
}

// Response Trả về ResultResponse
type Response struct {
	Result *ResultResponse `json:"result"`
}

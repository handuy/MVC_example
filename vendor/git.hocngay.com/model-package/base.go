package model

import (
	"fmt"
	"unicode"

	"git.hocngay.com/techmaster-example/config"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

const (
	// Loại item
	COURSE int32 = 1
	COMBO  int32 = 2
	TRACK  int32 = 3

	// Loại khoá học
	COURSE_ONLINE  int32 = 1
	COURSE_OFFLINE int32 = 2

	// Trạng thái lớp
	OPEN_STATUS   int32 = 1
	CLOSED_STATUS int32 = 2

	// Hình thức lớp
	CLASS_ONLINE  int32 = 1
	CLASS_OFFLINE int32 = 2

	DRAFT_COURSE  int32 = 0
	PUBLIC_COURSE int32 = 1
	HIDDEN_COURSE int32 = 2

	MIN_DESCRIPTION_LENGTH int = 100
	MAX_DESCRIPTION_LENGTH int = 3000
	MIN_TITLE_LENGTH       int = 10
	MAX_TITLE_LENGTH       int = 200

	// CURRICULUM_PATH string = "/app/resources/curriculum"

	CURRICULUM_PATH string = "./resources/curriculum"

	ROLE_ADMIN   int32 = 1
	ROLE_TRAINER int32 = 2

	// Trạng thái đăng ký
	REGISTER_STATUS int32 = 1
	DELIVERY_STATUS int32 = 2
	DEBT_STATUS     int32 = 3
	COMPLETE_STATUS int32 = 100

	// Hình thức hỗ trợ
	EMAIL_SUPPORT int32 = 1
	PHONE_SUPPORT int32 = 2

	// Hình thức thanh toán
	CASH           int32 = 1
	TRANSFER       int32 = 2
	CHARGE_AT_HOME int32 = 3

	// Trạng thái support
	NEED_SUPPORT_STATUS   int32 = 1
	SUPPORTING_STATUS     int32 = 2
	CLOSED_SUPPORT_STATUS int32 = 100

	// Date layout
	LayoutISO string = "2006-1-2"
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
	var schemas = []string{"public"}
	for _, schema := range schemas {
		_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + ";")
		if err != nil {
			return err
		}
	}

	var class Class
	err := createTable(&class, "public", "class", db, config)
	if err != nil {
		return err
	}

	var student Student
	err = createTable(&student, "public", "student", db, config)
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

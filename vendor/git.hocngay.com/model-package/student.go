package model

import (
	"errors"

	"github.com/go-pg/pg/orm"
	"github.com/rs/xid"
)

type Student struct {
	// Tạo bảng
	TableName []byte `json:"table_name" sql:"public.student"`
	// Id
	Id string `json:"id" sql:",pk"`
	// Tên
	Name string `json:"name"`
}

func (student Student) InsertStudent(DB orm.DB) error {
	// Validate thông tin
	if student.Name == "" {
		return errors.New("Tên không được để trống")
	}

	// Insert dữ liệu
	student.Id = xid.New().String()
	err := DB.Insert(&student)
	if err != nil {
		return err
	}

	return nil
}

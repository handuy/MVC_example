package model

import (
	"errors"

	"github.com/go-pg/pg/orm"
	"github.com/rs/xid"
)

type Class struct {
	// Tạo bảng
	TableName []byte `json:"table_name" sql:"public.class"`
	// Id class
	Id string `json:"id" sql:",pk"`
	// Tên lớp
	Name string `json:"name"`
	// Id Khóa học
	CourseId string `json:"course_id"`
	// Hình thức lớp: 1=online, 2=offline
	Type int32 `json:"enroll_status" sql:"default:1"`
	// Trạng thái lớp: 1=mở, 2=đóng
	ClassStatus int32 `json:"class_status" sql:"default:1"`
	// Số buổi học, với lớp online thì số buổi học = 0
	Sessions int32 `json:"sessions" sql:"default:0"`
	// Số sinh viên
	Students int32 `json:"students" sql:"default:0"`
}

func (class Class) GetClasses(DB orm.DB) ([]Class, error) {
	var classes []Class
	err := DB.Model(&classes).Select()
	if err != nil {
		return nil, errors.New("Không lấy được danh sách lớp")
	}
	return classes, nil
}

func (class Class) GetLastClass(DB orm.DB) (Class, error) {
	var lastClass Class
	err := DB.Model(&lastClass).Order("id ASC").Limit(1).Select()
	if err != nil {
		return lastClass, errors.New("Không lấy được danh sách lớp")
	}
	return lastClass, nil
}

func (class Class) InsertClass(DB orm.DB) error {
	// Validate thông tin
	if class.Name == "" {
		return errors.New("Tên lớp không được để trống")
	}
	if class.CourseId == "" {
		return errors.New("Lớp học chưa liên kết đến một khoá học cụ thể")
	}
	if class.Type != CLASS_OFFLINE && class.Type != CLASS_ONLINE {
		return errors.New("Hình thức lớp không hợp lệ")
	}
	if class.Type == CLASS_OFFLINE {
		if class.Sessions <= 0 {
			return errors.New("Số buổi học phải lớn hơn 0")
		}
	} else {
		class.Sessions = 0
	}

	// Insert dữ liệu
	class.Id = xid.New().String()
	class.ClassStatus = OPEN_STATUS
	err := DB.Insert(&class)
	if err != nil {
		return err
	}

	return nil
}

func (class Class) UpdateClass(DB orm.DB) error {
	// Validate thông tin
	if class.Name == "" {
		return errors.New("Tên lớp không được để trống")
	}
	if class.CourseId == "" {
		return errors.New("Lớp học chưa liên kết đến một khoá học cụ thể")
	}
	if class.Type != CLASS_OFFLINE && class.Type != CLASS_ONLINE {
		return errors.New("Hình thức lớp không hợp lệ")
	}
	if class.Type == CLASS_OFFLINE {
		if class.Sessions <= 0 {
			return errors.New("Số buổi học phải lớn hơn 0")
		}
	} else {
		class.Sessions = 0
	}

	// Update dữ liệu
	err := DB.Update(&class)
	if err != nil {
		return err
	}

	return nil
}

func (class Class) DeleteClass(DB orm.DB) error {
	// Delete dữ liệu
	err := DB.Delete(&class)
	if err != nil {
		return err
	}

	return nil
}

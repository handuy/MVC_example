# Cấu trúc thư mục

1. vendor/git.hocngay.com/model-package
- Phần code định nghĩa model được tách thành một package riêng và import vào trong vendor, có thể tái sử dụng ở các project khác nhau (admin, anonymous, user, teacher)
- Là nơi định nghĩa struct tạo bảng và các method tương ứng của từng struct. Các method này thực hiện các CRUD operation lên bảng tương ứng
- Với một số method CUD, trước khi chọc vào database thì sẽ có logic validate dữ liệu.
Ví dụ:
```go
func (class Class) InsertClass(DB orm.DB) error {
	// Validate thông tin
	if class.Name == "" {
		return errors.New("Tên lớp không được để trống")
	}
	if class.CourseId == "" {
		return errors.New("Lớp học chưa liên kết đến một khoá học cụ thể")
	}
	if class.Type != constant.CLASS_OFFLINE && class.Type != constant.CLASS_ONLINE {
		return errors.New("Hình thức lớp không hợp lệ")
	}
	if class.Type == constant.CLASS_OFFLINE {
		if class.Sessions <= 0 {
			return errors.New("Số buổi học phải lớn hơn 0")
		}
	} else {
		class.Sessions = 0
	}

	// Insert dữ liệu
	class.Id = xid.New().String()
	class.ClassStatus = constant.OPEN_STATUS
	err := DB.Insert(&class)
	if err != nil {
		return err
	}

	return nil
}
```

2. view
- Chứa file HTML để render dữ liệu từ controller gửi vào

3. controller
- Chứa các hàm xử lý. Các hàm này được chạy mỗi khi client gọi vào 1 router tương ứng
- Trong mỗi hàm ở controller: Đầu tiên sẽ khởi tạo struct tương ứng với bảng cần CRUD dữ liệu, sau đó từ struct này sẽ gọi đến method đã được định nghĩa trong model. Ví dụ:
```go
func (c *Controller) Create(ctx iris.Context) {
	var class model.Class
	class.Name = "HTML"
	class.CourseId = "1"
    class.Type = 1
    
	err := class.InsertClass(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.Redirect("/")
}
```

4. router
- Chứa các route và hàm xử lý
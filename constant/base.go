package constant

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

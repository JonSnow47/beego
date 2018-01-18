package common

const (
	ErrSucceed      = "Succeed"      // 成功
	ErrInvalidParam = "InvalidParam" // 参数错误
	ErrInvalidPerm  = "InvalidPerm"  // 权限错误

	ErrInvalidAdmin = "Invalid Admin" // Admin不存在
	ErrAdminExists  = "Admin Exists"  // Admin已存在

	ErrInvalidUser   = "Invalid User"       // 用户不存在
	ErrUserExists    = "User Exists"        // 用户已存在
	ErrPassword         = "Incorrect password" // 密码错误
	ErrLoginRequired = "Login required"     // 未登录

	ErrSession = "Session error" // Session 错误

	ErrNoData     = "Not found data" // 未查询到数据
	ErrMysqlQuery = "MySQL error"    // MySQL 错误
)

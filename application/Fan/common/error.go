package common

const (
	ErrSucceed      = "Succeed"      // 成功
	ErrInvalidParam = "InvalidParam" // 参数错误
	ErrInvalidPerm  = "InvalidPerm"  // 权限错误

	ErrInvalidUser   = "ErrInvalidUser"                 // 用户不存在
	ErrUserExists    = "ErrUserExists"                  // 用户已存在
	ErrLogin         = "Incorrect username or password" // 密码错误
	ErrLoginRequired = "Login required"                 // 未登录

	ErrSession = "Session error" // Session 错误

	ErrNoData     = "Not found data" // 未查询到数据
	ErrMysqlQuery = "MySQL error"    // MySQL 错误
)

package e

/**
本包定义状态码
*/

const (
	// 普通的问题
	SUCCESS = 200		// 成功
	ERROR = 500			// 服务器错误
	INVALID_PARAMS = 400  // 权限问题

	// 文章相关
	ERROR_EXIST_TAG = 10001		// 已存在
	ERROR_NOT_EXIST_TAG = 10002		// 不存在
	ERROR_NOT_EXIST_ARTICLE = 10003  // 文章不存在

	// 权限相关
	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001   // 权限失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002  // 超时
	ERROR_AUTH_TOKEN = 20003  			// token信息错误
	ERROR_AUTH = 20004  // 错误
)
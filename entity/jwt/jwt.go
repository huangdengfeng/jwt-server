package jwt

var StandardJwtKey = []string{"sub", "iss", "aud", "exp", "nbf", "iat", "jti"}

type JwtInfo struct {
	// 主题，一般放关键业务标识字段，如用户ID
	Sub string
	// 签发者，一般可以放业务标识，如登录，短信验证
	Iss string
	// 观众，一般代表接收方
	Aud []string
	// 过期时间戳（秒）
	Exp int64
	// 不早于某个时间使用，时间戳（秒）
	Nbf int64
	// 签发时间
	Iat int64
	// jwt token id
	Jti string
	// 自定义属性
	Attributes map[string]string
}

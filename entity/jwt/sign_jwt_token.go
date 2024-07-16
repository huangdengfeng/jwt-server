package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"jwt-server/entity/errs"
	"slices"
)

func Sign(jwtInfo *JwtInfo, secretKey []byte) (string, error) {
	// 标准推荐的数据
	claims := make(jwt.MapClaims)
	if len(jwtInfo.Sub) > 0 {
		claims["sub"] = jwtInfo.Sub
	}
	if len(jwtInfo.Iss) > 0 {
		claims["iss"] = jwtInfo.Iss
	}
	if len(jwtInfo.Aud) > 0 {
		claims["aud"] = jwtInfo.Aud
	}
	if jwtInfo.Exp > 0 {
		claims["exp"] = jwtInfo.Exp
	}
	if jwtInfo.Nbf > 0 {
		claims["nbf"] = jwtInfo.Nbf
	}
	if jwtInfo.Iat > 0 {
		claims["iat"] = jwtInfo.Iat
	}
	if len(jwtInfo.Jti) > 0 {
		claims["jti"] = jwtInfo.Jti
	}
	// 自定义数据
	for k, v := range jwtInfo.Attributes {
		if slices.Contains(StandardJwtKey, k) {
			return "", errs.AttrKeyLimit.Newf(k)
		}
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secretKey)
	if err != nil {
		log.Errorf("jwt sign string err [%s]", err)
		return "", errs.JwtError.Newf(err)
	}
	return signed, nil
}

package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"jwt-server/entity/errs"
	"slices"
)
import log "github.com/sirupsen/logrus"

func Verify(token string, secretKey []byte) (*JwtInfo, error) {
	// 过期或者签名不对都会有错误
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// 确保token的方法符合预期
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.JwtError.Newf(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return secretKey, nil
	})
	if err != nil {
		log.Errorf("parse token err [%s]", err)
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errs.JwtTokenExpired
		}
		return nil, errs.JwtError.Newf(err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, errs.JwtError.Newf("token syntax invalid")
	}

	jwtInfo := &JwtInfo{}
	subject, err := claims.GetSubject()
	if err != nil {
		return nil, errs.JwtError.Newf(err)
	}
	jwtInfo.Sub = subject

	issuer, err := claims.GetIssuer()
	if err != nil {
		return nil, errs.JwtError.Newf(err)
	}
	jwtInfo.Iss = issuer

	audience, err := claims.GetAudience()
	if err != nil {
		return nil, errs.JwtError.Newf(err)
	}
	jwtInfo.Aud = audience

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, errs.JwtError.Newf(err)
	}
	if exp != nil {
		jwtInfo.Exp = exp.Unix()
	}

	nbf, err := claims.GetNotBefore()
	if err != nil {
		return nil, errs.JwtError.Newf(err)
	}
	if nbf != nil {
		jwtInfo.Nbf = nbf.Unix()
	}

	iat, err := claims.GetIssuedAt()
	if err != nil {
		return nil, errs.JwtError.Newf(err)
	}
	if iat != nil {
		jwtInfo.Iat = iat.Unix()
	}

	attributes := make(map[string]string)
	for k, v := range claims {
		if !slices.Contains(StandardJwtKey, k) {
			if s, ok := v.(string); ok {
				attributes[k] = s
			}
		}
	}
	jwtInfo.Attributes = attributes

	return jwtInfo, nil
}

func getId(claims map[string]any) string {
	raw, ok := claims["id"]
	if !ok {
		return ""
	}
	if id, ok := raw.(string); ok {
		return id
	}
	return ""
}

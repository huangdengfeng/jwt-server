package logic

import (
	"jwt-server/entity/config"
	"jwt-server/entity/jwt"
)

type JwtService struct {
}

func (j *JwtService) Sign(info *jwt.JwtInfo) (string, error) {
	signed, err := jwt.Sign(info, []byte(config.Global.Jwt.SignKey))
	return signed, err
}

func (j *JwtService) Verify(token string) (*jwt.JwtInfo, error) {
	jwtInfo, err := jwt.Verify(token, []byte(config.Global.Jwt.SignKey))
	return jwtInfo, err
}

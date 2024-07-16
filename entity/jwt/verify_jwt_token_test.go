package jwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"jwt-server/entity/errs"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	secretKey := []byte("123456")
	signed, err := Sign(&JwtInfo{
		Sub: "test sub",
		Iss: "www.seezoon.com",
		Aud: []string{"aud1", "aud2"},
		Exp: time.Now().Add(time.Second * 5).Unix(),
		Nbf: 0,
		Iat: 0,
		Jti: "xx-xxx-xxx",
		Attributes: map[string]string{
			"c1": "c1",
		},
	}, secretKey)
	assert.NoError(t, err)
	fmt.Println(signed)
	jwtInfo, err := Verify(signed, secretKey)
	assert.NoError(t, err)
	fmt.Printf("jwt info %+v", jwtInfo)
}

func TestVerifyExpire(t *testing.T) {
	secretKey := []byte("123456")
	signed, err := Sign(&JwtInfo{
		Sub: "test sub",
		Iss: "www.seezoon.com",
		Aud: []string{"aud1", "aud2"},
		Exp: time.Now().Add(-time.Second * 5).Unix(),
		Nbf: 0,
		Iat: 0,
		Jti: "xx-xxx-xxx",
		Attributes: map[string]string{
			"c1": "c1",
		},
	}, secretKey)
	assert.NoError(t, err)
	fmt.Println(signed)
	jwtInfo, err := Verify(signed, secretKey)
	assert.ErrorIs(t, err, errs.JwtTokenExpired)
	assert.Nil(t, jwtInfo)
}

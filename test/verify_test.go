package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"jwt-server/entity/errs"
	"jwt-server/entity/pb"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	signResp, err := client.Sign(context.Background(), &pb.SignReq{JwtInfo: &pb.JwtInfo{
		Sub:        "",
		Iss:        "",
		Aud:        nil,
		Exp:        0,
		Nbf:        0,
		Iat:        0,
		Jti:        "",
		Attributes: nil,
	}})
	assert.NoError(t, err)
	verifyResp, err := client.Verify(context.Background(), &pb.VerifyReq{
		Token: signResp.Token,
	})
	assert.NoError(t, err)
	assert.NotNil(t, verifyResp)
}

func TestVerifyErr(t *testing.T) {

	verifyResp, err := client.Verify(context.Background(), &pb.VerifyReq{
		Token: "123",
	})
	assert.Error(t, err)
	assert.True(t, err.(*errs.Error).Code == errs.BasArgs.Code)
	assert.Nil(t, verifyResp)
}
func TestVerifyExpired(t *testing.T) {
	signResp, err := client.Sign(context.Background(), &pb.SignReq{JwtInfo: &pb.JwtInfo{
		Sub:        "",
		Iss:        "",
		Aud:        nil,
		Exp:        time.Now().Add(-10 * time.Second).Unix(),
		Nbf:        0,
		Iat:        0,
		Jti:        "",
		Attributes: nil,
	}})
	assert.NoError(t, err)
	resp, err := client.Verify(context.Background(), &pb.VerifyReq{
		Token: signResp.Token,
	})
	assert.NoError(t, err)
	assert.False(t, resp.Valid)
}

func BenchmarkVerify(b *testing.B) {
	signResp, err := client.Sign(context.Background(), &pb.SignReq{JwtInfo: &pb.JwtInfo{
		Sub: "test subject",
		Iss: "xxxxxxxxxxxxxxxxx",
		Aud: nil,
		Exp: time.Now().Add(10 * time.Minute).Unix(),
		Nbf: 0,
		Iat: 0,
		Jti: "xxxxxxxxxxxxxxxxx",
		Attributes: map[string]string{
			"checkSum": "xxxxxxxxxxxxxxxxx",
		},
	}})
	assert.NoError(b, err)
	token := signResp.Token
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Verify(context.Background(), &pb.VerifyReq{
			Token: token,
		})
		assert.NoError(b, err)
		assert.True(b, resp.Valid)
	}
}

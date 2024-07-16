package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"jwt-server/entity/pb"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	resp, err := client.Sign(context.Background(), &pb.SignReq{
		JwtInfo: &pb.JwtInfo{
			Sub:        "test",
			Iss:        "",
			Aud:        nil,
			Exp:        time.Now().Add(time.Hour * 2).Unix(),
			Nbf:        0,
			Iat:        0,
			Jti:        "",
			Attributes: nil,
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Token)
}

func TestSignErr(t *testing.T) {
	signed, err := client.Sign(context.Background(), &pb.SignReq{
		JwtInfo: &pb.JwtInfo{
			Sub: "test",
			Iss: "",
			Aud: nil,
			Exp: time.Now().Add(time.Hour * 2).Unix(),
			Nbf: 0,
			Iat: 0,
			Jti: "",
			Attributes: map[string]string{
				"sub": "exists sub",
			},
		},
	})
	assert.Error(t, err)
	assert.Nil(t, signed)
}

func BenchmarkSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := client.Sign(context.Background(), &pb.SignReq{
			JwtInfo: &pb.JwtInfo{
				Sub: "test sub",
				Iss: "xxxxxxxx",
				Aud: nil,
				Exp: time.Now().Add(time.Hour * 2).Unix(),
				Nbf: 0,
				Iat: 0,
				Jti: "",
				Attributes: map[string]string{
					"checkSum": "xxxxxxxxxx",
				},
			},
		})
		assert.NoError(b, err)
		assert.NotEmpty(b, resp.Token)
	}
}

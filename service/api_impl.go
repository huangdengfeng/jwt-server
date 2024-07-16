package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"jwt-server/entity/jwt"
	"jwt-server/entity/pb"
	"jwt-server/logic"
)

var jwtService = &logic.JwtService{}

type JwtServerImpl struct {
	pb.UnimplementedJwtServer
}

func (j *JwtServerImpl) Sign(ctx context.Context, req *pb.SignReq) (*pb.SignResp, error) {
	log.Debugf("sign recieve:%+v", req)
	jwtInfo := &jwt.JwtInfo{
		Sub:        req.JwtInfo.Sub,
		Iss:        req.JwtInfo.Iss,
		Aud:        req.JwtInfo.Aud,
		Exp:        req.JwtInfo.Exp,
		Nbf:        req.JwtInfo.Nbf,
		Iat:        req.JwtInfo.Iat,
		Jti:        req.JwtInfo.Jti,
		Attributes: req.JwtInfo.Attributes,
	}
	signed, err := jwtService.Sign(jwtInfo)
	if err != nil {
		return nil, err
	}

	return &pb.SignResp{Token: signed}, nil
}

func (j *JwtServerImpl) Verify(ctx context.Context, req *pb.VerifyReq) (*pb.VerifyResp, error) {
	jwtInfo, err := jwtService.Verify(req.Token)
	if err != nil {
		log.Errorf("Verify error [%s]", err)
		return &pb.VerifyResp{Valid: false}, nil
	}
	return &pb.VerifyResp{
		Valid: true,
		JwtInfo: &pb.JwtInfo{
			Sub:        jwtInfo.Sub,
			Iss:        jwtInfo.Iss,
			Aud:        jwtInfo.Aud,
			Exp:        jwtInfo.Exp,
			Nbf:        jwtInfo.Nbf,
			Iat:        jwtInfo.Iat,
			Jti:        jwtInfo.Jti,
			Attributes: jwtInfo.Attributes,
		},
	}, nil
}

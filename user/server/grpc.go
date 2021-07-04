package server

import (
	"context"
	pb "user/proto"
	"user/utils"
)

func (us *GRPCServer) CheckToken(ctx context.Context, userInfoRequest *pb.TokenCheckRequest, userInfoResponse *pb.TokenCheckResponse) error {
	token := userInfoRequest.JwtToken
	ua, err := utils.ParseJWT(token)
	if err != nil {
		return err
	}
	userInfoResponse.UserId = ua.UserID
	userInfoResponse.UserName = ua.UserName

	return nil
}

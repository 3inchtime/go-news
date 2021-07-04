package server

import (
	"context"
	pb "user/proto"
	"user/utils"
)


func (us *GRPCServer) TokenCheck (ctx context.Context, request *pb.TokenCheckRequest, response *pb.TokenCheckResponse) error {
	token := request.JwtToken
	ua, err := utils.ParseJWT(token)
	if err != nil {
		return err
	}
	response.UserId = ua.UserID
	response.UserName = ua.UserName

	return nil
}

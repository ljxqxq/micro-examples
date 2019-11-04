package handler

import (
	"context"

	user "github.com/entere/micro-examples/user/srv/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResponse) error {

	userID := req.UserId
	if userID == "" {
		userID = "xxxxxx"
	}
	// 模拟用户信息，TODO：从数据库中取出用户信息
	userInfo := &user.UserInfo{
		UserId:    userID,
		Nickname:  "匿名",
		Mobile:    "138********",
		AvatarUrl: "avatar.jpg",
		Gender:    1,
	}
	rsp.Code = 200
	rsp.Error = nil
	rsp.Data = userInfo

	return nil
}

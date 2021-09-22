package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sony/sonyflake"
	v1 "loginDemo/api/login/v1"
	"time"

	//v1 "login/api/login/v1"
)

type LoginService struct {

	v1.UnimplementedLoginServer
	sony *sonyflake.Sonyflake
	log *log.Helper
	tempDb map[string]interface{}
}
type User struct {
	id	int64
	username string
	password string

}

func NewLoginService(logger log.Logger) *LoginService {
	sonyflake.NewSonyflake(sonyflake.Settings{})
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf := sonyflake.NewSonyflake(st)

	return &LoginService{
		sony:sf,
		log: log.NewHelper(log.With(logger, "module", "interface/login")),
		tempDb: make(map[string] interface{}),
	}
}


func (s *LoginService) ListUser(ctx context.Context, req *v1.ListUserReq) (*v1.ListUserReply, error) {
	return &v1.ListUserReply{}, nil
}
func (s *LoginService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	return &v1.GetUserReply{}, nil
}
func (s *LoginService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}
func (s *LoginService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}
func (s *LoginService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	tempId,err := s.sony.NextID()
	if err != nil{
		s.log.Error(err)
	}
	tempUser := &User{username: req.GetUsername(),password: req.GetPassword(),id: int64(tempId)}
	s.tempDb[tempUser.username] = tempUser

	return &v1.RegisterReply{Id:tempUser.id,Username: tempUser.username}, nil
}
func (s *LoginService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	if s.tempDb[req.GetUsername()] == nil{
		return nil,errors.New("name not find")
	}
	tempUser := s.tempDb[req.GetUsername()].(*User)
	if  tempUser.username !=req.GetUsername() || tempUser.password != req.GetPassword() {
		return nil,errors.New("username or password is err")
	}

	return &v1.LoginReply{Token: "this is a token"}, nil
}
func (s *LoginService) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}




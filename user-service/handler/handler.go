package handler

import (
	"context"
	"fmt"

	pb "github.com/leor-w/laracom/user-service/proto/user"
	"github.com/leor-w/laracom/user-service/repo"
	"github.com/leor-w/laracom/user-service/service"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo  repo.Repository
	Token service.Authable
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	err = srv.Repo.Create(req)
	if err != nil {
		return err
	}
	resp.User = req
	return nil
}

func (srv *UserService) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	logrus.Infof("Logging in with : [%s] [%s]", req.Email, req.Password)

	user, err := srv.Repo.GetByEmail(req.Email)
	if err != nil {
		logrus.Errorf("GetUserByEmail failed : %s", err.Error())
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logrus.Errorf("CompareHashAndPassword failed : %s", err.Error())
		return err
	}

	token, err := srv.Token.Encode(user)
	if err != nil {
		logrus.Errorf("Encode Token failed : User [%v] error [%v]", user, err)
		return err
	}

	resp.Token = token
	return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	claims, err := srv.Token.Decode(req.Token)
	if err != nil {
		logrus.Errorf("Decode Token failed: Token [%v] error [%v]", req.Token, err)
		return err
	}

	if claims.User.Id == "" {
		logrus.Errorf("Invalid user: token has [%v]", req)
		return fmt.Errorf("Invalid user: token has [%v]", req)
	}

	resp.Valid = true

	return nil
}

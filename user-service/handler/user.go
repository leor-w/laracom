package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/leor-w/laracom/user-service/model"
	pb "github.com/leor-w/laracom/user-service/proto/user"
	"github.com/leor-w/laracom/user-service/repo"
	"github.com/leor-w/laracom/user-service/service"
	"github.com/micro/go-micro/v2/broker"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const topic = "password.reset"

type UserService struct {
	Repo      repo.UserRepositoryInterface
	ResetRepo repo.PasswordRepositoryInterface
	Token     service.Authable
	PubSub    broker.Broker
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	var (
		user    = &pb.User{}
		modUser *model.User
		id      uint64
		err     error
	)

	if req.Id != "" {
		id, err = strconv.ParseUint(req.Id, 10, 64)
		modUser, err = srv.Repo.Get(uint(id))
	} else if req.Email != "" {
		modUser, err = srv.Repo.GetByEmail(req.Email)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if modUser != nil {
		user, _ = modUser.ToProtoBuf()
	}
	resp.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	respUsers := []*pb.User{}
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		pbUser, _ := user.ToProtoBuf()
		respUsers = append(respUsers, pbUser)
	}
	resp.Users = respUsers
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	modUser := &model.User{}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	modUser.ToORM(req)
	err = srv.Repo.Create(modUser)
	if err != nil {
		return err
	}
	resp.User, _ = modUser.ToProtoBuf()
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

	if claims.User.Id <= 0 {
		logrus.Errorf("Invalid user: token has [%v]", req)
		return fmt.Errorf("Invalid user: token has [%v]", req)
	}

	resp.Valid = true

	return nil
}

func (srv *UserService) Update(ctx context.Context, req *pb.User, resp *pb.Response) error {
	model := &model.User{}
	if req.Id == "" {
		return fmt.Errorf("用户 [ID] 不能为空")
	}
	if req.Password != "" {
		hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			logrus.Errorf("Pass password hash failed, error was : %v", err)
			return err
		}
		req.Password = string(hashPass)
	}
	model.ToORM(req)
	if err := srv.Repo.Update(model); err != nil {
		logrus.Errorf("Update password failed, error was : %v", err)
		return err
	}
	resp.User, _ = model.ToProtoBuf()
	return nil
}

func (srv *UserService) CreatePasswordReset(ctx context.Context, req *pb.PasswordReset, resp *pb.PasswordResetResponse) error {
	passwordReset := &model.PasswordReset{}
	if req.Email == "" {
		logrus.Errorf("CreatePasswordReset failed : email field can not be empty!")
		return fmt.Errorf("邮箱不能为空")
	}
	passwordReset.ToORM(req)
	err := srv.ResetRepo.Create(passwordReset)
	if err != nil {
		logrus.Errorf("Insert PasswordReset failed : %v", err)
		return err
	}
	if passwordReset != nil {
		resp.PasswordReset, _ = passwordReset.ToProtoBuf()
		err := srv.publishEvent(resp.PasswordReset)
		if err != nil {
			return err
		}
	}
	return nil
}

func (srv UserService) ValidatePasswordResetToke(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	if req.Token == "" {
		logrus.Errorf("Token field can not be empty!")
		return fmt.Errorf("Token 不能为空")
	}
	_, err := srv.ResetRepo.GetByToken(req.Token)
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf("GetPasswordResetByToken failed : %v", err)
		return fmt.Errorf("数据库查询异常")
	}

	if err != gorm.ErrRecordNotFound {
		resp.Valid = true
	}
	return nil
}

func (srv *UserService) DeletePasswordReset(ctx context.Context, req *pb.PasswordReset, resp *pb.PasswordResetResponse) error {
	modPassword := &model.PasswordReset{}
	if req.Email == "" {
		return errors.New("Email 不能为空")
	}

	_, err := srv.ResetRepo.GetByEmail(req.Email)
	if err != nil {
		return errors.New("查询数据库出错")
	}

	modPassword.ToORM(req)
	err = srv.ResetRepo.DeletePasswordReset(modPassword)
	if err != nil {
		return err
	}

	resp.PasswordReset = nil
	return nil
}

func (srv *UserService) publishEvent(reset *pb.PasswordReset) error {
	body, err := json.Marshal(reset)
	if err != nil {
		logrus.Errorf("publishEvent marshal json faild %v", err)
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"email": reset.Email,
		},
		Body: body,
	}

	if err := srv.PubSub.Publish(topic, msg); err != nil {
		logrus.Errorf("publishEvent publish message failed : %v", err)
		return err
	}
	return nil
}

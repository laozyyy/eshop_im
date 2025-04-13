package rpc

import (
	"context"
	"eshop_im/kitex_gen/eshop/user_info"
	"eshop_im/kitex_gen/eshop/user_info/userservice"
	"eshop_im/log"
	"eshop_im/model"

	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func init() {
	var err error
	userClient, err = userservice.NewClient("hello", client.WithHostPorts("117.72.72.114:20000"))
	//userClient, err = userservice.NewClient("hello", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Errorf("error: %v", err)
	}
}

func GetOneUserByName(ctx context.Context, name string) (r *model.User, err error) {
	resp, err := userClient.GetOneUserByName(ctx, name)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if resp.User == nil {
		return nil, nil
	}
	return &model.User{
		UID:      resp.User.Uid,
		Name:     resp.User.Name,
		Phone:    *resp.User.Phone,
		Email:    *resp.User.Email,
		Password: resp.User.Password,
		Role:     int(resp.User.Role),
	}, nil
}

func GetOneUserById(ctx context.Context, uid string) (r *model.User, err error) {
	resp, err := userClient.GetOneUser(ctx, &user_info.GetOneUserRequest{Uid: uid})
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if resp.User == nil {
		return nil, nil
	}
	return &model.User{
		UID:      resp.User.Uid,
		Name:     resp.User.Name,
		Phone:    *resp.User.Phone,
		Email:    *resp.User.Email,
		Password: resp.User.Password,
		Role:     int(resp.User.Role),
	}, nil
}

func InsertOneUser(ctx context.Context, user *model.User) (err error) {
	uuser := &user_info.User{
		Uid:      user.UID,
		Name:     user.Name,
		Phone:    &user.Phone,
		Email:    &user.Email,
		Password: user.Password,
		Role:     int32(user.Role),
	}
	_, err = userClient.InsertOneUser(ctx, uuser)
	if err != nil {
		log.Errorf("error: %v", err)
		return err
	}
	return nil
}

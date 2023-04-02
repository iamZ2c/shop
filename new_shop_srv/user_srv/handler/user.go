package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"new_shop_srv/user_srv/global"
	"new_shop_srv/user_srv/model"
	"new_shop_srv/user_srv/proto"
	"strings"
	"time"
)

//GetUserList(context.Context, *PageInfo) (*UserInfoResponse, error)
//GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
//GetUserByID(context.Context, *IDRequest) (*UserInfoResponse, error)
//CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
//UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
//CheckUser(context.Context, *PasswordCheckInfo) (*CheckResponse, error)

type UserServer struct{}

func Model2UserListResp(user model.User) *proto.UserInfoResponse {
	// grpc message 中字段有默认值，不能为nil

	UserInfoResp := proto.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Password: user.Password,
		Role:     uint32(user.Role),
	}

	if user.Birthday != nil {
		UserInfoResp.BrithDay = uint64(user.Birthday.Unix())
	}
	return &UserInfoResp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (us *UserServer) GetUserList(ctx context.Context, pi *proto.PageInfo) (*proto.UserListResponse, error) {
	// 获取用户列表
	var users []model.User
	res := global.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	resp := &proto.UserListResponse{}
	resp.Total = int32(res.RowsAffected)

	global.DB.Scopes(Paginate(int(pi.Pn), int(pi.PSize))).Find(&users)

	for _, user := range users {
		userInfoResp := Model2UserListResp(user)
		resp.Data = append(resp.Data, userInfoResp)

	}
	return resp, nil
}

func (us *UserServer) GetUserByMobile(ctx context.Context, mReq *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	// 通过手机号查询用户
	var user model.User
	mobile := mReq.Mobile
	// 通过字符串查询
	res := global.DB.Where("mobile=?", mobile).First(&user)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	resp := Model2UserListResp(user)
	return resp, nil
}

func (us *UserServer) GetUserByID(ctx context.Context, req *proto.IDRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	id := req.Id
	// 主键查询方式
	res := global.DB.First(&user, id)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	resp := Model2UserListResp(user)
	return resp, nil
}

func (us *UserServer) CreateUser(ctx context.Context, CreateUserInfo *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	// 新建用户
	var user model.User
	NickName := CreateUserInfo.NickName
	Mobile := CreateUserInfo.Mobile
	Password := CreateUserInfo.Password
	// 先查询是否有手机号
	res := global.DB.Where("mobile=?", Mobile).First(&user)
	if res.Error != nil {
		// 处理查询不到的时候报错
		if res.Error != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, res.Error.Error())
		}
	}
	if res.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "user is created")
	}
	// 密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(Password, options)
	dbPassWord := fmt.Sprintf("$pbkdf2-sha256$%s$%s", salt, encodedPwd)

	user = model.User{NickName: NickName, Password: dbPassWord, Mobile: Mobile}
	res = global.DB.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	resp := Model2UserListResp(user)
	return resp, nil
}

func (us *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	res := global.DB.First(user, req.Id)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	user.NickName = req.NickName
	user.Gender = req.Gender
	Birthday := time.Unix(int64(req.BrithDay), 0)
	user.Birthday = &Birthday
	res = global.DB.Save(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (us *UserServer) CheckUser(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	dbPassWord := req.EncryptedPassword

	s := strings.Split(dbPassWord, "$")
	salt := s[2]
	encodedPwd := s[3]
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	check := password.Verify(req.Password, salt, encodedPwd, options)
	return &proto.CheckResponse{Success: check}, nil

}

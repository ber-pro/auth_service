package converter

import (
	"auth/internal/model"
	desc "auth/pkg/user_v1"
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelUserInfoToUserInfo(info *model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     desc.Role(info.Role),
	}
}

func ModelUserToUser(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}
	return &desc.User{
		Id:        user.Id,
		Info:      ModelUserInfoToUserInfo(&user.UserInfo),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func UserInfoToModeUserInfo(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     int32(info.Role),
	}
}

func UserToModelUser(user *desc.User) *model.User {
	var updateAt sql.NullTime
	_ = updateAt.Scan(user.UpdatedAt.AsTime())

	return &model.User{
		Id:        user.Id,
		UserInfo:  *UserInfoToModeUserInfo(user.Info),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: updateAt,
	}
}

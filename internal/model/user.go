package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	UserInfo  UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     int32
}

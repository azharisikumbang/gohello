package domain

import "time"

type User struct {
	Id          int
	Username    string
	Password    string
	UserGroupId string
	CreatedAt   *time.Time
}

package infra

import "time"

type UserModel struct {
	Id          uint64    `db:"id"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	UserGroupId uint64    `db:"user_group_id"`
}

type UserGroupModel struct {
	Id          uint64 `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Number      uint8  `db:"number"` // lower more powerfull, for example root will be number 1
}

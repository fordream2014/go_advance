package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type UserDao struct {
	db *MysqlBase
}

// new
func NewUserDao() *UserDao {
	return &UserDao{
		db: &MysqlBase{},
	}
}

// 查询用户信息
func (u *UserDao) GetUserInfo(ctx context.Context, uid string) (map[string]interface{}, error) {
	if uid == "" {
		return nil, nil
	}
	userInfo, err := u.db.QueryRow()
	if err == sql.ErrNoRows {
		fmt.Printf("query no row\n")
		return nil, nil
	}
	return userInfo, errors.Wrap(err, "query user info fail")
}

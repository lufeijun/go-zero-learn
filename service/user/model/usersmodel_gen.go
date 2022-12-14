// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByMobile(ctx context.Context, mobile string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`     // 用户姓名
		Gender     int64     `db:"gender"`   // 用户性别
		Mobile     string    `db:"mobile"`   // 用户电话
		Password   string    `db:"password"` // 用户密码
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUsersModel(conn sqlx.SqlConn) *defaultUsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByMobile(ctx context.Context, mobile string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Gender, data.Mobile, data.Password)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.Gender, newData.Mobile, newData.Password, newData.Id)
	return err
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}

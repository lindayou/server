// Code generated by goctl. DO NOT EDIT.

package user_auth

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserAuthorityFieldNames          = builder.RawFieldNames(&SysUserAuthority{})
	sysUserAuthorityRows                = strings.Join(sysUserAuthorityFieldNames, ",")
	sysUserAuthorityRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserAuthorityFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysUserAuthorityRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserAuthorityFieldNames, "`sys_user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysUserAuthorityModel interface {
		Insert(ctx context.Context, data *SysUserAuthority) (sql.Result, error)
		FindOne(ctx context.Context, sysUserId int64) (*SysUserAuthority, error)
		Update(ctx context.Context, data *SysUserAuthority) error
		Delete(ctx context.Context, sysUserId int64) error
	}

	defaultSysUserAuthorityModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUserAuthority struct {
		SysUserId               int64 `db:"sys_user_id"`
		SysAuthorityAuthorityId int64 `db:"sys_authority_authority_id"` // 角色ID
	}
)

func newSysUserAuthorityModel(conn sqlx.SqlConn) *defaultSysUserAuthorityModel {
	return &defaultSysUserAuthorityModel{
		conn:  conn,
		table: "`sys_user_authority`",
	}
}

func (m *defaultSysUserAuthorityModel) withSession(session sqlx.Session) *defaultSysUserAuthorityModel {
	return &defaultSysUserAuthorityModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`sys_user_authority`",
	}
}

func (m *defaultSysUserAuthorityModel) Delete(ctx context.Context, sysUserId int64) error {
	query := fmt.Sprintf("delete from %s where `sys_user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, sysUserId)
	return err
}

func (m *defaultSysUserAuthorityModel) FindOne(ctx context.Context, sysUserId int64) (*SysUserAuthority, error) {
	query := fmt.Sprintf("select %s from %s where `sys_user_id` = ? limit 1", sysUserAuthorityRows, m.table)
	var resp SysUserAuthority
	err := m.conn.QueryRowCtx(ctx, &resp, query, sysUserId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserAuthorityModel) Insert(ctx context.Context, data *SysUserAuthority) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysUserAuthorityRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SysUserId, data.SysAuthorityAuthorityId)
	return ret, err
}

func (m *defaultSysUserAuthorityModel) Update(ctx context.Context, data *SysUserAuthority) error {
	query := fmt.Sprintf("update %s set %s where `sys_user_id` = ?", m.table, sysUserAuthorityRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.SysAuthorityAuthorityId, data.SysUserId)
	return err
}

func (m *defaultSysUserAuthorityModel) tableName() string {
	return m.table
}

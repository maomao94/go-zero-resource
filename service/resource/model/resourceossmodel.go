package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	resourceOssFieldNames          = builder.RawFieldNames(&ResourceOss{})
	resourceOssRows                = strings.Join(resourceOssFieldNames, ",")
	resourceOssRowsExpectAutoSet   = strings.Join(stringx.Remove(resourceOssFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	resourceOssRowsWithPlaceHolder = strings.Join(stringx.Remove(resourceOssFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheResourceOssIdPrefix = "cache:resourceOss:id:"
)

type (
	ResourceOssModel interface {
		Insert(data ResourceOss) (sql.Result, error)
		FindOne(id int64) (*ResourceOss, error)
		Update(data ResourceOss) error
		Delete(id int64) error
	}

	defaultResourceOssModel struct {
		sqlc.CachedConn
		table string
	}

	ResourceOss struct {
		Id         int64          `db:"id"`
		CreatedAt  sql.NullTime   `db:"created_at"`
		UpdatedAt  sql.NullTime   `db:"updated_at"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
		TenantId   string         `db:"tenant_id"`   // 租户ID
		Category   sql.NullInt64  `db:"category"`    // 分类
		OssCode    sql.NullString `db:"oss_code"`    // 资源编号
		Endpoint   sql.NullString `db:"endpoint"`    // 资源地址
		AccessKey  sql.NullString `db:"access_key"`  // accessKey
		SecretKey  sql.NullString `db:"secret_key"`  // secretKey
		BucketName sql.NullString `db:"bucket_name"` // 空间名
		AppId      sql.NullString `db:"app_id"`      // 应用ID
		Region     sql.NullString `db:"region"`      // 地域简称
		Remark     sql.NullString `db:"remark"`      // 备注
		Status     sql.NullInt64  `db:"status"`      // 状态
	}
)

func NewResourceOssModel(conn sqlx.SqlConn, c cache.CacheConf) ResourceOssModel {
	return &defaultResourceOssModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`resource_oss`",
	}
}

func (m *defaultResourceOssModel) Insert(data ResourceOss) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, resourceOssRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.TenantId, data.Category, data.OssCode, data.Endpoint, data.AccessKey, data.SecretKey, data.BucketName, data.AppId, data.Region, data.Remark, data.Status)

	return ret, err
}

func (m *defaultResourceOssModel) FindOne(id int64) (*ResourceOss, error) {
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, id)
	var resp ResourceOss
	err := m.QueryRow(&resp, resourceOssIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceOssRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultResourceOssModel) Update(data ResourceOss) error {
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, resourceOssRowsWithPlaceHolder)
		return conn.Exec(query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.TenantId, data.Category, data.OssCode, data.Endpoint, data.AccessKey, data.SecretKey, data.BucketName, data.AppId, data.Region, data.Remark, data.Status, data.Id)
	}, resourceOssIdKey)
	return err
}

func (m *defaultResourceOssModel) Delete(id int64) error {

	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, resourceOssIdKey)
	return err
}

func (m *defaultResourceOssModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, primary)
}

func (m *defaultResourceOssModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceOssRows, m.table)
	return conn.QueryRow(v, query, primary)
}

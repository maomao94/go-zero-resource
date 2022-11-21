package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TOssModel = (*customTOssModel)(nil)

type (
	// TOssModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTOssModel.
	TOssModel interface {
		tOssModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*TOss, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TOss, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TOss, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TOss, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TOss, error)
		UpdateWithVersion(ctx context.Context, data *TOss) error
		DeleteSoft(ctx context.Context, data *TOss) error
	}

	customTOssModel struct {
		*defaultTOssModel
	}
)

// NewTOssModel returns a model for the database table.
func NewTOssModel(conn sqlx.SqlConn, c cache.CacheConf) TOssModel {
	return &customTOssModel{
		defaultTOssModel: newTOssModel(conn, c),
	}
}

func (m *defaultTOssModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultTOssModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(tOssRows).From(m.table)
}

func (m *defaultTOssModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultTOssModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultTOssModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*TOss, error) {
	query, values, err := rowBuilder.Where("del_state = ?", DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}
	var resp TOss
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTOssModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {
	query, values, err := sumBuilder.Where("del_state = ?", DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}
	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultTOssModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {
	query, values, err := countBuilder.Where("del_state = ?", DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}
	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultTOssModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TOss, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TOss
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTOssModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TOss, error) {
	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query, values, err := rowBuilder.Where("del_state = ?", DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*TOss
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTOssModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TOss, error) {
	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}
	query, values, err := rowBuilder.Where("del_state = ?", DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*TOss
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTOssModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TOss, error) {
	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}
	query, values, err := rowBuilder.Where("del_state = ?", DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*TOss
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTOssModel) UpdateWithVersion(ctx context.Context, newData *TOss) error {
	oldVersion := newData.Version
	newData.Version += 1
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	tOssIdKey := fmt.Sprintf("%s%v", cacheTOssIdPrefix, data.Id)
	tOssTenantIdOssCodeKey := fmt.Sprintf("%s%v:%v", cacheTOssTenantIdOssCodePrefix, data.TenantId, data.OssCode)
	sqlResult, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, tOssRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DelState, newData.Version, data.TenantId, data.Category, data.OssCode, data.Endpoint, data.AccessKey, data.SecretKey, data.BucketName, data.AppId, data.Region, data.Remark, data.Status, newData.Id, oldVersion)
	}, tOssIdKey, tOssTenantIdOssCodeKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return errors.New("UpdateWithVersion error.")
	}
	return err
}

func (m *defaultTOssModel) DeleteSoft(ctx context.Context, newData *TOss) error {
	newData.DelState = DelStateYes
	return m.UpdateWithVersion(ctx, newData)
}

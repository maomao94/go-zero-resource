package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TOssModel = (*customTOssModel)(nil)

type (
	// TOssModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTOssModel.
	TOssModel interface {
		tOssModel
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

package db

import (
	"context"

	"github.com/baisiyi/file_trans/config"
	"github.com/baisiyi/file_trans/models/table"
	db "github.com/baisiyi/gorm_helper/mysql/db"
	dbHelper "github.com/baisiyi/gorm_helper/mysql/db_helper"
)

type readerImpl struct {
	*dbHelper.DbHelper
}

func newReaderImpl() DbReader {
	return &readerImpl{
		db.GetDbHelper(config.GetDSN()),
	}
}

func (r *readerImpl) GetListByCondition(ctx context.Context,
	where string, args []interface{}) (fileList []*table.FileMeta, err error) {
	err = r.GetDb(ctx).Model(&table.FileMeta{}).Where(where, args...).Find(fileList).Error
	return
}

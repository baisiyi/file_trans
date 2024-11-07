package db

import (
	"context"

	"github.com/baisiyi/file_trans/config"
	"github.com/baisiyi/file_trans/models/table"
	db "github.com/baisiyi/gorm_helper/mysql/db"
	dbHelper "github.com/baisiyi/gorm_helper/mysql/db_helper"
)

type writerImpl struct {
	*dbHelper.DbHelper
}

func newWriterImpl() DbWriter {
	dsn := config.GetDSN()
	return &writerImpl{
		db.GetDbHelper(dsn),
	}
}

func (w *writerImpl) CreateFile(ctx context.Context, fileMeta *table.FileMeta) error {
	return w.GetDb(ctx).Model(&table.FileMeta{}).Create(fileMeta).Error
}

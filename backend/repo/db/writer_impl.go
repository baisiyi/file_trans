package db

import (
	"context"

	db "github.com/baisiyi/gorm_helper/mysql/db"
	dbHelper "github.com/baisiyi/gorm_helper/mysql/db_helper"
	"github.com/siyibai/file-transfer/config"
	"github.com/siyibai/file-transfer/models/table"
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

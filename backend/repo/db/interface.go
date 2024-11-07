package db

import (
	"context"

	"github.com/baisiyi/file_trans/models/table"
)

type DbReader interface {
	GetListByCondition(ctx context.Context, where string, args []interface{}) ([]*table.FileMeta, error)
}

type DbWriter interface {
	CreateFile(ctx context.Context, fileMeta *table.FileMeta) error
}

func NewDbWriter() DbWriter {
	return newWriterImpl()
}

func NewDbReader() DbReader {
	return newReaderImpl()
}

package db

type DbReader interface{}

type DbWriter interface{}

func NewDbWriter() DbWriter {
	return newWriterImpl()
}

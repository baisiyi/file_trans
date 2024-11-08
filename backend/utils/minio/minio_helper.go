package minioHelper

import (

	"github.com/minio/minio-go/v7"
)

type MinioHelper struct {
	cli *minio.Client
}

func NewMinioHelper(cli *minio.Client) *MinioHelper {
	return &MinioHelper{cli: cli}
}



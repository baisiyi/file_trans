package minio

import "github.com/minio/minio-go/v7"

type MinioRepo interface {
}

type minioRepo struct {
	minio.Client
}

func NewMinioRepo() MinioRepo {
	return &minioRepo{}
}

package minio

type MinioRepo interface {
}

type minioRepo struct {
}

func NewMinioRepo() MinioRepo {
	return &minioRepo{}
}

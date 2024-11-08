package minioHelper

import (
	"github.com/baisiyi/file_trans/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioPool struct {
	pool   chan *minio.Client
	config *config.MinioConfig
	size   int
}

func NewMinioPool(config *config.MinioConfig) (*MinioPool, error) {
	pool := &MinioPool{
		pool:   make(chan *minio.Client, config.PoolSize),
		config: config,
		size:   config.PoolSize,
	}
	for i := 0; i < config.PoolSize; i++ {
		cli, err := minio.New(config.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
			Secure: config.UseSSL,
		})
		if err != nil {
			panic("minio连接错误: " + err.Error())
		}
		pool.pool <- cli
	}
	return pool, nil
}

func (p *MinioPool) GetClient() (*minio.Client, error) {
	select {
	case client := <-p.pool:
		return client, nil
	default:
		return minio.New(
			p.config.Endpoint,
			&minio.Options{
				Creds:  credentials.NewStaticV4(p.config.AccessKeyID, p.config.SecretAccessKey, ""),
				Secure: p.config.UseSSL,
			},
		)
	}
}

func (p *MinioPool) ReleaseClient(client *minio.Client) {
	select {
	case p.pool <- client:
		// 成功放回连接池
	default:
	}
}

func GetMinioHelper() *MinioHelper {
	pool, err := NewMinioPool(&config.GetConfig().MinioCfg)
	if err != nil {
		return nil
	}
	cli, err := pool.GetClient()
	if err != nil {
		return nil
	}
	return NewMinioHelper(cli)
}

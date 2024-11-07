package fileServer

import (
	pb "github.com/siyibai/file-transfer/pb"
)

func NewFileService() pb.FileServiceServer {
	return &fileServiceImpl{
		fileLogic: filelogic.NewFileLogic(),
	}
}

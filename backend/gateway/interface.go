package fileServer

import (
	filelogic "github.com/baisiyi/file_trans/logic/file_logic"
	pb "github.com/baisiyi/file_trans/pb"
)

func NewFileService() pb.FileServiceServer {
	return &fileServiceImpl{
		FileLogic: filelogic.NewFileLogic(),
	}
}

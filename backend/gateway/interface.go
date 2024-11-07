package fileServer

import (
	pb "/github.com/baisiyi/file_transpb"

	filelogic "github.com/baisiyi/file_trans/logic/file_logic"
)

func NewFileService() pb.FileServiceServer {
	return &fileServiceImpl{
		fileLogic: filelogic.NewFileLogic(),
	}
}

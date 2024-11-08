package filelogic

import (
	"context"

	pb "github.com/baisiyi/file_trans/pb"
)

type FileLogic interface {
	GetFileList(ctx context.Context, req *pb.GetFileListRequest) (rsp *pb.GetFileListResponse, err error)
	UploadFile(ctx context.Context, req *pb.UploadFileRequest) (rsp *pb.UploadFileResponse, err error)
	DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (rsp *pb.DownloadFileResponse, err error)
}

func NewFileLogic() FileLogic {
	return &fileLogicInterface{}
}

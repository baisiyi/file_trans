package filelogic

import (
	"context"

	pb "github.com/siyibai/file-transfer/pb"
)

type FileLogic interface {
	GetFileList(ctx context.Context, req *pb.GetFileListRequest) (rsp *pb.GetFileListResponse, err error)
	UploadFile(ctx context.Context, req *pb.UploadFileRequest) (rsp *pb.UploadFileResponse, err error)
	DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (rsp *pb.DownloadFileResponse, err error)
}

type fileLogic struct {
}

func NewFileLogic() FileLogic {
	return &fileLogic{}
}

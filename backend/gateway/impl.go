package fileServer

import (
	"context"

	pb "github.com/siyibai/file-transfer/pb"
)

type fileServiceImpl struct {
	pb.UnimplementedFileServiceServer // 添加这一行
}

func (f fileServiceImpl) GetFileList(ctx context.Context, req *pb.GetFileListRequest) (rsp *pb.GetFileListResponse, err error) {
	rsp = &pb.GetFileListResponse{}
	rsp.Files = append(rsp.Files, &pb.FileInfo{
		Id:       "1",
		FileName: "test.txt",
		FileSize: 1024,
	})
	return
}

func (f fileServiceImpl) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (
	rsp *pb.UploadFileResponse, err error) {
	return
}

func (f fileServiceImpl) DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (
	rsp *pb.DownloadFileResponse, err error) {
	return
}

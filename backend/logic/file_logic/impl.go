package filelogic

import (
	"context"

	pb "github.com/baisiyi/file_trans/pb"
)

type fileLogicInterface struct {

}

func (f *fileLogicInterface) GetFileList(ctx context.Context, req *pb.GetFileListRequest) (rsp *pb.GetFileListResponse, err error) {
	// 从mysql中获取文件列表
	return
}

func (f *fileLogicInterface) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (rsp *pb.UploadFileResponse, err error) {
	// 写文件元信息到mysql

	// 上传文件到minio

	return
}

func (f *fileLogicInterface) DownloadFile(ctx context.Context, req *pb.DownloadFileRequest) (rsp *pb.DownloadFileResponse, err error) {
	return
}

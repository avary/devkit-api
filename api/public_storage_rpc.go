package api

import (
	"bytes"
	"context"

	"connectrpc.com/connect"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (api *Api) FilesDelete(ctx context.Context, req *connect.Request[devkitv1.FilesDeleteRequest]) (*connect.Response[devkitv1.FilesDeleteResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.FilesDelete(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) FilesList(ctx context.Context, req *connect.Request[devkitv1.FilesListRequest]) (*connect.Response[devkitv1.FilesListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.FilesList(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) BucketsList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[devkitv1.BucketsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	response, err := api.publicUsecase.BucketsList(ctx, req.Msg)
	return connect.NewResponse(response), err
}

func (api *Api) UploadFile(ctx context.Context, req *connect.Request[devkitv1.UploadFileRequest]) (*connect.Response[devkitv1.UploadFileResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.UploadFile(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UploadFiles(ctx context.Context, req *connect.Request[devkitv1.UploadFilesRequest]) (*connect.Response[devkitv1.UploadFileResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp, err := api.publicUsecase.UploadFiles(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) ImportTable(ctx context.Context, req *connect.Request[devkitv1.ImportTableRequest]) (*connect.Response[devkitv1.ImportTableResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	buffer := bytes.NewBuffer(req.Msg.Reader)
	resp, err := api.sqlSeeder.SeedFromExcel(*buffer, req.Msg.SchemaName, req.Msg.TableName, req.Msg.SchemaName)
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("resp", resp).Msg("respfrom import")
	return connect.NewResponse(&devkitv1.ImportTableResponse{
		Message: "imported",
	}), nil
}

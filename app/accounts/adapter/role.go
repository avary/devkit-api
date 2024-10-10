package adapter

import (
	"github.com/darwishdev/devkit-api/gen/db"
	devkitv1 "github.com/darwishdev/devkit-api/gen/pb/proto/devkit/v1"
	"github.com/jackc/pgx/v5/pgtype"
)

func (a *AccountsAdapter) RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *devkitv1.AccountsSchemaRole {
	return &devkitv1.AccountsSchemaRole{
		RoleId:          int32(resp.RoleID),
		RoleName:        resp.RoleName,
		RoleDescription: resp.RoleDescription.String,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
	}

}

func (a *AccountsAdapter) RoleCreateSqlFromGrpc(req *devkitv1.RoleCreateRequest) *db.RoleCreateParams {
	return &db.RoleCreateParams{
		RoleName:        req.RoleName,
		RoleDescription: pgtype.Text{String: req.RoleDescription, Valid: true},
	}
}
func (a *AccountsAdapter) RoleCreateGrpcFromSql(resp *db.AccountsSchemaRole) *devkitv1.RoleCreateResponse {
	return &devkitv1.RoleCreateResponse{
		Role: a.RoleEntityGrpcFromSql(resp),
	}
}
package repo

import (
	"context"

	"github.com/darwishdev/devkit-api/db"
)

func (repo *AccountsRepo) RolesList(ctx context.Context) ([]db.AccountsSchemaRole, error) {
	resp, err := repo.store.RolesList(context.Background())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (repo *AccountsRepo) RoleCreateUpdate(ctx context.Context, req db.RoleCreateUpdateParams) (*db.AccountsSchemaRole, error) {
	resp, err := repo.store.RoleCreateUpdate(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (repo *AccountsRepo) RolesDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.RolesDeleteRestore(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

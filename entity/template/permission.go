package gen

import (
	"context"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/template/internal"
	"taurus_go_demo/entity/template/user"
)

const PermissionTag = "usertag"

// Permission  is an struct of the database
type Permission struct {
	*internal.Config
	tracker entity.Tracker
	Users   *user.UserEntityBuilder
}

// NewPermission creates a new Permission instance.
func NewPermission() (*Permission, error) {
	config, err := internal.NewConfig(PermissionTag)
	if err != nil {
		return nil, err
	}
	permission := &Permission{
		Config:  config,
		tracker: &entity.Tracking{},
	}
	permission.init()
	return permission, nil
}

// Close closes the database.
func (d *Permission) Close() error {
	return d.Driver.Close()
}

// Save saves all changes to the database.
func (d *Permission) Save(ctx context.Context) error {
	tx, err := d.Config.MayTx(ctx)
	if err != nil {
		return err
	}
	if err := func() error {
		for _, m := range d.tracker.Mutators() {
			if err := m.Exec(ctx, tx); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return entitysql.Rollback(tx, err)
	}

	return tx.Commit()
}

func (d *Permission) init() {
	d.Users = user.NewUserEntityBuilder(d.Config, d.tracker)
}

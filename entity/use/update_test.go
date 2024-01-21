package use

import (
	"context"
	entity "taurus_go_demo/entity/gen"
	"testing"

	"github.com/yohobala/taurus_go/tlog"
)

func TestUpdate(t *testing.T) {
	db, err := entity.NewUser()
	defer db.Close()
	if err != nil {
		t.Errorf(err.Error())
	}
	ctx := context.Background()
	u, err := db.Users.First(ctx)
	if err != nil {
		t.Errorf(err.Error())
	}
	tlog.Print(u.State())
	u.Desc.Set("test")
	db.Save()
}

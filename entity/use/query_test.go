package use

import (
	"context"
	entity "taurus_go_demo/entity/gen"
	"testing"

	"github.com/yohobala/taurus_go/tlog"
)

func TestQuery(t *testing.T) {
	db, err := entity.NewUser()
	defer db.Close()
	if err != nil {
		t.Errorf(err.Error())
	}
	ctx := context.Background()

	t.Run("first", func(t *testing.T) {
		u, err := db.Users.First(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		tlog.Print(u)
	})

	t.Run("list", func(t *testing.T) {
		u, err := db.Users.ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		tlog.Print(u)
	})

	t.Run("where", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.Name.EQ("test"),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		tlog.Print(u)
	})
}

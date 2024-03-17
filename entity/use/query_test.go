package use

import (
	"context"
	entity "taurus_go_demo/entity/template"
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

	t.Run("eq", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.Name.EQ("test"),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("neq", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.Name.NEQ("test"),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("gt", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.ID.GT(20159078),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("gte", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.ID.GTE(20159078),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("lt", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.ID.LT(20157081),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("lte", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.ID.LTE(20157081),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("in", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.ID.IN(20157081, 20157083),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("not in", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.ID.NotIn(20157081, 20157083),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("like", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.Name.Like("test"),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("is null", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.Email.IsNull(),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("is not null", func(t *testing.T) {
		u, err := db.Users.Where(
			db.Users.Email.NotNull(),
		).ToList(ctx)
		tlog.Print(u)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

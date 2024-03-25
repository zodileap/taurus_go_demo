package use

import (
	"context"
	"testing"

	"github.com/yohobala/taurus_go/testutil/unit"
	"github.com/yohobala/taurus_go/tlog"
)

func TestUpdate(t *testing.T) {
	db := initDb()
	defer db.Close()
	ctx := context.Background()

	t.Run("single", func(t *testing.T) {
		u, err := db.Blogs.First(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		u.Desc.Set("single desc")
		err = db.Save(ctx)
		tlog.Print(*u.Desc.Get())
		unit.Must(err)
	})

	t.Run("multi", func(t *testing.T) {
		us, err := db.Blogs.Where(db.Blogs.Desc.EQ("single desc")).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		for _, u := range us {
			u.Desc.Set("multi desc")
		}
		err = db.Save(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

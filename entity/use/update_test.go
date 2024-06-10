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
		us, err := db.Blogs.Where(db.Blogs.Desc.EQ("desc")).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		for _, u := range us {
			u.Desc.Set("multi desc22")
		}
		err = db.Save(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("fieldDemo", func(t *testing.T) {
		u, err := db.FieldDemos.Where(db.FieldDemos.Int64F.EQ(1)).Single(ctx)
		if err != nil {
			tlog.Print(err.Error())
		}
		tlog.Print(*u.Int64F.Get())
		tlog.Print(u.TimeArrayF.Get())
		u.BoolF.Set(false)
		u.VarF.Set("world")
		u.IntArrayF.Set([]int64{2, 3, 4})
		err = db.Save(ctx)
		unit.Must(err)
	})
}

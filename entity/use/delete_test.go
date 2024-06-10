package use

import (
	"context"
	"testing"
	"time"

	"github.com/yohobala/taurus_go/testutil/unit"
	"github.com/yohobala/taurus_go/tlog"
)

func TestDelete(t *testing.T) {
	db := initDb()
	defer db.Close()
	ctx := context.Background()

	t.Run("single", func(t *testing.T) {
		u, err := db.Blogs.First(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		db.Blogs.Remove(u)
		db.Save(ctx)
	})

	t.Run("multi", func(t *testing.T) {
		starttime := time.Now()
		us, err := db.Blogs.Where(db.Blogs.Desc.Like("%desc%")).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		for _, u := range us {
			if err := db.Blogs.Remove(u); err != nil {
				unit.Must(err)
			}
		}
		err = db.Save(ctx)
		unit.Must(err)
		elapsedTime := time.Since(starttime)
		tlog.Printf("elapsedTime: %s", elapsedTime)
	})
}

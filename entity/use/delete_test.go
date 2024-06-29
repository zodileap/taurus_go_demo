package use

import (
	"context"
	"testing"
	"time"

	"github.com/yohobala/taurus_go/testutil/unit"
	"github.com/yohobala/taurus_go/tlog"
)

func TestDelete(t *testing.T) {
	// 删除单个实体。
	// Deleting a single entity.
	t.Run("Deleting a single entity.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.Blogs.First(ctx)
		unit.Must(t, err)

		db.Blogs.Remove(u)
		err = db.Save(ctx)
		unit.Must(t, err)
	})

	// 删除多个实体。
	// Deleting multiple entities.
	t.Run("Deleting multiple entities.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		starttime := time.Now()
		us, err := db.Blogs.Where(db.Blogs.Desc.Like("%lti desc%")).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		for _, u := range us {
			if err := db.Blogs.Remove(u); err != nil {
				unit.Must(t, err)
			}
		}
		tlog.Print(us)

		err = db.Save(ctx)
		unit.Must(t, err)
		elapsedTime := time.Since(starttime)
		tlog.Printf("elapsedTime: %s", elapsedTime)
	})
}

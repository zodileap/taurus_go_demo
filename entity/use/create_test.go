package use

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/yohobala/taurus_go/testutil/unit"
	"github.com/yohobala/taurus_go/tlog"
)

func TestCreate(t *testing.T) {
	db := initDb()
	defer db.Close()
	ctx := context.Background()

	t.Run("single", func(t *testing.T) {
		u, err := db.Blogs.New(
			uuid.New().String(),
			db.Blogs.WithUrl("http://test.com"),
		)
		unit.Must(err)
		err = db.Save(ctx)
		unit.Must(err)
		tlog.Print(u.ID)
		tlog.Print(u.CreatedTime)
	})

	t.Run("multi_1", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			_, err := db.Blogs.New(
				uuid.New().String(),
				db.Blogs.WithUrl("http://test.com"),
			)
			unit.Must(err)
		}
		err := db.Save(ctx)
		unit.Must(err)
	})
}

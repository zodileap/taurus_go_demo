package use

import (
	"context"
	"testing"
	"time"

	"github.com/yohobala/taurus_go/testutil/unit"
	"github.com/yohobala/taurus_go/tlog"
)

func TestUpdate(t *testing.T) {
	// 更新单个实体。
	// Updating a single entity.
	t.Run("Updating a single entity.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.Blogs.First(ctx)
		unit.Must(t, err)

		u.Description.Set("single desc")
		err = db.Save(ctx)
		unit.Must(t, err)
	})

	// 更新多个实体。
	// Updating multiple entities.
	t.Run("Updating multiple entities.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		us, err := db.Blogs.Where(db.Blogs.Description.EQ("desc")).ToList(ctx)
		unit.Must(t, err)

		for _, u := range us {
			u.Description.Set("multi desc")
		}
		err = db.Save(ctx)
		unit.Must(t, err)
	})

	// 全部的已有字段类型更新测试。
	// All existing field type update tests.
	t.Run("All existing field type update tests.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.FieldDemos.Where(db.FieldDemos.Int64F.EQ(1)).Single(ctx)
		unit.Must(t, err)

		u.VarF.Set("Update")
		u.BoolF.Set(false)
		u.IntArrayF.Set([]int64{2, 3, 4})
		u.Intarray2F.Set([][]int64{{2, 3}, {4, 5}})
		u.BoolArrayF.Set([]bool{false, false})
		u.TimeF.Set(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
		u.TimeArrayF.Set([]time.Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)})
		tlog.Print(u)
		err = db.Save(ctx)
		unit.Must(t, err)
	})
}

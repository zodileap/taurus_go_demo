package use

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/yohobala/taurus_go/testutil/unit"
)

func TestCreate(t *testing.T) {

	// 创建单个实体。
	// Creating a single entity.
	t.Run("Creating a single entity.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.Blogs.Create(
			uuid.New().String(),
			db.Blogs.WithCreatedTime(time.Now()),
		)
		unit.Must(t, err)

		err = db.Save(ctx)
		unit.Must(t, err)
		fmt.Print(u.ID)
	})

	// 创建多个实体。
	// Creating multiple entities.
	t.Run("Creating multiple entities.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		for i := 0; i < 2; i++ {
			_, err := db.Blogs.Create(
				uuid.New().String(),
				db.Blogs.WithDesc("desc"),
			)
			unit.Must(t, err)
		}
		err := db.Save(ctx)
		unit.Must(t, err)
	})

	// 全部的已有字段类型创建测试。
	// All existing field type creation tests.
	t.Run("All existing field type creation tests.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		_, err := db.FieldDemos.Create(
			1,
			"hello",
			true,
			[]int64{1, 2, 3},
			[][]int64{{1, 2}, {3, 4}},
			[]bool{true, true},
			time.Now(),
			[]time.Time{time.Now(), time.Now()},
		)
		unit.Must(t, err)

		err = db.Save(ctx)
		unit.Must(t, err)
	})
}

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
	db := initDb()
	defer db.Close()
	ctx := context.Background()

	t.Run("single", func(t *testing.T) {
		u, err := db.Blogs.Create(
			uuid.New().String(),
			db.Blogs.WithCreatedTime(time.Now()),
		)
		unit.Must(err)
		err = db.Save(ctx)
		unit.Must(err)
		fmt.Print(u.ID)
	})

	t.Run("multi_1", func(t *testing.T) {
		for i := 0; i < 2; i++ {
			_, err := db.Blogs.Create(
				uuid.New().String(),
				db.Blogs.WithDesc("desc"),
			)
			unit.Must(err)
		}
		err := db.Save(ctx)
		unit.Must(err)
	})

	t.Run("fieldDemo", func(t *testing.T) {
		_, err := db.FieldDemos.Create(
			1,
			"hello",
			true,
			[]int64{1, 2, 3},
			[][]int64{{1, 2}, {3, 4}},
			[]bool{true, false},
			time.Now(),
			[]time.Time{time.Now(), time.Now()},
		)
		unit.Must(err)
		err = db.Save(ctx)
		unit.Must(err)
	})
}

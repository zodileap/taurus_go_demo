package use

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/yohobala/taurus_go/testutil/unit"
)

func BenchmarkCreate(b *testing.B) {
	db := initDb()
	defer db.Close()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			_, err := db.Blogs.New(
				uuid.New().String(),
				db.Blogs.WithUrl("http://test.com"),
			)
			unit.Must(err)
		}
		err := db.Save(ctx)
		unit.Must(err)
	}
}

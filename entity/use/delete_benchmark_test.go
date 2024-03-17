package use

import (
	"context"
	"testing"

	"github.com/yohobala/taurus_go/testutil/unit"
)

func BenchmarkDelete(b *testing.B) {
	db := initDb()
	defer db.Close()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		us, err := db.Blogs.Where(db.Blogs.Url.EQ("http://test.com")).ToList(ctx)
		unit.Must(err)
		for _, u := range us {
			if err := db.Blogs.Remove(u); err != nil {
				unit.Must(err)
			}
		}
		err = db.Save(ctx)
		unit.Must(err)
	}
}

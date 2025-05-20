package use

// import (
// 	"context"
// 	"testing"

// 	"github.com/zodileap/taurus_go/testutil/unit"
// )

// func BenchmarkUpdate(b *testing.B) {
// 	db := initDb()
// 	defer db.Close()
// 	ctx := context.Background()
// 	for i := 0; i < b.N; i++ {
// 		us, err := db.Blogs.Where(db.Blogs.Desc.EQ("desc")).ToList(ctx)
// 		unit.Must(err)
// 		for _, u := range us {
// 			u.Desc.Set("BenchmarkUpdate")
// 		}
// 		err = db.Save(ctx)
// 		unit.Must(err)
// 	}
// }

package use

// import (
// 	"context"
// 	"testing"
// )

// func BenchmarkQuery(b *testing.B) {
// 	db := initDb()
// 	defer db.Close()
// 	ctx := context.Background()
// 	for i := 0; i < b.N; i++ {
// 		_, err := db.Blogs.Where(db.Blogs.Desc.EQ("desc")).ToList(ctx)
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

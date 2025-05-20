package use

import (
	"context"
	"fmt"
	"testing"

	"github.com/zodileap/taurus_go/testutil/unit"
)

func TestQuery(t *testing.T) {
	// 查询第一条数据。
	// Query the first entry.
	t.Run("Query the first entry.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.First(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 查询全部数据。
	// Query all entries.
	t.Run("Query all entries.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 相等条件查询，返回多条数据。
	// Query with equal conditions, and return multiple entries.
	t.Run("Query with equal conditions, and return multiple entries.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.EQ("single desc"),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 相等条件查询，返回单条数据。
	// Query with equal conditions, and return a single entry.
	t.Run("Query with equal conditions, and return a single entry.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.EQ("desc"),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 相等条件查询，返回多条数据，限制返回数量。
	// Query with equal conditions, and return multiple entries, with a limit on the number of entries returned.
	t.Run(" Query with equal conditions, and return multiple entries, with a limit on the number of entries returned.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.EQ("desc"),
		).Limit(3).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 不等条件查询。
	// Query with not equal conditions.
	t.Run("Query with not equal conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.NEQ("single desc"),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 大于条件查询。
	// Query with greater than conditions.
	t.Run("Query with greater than conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Id.GT(1),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 大于等于条件查询。
	// Query with greater than or equal conditions.
	t.Run("Query with greater than or equal conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Id.GTE(4068),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 小于条件查询。
	// Query with less than conditions.
	t.Run("Query with less than conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Id.LT(20157081),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 小于等于条件查询。
	// Query with less than or equal conditions.
	t.Run("Query with less than or equal conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Id.LTE(20157081),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 值是否在范围内查询。
	// Query with value in range conditions.
	t.Run("Query with value in range conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.In("single desc", "multi desc"),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 值是否不在范围内查询。
	// Query with value not in range conditions.
	t.Run("Query with value not in range conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Id.NotIn(1, 2),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 模糊查询。
	// Query with fuzzy conditions.
	t.Run("Query with fuzzy conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.Like("%mutil%"),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 是否为空查询。
	// Query with null conditions.
	t.Run("Query with null conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.IsNull(),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 是否不为空查询。
	// Query with not null conditions.
	t.Run("Query with not null conditions.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.Blogs.Where(
			db.Blogs.Description.NotNull(),
		).ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 对查询结果进行排序。
	// Sort the query results.
	t.Run("Sort the query results.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.Posts.Where(
			db.Posts.Id.EQ(1),
		).Order(db.Posts.ByID.NullsLast()).
			ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
	})

	// 查询全部的字段类型
	// Query all entries.
	t.Run("Query all field types.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		u, err := db.FieldDemos.ToList(ctx)
		unit.Must(t, err)
		fmt.Println(u)
		fmt.Println(u[0].JsonF.Get().Key)
	})
}

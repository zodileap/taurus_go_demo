package use

import (
	"context"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db := initDb()
	defer db.Close()
	ctx := context.Background()

	t.Run("first", func(t *testing.T) {
		u, err := db.Blogs.First(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("list", func(t *testing.T) {
		u, err := db.Blogs.ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("eq", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.EQ("single desc"),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("eq_single", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.EQ("desc"),
		).Single(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("eq_limit", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.EQ("desc"),
		).Limit(3).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("neq", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.NEQ("single desc"),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("gt", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.ID.GT(1),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("gte", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.ID.GTE(4068),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("lt", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.ID.LT(20157081),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("lte", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.ID.LTE(20157081),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("in", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.In("single desc", "multi desc"),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("not in", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.ID.NotIn(1, 2),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("like", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.Like("%mutil%"),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("is null", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.IsNull(),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	t.Run("is not null", func(t *testing.T) {
		u, err := db.Blogs.Where(
			db.Blogs.Desc.NotNull(),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})
}

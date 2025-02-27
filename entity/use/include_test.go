package use

import (
	"context"
	"fmt"
	"testing"
)

func TestRel(t *testing.T) {

	// Include只包含自身关联表，返回单条数据。
	// Include only the self-referencing table and returns a single entry.
	t.Run("Include only the self-referencing table and returns a single entry.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		b, err := db.Blogs.Where(
			db.Blogs.ID.EQ(4077),
		).
			Include(
				db.Blogs.Posts,
			).Single(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(b)
	})

	// Include只包含自身关联表，返回多条数据。
	// Include only the self-referencing table and returns multiple entries.
	t.Run("Include only the self-referencing table and returns multiple entries.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		b, err := db.Blogs.Where(
			db.Blogs.ID.EQ(4077),
		).
			Include(
				db.Blogs.Posts,
			).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(b)
	})

	// Include没有包含逻辑运算符，默认添加And。
	// Include does not contain logical operators, and And is added by default.
	t.Run("Include does not contain logical operators, and And is added by default.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		b, err := db.Blogs.Where(
			db.Blogs.Description.EQ("single desc"),
		).
			Include(
				db.Blogs.Posts.Where(
					db.Posts.Content.EQ("single content"),
				),
			).Single(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(b)
	})

	// Include包含多级的关联表，返回全部数据。
	// Include multiple levels of associated tables and return a single entry.
	t.Run("Include multiple levels of associated tables and return a single entry.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.Posts.Include(
			db.Posts.Blog,
			db.Posts.Author,
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})

	// Include包含子Include，返回全部数据。
	// Include contains sub-Include and returns a single entry.
	t.Run("Include contains sub-Include and returns a single entry.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		u, err := db.Posts.Include(
			db.Posts.Blog.Include(
				db.Blogs.Posts,
			),
		).ToList(ctx)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(u)
	})
}

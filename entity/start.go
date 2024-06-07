package main

import (
	"context"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	e "github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"

	entity "taurus_go_demo/entity/new/entity"
)

func main() {
	e.AddConnection(e.ConnectionConfig{
		Driver:   dialect.PostgreSQL,
		Tag:      "User",
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "test",
		DBName:   "test",
	})
	db, err := entity.NewUser()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()
	_, err = db.Blogs.Create(
		uuid.New().String(),
		db.Blogs.WithDesc("http://test.com"),
	)
	if err != nil {
		panic(err)
	}
	// Save the new data to Blog
	db.Save(ctx)
}

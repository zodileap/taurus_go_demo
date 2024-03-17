package use

import (
	entity "taurus_go_demo/entity/new/entity"

	e "github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"

	_ "github.com/lib/pq"
)

var conn = e.ConnectionConfig{
	
		Driver:   dialect.PostgreSQL,
		Tag:      "User",
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "test",
		DBName:   "test",
	},


func init() {
	for _, conn := range testConnections {
		err := e.AddConnection(conn)
		if err != nil {
			panic(err)
		}
	}
}

func initDb() *entity.User {
	db, err := entity.NewUser()
	if err != nil {
		panic(err)
	}
	return db
}

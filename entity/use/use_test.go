package use

import (
	e "github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"

	_ "github.com/lib/pq"
)

var testConnections = []e.ConnectionConfig{
	{
		Driver:   dialect.PostgreSQL,
		Tag:      "usertag",
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "test",
		DBName:   "test",
	},
}

func init() {
	for _, conn := range testConnections {
		err := e.AddConnection(conn)
		if err != nil {
			panic(err)
		}
	}
}

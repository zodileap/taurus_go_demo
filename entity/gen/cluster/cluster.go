package cluster

import (
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
)

type User struct {
	entity.Database
	User UserEntity
}

func (d User) Config() entity.DbConfig {
	return entity.DbConfig{
		Tag:  "usertag",
		Type: dialect.PostgreSQL,
	}
}

type Permission struct {
	entity.Database
	User UserEntity
}

func (d Permission) Config() entity.DbConfig {
	return entity.DbConfig{
		Tag:  "usertag",
		Type: dialect.PostgreSQL,
	}
}

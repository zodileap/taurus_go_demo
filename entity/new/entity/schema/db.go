package schema

import (
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
)

type User struct {
	entity.Database
	Blog BlogEntity
}

func (d User) Config() entity.DbConfig {
	return entity.DbConfig{
		Name: "user",
		Tag:  "UserTag",
		Type: dialect.PostgreSQL,
	}
}

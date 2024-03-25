package cluster

import (
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type UserEntity struct {
	entity.Entity
	ID          field.Int64
	UUID        field.UUID
	Name        field.Varchar
	Email       field.Varchar
	CreatedTime field.Timestamptz
	Desc        field.Varchar
}

func (u UserEntity) Config() entity.EntityConfig {
	return entity.EntityConfig{
		AttrName: "user",
	}
}

func (u UserEntity) Fields() []entity.FieldBuilder {
	return []entity.FieldBuilder{
		u.ID.Name("id").Primary(1).Sequence(entity.NewSequence("user_id_seq")).Comment("用户的id"),
		u.UUID.Primary(2).Name("uuid"),
		u.Name.Required(),
		u.CreatedTime.Required().Default("CURRENT_TIMESTAMP"),
		u.Desc.Default(`'desc'`),
	}
}

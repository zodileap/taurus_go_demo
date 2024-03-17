package schema

import (
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type BlogEntity struct {
	entity.Entity
	ID          field.Int64
	UUID        field.UUID
	Desc        field.Varchar
	CreatedTime field.Timestamptz
}

func (e BlogEntity) Config() entity.EntityConfig {
	return entity.EntityConfig{
		AttrName: "blog",
	}
}

func (e BlogEntity) Fields() []entity.FieldBuilder {
	return []entity.FieldBuilder{
		e.ID.Primary(1).Sequence(entity.NewSequence("blog_id_seq")).Comment("Blog primary key"),
		e.UUID.Required().Name("uuid"),
		e.CreatedTime.Required().Default("CURRENT_TIMESTAMP"),
	}
}

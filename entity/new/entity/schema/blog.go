package schema

import (
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type BlogEntity struct {
	entity.Entity
	ID          *field.Int64
	UUID        *field.UUID
	Desc        *field.Varchar
	CreatedTime *field.Timestamptz
}

func (e *BlogEntity) Config() entity.EntityConfig {
	return entity.EntityConfig{
		AttrName: "blog",
	}
}

func (e *BlogEntity) Fields() []entity.FieldBuilder {
	return []entity.FieldBuilder{
		e.ID.Name("id").Primary(1).Sequence(entity.NewSequence("blog_id_seq")).Comment("Blog primary key").Locked(),
		e.UUID.Required().Name("uuid"),
		e.CreatedTime.Default("CURRENT_TIMESTAMP"),
	}
}

type PostEntity struct {
	entity.Entity
	ID       *field.Int64
	Content  *field.Varchar
	BlogID   *field.Int64
	AuthorID *field.Int64
}

func (e *PostEntity) Config() entity.EntityConfig {
	return entity.EntityConfig{
		AttrName: "post",
	}
}

func (e *PostEntity) Fields() []entity.FieldBuilder {
	return []entity.FieldBuilder{
		e.ID.Name("id").Primary(1).Sequence(entity.NewSequence("post_id_seq")).Comment("Post primary key").Locked(),
		e.Content.Required().Name("content"),
		e.BlogID.Required().Name("blog_id"),
		e.AuthorID.Required().Name("author_id"),
	}
}

type AuthorEntity struct {
	entity.Entity
	ID   *field.Int64
	Name *field.Varchar
}

func (e *AuthorEntity) Config() entity.EntityConfig {
	return entity.EntityConfig{
		AttrName: "author",
	}
}

func (e *AuthorEntity) Fields() []entity.FieldBuilder {
	return []entity.FieldBuilder{
		e.ID.Name("id").Primary(1).Sequence(entity.NewSequence("author_id_seq")).Comment("Author primary key").Locked(),
		e.Name.Required().Name("name"),
	}
}

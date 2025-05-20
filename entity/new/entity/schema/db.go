package schema

import (
	"github.com/zodileap/taurus_go/entity"
	"github.com/zodileap/taurus_go/entity/dialect"
)

type User struct {
	entity.Database
	Blog      *BlogEntity
	Post      *PostEntity
	Author    *AuthorEntity
	FieldDemo *FieldDemoEntity
	Geo       *GeoEntity
}

func (d *User) Config() entity.DbConfig {
	return entity.DbConfig{
		Name: "user",
		Tag:  "UserTag",
		Type: dialect.PostgreSQL,
	}
}

func (u *User) Relationships() []entity.RelationshipBuilder {
	return []entity.RelationshipBuilder{
		entity.InitRelationship().
			HasMany(u.Post).
			WithOne(u.Blog).
			ReferenceKey(u.Blog.Id).
			ForeignKey(u.Post.BlogID).
			Update(entity.Cascade).
			ConstraintName("fk_blog_id"),
		entity.InitRelationship().
			HasOne(u.Author).
			WithMany(u.Post).
			ReferenceKey(u.Author.Id).
			ForeignKey(u.Post.AuthorID).
			Update(entity.Cascade).
			ConstraintName("fk_author_id"),
	}
}

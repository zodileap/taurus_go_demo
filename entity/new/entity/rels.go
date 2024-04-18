package entity

import (
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type Rel interface {
	Desc() (entitysql.RelationDesc, []Rel, internal.EntityConfig)
	reset()
}

type BlogEntityRel interface {
	Rel
}

type PostEntityRel interface {
	Rel
}

type AuthorEntityRel interface {
	Rel
}

package entity

import (
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type AuthorEntityRelation struct {
	PostEntityRel
	Config   internal.EntityConfig
	relation entitysql.RelationDesc
	children []Rel
}

func NewAuthorEntityRelation(config internal.EntityConfig, desc entitysql.RelationDesc) *AuthorEntityRelation {
	return &AuthorEntityRelation{
		Config:   config,
		relation: desc,
		children: []Rel{},
	}
}

func (r *AuthorEntityRelation) Where(predicates ...func(*entitysql.Predicate)) *AuthorEntityRelation {
	r.relation.Predicates = append(r.relation.Predicates, predicates...)
	return r
}

func (r *AuthorEntityRelation) Include(rels ...BlogEntityRel) *AuthorEntityRelation {
	newRels := make([]Rel, len(rels)) // 创建一个长度与 r.children 相同的 Rel 类型切片
	for i, r := range rels {
		newRels[i] = Rel(r) // 将每个 AuthorEntityRel 转换为 Rel 并存储在新切片中
	}
	r.children = append(r.children, newRels...)
	return r
}

func (r AuthorEntityRelation) Desc() (entitysql.RelationDesc, []Rel, internal.EntityConfig) {
	return r.relation, r.children, r.Config
}

func (r *AuthorEntityRelation) reset() {
	for _, child := range r.children {
		child.reset()
	}
	r.relation.Reset()
	r.children = []Rel{}
}

package entity

import (
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/new/entity/internal"
)

type BlogEntityRelation struct {
	PostEntityRel
	Config   internal.EntityConfig
	relation *entitysql.RelationDesc
	children []Rel
}

func NewBlogEntityRelation(config internal.EntityConfig, desc entitysql.RelationDesc) *BlogEntityRelation {
	return &BlogEntityRelation{
		Config:   config,
		relation: &desc,
		children: []Rel{},
	}
}

func (r *BlogEntityRelation) Where(predicates ...func(*entitysql.Predicate)) *BlogEntityRelation {
	r.relation.Predicates = append(r.relation.Predicates, predicates...)
	return r
}

func (r *BlogEntityRelation) Include(rels ...BlogEntityRel) *BlogEntityRelation {
	newRels := make([]Rel, len(rels)) // 创建一个长度与 r.children 相同的 Rel 类型切片
	for i, r := range rels {
		newRels[i] = Rel(r) // 将每个 AuthorEntityRel 转换为 Rel 并存储在新切片中
	}
	r.children = append(r.children, newRels...)
	return r
}

func (r BlogEntityRelation) Desc() (entitysql.RelationDesc, []Rel, internal.EntityConfig) {
	return *r.relation, r.children, r.Config
}

func (r *BlogEntityRelation) reset() {
	for _, child := range r.children {
		child.reset()
	}
	r.relation.Reset()
	r.children = []Rel{}
}

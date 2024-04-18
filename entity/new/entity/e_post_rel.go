package entity

import (
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type PostEntityRelation struct {
	BlogEntityRel
	Config   internal.EntityConfig
	relation entitysql.RelationDesc
	children []Rel
}

func NewPostEntityRelation(config internal.EntityConfig, desc entitysql.RelationDesc) *PostEntityRelation {
	return &PostEntityRelation{
		Config:   config,
		relation: desc,
		children: []Rel{},
	}
}

func (r *PostEntityRelation) Where(predicates ...func(*entitysql.Predicate)) *PostEntityRelation {
	r.relation.Predicates = append(r.relation.Predicates, predicates...)
	return r
}

func (r *PostEntityRelation) Include(rels ...PostEntityRel) *PostEntityRelation {
	newRels := make([]Rel, len(rels)) // 创建一个长度与 r.children 相同的 Rel 类型切片
	for i, r := range rels {
		newRels[i] = Rel(r) // 将每个 AuthorEntityRel 转换为 Rel 并存储在新切片中
	}
	r.children = append(r.children, newRels...)
	return r
}

func (r PostEntityRelation) Desc() (entitysql.RelationDesc, []Rel, internal.EntityConfig) {
	return r.relation, r.children, r.Config
}

func (r *PostEntityRelation) reset() {
	for _, child := range r.children {
		child.reset()
	}
	r.relation.Reset()
	r.children = []Rel{}
}

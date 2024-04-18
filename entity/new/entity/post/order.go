package post

import (
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type OrderOption func(*entitysql.Order)

func ByPrimary(o *entitysql.Order) {
	(&ByID{}).Apply(o)
}

type OrderTerm interface {
	Apply(*entitysql.Order)
}

type ByID struct {
	OrderTerm
	Options []OrderOption
	Field   string
}

func (b *ByID) Apply(o *entitysql.Order) {
	o.SetColumn(FieldID.Name.String())
	if len(b.Options) == 0 {
		b.Asc()
	}
	for _, opt := range b.Options {
		opt(o)
	}
}

func (b *ByID) Desc() *ByID {
	b.Options = append(b.Options, func(o *entitysql.Order) {
		o.Desc()
	})
	return b
}

func (o *ByID) Asc() *ByID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.Asc()
	})
	return o
}

func (o *ByID) NullsFirst() *ByID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsFirst()
	})
	return o
}

func (o *ByID) NullsLast() *ByID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsLast()
	})
	return o
}

type ByContent struct {
	OrderTerm
	Options []OrderOption
	Field   string
}

func (b *ByContent) Apply(o *entitysql.Order) {
	o.SetColumn(FieldContent.Name.String())
	if len(b.Options) == 0 {
		b.Asc()
	}
	for _, opt := range b.Options {
		opt(o)
	}
}

func (b *ByContent) Desc() *ByContent {
	b.Options = append(b.Options, func(o *entitysql.Order) {
		o.Desc()
	})
	return b
}

func (o *ByContent) Asc() *ByContent {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.Asc()
	})
	return o
}

func (o *ByContent) NullsFirst() *ByContent {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsFirst()
	})
	return o
}

func (o *ByContent) NullsLast() *ByContent {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsLast()
	})
	return o
}

type ByBlogID struct {
	OrderTerm
	Options []OrderOption
	Field   string
}

func (b *ByBlogID) Apply(o *entitysql.Order) {
	o.SetColumn(FieldBlogID.Name.String())
	if len(b.Options) == 0 {
		b.Asc()
	}
	for _, opt := range b.Options {
		opt(o)
	}
}

func (b *ByBlogID) Desc() *ByBlogID {
	b.Options = append(b.Options, func(o *entitysql.Order) {
		o.Desc()
	})
	return b
}

func (o *ByBlogID) Asc() *ByBlogID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.Asc()
	})
	return o
}

func (o *ByBlogID) NullsFirst() *ByBlogID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsFirst()
	})
	return o
}

func (o *ByBlogID) NullsLast() *ByBlogID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsLast()
	})
	return o
}

type ByAuthorID struct {
	OrderTerm
	Options []OrderOption
	Field   string
}

func (b *ByAuthorID) Apply(o *entitysql.Order) {
	o.SetColumn(FieldAuthorID.Name.String())
	if len(b.Options) == 0 {
		b.Asc()
	}
	for _, opt := range b.Options {
		opt(o)
	}
}

func (b *ByAuthorID) Desc() *ByAuthorID {
	b.Options = append(b.Options, func(o *entitysql.Order) {
		o.Desc()
	})
	return b
}

func (o *ByAuthorID) Asc() *ByAuthorID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.Asc()
	})
	return o
}

func (o *ByAuthorID) NullsFirst() *ByAuthorID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsFirst()
	})
	return o
}

func (o *ByAuthorID) NullsLast() *ByAuthorID {
	o.Options = append(o.Options, func(o *entitysql.Order) {
		o.NullsLast()
	})
	return o
}

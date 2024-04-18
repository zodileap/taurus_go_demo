package author

import "github.com/yohobala/taurus_go/entity/entitysql"

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

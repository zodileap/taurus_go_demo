package user

import (
	"context"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

// UserEntityUpdate is the update action for the UserEntity.
type UserEntityUpdate struct {
	*internal.Config
	ctx        *entitysql.QueryContext
	tracker    entity.Tracker
	e          *UserEntity
	predicates []func(*entitysql.Predicate)
}

// NewUserEntityUpdate creates a new UserEntityUpdate.
func NewUserEntityUpdate(c *internal.Config, e *UserEntity, t entity.Tracker) *UserEntityUpdate {
	return &UserEntityUpdate{
		Config:  c,
		ctx:     &entitysql.QueryContext{},
		e:       e,
		tracker: t,
	}
}

// Where 更新的条件，会默认添加主键的值作为条件，所以如果有添加主键的值，会被覆盖
func (o *UserEntityUpdate) Where(predicates ...func(*entitysql.Predicate)) *UserEntityUpdate {
	o.predicates = append(o.predicates, predicates...)
	return o
}

func (o *UserEntityUpdate) update(ctx context.Context) (*UserEntity, error) {
	return o.sqlUpdate(ctx)
}

func (o *UserEntityUpdate) sqlUpdate(ctx context.Context) (*UserEntity, error) {
	var (
		spec, err = o.updateSpec()
		res       = o.e
	)
	if err != nil {
		return nil, err
	}
	spec.Scan = func(rows dialect.Rows, fields []entitysql.FieldName) error {
		return scan(res, fields, rows)
	}
	if err := entitysql.NewUpdate(ctx, o.Driver, spec); err != nil {
		return nil, err
	}
	setUnchanged(o.tracker, res)
	return res, nil
}

func (o *UserEntityUpdate) updateSpec() (*entitysql.UpdateSpec, error) {
	spec := entitysql.NewUpdateSpec(Entity, o.ctx.Fields)
	fields := o.e.Mutation.Fields()
	if len(fields) == 0 {
		return nil, entity.Err_0100030002.Sprintf(o.e.Name)
	}
	for _, f := range fields {
		switch f {
		case FieldID.Name.String():
			spec.Sets = append(spec.Sets, &entitysql.FieldSpec{
				Column: FieldID.Name.String(),
				Value:  o.e.ID.Value(),
			})
		case FieldUUID.Name.String():
			spec.Sets = append(spec.Sets, &entitysql.FieldSpec{
				Column: FieldUUID.Name.String(),
				Value:  o.e.UUID.Value(),
			})
		case FieldName.Name.String():
			spec.Sets = append(spec.Sets, &entitysql.FieldSpec{
				Column: FieldName.Name.String(),
				Value:  o.e.Name.Value(),
			})
		case FieldEmail.Name.String():
			spec.Sets = append(spec.Sets, &entitysql.FieldSpec{
				Column: FieldEmail.Name.String(),
				Value:  o.e.Email.Value(),
			})
		case FieldCreatedTime.Name.String():
			spec.Sets = append(spec.Sets, &entitysql.FieldSpec{
				Column: FieldCreatedTime.Name.String(),
				Value:  o.e.CreatedTime.Value(),
			})
		case FieldDesc.Name.String():
			spec.Sets = append(spec.Sets, &entitysql.FieldSpec{
				Column: FieldDesc.Name.String(),
				Value:  o.e.Desc.Value(),
			})
		}
	}
	id := &PredID{}
	o.predicates = append(o.predicates, id.EQ(*o.e.ID.Get()))
	uuid := &PredUUID{}
	o.predicates = append(o.predicates, entitysql.OpAnd, uuid.EQ(*o.e.UUID.Get()))
	if ps := o.predicates; len(ps) > 0 {
		spec.Predicate = func(p *entitysql.Predicate) {
			for _, f := range ps {
				f(p)
			}
		}
	}
	return spec, nil
}

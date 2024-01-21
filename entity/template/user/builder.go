package user

import (
	"context"
	"taurus_go_demo/entity/gen/user"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type UserEntityBuilder struct {
	*internal.Config
	tracker     entity.Tracker
	ID          PredID
	UUID        PredUUID
	Name        PredName
	Email       PredEmail
	CreatedTime PredCreatedTime
	Desc        PredDesc
}

func NewUserEntityBuilder(c *internal.Config, t entity.Tracker) *UserEntityBuilder {
	return &UserEntityBuilder{
		Config:  c,
		tracker: t,
	}
}

// New creates a new UserEntity，and add it to the tracker.
// Required parameters are fields that have no default value but are required,
// and options are fields that can be left empty by calling WithFieldName.
func (b *UserEntityBuilder) New(uuid string, name string, options ...func(*UserEntity)) (*UserEntity, error) {
	e := New(b.Config, b.tracker)
	b.tracker.Add(e)
	return e.create(uuid, name, options...)
}

func (b *UserEntityBuilder) Remove(e *UserEntity) error {
	if e.Mutation == nil {
		return nil
	}
	return e.remove()
}

// First returns the first UserEntity.
func (s *UserEntityBuilder) First(ctx context.Context) (*UserEntity, error) {
	query := s.initQuery()
	return query.First(ctx)
}

func (s *UserEntityBuilder) ToList(ctx context.Context) ([]*UserEntity, error) {
	query := s.initQuery()
	return query.ToList(ctx)
}

func (s *UserEntityBuilder) Where(conditions ...func(*entitysql.Predicate)) *UserEntityQuery {
	query := s.initQuery()
	return query.Where(conditions...)
}

// WithEmail sets the "email" field of the UserEntity.
func (s *UserEntityBuilder) WithEmail(email string) func(*user.UserEntity) {
	return func(e *user.UserEntity) {
		e.Email.Set(email)
	}
}

// WithDesc sets the "desc" field of the UserEntity.
func (s *UserEntityBuilder) WithDesc(desc string) func(*user.UserEntity) {
	return func(e *user.UserEntity) {
		e.Desc.Set(desc)
	}
}

func (s *UserEntityBuilder) initQuery() *UserEntityQuery {
	return NewUserEntityQuery(s.Config, s.tracker)
}

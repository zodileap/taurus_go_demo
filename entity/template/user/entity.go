package user

import (
	"fmt"
	"taurus_go_demo/entity/template/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"
)

type UserEntity struct {
	config      *UserEntityConfig
	ID          *IDType
	UUID        *UUIDType
	Name        *NameType
	Email       *EmailType
	CreatedTime *CreatedTimeType
	Desc        *DescType
}

type UserEntityConfig struct {
	*mutations
	*entity.Mutation
	*internal.Config
}

// New creates a new UserEntity, but does not add tracking.
func New(c *internal.Config, ms *mutations) *UserEntity {
	b := entity.NewMutation(entity.Detached)
	e := &UserEntity{
		config: &UserEntityConfig{
			Mutation:  b,
			Config:    c,
			mutations: ms,
		},
	}
	e.setState(entity.Detached)
	e.ID = newIDType(e.config)
	e.UUID = newUUIDType(e.config)
	e.Name = newNameType(e.config)
	e.Email = newEmailType(e.config)
	e.CreatedTime = newCreatedTimeType(e.config)
	e.Desc = newDescType(e.config)
	return e
}

// String implements the fmt.Stringer interface.
func (e *UserEntity) String() string {
	return fmt.Sprintf("{ ID: %v, UUID: %v, Name: %v, Email: %v, CreatedTime: %v, Desc: %v }",
		e.ID,
		e.UUID,
		e.Name,
		e.Email,
		e.CreatedTime,
		e.Desc,
	)
}

// State returns the state of the UserEntity.
func (e *UserEntity) State() entity.EntityState {
	return e.config.State()
}

func (e *UserEntity) remove() error {
	return e.setState(entity.Deleted)
}

// create creates a new UserEntity and adds tracking.
func (e *UserEntity) create(uuid string, name string, options ...func(*UserEntity)) (*UserEntity, error) {
	e.setState(entity.Added)
	e.UUID.Set(uuid)
	e.Name.Set(name)
	for _, option := range options {
		option(e)
	}
	return e, nil
}

func (e *UserEntity) setUnchanged() error {
	return e.setState(entity.Unchanged)
}

func (e *UserEntity) setState(state entity.EntityState) error {
	return e.config.mutations.SetEntityState(e, state)
}

func scan(e *UserEntity, fields []entitysql.FieldName, rows dialect.Rows) error {
	args := make([]interface{}, len(fields))
	for i := range fields {
		switch fields[i] {
		case FieldID.Name:
			args[i] = e.ID
		case FieldUUID.Name:
			args[i] = e.UUID
		case FieldName.Name:
			args[i] = e.Name
		case FieldEmail.Name:
			args[i] = e.Email
		case FieldCreatedTime.Name:
			args[i] = e.CreatedTime
		case FieldDesc.Name:
			args[i] = e.Desc
		}
	}

	if err := rows.Scan(args...); err != nil {
		return err
	}
	return nil
}

package user

import (
	"time"

	"github.com/yohobala/taurus_go/entity/entitysql"
)

type PredID struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredID) EQ(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldID.Name.String(), id)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredID) NEQ(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldID.Name.String(), id)
	}
}

func (f *PredID) GT(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GT(FieldID.Name.String(), id)
	}
}

func (f *PredID) GTE(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldID.Name.String(), id)
	}
}

func (f *PredID) LT(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LT(FieldID.Name.String(), id)
	}
}

func (f *PredID) LTE(id int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldID.Name.String(), id)
	}
}

func (f *PredID) IN(ids ...int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		p.In(FieldID.Name.String(), v...)
	}
}

func (f *PredID) NotIn(ids ...int64) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		p.NotIn(FieldID.Name.String(), v...)
	}
}

func (f *PredID) Like(id string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.Like(FieldID.Name.String(), id)
	}
}

type PredUUID struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredUUID) EQ(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldUUID.Name.String(), uuid)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredUUID) NEQ(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldUUID.Name.String(), uuid)
	}
}

func (f *PredUUID) GT(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GT(FieldUUID.Name.String(), uuid)
	}
}

func (f *PredUUID) GTE(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldUUID.Name.String(), uuid)
	}
}

func (f *PredUUID) LT(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LT(FieldUUID.Name.String(), uuid)
	}
}

func (f *PredUUID) LTE(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldUUID.Name.String(), uuid)
	}
}

func (f *PredUUID) IN(uuids ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(uuids))
		for i := range v {
			v[i] = uuids[i]
		}
		p.In(FieldUUID.Name.String(), v...)
	}
}

func (f *PredUUID) NotIn(uuids ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(uuids))
		for i := range v {
			v[i] = uuids[i]
		}
		p.NotIn(FieldUUID.Name.String(), v...)
	}
}

func (f *PredUUID) Like(uuid string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.Like(FieldUUID.Name.String(), uuid)
	}
}

type PredName struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredName) EQ(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldName.Name.String(), name)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredName) NEQ(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldName.Name.String(), name)
	}
}

func (f *PredName) GT(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GT(FieldName.Name.String(), name)
	}
}

func (f *PredName) GTE(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldName.Name.String(), name)
	}
}

func (f *PredName) LT(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LT(FieldName.Name.String(), name)
	}
}

func (f *PredName) LTE(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldName.Name.String(), name)
	}
}

func (f *PredName) IN(names ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(names))
		for i := range v {
			v[i] = names[i]
		}
		p.In(FieldName.Name.String(), v...)
	}
}

func (f *PredName) NotIn(names ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(names))
		for i := range v {
			v[i] = names[i]
		}
		p.NotIn(FieldName.Name.String(), v...)
	}
}

func (f *PredName) Like(name string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.Like(FieldName.Name.String(), name)
	}
}

type PredEmail struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredEmail) EQ(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldEmail.Name.String(), email)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredEmail) NEQ(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldEmail.Name.String(), email)
	}
}

func (f *PredEmail) GT(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GT(FieldEmail.Name.String(), email)
	}
}

func (f *PredEmail) GTE(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldEmail.Name.String(), email)
	}
}

func (f *PredEmail) LT(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LT(FieldEmail.Name.String(), email)
	}
}

func (f *PredEmail) LTE(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldEmail.Name.String(), email)
	}
}

func (f *PredEmail) IN(emails ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(emails))
		for i := range v {
			v[i] = emails[i]
		}
		p.In(FieldEmail.Name.String(), v...)
	}
}

func (f *PredEmail) NotIn(emails ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(emails))
		for i := range v {
			v[i] = emails[i]
		}
		p.NotIn(FieldEmail.Name.String(), v...)
	}
}

func (f *PredEmail) Like(email string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.Like(FieldEmail.Name.String(), email)
	}
}

func (f *PredEmail) IsNull() func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.IsNull(FieldEmail.Name.String())
	}
}

func (f *PredEmail) NotNull() func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NotNull(FieldEmail.Name.String())
	}
}

type PredCreatedTime struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredCreatedTime) EQ(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldCreatedTime.Name.String(), created_time)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredCreatedTime) NEQ(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldCreatedTime.Name.String(), created_time)
	}
}

func (f *PredCreatedTime) GT(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GT(FieldCreatedTime.Name.String(), created_time)
	}
}

func (f *PredCreatedTime) GTE(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldCreatedTime.Name.String(), created_time)
	}
}

func (f *PredCreatedTime) LT(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LT(FieldCreatedTime.Name.String(), created_time)
	}
}

func (f *PredCreatedTime) LTE(created_time time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldCreatedTime.Name.String(), created_time)
	}
}

func (f *PredCreatedTime) IN(created_times ...time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(created_times))
		for i := range v {
			v[i] = created_times[i]
		}
		p.In(FieldCreatedTime.Name.String(), v...)
	}
}

func (f *PredCreatedTime) NotIn(created_times ...time.Time) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(created_times))
		for i := range v {
			v[i] = created_times[i]
		}
		p.NotIn(FieldCreatedTime.Name.String(), v...)
	}
}

func (f *PredCreatedTime) Like(created_time string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.Like(FieldCreatedTime.Name.String(), created_time)
	}
}

type PredDesc struct {
}

// EQ returns a function that sets the predicate to check if the field is equal to the given value.
// Operator "="
func (f *PredDesc) EQ(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.EQ(FieldDesc.Name.String(), desc)
	}
}

// NEQ returns a function that sets the predicate to check if the field is not equal to the given value.
// Operator "<>"
func (f *PredDesc) NEQ(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NEQ(FieldDesc.Name.String(), desc)
	}
}

func (f *PredDesc) GT(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GT(FieldDesc.Name.String(), desc)
	}
}

func (f *PredDesc) GTE(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.GTE(FieldDesc.Name.String(), desc)
	}
}

func (f *PredDesc) LT(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LT(FieldDesc.Name.String(), desc)
	}
}

func (f *PredDesc) LTE(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.LTE(FieldDesc.Name.String(), desc)
	}
}

func (f *PredDesc) IN(descs ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(descs))
		for i := range v {
			v[i] = descs[i]
		}
		p.In(FieldDesc.Name.String(), v...)
	}
}

func (f *PredDesc) NotIn(descs ...string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		v := make([]any, len(descs))
		for i := range v {
			v[i] = descs[i]
		}
		p.NotIn(FieldDesc.Name.String(), v...)
	}
}

func (f *PredDesc) Like(desc string) func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.Like(FieldDesc.Name.String(), desc)
	}
}

func (f *PredDesc) IsNull() func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.IsNull(FieldDesc.Name.String())
	}
}

func (f *PredDesc) NotNull() func(*entitysql.Predicate) {
	return func(p *entitysql.Predicate) {
		p.NotNull(FieldDesc.Name.String())
	}
}

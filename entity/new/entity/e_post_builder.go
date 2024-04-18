package entity

import (
	"context"
	"taurus_go_demo/entity/new/entity/internal"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/new/entity/post"
)

// PostEntityBuilder is a builder for the PostEntity entity.
//
// The builder is used to create, update, and delete PostEntity entities.
type PostEntityBuilder struct {
	config  *PostEntityConfig
	tracker entity.Tracker

	// ID Post primary key
	ID post.PredID

	Content post.PredContent

	BlogID post.PredBlogID

	AuthorID post.PredAuthorID

	Blogs *BlogEntityRelation

	Authors *AuthorEntityRelation

	ByID post.ByID

	ByContent post.ByContent

	ByBlogID post.ByBlogID

	ByAuthorID post.ByAuthorID
}

// NewPostEntityBuilder creates a new PostEntityBuilder.
func NewPostEntityBuilder(c *PostEntityConfig, t entity.Tracker, blogs BlogEntityRelation, authors AuthorEntityRelation) *PostEntityBuilder {
	return &PostEntityBuilder{
		config:  c,
		tracker: t,
		Blogs:   &blogs,
		Authors: &authors,
	}
}

// Create creates a new UserEntity，and add it to the tracker.
// Required parameters are fields that have no default value but are required,
// and options are fields that can be left empty by calling WithFieldName.
func (b *PostEntityBuilder) Create(content string, blog_id int64, author_id int64, options ...func(*PostEntity)) (*PostEntity, error) {
	e := b.config.New()
	switch t := e.(type) {
	case *PostEntity:
		return t.create(content, blog_id, author_id, options...)
	default:
		return nil, entity.Err_0100030006
	}

}

func (b *PostEntityBuilder) Remove(e *PostEntity) error {
	if e.config.Mutation == nil {
		return nil
	}
	return e.remove()
}

// First returns the first PostEntity.
func (s *PostEntityBuilder) First(ctx context.Context) (*PostEntity, error) {
	query := s.initQuery()
	return query.First(ctx)
}

func (s *PostEntityBuilder) ToList(ctx context.Context) ([]*PostEntity, error) {
	query := s.initQuery()
	return query.ToList(ctx)
}

func (s *PostEntityBuilder) Include(rels ...PostEntityRel) *PostEntityQuery {
	query := s.initQuery()
	return query.Include(rels...)
}

func (s *PostEntityBuilder) Order(o ...post.OrderTerm) *PostEntityQuery {
	query := s.initQuery()
	return query.Order(o...)
}

func (s *PostEntityBuilder) Where(conditions ...func(*entitysql.Predicate)) *PostEntityQuery {
	query := s.initQuery()
	return query.Where(conditions...)
}

// Exec executes all the mutations for the PostEntity.
func (s *PostEntityBuilder) Exec(ctx context.Context, tx dialect.Tx) error {
	if len(s.config.postMutations.Addeds) > 0 {
		e := s.config.postMutations.Get(entity.Added)
		n := NewPostEntityCreate(s.config.Dialect, e...)
		if err := n.create(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.config.postMutations.Modifieds) > 0 {
		e := s.config.postMutations.Get(entity.Modified)
		n := NewPostEntityUpdate(s.config.Dialect, e...)
		if err := n.update(ctx, tx); err != nil {
			return err
		}
	}
	if len(s.config.postMutations.Deleteds) > 0 {
		e := s.config.postMutations.Get(entity.Deleted)
		n := NewPostEntityDelete(s.config.Dialect, e...)
		if err := n.delete(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *PostEntityBuilder) initQuery() *PostEntityQuery {
	return NewPostEntityQuery(s.config.Dialect, s.tracker, s.config.postMutations)
}

// postMutations is a collection of PostEntity mutation.
type postMutations struct {
	Detacheds  map[string]*PostEntity
	Unchangeds map[string]*PostEntity
	Deleteds   map[string]*PostEntity
	Modifieds  map[string]*PostEntity
	Addeds     map[string]*PostEntity
}

// newPostMutations creates a new mutations.
func newPostMutations() *postMutations {
	return &postMutations{
		Detacheds:  make(map[string]*PostEntity),
		Unchangeds: make(map[string]*PostEntity),
		Deleteds:   make(map[string]*PostEntity),
		Modifieds:  make(map[string]*PostEntity),
		Addeds:     make(map[string]*PostEntity),
	}
}

// Get returns all the PostEntity in the specified state.
func (ms *postMutations) Get(state entity.EntityState) []*PostEntity {
	switch state {
	case entity.Detached:
		s := make([]*PostEntity, 0, len(ms.Detacheds))
		for _, m := range ms.Detacheds {
			s = append(s, m)
		}
		return s
	case entity.Unchanged:
		s := make([]*PostEntity, 0, len(ms.Unchangeds))
		for _, m := range ms.Unchangeds {
			s = append(s, m)
		}
		return s
	case entity.Deleted:
		s := make([]*PostEntity, 0, len(ms.Deleteds))
		for _, m := range ms.Deleteds {
			s = append(s, m)
		}
		return s
	case entity.Modified:
		s := make([]*PostEntity, 0, len(ms.Modifieds))
		for _, m := range ms.Modifieds {
			s = append(s, m)
		}
		return s
	case entity.Added:
		s := make([]*PostEntity, 0, len(ms.Addeds))
		for _, m := range ms.Addeds {
			s = append(s, m)
		}
		return s
	}
	return nil
}

// SetEntityState sets the state of the entity.
func (ms *postMutations) SetEntityState(e *PostEntity, state entity.EntityState) error {
	m := e.config.Mutation
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return err
	}
	return nil
}

// ChangeEntityState attempts to set the desired entity state,
// but will not do so if the conditions are not met.
func (ms *postMutations) ChangeEntityState(m *entity.Mutation, state entity.EntityState) {
	e := ms.getEntity(m)
	ms.set(e, state)
	if err := internal.SetEntityState(m, state); err != nil {
		return
	}
}

// getEntity returns the entity in the specified state.
func (ms *postMutations) getEntity(m *entity.Mutation) *PostEntity {
	key := m.Key()
	switch m.State() {
	case entity.Detached:
		return ms.Detacheds[key]
	case entity.Unchanged:
		return ms.Unchangeds[key]
	case entity.Deleted:
		return ms.Deleteds[key]
	case entity.Modified:
		return ms.Modifieds[key]
	case entity.Added:
		return ms.Addeds[key]
	}
	return nil
}

// Set sets the entity in the specified state.
func (ms *postMutations) set(e *PostEntity, state entity.EntityState) {
	m := e.config.Mutation
	key := m.Key()
	switch m.State() {
	case entity.Detached:
		delete(ms.Detacheds, key)
	case entity.Unchanged:
		delete(ms.Unchangeds, key)
	case entity.Deleted:
		delete(ms.Deleteds, key)
	case entity.Modified:
		delete(ms.Modifieds, key)
	case entity.Added:
		delete(ms.Addeds, key)
	}
	if state >= 0 {
		switch state {
		case entity.Detached:
			ms.Detacheds[key] = e
		case entity.Unchanged:
			ms.Unchangeds[key] = e
		case entity.Deleted:
			ms.Deleteds[key] = e
		case entity.Modified:
			ms.Modifieds[key] = e
		case entity.Added:
			ms.Addeds[key] = e
		}
	}
}

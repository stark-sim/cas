// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/stark-sim/cas/pkg/ent/role"
	"github.com/stark-sim/cas/pkg/ent/user"
	"github.com/stark-sim/cas/pkg/ent/userrole"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int64 `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// RoleEdge is the edge representation of Role.
type RoleEdge struct {
	Node   *Role  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// RoleConnection is the connection containing edges to Role.
type RoleConnection struct {
	Edges      []*RoleEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *RoleConnection) build(nodes []*Role, pager *rolePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Role
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Role {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Role {
			return nodes[i]
		}
	}
	c.Edges = make([]*RoleEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &RoleEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// RolePaginateOption enables pagination customization.
type RolePaginateOption func(*rolePager) error

// WithRoleOrder configures pagination ordering.
func WithRoleOrder(order *RoleOrder) RolePaginateOption {
	if order == nil {
		order = DefaultRoleOrder
	}
	o := *order
	return func(pager *rolePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultRoleOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithRoleFilter configures pagination filter.
func WithRoleFilter(filter func(*RoleQuery) (*RoleQuery, error)) RolePaginateOption {
	return func(pager *rolePager) error {
		if filter == nil {
			return errors.New("RoleQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type rolePager struct {
	order  *RoleOrder
	filter func(*RoleQuery) (*RoleQuery, error)
}

func newRolePager(opts []RolePaginateOption) (*rolePager, error) {
	pager := &rolePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRoleOrder
	}
	return pager, nil
}

func (p *rolePager) applyFilter(query *RoleQuery) (*RoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *rolePager) toCursor(r *Role) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *rolePager) applyCursors(query *RoleQuery, after, before *Cursor) *RoleQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultRoleOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *rolePager) applyOrder(query *RoleQuery, reverse bool) *RoleQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultRoleOrder.Field {
		query = query.Order(direction.orderFunc(DefaultRoleOrder.Field.field))
	}
	return query
}

func (p *rolePager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultRoleOrder.Field {
			b.Comma().Ident(DefaultRoleOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Role.
func (r *RoleQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...RolePaginateOption,
) (*RoleConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newRolePager(opts)
	if err != nil {
		return nil, err
	}
	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}
	conn := &RoleConnection{Edges: []*RoleEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = r.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	r = pager.applyCursors(r, after, before)
	r = pager.applyOrder(r, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		r.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// RoleOrderFieldCreatedAt orders Role by created_at.
	RoleOrderFieldCreatedAt = &RoleOrderField{
		field: role.FieldCreatedAt,
		toCursor: func(r *Role) Cursor {
			return Cursor{
				ID:    r.ID,
				Value: r.CreatedAt,
			}
		},
	}
	// RoleOrderFieldUpdatedAt orders Role by updated_at.
	RoleOrderFieldUpdatedAt = &RoleOrderField{
		field: role.FieldUpdatedAt,
		toCursor: func(r *Role) Cursor {
			return Cursor{
				ID:    r.ID,
				Value: r.UpdatedAt,
			}
		},
	}
	// RoleOrderFieldDeletedAt orders Role by deleted_at.
	RoleOrderFieldDeletedAt = &RoleOrderField{
		field: role.FieldDeletedAt,
		toCursor: func(r *Role) Cursor {
			return Cursor{
				ID:    r.ID,
				Value: r.DeletedAt,
			}
		},
	}
	// RoleOrderFieldName orders Role by name.
	RoleOrderFieldName = &RoleOrderField{
		field: role.FieldName,
		toCursor: func(r *Role) Cursor {
			return Cursor{
				ID:    r.ID,
				Value: r.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f RoleOrderField) String() string {
	var str string
	switch f.field {
	case role.FieldCreatedAt:
		str = "CREATED_AT"
	case role.FieldUpdatedAt:
		str = "UPDATED_AT"
	case role.FieldDeletedAt:
		str = "DELETED_AT"
	case role.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f RoleOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *RoleOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("RoleOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *RoleOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *RoleOrderFieldUpdatedAt
	case "DELETED_AT":
		*f = *RoleOrderFieldDeletedAt
	case "NAME":
		*f = *RoleOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid RoleOrderField", str)
	}
	return nil
}

// RoleOrderField defines the ordering field of Role.
type RoleOrderField struct {
	field    string
	toCursor func(*Role) Cursor
}

// RoleOrder defines the ordering of Role.
type RoleOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *RoleOrderField `json:"field"`
}

// DefaultRoleOrder is the default ordering of Role.
var DefaultRoleOrder = &RoleOrder{
	Direction: OrderDirectionAsc,
	Field: &RoleOrderField{
		field: role.FieldID,
		toCursor: func(r *Role) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Role into RoleEdge.
func (r *Role) ToEdge(order *RoleOrder) *RoleEdge {
	if order == nil {
		order = DefaultRoleOrder
	}
	return &RoleEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

func (p *userPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldCreatedAt orders User by created_at.
	UserOrderFieldCreatedAt = &UserOrderField{
		field: user.FieldCreatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.CreatedAt,
			}
		},
	}
	// UserOrderFieldUpdatedAt orders User by updated_at.
	UserOrderFieldUpdatedAt = &UserOrderField{
		field: user.FieldUpdatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.UpdatedAt,
			}
		},
	}
	// UserOrderFieldDeletedAt orders User by deleted_at.
	UserOrderFieldDeletedAt = &UserOrderField{
		field: user.FieldDeletedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.DeletedAt,
			}
		},
	}
	// UserOrderFieldName orders User by name.
	UserOrderFieldName = &UserOrderField{
		field: user.FieldName,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Name,
			}
		},
	}
	// UserOrderFieldPhone orders User by phone.
	UserOrderFieldPhone = &UserOrderField{
		field: user.FieldPhone,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Phone,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.field {
	case user.FieldCreatedAt:
		str = "CREATED_AT"
	case user.FieldUpdatedAt:
		str = "UPDATED_AT"
	case user.FieldDeletedAt:
		str = "DELETED_AT"
	case user.FieldName:
		str = "NAME"
	case user.FieldPhone:
		str = "PHONE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *UserOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *UserOrderFieldUpdatedAt
	case "DELETED_AT":
		*f = *UserOrderFieldDeletedAt
	case "NAME":
		*f = *UserOrderFieldName
	case "PHONE":
		*f = *UserOrderFieldPhone
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}

// UserRoleEdge is the edge representation of UserRole.
type UserRoleEdge struct {
	Node   *UserRole `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// UserRoleConnection is the connection containing edges to UserRole.
type UserRoleConnection struct {
	Edges      []*UserRoleEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *UserRoleConnection) build(nodes []*UserRole, pager *userrolePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *UserRole
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *UserRole {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *UserRole {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserRoleEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserRoleEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserRolePaginateOption enables pagination customization.
type UserRolePaginateOption func(*userrolePager) error

// WithUserRoleOrder configures pagination ordering.
func WithUserRoleOrder(order *UserRoleOrder) UserRolePaginateOption {
	if order == nil {
		order = DefaultUserRoleOrder
	}
	o := *order
	return func(pager *userrolePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserRoleOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserRoleFilter configures pagination filter.
func WithUserRoleFilter(filter func(*UserRoleQuery) (*UserRoleQuery, error)) UserRolePaginateOption {
	return func(pager *userrolePager) error {
		if filter == nil {
			return errors.New("UserRoleQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userrolePager struct {
	order  *UserRoleOrder
	filter func(*UserRoleQuery) (*UserRoleQuery, error)
}

func newUserRolePager(opts []UserRolePaginateOption) (*userrolePager, error) {
	pager := &userrolePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserRoleOrder
	}
	return pager, nil
}

func (p *userrolePager) applyFilter(query *UserRoleQuery) (*UserRoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userrolePager) toCursor(ur *UserRole) Cursor {
	return p.order.Field.toCursor(ur)
}

func (p *userrolePager) applyCursors(query *UserRoleQuery, after, before *Cursor) *UserRoleQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserRoleOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userrolePager) applyOrder(query *UserRoleQuery, reverse bool) *UserRoleQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserRoleOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserRoleOrder.Field.field))
	}
	return query
}

func (p *userrolePager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserRoleOrder.Field {
			b.Comma().Ident(DefaultUserRoleOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to UserRole.
func (ur *UserRoleQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserRolePaginateOption,
) (*UserRoleConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserRolePager(opts)
	if err != nil {
		return nil, err
	}
	if ur, err = pager.applyFilter(ur); err != nil {
		return nil, err
	}
	conn := &UserRoleConnection{Edges: []*UserRoleEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = ur.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	ur = pager.applyCursors(ur, after, before)
	ur = pager.applyOrder(ur, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		ur.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ur.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := ur.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserRoleOrderFieldCreatedAt orders UserRole by created_at.
	UserRoleOrderFieldCreatedAt = &UserRoleOrderField{
		field: userrole.FieldCreatedAt,
		toCursor: func(ur *UserRole) Cursor {
			return Cursor{
				ID:    ur.ID,
				Value: ur.CreatedAt,
			}
		},
	}
	// UserRoleOrderFieldUpdatedAt orders UserRole by updated_at.
	UserRoleOrderFieldUpdatedAt = &UserRoleOrderField{
		field: userrole.FieldUpdatedAt,
		toCursor: func(ur *UserRole) Cursor {
			return Cursor{
				ID:    ur.ID,
				Value: ur.UpdatedAt,
			}
		},
	}
	// UserRoleOrderFieldDeletedAt orders UserRole by deleted_at.
	UserRoleOrderFieldDeletedAt = &UserRoleOrderField{
		field: userrole.FieldDeletedAt,
		toCursor: func(ur *UserRole) Cursor {
			return Cursor{
				ID:    ur.ID,
				Value: ur.DeletedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserRoleOrderField) String() string {
	var str string
	switch f.field {
	case userrole.FieldCreatedAt:
		str = "CREATED_AT"
	case userrole.FieldUpdatedAt:
		str = "UPDATED_AT"
	case userrole.FieldDeletedAt:
		str = "DELETED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserRoleOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserRoleOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserRoleOrderField %T must be a string", v)
	}
	switch str {
	case "CREATED_AT":
		*f = *UserRoleOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *UserRoleOrderFieldUpdatedAt
	case "DELETED_AT":
		*f = *UserRoleOrderFieldDeletedAt
	default:
		return fmt.Errorf("%s is not a valid UserRoleOrderField", str)
	}
	return nil
}

// UserRoleOrderField defines the ordering field of UserRole.
type UserRoleOrderField struct {
	field    string
	toCursor func(*UserRole) Cursor
}

// UserRoleOrder defines the ordering of UserRole.
type UserRoleOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *UserRoleOrderField `json:"field"`
}

// DefaultUserRoleOrder is the default ordering of UserRole.
var DefaultUserRoleOrder = &UserRoleOrder{
	Direction: OrderDirectionAsc,
	Field: &UserRoleOrderField{
		field: userrole.FieldID,
		toCursor: func(ur *UserRole) Cursor {
			return Cursor{ID: ur.ID}
		},
	},
}

// ToEdge converts UserRole into UserRoleEdge.
func (ur *UserRole) ToEdge(order *UserRoleOrder) *UserRoleEdge {
	if order == nil {
		order = DefaultUserRoleOrder
	}
	return &UserRoleEdge{
		Node:   ur,
		Cursor: order.Field.toCursor(ur),
	}
}

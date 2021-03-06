// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// RoleUser is an object representing the database table.
type RoleUser struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UserID    null.Int  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	RoleID    null.Int  `boil:"role_id" json:"role_id,omitempty" toml:"role_id" yaml:"role_id,omitempty"`

	R *roleUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L roleUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoleUserColumns = struct {
	ID        string
	UpdatedAt string
	CreatedAt string
	UserID    string
	RoleID    string
}{
	ID:        "id",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
	UserID:    "user_id",
	RoleID:    "role_id",
}

// Generated where

var RoleUserWhere = struct {
	ID        whereHelperint
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	UserID    whereHelpernull_Int
	RoleID    whereHelpernull_Int
}{
	ID:        whereHelperint{field: "\"role_user\".\"id\""},
	UpdatedAt: whereHelpertime_Time{field: "\"role_user\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"role_user\".\"created_at\""},
	UserID:    whereHelpernull_Int{field: "\"role_user\".\"user_id\""},
	RoleID:    whereHelpernull_Int{field: "\"role_user\".\"role_id\""},
}

// RoleUserRels is where relationship names are stored.
var RoleUserRels = struct {
	Role string
	User string
}{
	Role: "Role",
	User: "User",
}

// roleUserR is where relationships are stored.
type roleUserR struct {
	Role *Role    `boil:"Role" json:"Role" toml:"Role" yaml:"Role"`
	User *Account `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*roleUserR) NewStruct() *roleUserR {
	return &roleUserR{}
}

// roleUserL is where Load methods for each relationship are stored.
type roleUserL struct{}

var (
	roleUserAllColumns            = []string{"id", "updated_at", "created_at", "user_id", "role_id"}
	roleUserColumnsWithoutDefault = []string{"created_at", "user_id", "role_id"}
	roleUserColumnsWithDefault    = []string{"id", "updated_at"}
	roleUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// RoleUserSlice is an alias for a slice of pointers to RoleUser.
	// This should generally be used opposed to []RoleUser.
	RoleUserSlice []*RoleUser
	// RoleUserHook is the signature for custom RoleUser hook methods
	RoleUserHook func(context.Context, boil.ContextExecutor, *RoleUser) error

	roleUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	roleUserType                 = reflect.TypeOf(&RoleUser{})
	roleUserMapping              = queries.MakeStructMapping(roleUserType)
	roleUserPrimaryKeyMapping, _ = queries.BindMapping(roleUserType, roleUserMapping, roleUserPrimaryKeyColumns)
	roleUserInsertCacheMut       sync.RWMutex
	roleUserInsertCache          = make(map[string]insertCache)
	roleUserUpdateCacheMut       sync.RWMutex
	roleUserUpdateCache          = make(map[string]updateCache)
	roleUserUpsertCacheMut       sync.RWMutex
	roleUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var roleUserBeforeInsertHooks []RoleUserHook
var roleUserBeforeUpdateHooks []RoleUserHook
var roleUserBeforeDeleteHooks []RoleUserHook
var roleUserBeforeUpsertHooks []RoleUserHook

var roleUserAfterInsertHooks []RoleUserHook
var roleUserAfterSelectHooks []RoleUserHook
var roleUserAfterUpdateHooks []RoleUserHook
var roleUserAfterDeleteHooks []RoleUserHook
var roleUserAfterUpsertHooks []RoleUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RoleUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RoleUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RoleUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RoleUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RoleUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RoleUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RoleUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RoleUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RoleUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRoleUserHook registers your hook function for all future operations.
func AddRoleUserHook(hookPoint boil.HookPoint, roleUserHook RoleUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		roleUserBeforeInsertHooks = append(roleUserBeforeInsertHooks, roleUserHook)
	case boil.BeforeUpdateHook:
		roleUserBeforeUpdateHooks = append(roleUserBeforeUpdateHooks, roleUserHook)
	case boil.BeforeDeleteHook:
		roleUserBeforeDeleteHooks = append(roleUserBeforeDeleteHooks, roleUserHook)
	case boil.BeforeUpsertHook:
		roleUserBeforeUpsertHooks = append(roleUserBeforeUpsertHooks, roleUserHook)
	case boil.AfterInsertHook:
		roleUserAfterInsertHooks = append(roleUserAfterInsertHooks, roleUserHook)
	case boil.AfterSelectHook:
		roleUserAfterSelectHooks = append(roleUserAfterSelectHooks, roleUserHook)
	case boil.AfterUpdateHook:
		roleUserAfterUpdateHooks = append(roleUserAfterUpdateHooks, roleUserHook)
	case boil.AfterDeleteHook:
		roleUserAfterDeleteHooks = append(roleUserAfterDeleteHooks, roleUserHook)
	case boil.AfterUpsertHook:
		roleUserAfterUpsertHooks = append(roleUserAfterUpsertHooks, roleUserHook)
	}
}

// One returns a single roleUser record from the query.
func (q roleUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RoleUser, error) {
	o := &RoleUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for role_user")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RoleUser records from the query.
func (q roleUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoleUserSlice, error) {
	var o []*RoleUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RoleUser slice")
	}

	if len(roleUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RoleUser records in the query.
func (q roleUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count role_user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q roleUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if role_user exists")
	}

	return count > 0, nil
}

// Role pointed to by the foreign key.
func (o *RoleUser) Role(mods ...qm.QueryMod) roleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RoleID),
	}

	queryMods = append(queryMods, mods...)

	query := Roles(queryMods...)
	queries.SetFrom(query.Query, "\"roles\"")

	return query
}

// User pointed to by the foreign key.
func (o *RoleUser) User(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"accounts\"")

	return query
}

// LoadRole allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (roleUserL) LoadRole(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRoleUser interface{}, mods queries.Applicator) error {
	var slice []*RoleUser
	var object *RoleUser

	if singular {
		object = maybeRoleUser.(*RoleUser)
	} else {
		slice = *maybeRoleUser.(*[]*RoleUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleUserR{}
		}
		if !queries.IsNil(object.RoleID) {
			args = append(args, object.RoleID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleUserR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.RoleID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.RoleID) {
				args = append(args, obj.RoleID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`roles`),
		qm.WhereIn(`roles.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Role")
	}

	var resultSlice []*Role
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Role")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for roles")
	}

	if len(roleUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Role = foreign
		if foreign.R == nil {
			foreign.R = &roleR{}
		}
		foreign.R.RoleUsers = append(foreign.R.RoleUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.RoleID, foreign.ID) {
				local.R.Role = foreign
				if foreign.R == nil {
					foreign.R = &roleR{}
				}
				foreign.R.RoleUsers = append(foreign.R.RoleUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (roleUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRoleUser interface{}, mods queries.Applicator) error {
	var slice []*RoleUser
	var object *RoleUser

	if singular {
		object = maybeRoleUser.(*RoleUser)
	} else {
		slice = *maybeRoleUser.(*[]*RoleUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleUserR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleUserR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`accounts`),
		qm.WhereIn(`accounts.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for accounts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for accounts")
	}

	if len(roleUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.UserRoleUsers = append(foreign.R.UserRoleUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.UserRoleUsers = append(foreign.R.UserRoleUsers, local)
				break
			}
		}
	}

	return nil
}

// SetRole of the roleUser to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleUsers.
func (o *RoleUser) SetRole(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Role) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"role_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
		strmangle.WhereClause("\"", "\"", 2, roleUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.RoleID, related.ID)
	if o.R == nil {
		o.R = &roleUserR{
			Role: related,
		}
	} else {
		o.R.Role = related
	}

	if related.R == nil {
		related.R = &roleR{
			RoleUsers: RoleUserSlice{o},
		}
	} else {
		related.R.RoleUsers = append(related.R.RoleUsers, o)
	}

	return nil
}

// RemoveRole relationship.
// Sets o.R.Role to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *RoleUser) RemoveRole(ctx context.Context, exec boil.ContextExecutor, related *Role) error {
	var err error

	queries.SetScanner(&o.RoleID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("role_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Role = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.RoleUsers {
		if queries.Equal(o.RoleID, ri.RoleID) {
			continue
		}

		ln := len(related.R.RoleUsers)
		if ln > 1 && i < ln-1 {
			related.R.RoleUsers[i] = related.R.RoleUsers[ln-1]
		}
		related.R.RoleUsers = related.R.RoleUsers[:ln-1]
		break
	}
	return nil
}

// SetUser of the roleUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserRoleUsers.
func (o *RoleUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"role_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, roleUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &roleUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &accountR{
			UserRoleUsers: RoleUserSlice{o},
		}
	} else {
		related.R.UserRoleUsers = append(related.R.UserRoleUsers, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *RoleUser) RemoveUser(ctx context.Context, exec boil.ContextExecutor, related *Account) error {
	var err error

	queries.SetScanner(&o.UserID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.User = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.UserRoleUsers {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.UserRoleUsers)
		if ln > 1 && i < ln-1 {
			related.R.UserRoleUsers[i] = related.R.UserRoleUsers[ln-1]
		}
		related.R.UserRoleUsers = related.R.UserRoleUsers[:ln-1]
		break
	}
	return nil
}

// RoleUsers retrieves all the records using an executor.
func RoleUsers(mods ...qm.QueryMod) roleUserQuery {
	mods = append(mods, qm.From("\"role_user\""))
	return roleUserQuery{NewQuery(mods...)}
}

// FindRoleUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRoleUser(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RoleUser, error) {
	roleUserObj := &RoleUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"role_user\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, roleUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from role_user")
	}

	return roleUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RoleUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no role_user provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	roleUserInsertCacheMut.RLock()
	cache, cached := roleUserInsertCache[key]
	roleUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			roleUserAllColumns,
			roleUserColumnsWithDefault,
			roleUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(roleUserType, roleUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(roleUserType, roleUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"role_user\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"role_user\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into role_user")
	}

	if !cached {
		roleUserInsertCacheMut.Lock()
		roleUserInsertCache[key] = cache
		roleUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RoleUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RoleUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	roleUserUpdateCacheMut.RLock()
	cache, cached := roleUserUpdateCache[key]
	roleUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			roleUserAllColumns,
			roleUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update role_user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"role_user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, roleUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(roleUserType, roleUserMapping, append(wl, roleUserPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update role_user row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for role_user")
	}

	if !cached {
		roleUserUpdateCacheMut.Lock()
		roleUserUpdateCache[key] = cache
		roleUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q roleUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for role_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for role_user")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoleUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roleUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"role_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, roleUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in roleUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all roleUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RoleUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no role_user provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	roleUserUpsertCacheMut.RLock()
	cache, cached := roleUserUpsertCache[key]
	roleUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			roleUserAllColumns,
			roleUserColumnsWithDefault,
			roleUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			roleUserAllColumns,
			roleUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert role_user, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(roleUserPrimaryKeyColumns))
			copy(conflict, roleUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"role_user\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(roleUserType, roleUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(roleUserType, roleUserMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert role_user")
	}

	if !cached {
		roleUserUpsertCacheMut.Lock()
		roleUserUpsertCache[key] = cache
		roleUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RoleUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RoleUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RoleUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), roleUserPrimaryKeyMapping)
	sql := "DELETE FROM \"role_user\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from role_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for role_user")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q roleUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no roleUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from role_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for role_user")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoleUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(roleUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roleUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"role_user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, roleUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from roleUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for role_user")
	}

	if len(roleUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RoleUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRoleUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoleUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoleUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), roleUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"role_user\".* FROM \"role_user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, roleUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RoleUserSlice")
	}

	*o = slice

	return nil
}

// RoleUserExists checks if the RoleUser row exists.
func RoleUserExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"role_user\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if role_user exists")
	}

	return exists, nil
}

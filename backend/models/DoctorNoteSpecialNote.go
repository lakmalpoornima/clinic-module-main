// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// DoctorNoteSpecialNote is an object representing the database table.
type DoctorNoteSpecialNote struct {
	ID                 int         `boil:"ID" json:"ID" toml:"ID" yaml:"ID"`
	CompanyCode        string      `boil:"CompanyCode" json:"CompanyCode" toml:"CompanyCode" yaml:"CompanyCode"`
	DocumentID         null.String `boil:"DocumentID" json:"DocumentID,omitempty" toml:"DocumentID" yaml:"DocumentID,omitempty"`
	RegistrationID     null.String `boil:"RegistrationID" json:"RegistrationID,omitempty" toml:"RegistrationID" yaml:"RegistrationID,omitempty"`
	ProblemCode        null.String `boil:"ProblemCode" json:"ProblemCode,omitempty" toml:"ProblemCode" yaml:"ProblemCode,omitempty"`
	ProblemDescription null.String `boil:"ProblemDescription" json:"ProblemDescription,omitempty" toml:"ProblemDescription" yaml:"ProblemDescription,omitempty"`
	Note               null.String `boil:"Note" json:"Note,omitempty" toml:"Note" yaml:"Note,omitempty"`
	Status             null.String `boil:"Status" json:"Status,omitempty" toml:"Status" yaml:"Status,omitempty"`
	Priority           null.Int    `boil:"Priority" json:"Priority,omitempty" toml:"Priority" yaml:"Priority,omitempty"`
	CreatedDate        time.Time   `boil:"CreatedDate" json:"CreatedDate" toml:"CreatedDate" yaml:"CreatedDate"`

	R *doctorNoteSpecialNoteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L doctorNoteSpecialNoteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DoctorNoteSpecialNoteColumns = struct {
	ID                 string
	CompanyCode        string
	DocumentID         string
	RegistrationID     string
	ProblemCode        string
	ProblemDescription string
	Note               string
	Status             string
	Priority           string
	CreatedDate        string
}{
	ID:                 "ID",
	CompanyCode:        "CompanyCode",
	DocumentID:         "DocumentID",
	RegistrationID:     "RegistrationID",
	ProblemCode:        "ProblemCode",
	ProblemDescription: "ProblemDescription",
	Note:               "Note",
	Status:             "Status",
	Priority:           "Priority",
	CreatedDate:        "CreatedDate",
}

var DoctorNoteSpecialNoteTableColumns = struct {
	ID                 string
	CompanyCode        string
	DocumentID         string
	RegistrationID     string
	ProblemCode        string
	ProblemDescription string
	Note               string
	Status             string
	Priority           string
	CreatedDate        string
}{
	ID:                 "DoctorNoteSpecialNote.ID",
	CompanyCode:        "DoctorNoteSpecialNote.CompanyCode",
	DocumentID:         "DoctorNoteSpecialNote.DocumentID",
	RegistrationID:     "DoctorNoteSpecialNote.RegistrationID",
	ProblemCode:        "DoctorNoteSpecialNote.ProblemCode",
	ProblemDescription: "DoctorNoteSpecialNote.ProblemDescription",
	Note:               "DoctorNoteSpecialNote.Note",
	Status:             "DoctorNoteSpecialNote.Status",
	Priority:           "DoctorNoteSpecialNote.Priority",
	CreatedDate:        "DoctorNoteSpecialNote.CreatedDate",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var DoctorNoteSpecialNoteWhere = struct {
	ID                 whereHelperint
	CompanyCode        whereHelperstring
	DocumentID         whereHelpernull_String
	RegistrationID     whereHelpernull_String
	ProblemCode        whereHelpernull_String
	ProblemDescription whereHelpernull_String
	Note               whereHelpernull_String
	Status             whereHelpernull_String
	Priority           whereHelpernull_Int
	CreatedDate        whereHelpertime_Time
}{
	ID:                 whereHelperint{field: "[dbo].[DoctorNoteSpecialNote].[ID]"},
	CompanyCode:        whereHelperstring{field: "[dbo].[DoctorNoteSpecialNote].[CompanyCode]"},
	DocumentID:         whereHelpernull_String{field: "[dbo].[DoctorNoteSpecialNote].[DocumentID]"},
	RegistrationID:     whereHelpernull_String{field: "[dbo].[DoctorNoteSpecialNote].[RegistrationID]"},
	ProblemCode:        whereHelpernull_String{field: "[dbo].[DoctorNoteSpecialNote].[ProblemCode]"},
	ProblemDescription: whereHelpernull_String{field: "[dbo].[DoctorNoteSpecialNote].[ProblemDescription]"},
	Note:               whereHelpernull_String{field: "[dbo].[DoctorNoteSpecialNote].[Note]"},
	Status:             whereHelpernull_String{field: "[dbo].[DoctorNoteSpecialNote].[Status]"},
	Priority:           whereHelpernull_Int{field: "[dbo].[DoctorNoteSpecialNote].[Priority]"},
	CreatedDate:        whereHelpertime_Time{field: "[dbo].[DoctorNoteSpecialNote].[CreatedDate]"},
}

// DoctorNoteSpecialNoteRels is where relationship names are stored.
var DoctorNoteSpecialNoteRels = struct {
}{}

// doctorNoteSpecialNoteR is where relationships are stored.
type doctorNoteSpecialNoteR struct {
}

// NewStruct creates a new relationship struct
func (*doctorNoteSpecialNoteR) NewStruct() *doctorNoteSpecialNoteR {
	return &doctorNoteSpecialNoteR{}
}

// doctorNoteSpecialNoteL is where Load methods for each relationship are stored.
type doctorNoteSpecialNoteL struct{}

var (
	doctorNoteSpecialNoteAllColumns            = []string{"ID", "CompanyCode", "DocumentID", "RegistrationID", "ProblemCode", "ProblemDescription", "Note", "Status", "Priority", "CreatedDate"}
	doctorNoteSpecialNoteColumnsWithAuto       = []string{}
	doctorNoteSpecialNoteColumnsWithoutDefault = []string{"RegistrationID", "ProblemCode", "ProblemDescription", "Note", "Status", "Priority"}
	doctorNoteSpecialNoteColumnsWithDefault    = []string{"ID", "CompanyCode", "DocumentID", "CreatedDate"}
	doctorNoteSpecialNotePrimaryKeyColumns     = []string{"ID"}
)

type (
	// DoctorNoteSpecialNoteSlice is an alias for a slice of pointers to DoctorNoteSpecialNote.
	// This should almost always be used instead of []DoctorNoteSpecialNote.
	DoctorNoteSpecialNoteSlice []*DoctorNoteSpecialNote
	// DoctorNoteSpecialNoteHook is the signature for custom DoctorNoteSpecialNote hook methods
	DoctorNoteSpecialNoteHook func(context.Context, boil.ContextExecutor, *DoctorNoteSpecialNote) error

	doctorNoteSpecialNoteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	doctorNoteSpecialNoteType                 = reflect.TypeOf(&DoctorNoteSpecialNote{})
	doctorNoteSpecialNoteMapping              = queries.MakeStructMapping(doctorNoteSpecialNoteType)
	doctorNoteSpecialNotePrimaryKeyMapping, _ = queries.BindMapping(doctorNoteSpecialNoteType, doctorNoteSpecialNoteMapping, doctorNoteSpecialNotePrimaryKeyColumns)
	doctorNoteSpecialNoteInsertCacheMut       sync.RWMutex
	doctorNoteSpecialNoteInsertCache          = make(map[string]insertCache)
	doctorNoteSpecialNoteUpdateCacheMut       sync.RWMutex
	doctorNoteSpecialNoteUpdateCache          = make(map[string]updateCache)
	doctorNoteSpecialNoteUpsertCacheMut       sync.RWMutex
	doctorNoteSpecialNoteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var doctorNoteSpecialNoteBeforeInsertHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteBeforeUpdateHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteBeforeDeleteHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteBeforeUpsertHooks []DoctorNoteSpecialNoteHook

var doctorNoteSpecialNoteAfterInsertHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteAfterSelectHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteAfterUpdateHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteAfterDeleteHooks []DoctorNoteSpecialNoteHook
var doctorNoteSpecialNoteAfterUpsertHooks []DoctorNoteSpecialNoteHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DoctorNoteSpecialNote) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DoctorNoteSpecialNote) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DoctorNoteSpecialNote) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DoctorNoteSpecialNote) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DoctorNoteSpecialNote) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DoctorNoteSpecialNote) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DoctorNoteSpecialNote) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DoctorNoteSpecialNote) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DoctorNoteSpecialNote) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteSpecialNoteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDoctorNoteSpecialNoteHook registers your hook function for all future operations.
func AddDoctorNoteSpecialNoteHook(hookPoint boil.HookPoint, doctorNoteSpecialNoteHook DoctorNoteSpecialNoteHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		doctorNoteSpecialNoteBeforeInsertHooks = append(doctorNoteSpecialNoteBeforeInsertHooks, doctorNoteSpecialNoteHook)
	case boil.BeforeUpdateHook:
		doctorNoteSpecialNoteBeforeUpdateHooks = append(doctorNoteSpecialNoteBeforeUpdateHooks, doctorNoteSpecialNoteHook)
	case boil.BeforeDeleteHook:
		doctorNoteSpecialNoteBeforeDeleteHooks = append(doctorNoteSpecialNoteBeforeDeleteHooks, doctorNoteSpecialNoteHook)
	case boil.BeforeUpsertHook:
		doctorNoteSpecialNoteBeforeUpsertHooks = append(doctorNoteSpecialNoteBeforeUpsertHooks, doctorNoteSpecialNoteHook)
	case boil.AfterInsertHook:
		doctorNoteSpecialNoteAfterInsertHooks = append(doctorNoteSpecialNoteAfterInsertHooks, doctorNoteSpecialNoteHook)
	case boil.AfterSelectHook:
		doctorNoteSpecialNoteAfterSelectHooks = append(doctorNoteSpecialNoteAfterSelectHooks, doctorNoteSpecialNoteHook)
	case boil.AfterUpdateHook:
		doctorNoteSpecialNoteAfterUpdateHooks = append(doctorNoteSpecialNoteAfterUpdateHooks, doctorNoteSpecialNoteHook)
	case boil.AfterDeleteHook:
		doctorNoteSpecialNoteAfterDeleteHooks = append(doctorNoteSpecialNoteAfterDeleteHooks, doctorNoteSpecialNoteHook)
	case boil.AfterUpsertHook:
		doctorNoteSpecialNoteAfterUpsertHooks = append(doctorNoteSpecialNoteAfterUpsertHooks, doctorNoteSpecialNoteHook)
	}
}

// One returns a single doctorNoteSpecialNote record from the query.
func (q doctorNoteSpecialNoteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DoctorNoteSpecialNote, error) {
	o := &DoctorNoteSpecialNote{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for DoctorNoteSpecialNote")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DoctorNoteSpecialNote records from the query.
func (q doctorNoteSpecialNoteQuery) All(ctx context.Context, exec boil.ContextExecutor) (DoctorNoteSpecialNoteSlice, error) {
	var o []*DoctorNoteSpecialNote

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DoctorNoteSpecialNote slice")
	}

	if len(doctorNoteSpecialNoteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DoctorNoteSpecialNote records in the query.
func (q doctorNoteSpecialNoteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count DoctorNoteSpecialNote rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q doctorNoteSpecialNoteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if DoctorNoteSpecialNote exists")
	}

	return count > 0, nil
}

// DoctorNoteSpecialNotes retrieves all the records using an executor.
func DoctorNoteSpecialNotes(mods ...qm.QueryMod) doctorNoteSpecialNoteQuery {
	mods = append(mods, qm.From("[dbo].[DoctorNoteSpecialNote]"))
	return doctorNoteSpecialNoteQuery{NewQuery(mods...)}
}

// FindDoctorNoteSpecialNote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDoctorNoteSpecialNote(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DoctorNoteSpecialNote, error) {
	doctorNoteSpecialNoteObj := &DoctorNoteSpecialNote{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[DoctorNoteSpecialNote] where [ID]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, doctorNoteSpecialNoteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from DoctorNoteSpecialNote")
	}

	if err = doctorNoteSpecialNoteObj.doAfterSelectHooks(ctx, exec); err != nil {
		return doctorNoteSpecialNoteObj, err
	}

	return doctorNoteSpecialNoteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DoctorNoteSpecialNote) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no DoctorNoteSpecialNote provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(doctorNoteSpecialNoteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	doctorNoteSpecialNoteInsertCacheMut.RLock()
	cache, cached := doctorNoteSpecialNoteInsertCache[key]
	doctorNoteSpecialNoteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			doctorNoteSpecialNoteAllColumns,
			doctorNoteSpecialNoteColumnsWithDefault,
			doctorNoteSpecialNoteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(doctorNoteSpecialNoteType, doctorNoteSpecialNoteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(doctorNoteSpecialNoteType, doctorNoteSpecialNoteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[DoctorNoteSpecialNote] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[DoctorNoteSpecialNote] %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryOutput = fmt.Sprintf("OUTPUT INSERTED.[%s] ", strings.Join(returnColumns, "],INSERTED.["))
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
		return errors.Wrap(err, "models: unable to insert into DoctorNoteSpecialNote")
	}

	if !cached {
		doctorNoteSpecialNoteInsertCacheMut.Lock()
		doctorNoteSpecialNoteInsertCache[key] = cache
		doctorNoteSpecialNoteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DoctorNoteSpecialNote.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DoctorNoteSpecialNote) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	doctorNoteSpecialNoteUpdateCacheMut.RLock()
	cache, cached := doctorNoteSpecialNoteUpdateCache[key]
	doctorNoteSpecialNoteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			doctorNoteSpecialNoteAllColumns,
			doctorNoteSpecialNotePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, doctorNoteSpecialNoteColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update DoctorNoteSpecialNote, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[DoctorNoteSpecialNote] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, doctorNoteSpecialNotePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(doctorNoteSpecialNoteType, doctorNoteSpecialNoteMapping, append(wl, doctorNoteSpecialNotePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update DoctorNoteSpecialNote row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for DoctorNoteSpecialNote")
	}

	if !cached {
		doctorNoteSpecialNoteUpdateCacheMut.Lock()
		doctorNoteSpecialNoteUpdateCache[key] = cache
		doctorNoteSpecialNoteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q doctorNoteSpecialNoteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for DoctorNoteSpecialNote")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for DoctorNoteSpecialNote")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DoctorNoteSpecialNoteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorNoteSpecialNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[DoctorNoteSpecialNote] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, doctorNoteSpecialNotePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in doctorNoteSpecialNote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all doctorNoteSpecialNote")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DoctorNoteSpecialNote) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no DoctorNoteSpecialNote provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(doctorNoteSpecialNoteColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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

	doctorNoteSpecialNoteUpsertCacheMut.RLock()
	cache, cached := doctorNoteSpecialNoteUpsertCache[key]
	doctorNoteSpecialNoteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			doctorNoteSpecialNoteAllColumns,
			doctorNoteSpecialNoteColumnsWithDefault,
			doctorNoteSpecialNoteColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, doctorNoteSpecialNoteColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(doctorNoteSpecialNotePrimaryKeyColumns, v) && strmangle.ContainsAny(doctorNoteSpecialNoteColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert DoctorNoteSpecialNote, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, doctorNoteSpecialNoteColumnsWithAuto)
		ret = strmangle.SetMerge(ret, doctorNoteSpecialNoteColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			doctorNoteSpecialNoteAllColumns,
			doctorNoteSpecialNotePrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, doctorNoteSpecialNoteColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert DoctorNoteSpecialNote, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[DoctorNoteSpecialNote]", doctorNoteSpecialNotePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(doctorNoteSpecialNotePrimaryKeyColumns))
		copy(whitelist, doctorNoteSpecialNotePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(doctorNoteSpecialNoteType, doctorNoteSpecialNoteMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(doctorNoteSpecialNoteType, doctorNoteSpecialNoteMapping, ret)
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
			err = nil // MSSQL doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert DoctorNoteSpecialNote")
	}

	if !cached {
		doctorNoteSpecialNoteUpsertCacheMut.Lock()
		doctorNoteSpecialNoteUpsertCache[key] = cache
		doctorNoteSpecialNoteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DoctorNoteSpecialNote record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DoctorNoteSpecialNote) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DoctorNoteSpecialNote provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), doctorNoteSpecialNotePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[DoctorNoteSpecialNote] WHERE [ID]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from DoctorNoteSpecialNote")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for DoctorNoteSpecialNote")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q doctorNoteSpecialNoteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no doctorNoteSpecialNoteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from DoctorNoteSpecialNote")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for DoctorNoteSpecialNote")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DoctorNoteSpecialNoteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(doctorNoteSpecialNoteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorNoteSpecialNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[DoctorNoteSpecialNote] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, doctorNoteSpecialNotePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from doctorNoteSpecialNote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for DoctorNoteSpecialNote")
	}

	if len(doctorNoteSpecialNoteAfterDeleteHooks) != 0 {
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
func (o *DoctorNoteSpecialNote) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDoctorNoteSpecialNote(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DoctorNoteSpecialNoteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DoctorNoteSpecialNoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorNoteSpecialNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[DoctorNoteSpecialNote].* FROM [dbo].[DoctorNoteSpecialNote] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, doctorNoteSpecialNotePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DoctorNoteSpecialNoteSlice")
	}

	*o = slice

	return nil
}

// DoctorNoteSpecialNoteExists checks if the DoctorNoteSpecialNote row exists.
func DoctorNoteSpecialNoteExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[DoctorNoteSpecialNote] where [ID]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if DoctorNoteSpecialNote exists")
	}

	return exists, nil
}
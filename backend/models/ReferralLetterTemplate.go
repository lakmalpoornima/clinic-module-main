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

// ReferralLetterTemplate is an object representing the database table.
type ReferralLetterTemplate struct {
	TemplateID   int         `boil:"TemplateID" json:"TemplateID" toml:"TemplateID" yaml:"TemplateID"`
	TemplateName null.String `boil:"TemplateName" json:"TemplateName,omitempty" toml:"TemplateName" yaml:"TemplateName,omitempty"`
	DoctorID     string      `boil:"DoctorID" json:"DoctorID" toml:"DoctorID" yaml:"DoctorID"`
	Letter       null.String `boil:"Letter" json:"Letter,omitempty" toml:"Letter" yaml:"Letter,omitempty"`
	ClinicID     null.String `boil:"ClinicID" json:"ClinicID,omitempty" toml:"ClinicID" yaml:"ClinicID,omitempty"`

	R *referralLetterTemplateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L referralLetterTemplateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReferralLetterTemplateColumns = struct {
	TemplateID   string
	TemplateName string
	DoctorID     string
	Letter       string
	ClinicID     string
}{
	TemplateID:   "TemplateID",
	TemplateName: "TemplateName",
	DoctorID:     "DoctorID",
	Letter:       "Letter",
	ClinicID:     "ClinicID",
}

var ReferralLetterTemplateTableColumns = struct {
	TemplateID   string
	TemplateName string
	DoctorID     string
	Letter       string
	ClinicID     string
}{
	TemplateID:   "ReferralLetterTemplate.TemplateID",
	TemplateName: "ReferralLetterTemplate.TemplateName",
	DoctorID:     "ReferralLetterTemplate.DoctorID",
	Letter:       "ReferralLetterTemplate.Letter",
	ClinicID:     "ReferralLetterTemplate.ClinicID",
}

// Generated where

var ReferralLetterTemplateWhere = struct {
	TemplateID   whereHelperint
	TemplateName whereHelpernull_String
	DoctorID     whereHelperstring
	Letter       whereHelpernull_String
	ClinicID     whereHelpernull_String
}{
	TemplateID:   whereHelperint{field: "[dbo].[ReferralLetterTemplate].[TemplateID]"},
	TemplateName: whereHelpernull_String{field: "[dbo].[ReferralLetterTemplate].[TemplateName]"},
	DoctorID:     whereHelperstring{field: "[dbo].[ReferralLetterTemplate].[DoctorID]"},
	Letter:       whereHelpernull_String{field: "[dbo].[ReferralLetterTemplate].[Letter]"},
	ClinicID:     whereHelpernull_String{field: "[dbo].[ReferralLetterTemplate].[ClinicID]"},
}

// ReferralLetterTemplateRels is where relationship names are stored.
var ReferralLetterTemplateRels = struct {
}{}

// referralLetterTemplateR is where relationships are stored.
type referralLetterTemplateR struct {
}

// NewStruct creates a new relationship struct
func (*referralLetterTemplateR) NewStruct() *referralLetterTemplateR {
	return &referralLetterTemplateR{}
}

// referralLetterTemplateL is where Load methods for each relationship are stored.
type referralLetterTemplateL struct{}

var (
	referralLetterTemplateAllColumns            = []string{"TemplateID", "TemplateName", "DoctorID", "Letter", "ClinicID"}
	referralLetterTemplateColumnsWithAuto       = []string{}
	referralLetterTemplateColumnsWithoutDefault = []string{"DoctorID", "Letter"}
	referralLetterTemplateColumnsWithDefault    = []string{"TemplateID", "TemplateName", "ClinicID"}
	referralLetterTemplatePrimaryKeyColumns     = []string{"TemplateID"}
)

type (
	// ReferralLetterTemplateSlice is an alias for a slice of pointers to ReferralLetterTemplate.
	// This should almost always be used instead of []ReferralLetterTemplate.
	ReferralLetterTemplateSlice []*ReferralLetterTemplate
	// ReferralLetterTemplateHook is the signature for custom ReferralLetterTemplate hook methods
	ReferralLetterTemplateHook func(context.Context, boil.ContextExecutor, *ReferralLetterTemplate) error

	referralLetterTemplateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	referralLetterTemplateType                 = reflect.TypeOf(&ReferralLetterTemplate{})
	referralLetterTemplateMapping              = queries.MakeStructMapping(referralLetterTemplateType)
	referralLetterTemplatePrimaryKeyMapping, _ = queries.BindMapping(referralLetterTemplateType, referralLetterTemplateMapping, referralLetterTemplatePrimaryKeyColumns)
	referralLetterTemplateInsertCacheMut       sync.RWMutex
	referralLetterTemplateInsertCache          = make(map[string]insertCache)
	referralLetterTemplateUpdateCacheMut       sync.RWMutex
	referralLetterTemplateUpdateCache          = make(map[string]updateCache)
	referralLetterTemplateUpsertCacheMut       sync.RWMutex
	referralLetterTemplateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var referralLetterTemplateBeforeInsertHooks []ReferralLetterTemplateHook
var referralLetterTemplateBeforeUpdateHooks []ReferralLetterTemplateHook
var referralLetterTemplateBeforeDeleteHooks []ReferralLetterTemplateHook
var referralLetterTemplateBeforeUpsertHooks []ReferralLetterTemplateHook

var referralLetterTemplateAfterInsertHooks []ReferralLetterTemplateHook
var referralLetterTemplateAfterSelectHooks []ReferralLetterTemplateHook
var referralLetterTemplateAfterUpdateHooks []ReferralLetterTemplateHook
var referralLetterTemplateAfterDeleteHooks []ReferralLetterTemplateHook
var referralLetterTemplateAfterUpsertHooks []ReferralLetterTemplateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ReferralLetterTemplate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ReferralLetterTemplate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ReferralLetterTemplate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ReferralLetterTemplate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ReferralLetterTemplate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ReferralLetterTemplate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ReferralLetterTemplate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ReferralLetterTemplate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ReferralLetterTemplate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range referralLetterTemplateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddReferralLetterTemplateHook registers your hook function for all future operations.
func AddReferralLetterTemplateHook(hookPoint boil.HookPoint, referralLetterTemplateHook ReferralLetterTemplateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		referralLetterTemplateBeforeInsertHooks = append(referralLetterTemplateBeforeInsertHooks, referralLetterTemplateHook)
	case boil.BeforeUpdateHook:
		referralLetterTemplateBeforeUpdateHooks = append(referralLetterTemplateBeforeUpdateHooks, referralLetterTemplateHook)
	case boil.BeforeDeleteHook:
		referralLetterTemplateBeforeDeleteHooks = append(referralLetterTemplateBeforeDeleteHooks, referralLetterTemplateHook)
	case boil.BeforeUpsertHook:
		referralLetterTemplateBeforeUpsertHooks = append(referralLetterTemplateBeforeUpsertHooks, referralLetterTemplateHook)
	case boil.AfterInsertHook:
		referralLetterTemplateAfterInsertHooks = append(referralLetterTemplateAfterInsertHooks, referralLetterTemplateHook)
	case boil.AfterSelectHook:
		referralLetterTemplateAfterSelectHooks = append(referralLetterTemplateAfterSelectHooks, referralLetterTemplateHook)
	case boil.AfterUpdateHook:
		referralLetterTemplateAfterUpdateHooks = append(referralLetterTemplateAfterUpdateHooks, referralLetterTemplateHook)
	case boil.AfterDeleteHook:
		referralLetterTemplateAfterDeleteHooks = append(referralLetterTemplateAfterDeleteHooks, referralLetterTemplateHook)
	case boil.AfterUpsertHook:
		referralLetterTemplateAfterUpsertHooks = append(referralLetterTemplateAfterUpsertHooks, referralLetterTemplateHook)
	}
}

// One returns a single referralLetterTemplate record from the query.
func (q referralLetterTemplateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ReferralLetterTemplate, error) {
	o := &ReferralLetterTemplate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ReferralLetterTemplate")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ReferralLetterTemplate records from the query.
func (q referralLetterTemplateQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReferralLetterTemplateSlice, error) {
	var o []*ReferralLetterTemplate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ReferralLetterTemplate slice")
	}

	if len(referralLetterTemplateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ReferralLetterTemplate records in the query.
func (q referralLetterTemplateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ReferralLetterTemplate rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q referralLetterTemplateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ReferralLetterTemplate exists")
	}

	return count > 0, nil
}

// ReferralLetterTemplates retrieves all the records using an executor.
func ReferralLetterTemplates(mods ...qm.QueryMod) referralLetterTemplateQuery {
	mods = append(mods, qm.From("[dbo].[ReferralLetterTemplate]"))
	return referralLetterTemplateQuery{NewQuery(mods...)}
}

// FindReferralLetterTemplate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReferralLetterTemplate(ctx context.Context, exec boil.ContextExecutor, templateID int, selectCols ...string) (*ReferralLetterTemplate, error) {
	referralLetterTemplateObj := &ReferralLetterTemplate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[ReferralLetterTemplate] where [TemplateID]=$1", sel,
	)

	q := queries.Raw(query, templateID)

	err := q.Bind(ctx, exec, referralLetterTemplateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ReferralLetterTemplate")
	}

	if err = referralLetterTemplateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return referralLetterTemplateObj, err
	}

	return referralLetterTemplateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ReferralLetterTemplate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ReferralLetterTemplate provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(referralLetterTemplateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	referralLetterTemplateInsertCacheMut.RLock()
	cache, cached := referralLetterTemplateInsertCache[key]
	referralLetterTemplateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			referralLetterTemplateAllColumns,
			referralLetterTemplateColumnsWithDefault,
			referralLetterTemplateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(referralLetterTemplateType, referralLetterTemplateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(referralLetterTemplateType, referralLetterTemplateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[ReferralLetterTemplate] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[ReferralLetterTemplate] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into ReferralLetterTemplate")
	}

	if !cached {
		referralLetterTemplateInsertCacheMut.Lock()
		referralLetterTemplateInsertCache[key] = cache
		referralLetterTemplateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ReferralLetterTemplate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ReferralLetterTemplate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	referralLetterTemplateUpdateCacheMut.RLock()
	cache, cached := referralLetterTemplateUpdateCache[key]
	referralLetterTemplateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			referralLetterTemplateAllColumns,
			referralLetterTemplatePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, referralLetterTemplateColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ReferralLetterTemplate, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[ReferralLetterTemplate] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, referralLetterTemplatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(referralLetterTemplateType, referralLetterTemplateMapping, append(wl, referralLetterTemplatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ReferralLetterTemplate row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ReferralLetterTemplate")
	}

	if !cached {
		referralLetterTemplateUpdateCacheMut.Lock()
		referralLetterTemplateUpdateCache[key] = cache
		referralLetterTemplateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q referralLetterTemplateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ReferralLetterTemplate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ReferralLetterTemplate")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReferralLetterTemplateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), referralLetterTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[ReferralLetterTemplate] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, referralLetterTemplatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in referralLetterTemplate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all referralLetterTemplate")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ReferralLetterTemplate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ReferralLetterTemplate provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(referralLetterTemplateColumnsWithDefault, o)

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

	referralLetterTemplateUpsertCacheMut.RLock()
	cache, cached := referralLetterTemplateUpsertCache[key]
	referralLetterTemplateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			referralLetterTemplateAllColumns,
			referralLetterTemplateColumnsWithDefault,
			referralLetterTemplateColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, referralLetterTemplateColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(referralLetterTemplatePrimaryKeyColumns, v) && strmangle.ContainsAny(referralLetterTemplateColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert ReferralLetterTemplate, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, referralLetterTemplateColumnsWithAuto)
		ret = strmangle.SetMerge(ret, referralLetterTemplateColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			referralLetterTemplateAllColumns,
			referralLetterTemplatePrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, referralLetterTemplateColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert ReferralLetterTemplate, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[ReferralLetterTemplate]", referralLetterTemplatePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(referralLetterTemplatePrimaryKeyColumns))
		copy(whitelist, referralLetterTemplatePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(referralLetterTemplateType, referralLetterTemplateMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(referralLetterTemplateType, referralLetterTemplateMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert ReferralLetterTemplate")
	}

	if !cached {
		referralLetterTemplateUpsertCacheMut.Lock()
		referralLetterTemplateUpsertCache[key] = cache
		referralLetterTemplateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ReferralLetterTemplate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReferralLetterTemplate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ReferralLetterTemplate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), referralLetterTemplatePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[ReferralLetterTemplate] WHERE [TemplateID]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ReferralLetterTemplate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ReferralLetterTemplate")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q referralLetterTemplateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no referralLetterTemplateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ReferralLetterTemplate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ReferralLetterTemplate")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReferralLetterTemplateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(referralLetterTemplateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), referralLetterTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[ReferralLetterTemplate] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, referralLetterTemplatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from referralLetterTemplate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ReferralLetterTemplate")
	}

	if len(referralLetterTemplateAfterDeleteHooks) != 0 {
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
func (o *ReferralLetterTemplate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReferralLetterTemplate(ctx, exec, o.TemplateID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReferralLetterTemplateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReferralLetterTemplateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), referralLetterTemplatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[ReferralLetterTemplate].* FROM [dbo].[ReferralLetterTemplate] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, referralLetterTemplatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReferralLetterTemplateSlice")
	}

	*o = slice

	return nil
}

// ReferralLetterTemplateExists checks if the ReferralLetterTemplate row exists.
func ReferralLetterTemplateExists(ctx context.Context, exec boil.ContextExecutor, templateID int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[ReferralLetterTemplate] where [TemplateID]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, templateID)
	}
	row := exec.QueryRowContext(ctx, sql, templateID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ReferralLetterTemplate exists")
	}

	return exists, nil
}
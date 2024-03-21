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

// ClinicalProcedure is an object representing the database table.
type ClinicalProcedure struct {
	ClinicCode  string      `boil:"ClinicCode" json:"ClinicCode" toml:"ClinicCode" yaml:"ClinicCode"`
	Description null.String `boil:"Description" json:"Description,omitempty" toml:"Description" yaml:"Description,omitempty"`

	R *clinicalProcedureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clinicalProcedureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClinicalProcedureColumns = struct {
	ClinicCode  string
	Description string
}{
	ClinicCode:  "ClinicCode",
	Description: "Description",
}

var ClinicalProcedureTableColumns = struct {
	ClinicCode  string
	Description string
}{
	ClinicCode:  "ClinicalProcedures.ClinicCode",
	Description: "ClinicalProcedures.Description",
}

// Generated where

var ClinicalProcedureWhere = struct {
	ClinicCode  whereHelperstring
	Description whereHelpernull_String
}{
	ClinicCode:  whereHelperstring{field: "[dbo].[ClinicalProcedures].[ClinicCode]"},
	Description: whereHelpernull_String{field: "[dbo].[ClinicalProcedures].[Description]"},
}

// ClinicalProcedureRels is where relationship names are stored.
var ClinicalProcedureRels = struct {
}{}

// clinicalProcedureR is where relationships are stored.
type clinicalProcedureR struct {
}

// NewStruct creates a new relationship struct
func (*clinicalProcedureR) NewStruct() *clinicalProcedureR {
	return &clinicalProcedureR{}
}

// clinicalProcedureL is where Load methods for each relationship are stored.
type clinicalProcedureL struct{}

var (
	clinicalProcedureAllColumns            = []string{"ClinicCode", "Description"}
	clinicalProcedureColumnsWithAuto       = []string{}
	clinicalProcedureColumnsWithoutDefault = []string{"ClinicCode", "Description"}
	clinicalProcedureColumnsWithDefault    = []string{}
	clinicalProcedurePrimaryKeyColumns     = []string{"ClinicCode"}
)

type (
	// ClinicalProcedureSlice is an alias for a slice of pointers to ClinicalProcedure.
	// This should almost always be used instead of []ClinicalProcedure.
	ClinicalProcedureSlice []*ClinicalProcedure
	// ClinicalProcedureHook is the signature for custom ClinicalProcedure hook methods
	ClinicalProcedureHook func(context.Context, boil.ContextExecutor, *ClinicalProcedure) error

	clinicalProcedureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clinicalProcedureType                 = reflect.TypeOf(&ClinicalProcedure{})
	clinicalProcedureMapping              = queries.MakeStructMapping(clinicalProcedureType)
	clinicalProcedurePrimaryKeyMapping, _ = queries.BindMapping(clinicalProcedureType, clinicalProcedureMapping, clinicalProcedurePrimaryKeyColumns)
	clinicalProcedureInsertCacheMut       sync.RWMutex
	clinicalProcedureInsertCache          = make(map[string]insertCache)
	clinicalProcedureUpdateCacheMut       sync.RWMutex
	clinicalProcedureUpdateCache          = make(map[string]updateCache)
	clinicalProcedureUpsertCacheMut       sync.RWMutex
	clinicalProcedureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clinicalProcedureBeforeInsertHooks []ClinicalProcedureHook
var clinicalProcedureBeforeUpdateHooks []ClinicalProcedureHook
var clinicalProcedureBeforeDeleteHooks []ClinicalProcedureHook
var clinicalProcedureBeforeUpsertHooks []ClinicalProcedureHook

var clinicalProcedureAfterInsertHooks []ClinicalProcedureHook
var clinicalProcedureAfterSelectHooks []ClinicalProcedureHook
var clinicalProcedureAfterUpdateHooks []ClinicalProcedureHook
var clinicalProcedureAfterDeleteHooks []ClinicalProcedureHook
var clinicalProcedureAfterUpsertHooks []ClinicalProcedureHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ClinicalProcedure) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ClinicalProcedure) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ClinicalProcedure) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ClinicalProcedure) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ClinicalProcedure) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ClinicalProcedure) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ClinicalProcedure) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ClinicalProcedure) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ClinicalProcedure) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clinicalProcedureAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClinicalProcedureHook registers your hook function for all future operations.
func AddClinicalProcedureHook(hookPoint boil.HookPoint, clinicalProcedureHook ClinicalProcedureHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		clinicalProcedureBeforeInsertHooks = append(clinicalProcedureBeforeInsertHooks, clinicalProcedureHook)
	case boil.BeforeUpdateHook:
		clinicalProcedureBeforeUpdateHooks = append(clinicalProcedureBeforeUpdateHooks, clinicalProcedureHook)
	case boil.BeforeDeleteHook:
		clinicalProcedureBeforeDeleteHooks = append(clinicalProcedureBeforeDeleteHooks, clinicalProcedureHook)
	case boil.BeforeUpsertHook:
		clinicalProcedureBeforeUpsertHooks = append(clinicalProcedureBeforeUpsertHooks, clinicalProcedureHook)
	case boil.AfterInsertHook:
		clinicalProcedureAfterInsertHooks = append(clinicalProcedureAfterInsertHooks, clinicalProcedureHook)
	case boil.AfterSelectHook:
		clinicalProcedureAfterSelectHooks = append(clinicalProcedureAfterSelectHooks, clinicalProcedureHook)
	case boil.AfterUpdateHook:
		clinicalProcedureAfterUpdateHooks = append(clinicalProcedureAfterUpdateHooks, clinicalProcedureHook)
	case boil.AfterDeleteHook:
		clinicalProcedureAfterDeleteHooks = append(clinicalProcedureAfterDeleteHooks, clinicalProcedureHook)
	case boil.AfterUpsertHook:
		clinicalProcedureAfterUpsertHooks = append(clinicalProcedureAfterUpsertHooks, clinicalProcedureHook)
	}
}

// One returns a single clinicalProcedure record from the query.
func (q clinicalProcedureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClinicalProcedure, error) {
	o := &ClinicalProcedure{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ClinicalProcedures")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ClinicalProcedure records from the query.
func (q clinicalProcedureQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClinicalProcedureSlice, error) {
	var o []*ClinicalProcedure

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClinicalProcedure slice")
	}

	if len(clinicalProcedureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ClinicalProcedure records in the query.
func (q clinicalProcedureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ClinicalProcedures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clinicalProcedureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ClinicalProcedures exists")
	}

	return count > 0, nil
}

// ClinicalProcedures retrieves all the records using an executor.
func ClinicalProcedures(mods ...qm.QueryMod) clinicalProcedureQuery {
	mods = append(mods, qm.From("[dbo].[ClinicalProcedures]"))
	return clinicalProcedureQuery{NewQuery(mods...)}
}

// FindClinicalProcedure retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClinicalProcedure(ctx context.Context, exec boil.ContextExecutor, clinicCode string, selectCols ...string) (*ClinicalProcedure, error) {
	clinicalProcedureObj := &ClinicalProcedure{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[ClinicalProcedures] where [ClinicCode]=$1", sel,
	)

	q := queries.Raw(query, clinicCode)

	err := q.Bind(ctx, exec, clinicalProcedureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ClinicalProcedures")
	}

	if err = clinicalProcedureObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clinicalProcedureObj, err
	}

	return clinicalProcedureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ClinicalProcedure) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ClinicalProcedures provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clinicalProcedureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clinicalProcedureInsertCacheMut.RLock()
	cache, cached := clinicalProcedureInsertCache[key]
	clinicalProcedureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clinicalProcedureAllColumns,
			clinicalProcedureColumnsWithDefault,
			clinicalProcedureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clinicalProcedureType, clinicalProcedureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clinicalProcedureType, clinicalProcedureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[ClinicalProcedures] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[ClinicalProcedures] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into ClinicalProcedures")
	}

	if !cached {
		clinicalProcedureInsertCacheMut.Lock()
		clinicalProcedureInsertCache[key] = cache
		clinicalProcedureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ClinicalProcedure.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ClinicalProcedure) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clinicalProcedureUpdateCacheMut.RLock()
	cache, cached := clinicalProcedureUpdateCache[key]
	clinicalProcedureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clinicalProcedureAllColumns,
			clinicalProcedurePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, clinicalProcedureColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ClinicalProcedures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[ClinicalProcedures] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, clinicalProcedurePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clinicalProcedureType, clinicalProcedureMapping, append(wl, clinicalProcedurePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ClinicalProcedures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ClinicalProcedures")
	}

	if !cached {
		clinicalProcedureUpdateCacheMut.Lock()
		clinicalProcedureUpdateCache[key] = cache
		clinicalProcedureUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q clinicalProcedureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ClinicalProcedures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ClinicalProcedures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClinicalProcedureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clinicalProcedurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[ClinicalProcedures] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, clinicalProcedurePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clinicalProcedure slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clinicalProcedure")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ClinicalProcedure) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ClinicalProcedures provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clinicalProcedureColumnsWithDefault, o)

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

	clinicalProcedureUpsertCacheMut.RLock()
	cache, cached := clinicalProcedureUpsertCache[key]
	clinicalProcedureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clinicalProcedureAllColumns,
			clinicalProcedureColumnsWithDefault,
			clinicalProcedureColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, clinicalProcedureColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(clinicalProcedurePrimaryKeyColumns, v) && strmangle.ContainsAny(clinicalProcedureColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert ClinicalProcedures, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, clinicalProcedureColumnsWithAuto)
		ret = strmangle.SetMerge(ret, clinicalProcedureColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			clinicalProcedureAllColumns,
			clinicalProcedurePrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, clinicalProcedureColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert ClinicalProcedures, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[ClinicalProcedures]", clinicalProcedurePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(clinicalProcedurePrimaryKeyColumns))
		copy(whitelist, clinicalProcedurePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(clinicalProcedureType, clinicalProcedureMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clinicalProcedureType, clinicalProcedureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert ClinicalProcedures")
	}

	if !cached {
		clinicalProcedureUpsertCacheMut.Lock()
		clinicalProcedureUpsertCache[key] = cache
		clinicalProcedureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ClinicalProcedure record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ClinicalProcedure) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClinicalProcedure provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clinicalProcedurePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[ClinicalProcedures] WHERE [ClinicCode]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ClinicalProcedures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ClinicalProcedures")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clinicalProcedureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clinicalProcedureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ClinicalProcedures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ClinicalProcedures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClinicalProcedureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clinicalProcedureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clinicalProcedurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[ClinicalProcedures] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clinicalProcedurePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clinicalProcedure slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ClinicalProcedures")
	}

	if len(clinicalProcedureAfterDeleteHooks) != 0 {
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
func (o *ClinicalProcedure) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClinicalProcedure(ctx, exec, o.ClinicCode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClinicalProcedureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClinicalProcedureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clinicalProcedurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[ClinicalProcedures].* FROM [dbo].[ClinicalProcedures] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clinicalProcedurePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClinicalProcedureSlice")
	}

	*o = slice

	return nil
}

// ClinicalProcedureExists checks if the ClinicalProcedure row exists.
func ClinicalProcedureExists(ctx context.Context, exec boil.ContextExecutor, clinicCode string) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[ClinicalProcedures] where [ClinicCode]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, clinicCode)
	}
	row := exec.QueryRowContext(ctx, sql, clinicCode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ClinicalProcedures exists")
	}

	return exists, nil
}
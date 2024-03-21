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

// Classification is an object representing the database table.
type Classification struct {
	ID          int         `boil:"ID" json:"ID" toml:"ID" yaml:"ID"`
	Description null.String `boil:"Description" json:"Description,omitempty" toml:"Description" yaml:"Description,omitempty"`

	R *classificationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L classificationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClassificationColumns = struct {
	ID          string
	Description string
}{
	ID:          "ID",
	Description: "Description",
}

var ClassificationTableColumns = struct {
	ID          string
	Description string
}{
	ID:          "Classification.ID",
	Description: "Classification.Description",
}

// Generated where

var ClassificationWhere = struct {
	ID          whereHelperint
	Description whereHelpernull_String
}{
	ID:          whereHelperint{field: "[dbo].[Classification].[ID]"},
	Description: whereHelpernull_String{field: "[dbo].[Classification].[Description]"},
}

// ClassificationRels is where relationship names are stored.
var ClassificationRels = struct {
}{}

// classificationR is where relationships are stored.
type classificationR struct {
}

// NewStruct creates a new relationship struct
func (*classificationR) NewStruct() *classificationR {
	return &classificationR{}
}

// classificationL is where Load methods for each relationship are stored.
type classificationL struct{}

var (
	classificationAllColumns            = []string{"ID", "Description"}
	classificationColumnsWithAuto       = []string{}
	classificationColumnsWithoutDefault = []string{"Description"}
	classificationColumnsWithDefault    = []string{"ID"}
	classificationPrimaryKeyColumns     = []string{"ID"}
)

type (
	// ClassificationSlice is an alias for a slice of pointers to Classification.
	// This should almost always be used instead of []Classification.
	ClassificationSlice []*Classification
	// ClassificationHook is the signature for custom Classification hook methods
	ClassificationHook func(context.Context, boil.ContextExecutor, *Classification) error

	classificationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	classificationType                 = reflect.TypeOf(&Classification{})
	classificationMapping              = queries.MakeStructMapping(classificationType)
	classificationPrimaryKeyMapping, _ = queries.BindMapping(classificationType, classificationMapping, classificationPrimaryKeyColumns)
	classificationInsertCacheMut       sync.RWMutex
	classificationInsertCache          = make(map[string]insertCache)
	classificationUpdateCacheMut       sync.RWMutex
	classificationUpdateCache          = make(map[string]updateCache)
	classificationUpsertCacheMut       sync.RWMutex
	classificationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var classificationBeforeInsertHooks []ClassificationHook
var classificationBeforeUpdateHooks []ClassificationHook
var classificationBeforeDeleteHooks []ClassificationHook
var classificationBeforeUpsertHooks []ClassificationHook

var classificationAfterInsertHooks []ClassificationHook
var classificationAfterSelectHooks []ClassificationHook
var classificationAfterUpdateHooks []ClassificationHook
var classificationAfterDeleteHooks []ClassificationHook
var classificationAfterUpsertHooks []ClassificationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Classification) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Classification) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Classification) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Classification) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Classification) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Classification) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Classification) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Classification) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Classification) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range classificationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClassificationHook registers your hook function for all future operations.
func AddClassificationHook(hookPoint boil.HookPoint, classificationHook ClassificationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		classificationBeforeInsertHooks = append(classificationBeforeInsertHooks, classificationHook)
	case boil.BeforeUpdateHook:
		classificationBeforeUpdateHooks = append(classificationBeforeUpdateHooks, classificationHook)
	case boil.BeforeDeleteHook:
		classificationBeforeDeleteHooks = append(classificationBeforeDeleteHooks, classificationHook)
	case boil.BeforeUpsertHook:
		classificationBeforeUpsertHooks = append(classificationBeforeUpsertHooks, classificationHook)
	case boil.AfterInsertHook:
		classificationAfterInsertHooks = append(classificationAfterInsertHooks, classificationHook)
	case boil.AfterSelectHook:
		classificationAfterSelectHooks = append(classificationAfterSelectHooks, classificationHook)
	case boil.AfterUpdateHook:
		classificationAfterUpdateHooks = append(classificationAfterUpdateHooks, classificationHook)
	case boil.AfterDeleteHook:
		classificationAfterDeleteHooks = append(classificationAfterDeleteHooks, classificationHook)
	case boil.AfterUpsertHook:
		classificationAfterUpsertHooks = append(classificationAfterUpsertHooks, classificationHook)
	}
}

// One returns a single classification record from the query.
func (q classificationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Classification, error) {
	o := &Classification{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Classification")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Classification records from the query.
func (q classificationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClassificationSlice, error) {
	var o []*Classification

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Classification slice")
	}

	if len(classificationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Classification records in the query.
func (q classificationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Classification rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q classificationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Classification exists")
	}

	return count > 0, nil
}

// Classifications retrieves all the records using an executor.
func Classifications(mods ...qm.QueryMod) classificationQuery {
	mods = append(mods, qm.From("[dbo].[Classification]"))
	return classificationQuery{NewQuery(mods...)}
}

// FindClassification retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClassification(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Classification, error) {
	classificationObj := &Classification{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[Classification] where [ID]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, classificationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Classification")
	}

	if err = classificationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return classificationObj, err
	}

	return classificationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Classification) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Classification provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(classificationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	classificationInsertCacheMut.RLock()
	cache, cached := classificationInsertCache[key]
	classificationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			classificationAllColumns,
			classificationColumnsWithDefault,
			classificationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(classificationType, classificationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(classificationType, classificationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[Classification] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[Classification] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into Classification")
	}

	if !cached {
		classificationInsertCacheMut.Lock()
		classificationInsertCache[key] = cache
		classificationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Classification.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Classification) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	classificationUpdateCacheMut.RLock()
	cache, cached := classificationUpdateCache[key]
	classificationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			classificationAllColumns,
			classificationPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, classificationColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Classification, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[Classification] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, classificationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(classificationType, classificationMapping, append(wl, classificationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Classification row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Classification")
	}

	if !cached {
		classificationUpdateCacheMut.Lock()
		classificationUpdateCache[key] = cache
		classificationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q classificationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Classification")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Classification")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClassificationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[Classification] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, classificationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in classification slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all classification")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Classification) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Classification provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(classificationColumnsWithDefault, o)

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

	classificationUpsertCacheMut.RLock()
	cache, cached := classificationUpsertCache[key]
	classificationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			classificationAllColumns,
			classificationColumnsWithDefault,
			classificationColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, classificationColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(classificationPrimaryKeyColumns, v) && strmangle.ContainsAny(classificationColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert Classification, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, classificationColumnsWithAuto)
		ret = strmangle.SetMerge(ret, classificationColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			classificationAllColumns,
			classificationPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, classificationColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert Classification, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[Classification]", classificationPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(classificationPrimaryKeyColumns))
		copy(whitelist, classificationPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(classificationType, classificationMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(classificationType, classificationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert Classification")
	}

	if !cached {
		classificationUpsertCacheMut.Lock()
		classificationUpsertCache[key] = cache
		classificationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Classification record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Classification) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Classification provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), classificationPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[Classification] WHERE [ID]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Classification")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Classification")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q classificationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no classificationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Classification")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Classification")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClassificationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(classificationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[Classification] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, classificationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from classification slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Classification")
	}

	if len(classificationAfterDeleteHooks) != 0 {
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
func (o *Classification) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClassification(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClassificationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClassificationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[Classification].* FROM [dbo].[Classification] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, classificationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClassificationSlice")
	}

	*o = slice

	return nil
}

// ClassificationExists checks if the Classification row exists.
func ClassificationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[Classification] where [ID]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Classification exists")
	}

	return exists, nil
}

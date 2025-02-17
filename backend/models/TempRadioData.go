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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TempRadioDatum is an object representing the database table.
type TempRadioDatum struct {
	CRegistrationNo string    `boil:"cRegistrationNo" json:"cRegistrationNo" toml:"cRegistrationNo" yaml:"cRegistrationNo"`
	VcPatientName   string    `boil:"vcPatientName" json:"vcPatientName" toml:"vcPatientName" yaml:"vcPatientName"`
	DtpDate         time.Time `boil:"dtpDate" json:"dtpDate" toml:"dtpDate" yaml:"dtpDate"`
	CReferenceNo    string    `boil:"cReferenceNo" json:"cReferenceNo" toml:"cReferenceNo" yaml:"cReferenceNo"`
	CItemCode       string    `boil:"cItemCode" json:"cItemCode" toml:"cItemCode" yaml:"cItemCode"`
	VcDescription   string    `boil:"vcDescription" json:"vcDescription" toml:"vcDescription" yaml:"vcDescription"`
	ID              int       `boil:"iD" json:"iD" toml:"iD" yaml:"iD"`

	R *tempRadioDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tempRadioDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TempRadioDatumColumns = struct {
	CRegistrationNo string
	VcPatientName   string
	DtpDate         string
	CReferenceNo    string
	CItemCode       string
	VcDescription   string
	ID              string
}{
	CRegistrationNo: "cRegistrationNo",
	VcPatientName:   "vcPatientName",
	DtpDate:         "dtpDate",
	CReferenceNo:    "cReferenceNo",
	CItemCode:       "cItemCode",
	VcDescription:   "vcDescription",
	ID:              "iD",
}

var TempRadioDatumTableColumns = struct {
	CRegistrationNo string
	VcPatientName   string
	DtpDate         string
	CReferenceNo    string
	CItemCode       string
	VcDescription   string
	ID              string
}{
	CRegistrationNo: "TempRadioData.cRegistrationNo",
	VcPatientName:   "TempRadioData.vcPatientName",
	DtpDate:         "TempRadioData.dtpDate",
	CReferenceNo:    "TempRadioData.cReferenceNo",
	CItemCode:       "TempRadioData.cItemCode",
	VcDescription:   "TempRadioData.vcDescription",
	ID:              "TempRadioData.iD",
}

// Generated where

var TempRadioDatumWhere = struct {
	CRegistrationNo whereHelperstring
	VcPatientName   whereHelperstring
	DtpDate         whereHelpertime_Time
	CReferenceNo    whereHelperstring
	CItemCode       whereHelperstring
	VcDescription   whereHelperstring
	ID              whereHelperint
}{
	CRegistrationNo: whereHelperstring{field: "[dbo].[TempRadioData].[cRegistrationNo]"},
	VcPatientName:   whereHelperstring{field: "[dbo].[TempRadioData].[vcPatientName]"},
	DtpDate:         whereHelpertime_Time{field: "[dbo].[TempRadioData].[dtpDate]"},
	CReferenceNo:    whereHelperstring{field: "[dbo].[TempRadioData].[cReferenceNo]"},
	CItemCode:       whereHelperstring{field: "[dbo].[TempRadioData].[cItemCode]"},
	VcDescription:   whereHelperstring{field: "[dbo].[TempRadioData].[vcDescription]"},
	ID:              whereHelperint{field: "[dbo].[TempRadioData].[iD]"},
}

// TempRadioDatumRels is where relationship names are stored.
var TempRadioDatumRels = struct {
}{}

// tempRadioDatumR is where relationships are stored.
type tempRadioDatumR struct {
}

// NewStruct creates a new relationship struct
func (*tempRadioDatumR) NewStruct() *tempRadioDatumR {
	return &tempRadioDatumR{}
}

// tempRadioDatumL is where Load methods for each relationship are stored.
type tempRadioDatumL struct{}

var (
	tempRadioDatumAllColumns            = []string{"cRegistrationNo", "vcPatientName", "dtpDate", "cReferenceNo", "cItemCode", "vcDescription", "iD"}
	tempRadioDatumColumnsWithAuto       = []string{}
	tempRadioDatumColumnsWithoutDefault = []string{}
	tempRadioDatumColumnsWithDefault    = []string{"cRegistrationNo", "vcPatientName", "dtpDate", "cReferenceNo", "cItemCode", "vcDescription", "iD"}
	tempRadioDatumPrimaryKeyColumns     = []string{"iD"}
)

type (
	// TempRadioDatumSlice is an alias for a slice of pointers to TempRadioDatum.
	// This should almost always be used instead of []TempRadioDatum.
	TempRadioDatumSlice []*TempRadioDatum
	// TempRadioDatumHook is the signature for custom TempRadioDatum hook methods
	TempRadioDatumHook func(context.Context, boil.ContextExecutor, *TempRadioDatum) error

	tempRadioDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tempRadioDatumType                 = reflect.TypeOf(&TempRadioDatum{})
	tempRadioDatumMapping              = queries.MakeStructMapping(tempRadioDatumType)
	tempRadioDatumPrimaryKeyMapping, _ = queries.BindMapping(tempRadioDatumType, tempRadioDatumMapping, tempRadioDatumPrimaryKeyColumns)
	tempRadioDatumInsertCacheMut       sync.RWMutex
	tempRadioDatumInsertCache          = make(map[string]insertCache)
	tempRadioDatumUpdateCacheMut       sync.RWMutex
	tempRadioDatumUpdateCache          = make(map[string]updateCache)
	tempRadioDatumUpsertCacheMut       sync.RWMutex
	tempRadioDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tempRadioDatumBeforeInsertHooks []TempRadioDatumHook
var tempRadioDatumBeforeUpdateHooks []TempRadioDatumHook
var tempRadioDatumBeforeDeleteHooks []TempRadioDatumHook
var tempRadioDatumBeforeUpsertHooks []TempRadioDatumHook

var tempRadioDatumAfterInsertHooks []TempRadioDatumHook
var tempRadioDatumAfterSelectHooks []TempRadioDatumHook
var tempRadioDatumAfterUpdateHooks []TempRadioDatumHook
var tempRadioDatumAfterDeleteHooks []TempRadioDatumHook
var tempRadioDatumAfterUpsertHooks []TempRadioDatumHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TempRadioDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TempRadioDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TempRadioDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TempRadioDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TempRadioDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TempRadioDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TempRadioDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TempRadioDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TempRadioDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tempRadioDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTempRadioDatumHook registers your hook function for all future operations.
func AddTempRadioDatumHook(hookPoint boil.HookPoint, tempRadioDatumHook TempRadioDatumHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tempRadioDatumBeforeInsertHooks = append(tempRadioDatumBeforeInsertHooks, tempRadioDatumHook)
	case boil.BeforeUpdateHook:
		tempRadioDatumBeforeUpdateHooks = append(tempRadioDatumBeforeUpdateHooks, tempRadioDatumHook)
	case boil.BeforeDeleteHook:
		tempRadioDatumBeforeDeleteHooks = append(tempRadioDatumBeforeDeleteHooks, tempRadioDatumHook)
	case boil.BeforeUpsertHook:
		tempRadioDatumBeforeUpsertHooks = append(tempRadioDatumBeforeUpsertHooks, tempRadioDatumHook)
	case boil.AfterInsertHook:
		tempRadioDatumAfterInsertHooks = append(tempRadioDatumAfterInsertHooks, tempRadioDatumHook)
	case boil.AfterSelectHook:
		tempRadioDatumAfterSelectHooks = append(tempRadioDatumAfterSelectHooks, tempRadioDatumHook)
	case boil.AfterUpdateHook:
		tempRadioDatumAfterUpdateHooks = append(tempRadioDatumAfterUpdateHooks, tempRadioDatumHook)
	case boil.AfterDeleteHook:
		tempRadioDatumAfterDeleteHooks = append(tempRadioDatumAfterDeleteHooks, tempRadioDatumHook)
	case boil.AfterUpsertHook:
		tempRadioDatumAfterUpsertHooks = append(tempRadioDatumAfterUpsertHooks, tempRadioDatumHook)
	}
}

// One returns a single tempRadioDatum record from the query.
func (q tempRadioDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TempRadioDatum, error) {
	o := &TempRadioDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TempRadioData")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TempRadioDatum records from the query.
func (q tempRadioDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (TempRadioDatumSlice, error) {
	var o []*TempRadioDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TempRadioDatum slice")
	}

	if len(tempRadioDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TempRadioDatum records in the query.
func (q tempRadioDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TempRadioData rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tempRadioDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TempRadioData exists")
	}

	return count > 0, nil
}

// TempRadioData retrieves all the records using an executor.
func TempRadioData(mods ...qm.QueryMod) tempRadioDatumQuery {
	mods = append(mods, qm.From("[dbo].[TempRadioData]"))
	return tempRadioDatumQuery{NewQuery(mods...)}
}

// FindTempRadioDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTempRadioDatum(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TempRadioDatum, error) {
	tempRadioDatumObj := &TempRadioDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[TempRadioData] where [iD]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tempRadioDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TempRadioData")
	}

	if err = tempRadioDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tempRadioDatumObj, err
	}

	return tempRadioDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TempRadioDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TempRadioData provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tempRadioDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tempRadioDatumInsertCacheMut.RLock()
	cache, cached := tempRadioDatumInsertCache[key]
	tempRadioDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tempRadioDatumAllColumns,
			tempRadioDatumColumnsWithDefault,
			tempRadioDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tempRadioDatumType, tempRadioDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tempRadioDatumType, tempRadioDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[TempRadioData] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[TempRadioData] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into TempRadioData")
	}

	if !cached {
		tempRadioDatumInsertCacheMut.Lock()
		tempRadioDatumInsertCache[key] = cache
		tempRadioDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TempRadioDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TempRadioDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tempRadioDatumUpdateCacheMut.RLock()
	cache, cached := tempRadioDatumUpdateCache[key]
	tempRadioDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tempRadioDatumAllColumns,
			tempRadioDatumPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, tempRadioDatumColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TempRadioData, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[TempRadioData] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, tempRadioDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tempRadioDatumType, tempRadioDatumMapping, append(wl, tempRadioDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update TempRadioData row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TempRadioData")
	}

	if !cached {
		tempRadioDatumUpdateCacheMut.Lock()
		tempRadioDatumUpdateCache[key] = cache
		tempRadioDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tempRadioDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TempRadioData")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TempRadioData")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TempRadioDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tempRadioDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[TempRadioData] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tempRadioDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tempRadioDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tempRadioDatum")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *TempRadioDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TempRadioData provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tempRadioDatumColumnsWithDefault, o)

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

	tempRadioDatumUpsertCacheMut.RLock()
	cache, cached := tempRadioDatumUpsertCache[key]
	tempRadioDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tempRadioDatumAllColumns,
			tempRadioDatumColumnsWithDefault,
			tempRadioDatumColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, tempRadioDatumColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(tempRadioDatumPrimaryKeyColumns, v) && strmangle.ContainsAny(tempRadioDatumColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert TempRadioData, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, tempRadioDatumColumnsWithAuto)
		ret = strmangle.SetMerge(ret, tempRadioDatumColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			tempRadioDatumAllColumns,
			tempRadioDatumPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, tempRadioDatumColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert TempRadioData, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[TempRadioData]", tempRadioDatumPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(tempRadioDatumPrimaryKeyColumns))
		copy(whitelist, tempRadioDatumPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(tempRadioDatumType, tempRadioDatumMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tempRadioDatumType, tempRadioDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert TempRadioData")
	}

	if !cached {
		tempRadioDatumUpsertCacheMut.Lock()
		tempRadioDatumUpsertCache[key] = cache
		tempRadioDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TempRadioDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TempRadioDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TempRadioDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tempRadioDatumPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[TempRadioData] WHERE [iD]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TempRadioData")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TempRadioData")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tempRadioDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tempRadioDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TempRadioData")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TempRadioData")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TempRadioDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tempRadioDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tempRadioDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[TempRadioData] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tempRadioDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tempRadioDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TempRadioData")
	}

	if len(tempRadioDatumAfterDeleteHooks) != 0 {
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
func (o *TempRadioDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTempRadioDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TempRadioDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TempRadioDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tempRadioDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[TempRadioData].* FROM [dbo].[TempRadioData] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tempRadioDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TempRadioDatumSlice")
	}

	*o = slice

	return nil
}

// TempRadioDatumExists checks if the TempRadioDatum row exists.
func TempRadioDatumExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[TempRadioData] where [iD]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TempRadioData exists")
	}

	return exists, nil
}

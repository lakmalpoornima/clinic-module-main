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

// FNClinicWisePatient is an object representing the database table.
type FNClinicWisePatient struct {
	ID             int         `boil:"Id" json:"Id" toml:"Id" yaml:"Id"`
	PatientName    null.String `boil:"PatientName" json:"PatientName,omitempty" toml:"PatientName" yaml:"PatientName,omitempty"`
	RegistrationNo null.String `boil:"RegistrationNo" json:"RegistrationNo,omitempty" toml:"RegistrationNo" yaml:"RegistrationNo,omitempty"`
	Category       null.String `boil:"Category" json:"Category,omitempty" toml:"Category" yaml:"Category,omitempty"`
	Gender         null.String `boil:"Gender" json:"Gender,omitempty" toml:"Gender" yaml:"Gender,omitempty"`
	Age            null.String `boil:"Age" json:"Age,omitempty" toml:"Age" yaml:"Age,omitempty"`
	ReferenceNo    null.String `boil:"ReferenceNo" json:"ReferenceNo,omitempty" toml:"ReferenceNo" yaml:"ReferenceNo,omitempty"`
	AppNo          null.Int    `boil:"AppNo" json:"AppNo,omitempty" toml:"AppNo" yaml:"AppNo,omitempty"`
	AppDate        null.Time   `boil:"AppDate" json:"AppDate,omitempty" toml:"AppDate" yaml:"AppDate,omitempty"`
	Clinic         null.String `boil:"Clinic" json:"Clinic,omitempty" toml:"Clinic" yaml:"Clinic,omitempty"`
	Address        null.String `boil:"Address" json:"Address,omitempty" toml:"Address" yaml:"Address,omitempty"`
	DOB            null.Time   `boil:"DOB" json:"DOB,omitempty" toml:"DOB" yaml:"DOB,omitempty"`

	R *fnClinicWisePatientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fnClinicWisePatientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FNClinicWisePatientColumns = struct {
	ID             string
	PatientName    string
	RegistrationNo string
	Category       string
	Gender         string
	Age            string
	ReferenceNo    string
	AppNo          string
	AppDate        string
	Clinic         string
	Address        string
	DOB            string
}{
	ID:             "Id",
	PatientName:    "PatientName",
	RegistrationNo: "RegistrationNo",
	Category:       "Category",
	Gender:         "Gender",
	Age:            "Age",
	ReferenceNo:    "ReferenceNo",
	AppNo:          "AppNo",
	AppDate:        "AppDate",
	Clinic:         "Clinic",
	Address:        "Address",
	DOB:            "DOB",
}

var FNClinicWisePatientTableColumns = struct {
	ID             string
	PatientName    string
	RegistrationNo string
	Category       string
	Gender         string
	Age            string
	ReferenceNo    string
	AppNo          string
	AppDate        string
	Clinic         string
	Address        string
	DOB            string
}{
	ID:             "fn_ClinicWisePatients.Id",
	PatientName:    "fn_ClinicWisePatients.PatientName",
	RegistrationNo: "fn_ClinicWisePatients.RegistrationNo",
	Category:       "fn_ClinicWisePatients.Category",
	Gender:         "fn_ClinicWisePatients.Gender",
	Age:            "fn_ClinicWisePatients.Age",
	ReferenceNo:    "fn_ClinicWisePatients.ReferenceNo",
	AppNo:          "fn_ClinicWisePatients.AppNo",
	AppDate:        "fn_ClinicWisePatients.AppDate",
	Clinic:         "fn_ClinicWisePatients.Clinic",
	Address:        "fn_ClinicWisePatients.Address",
	DOB:            "fn_ClinicWisePatients.DOB",
}

// Generated where

var FNClinicWisePatientWhere = struct {
	ID             whereHelperint
	PatientName    whereHelpernull_String
	RegistrationNo whereHelpernull_String
	Category       whereHelpernull_String
	Gender         whereHelpernull_String
	Age            whereHelpernull_String
	ReferenceNo    whereHelpernull_String
	AppNo          whereHelpernull_Int
	AppDate        whereHelpernull_Time
	Clinic         whereHelpernull_String
	Address        whereHelpernull_String
	DOB            whereHelpernull_Time
}{
	ID:             whereHelperint{field: "[dbo].[fn_ClinicWisePatients].[Id]"},
	PatientName:    whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[PatientName]"},
	RegistrationNo: whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[RegistrationNo]"},
	Category:       whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[Category]"},
	Gender:         whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[Gender]"},
	Age:            whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[Age]"},
	ReferenceNo:    whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[ReferenceNo]"},
	AppNo:          whereHelpernull_Int{field: "[dbo].[fn_ClinicWisePatients].[AppNo]"},
	AppDate:        whereHelpernull_Time{field: "[dbo].[fn_ClinicWisePatients].[AppDate]"},
	Clinic:         whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[Clinic]"},
	Address:        whereHelpernull_String{field: "[dbo].[fn_ClinicWisePatients].[Address]"},
	DOB:            whereHelpernull_Time{field: "[dbo].[fn_ClinicWisePatients].[DOB]"},
}

// FNClinicWisePatientRels is where relationship names are stored.
var FNClinicWisePatientRels = struct {
}{}

// fnClinicWisePatientR is where relationships are stored.
type fnClinicWisePatientR struct {
}

// NewStruct creates a new relationship struct
func (*fnClinicWisePatientR) NewStruct() *fnClinicWisePatientR {
	return &fnClinicWisePatientR{}
}

// fnClinicWisePatientL is where Load methods for each relationship are stored.
type fnClinicWisePatientL struct{}

var (
	fnClinicWisePatientAllColumns            = []string{"Id", "PatientName", "RegistrationNo", "Category", "Gender", "Age", "ReferenceNo", "AppNo", "AppDate", "Clinic", "Address", "DOB"}
	fnClinicWisePatientColumnsWithAuto       = []string{}
	fnClinicWisePatientColumnsWithoutDefault = []string{"Id", "PatientName", "RegistrationNo", "Category", "Gender", "Age", "ReferenceNo", "AppNo", "AppDate", "Clinic", "Address"}
	fnClinicWisePatientColumnsWithDefault    = []string{"DOB"}
	fnClinicWisePatientPrimaryKeyColumns     = []string{"Id"}
)

type (
	// FNClinicWisePatientSlice is an alias for a slice of pointers to FNClinicWisePatient.
	// This should almost always be used instead of []FNClinicWisePatient.
	FNClinicWisePatientSlice []*FNClinicWisePatient
	// FNClinicWisePatientHook is the signature for custom FNClinicWisePatient hook methods
	FNClinicWisePatientHook func(context.Context, boil.ContextExecutor, *FNClinicWisePatient) error

	fnClinicWisePatientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fnClinicWisePatientType                 = reflect.TypeOf(&FNClinicWisePatient{})
	fnClinicWisePatientMapping              = queries.MakeStructMapping(fnClinicWisePatientType)
	fnClinicWisePatientPrimaryKeyMapping, _ = queries.BindMapping(fnClinicWisePatientType, fnClinicWisePatientMapping, fnClinicWisePatientPrimaryKeyColumns)
	fnClinicWisePatientInsertCacheMut       sync.RWMutex
	fnClinicWisePatientInsertCache          = make(map[string]insertCache)
	fnClinicWisePatientUpdateCacheMut       sync.RWMutex
	fnClinicWisePatientUpdateCache          = make(map[string]updateCache)
	fnClinicWisePatientUpsertCacheMut       sync.RWMutex
	fnClinicWisePatientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fnClinicWisePatientBeforeInsertHooks []FNClinicWisePatientHook
var fnClinicWisePatientBeforeUpdateHooks []FNClinicWisePatientHook
var fnClinicWisePatientBeforeDeleteHooks []FNClinicWisePatientHook
var fnClinicWisePatientBeforeUpsertHooks []FNClinicWisePatientHook

var fnClinicWisePatientAfterInsertHooks []FNClinicWisePatientHook
var fnClinicWisePatientAfterSelectHooks []FNClinicWisePatientHook
var fnClinicWisePatientAfterUpdateHooks []FNClinicWisePatientHook
var fnClinicWisePatientAfterDeleteHooks []FNClinicWisePatientHook
var fnClinicWisePatientAfterUpsertHooks []FNClinicWisePatientHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FNClinicWisePatient) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FNClinicWisePatient) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FNClinicWisePatient) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FNClinicWisePatient) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FNClinicWisePatient) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FNClinicWisePatient) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FNClinicWisePatient) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FNClinicWisePatient) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FNClinicWisePatient) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fnClinicWisePatientAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFNClinicWisePatientHook registers your hook function for all future operations.
func AddFNClinicWisePatientHook(hookPoint boil.HookPoint, fnClinicWisePatientHook FNClinicWisePatientHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fnClinicWisePatientBeforeInsertHooks = append(fnClinicWisePatientBeforeInsertHooks, fnClinicWisePatientHook)
	case boil.BeforeUpdateHook:
		fnClinicWisePatientBeforeUpdateHooks = append(fnClinicWisePatientBeforeUpdateHooks, fnClinicWisePatientHook)
	case boil.BeforeDeleteHook:
		fnClinicWisePatientBeforeDeleteHooks = append(fnClinicWisePatientBeforeDeleteHooks, fnClinicWisePatientHook)
	case boil.BeforeUpsertHook:
		fnClinicWisePatientBeforeUpsertHooks = append(fnClinicWisePatientBeforeUpsertHooks, fnClinicWisePatientHook)
	case boil.AfterInsertHook:
		fnClinicWisePatientAfterInsertHooks = append(fnClinicWisePatientAfterInsertHooks, fnClinicWisePatientHook)
	case boil.AfterSelectHook:
		fnClinicWisePatientAfterSelectHooks = append(fnClinicWisePatientAfterSelectHooks, fnClinicWisePatientHook)
	case boil.AfterUpdateHook:
		fnClinicWisePatientAfterUpdateHooks = append(fnClinicWisePatientAfterUpdateHooks, fnClinicWisePatientHook)
	case boil.AfterDeleteHook:
		fnClinicWisePatientAfterDeleteHooks = append(fnClinicWisePatientAfterDeleteHooks, fnClinicWisePatientHook)
	case boil.AfterUpsertHook:
		fnClinicWisePatientAfterUpsertHooks = append(fnClinicWisePatientAfterUpsertHooks, fnClinicWisePatientHook)
	}
}

// One returns a single fnClinicWisePatient record from the query.
func (q fnClinicWisePatientQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FNClinicWisePatient, error) {
	o := &FNClinicWisePatient{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for fn_ClinicWisePatients")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FNClinicWisePatient records from the query.
func (q fnClinicWisePatientQuery) All(ctx context.Context, exec boil.ContextExecutor) (FNClinicWisePatientSlice, error) {
	var o []*FNClinicWisePatient

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FNClinicWisePatient slice")
	}

	if len(fnClinicWisePatientAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FNClinicWisePatient records in the query.
func (q fnClinicWisePatientQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count fn_ClinicWisePatients rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fnClinicWisePatientQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if fn_ClinicWisePatients exists")
	}

	return count > 0, nil
}

// FNClinicWisePatients retrieves all the records using an executor.
func FNClinicWisePatients(mods ...qm.QueryMod) fnClinicWisePatientQuery {
	mods = append(mods, qm.From("[dbo].[fn_ClinicWisePatients]"))
	return fnClinicWisePatientQuery{NewQuery(mods...)}
}

// FindFNClinicWisePatient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFNClinicWisePatient(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*FNClinicWisePatient, error) {
	fnClinicWisePatientObj := &FNClinicWisePatient{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[fn_ClinicWisePatients] where [Id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, fnClinicWisePatientObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from fn_ClinicWisePatients")
	}

	if err = fnClinicWisePatientObj.doAfterSelectHooks(ctx, exec); err != nil {
		return fnClinicWisePatientObj, err
	}

	return fnClinicWisePatientObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FNClinicWisePatient) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no fn_ClinicWisePatients provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fnClinicWisePatientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fnClinicWisePatientInsertCacheMut.RLock()
	cache, cached := fnClinicWisePatientInsertCache[key]
	fnClinicWisePatientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fnClinicWisePatientAllColumns,
			fnClinicWisePatientColumnsWithDefault,
			fnClinicWisePatientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fnClinicWisePatientType, fnClinicWisePatientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fnClinicWisePatientType, fnClinicWisePatientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[fn_ClinicWisePatients] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[fn_ClinicWisePatients] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into fn_ClinicWisePatients")
	}

	if !cached {
		fnClinicWisePatientInsertCacheMut.Lock()
		fnClinicWisePatientInsertCache[key] = cache
		fnClinicWisePatientInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FNClinicWisePatient.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FNClinicWisePatient) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fnClinicWisePatientUpdateCacheMut.RLock()
	cache, cached := fnClinicWisePatientUpdateCache[key]
	fnClinicWisePatientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fnClinicWisePatientAllColumns,
			fnClinicWisePatientPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, fnClinicWisePatientColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update fn_ClinicWisePatients, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[fn_ClinicWisePatients] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, fnClinicWisePatientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fnClinicWisePatientType, fnClinicWisePatientMapping, append(wl, fnClinicWisePatientPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update fn_ClinicWisePatients row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for fn_ClinicWisePatients")
	}

	if !cached {
		fnClinicWisePatientUpdateCacheMut.Lock()
		fnClinicWisePatientUpdateCache[key] = cache
		fnClinicWisePatientUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fnClinicWisePatientQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for fn_ClinicWisePatients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for fn_ClinicWisePatients")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FNClinicWisePatientSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fnClinicWisePatientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[fn_ClinicWisePatients] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, fnClinicWisePatientPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in fnClinicWisePatient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all fnClinicWisePatient")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FNClinicWisePatient) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no fn_ClinicWisePatients provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fnClinicWisePatientColumnsWithDefault, o)

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

	fnClinicWisePatientUpsertCacheMut.RLock()
	cache, cached := fnClinicWisePatientUpsertCache[key]
	fnClinicWisePatientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fnClinicWisePatientAllColumns,
			fnClinicWisePatientColumnsWithDefault,
			fnClinicWisePatientColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, fnClinicWisePatientColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(fnClinicWisePatientPrimaryKeyColumns, v) && strmangle.ContainsAny(fnClinicWisePatientColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert fn_ClinicWisePatients, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, fnClinicWisePatientColumnsWithAuto)
		ret = strmangle.SetMerge(ret, fnClinicWisePatientColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			fnClinicWisePatientAllColumns,
			fnClinicWisePatientPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, fnClinicWisePatientColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert fn_ClinicWisePatients, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[fn_ClinicWisePatients]", fnClinicWisePatientPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(fnClinicWisePatientPrimaryKeyColumns))
		copy(whitelist, fnClinicWisePatientPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(fnClinicWisePatientType, fnClinicWisePatientMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fnClinicWisePatientType, fnClinicWisePatientMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert fn_ClinicWisePatients")
	}

	if !cached {
		fnClinicWisePatientUpsertCacheMut.Lock()
		fnClinicWisePatientUpsertCache[key] = cache
		fnClinicWisePatientUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single FNClinicWisePatient record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FNClinicWisePatient) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FNClinicWisePatient provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fnClinicWisePatientPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[fn_ClinicWisePatients] WHERE [Id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from fn_ClinicWisePatients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for fn_ClinicWisePatients")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fnClinicWisePatientQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fnClinicWisePatientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fn_ClinicWisePatients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for fn_ClinicWisePatients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FNClinicWisePatientSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(fnClinicWisePatientBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fnClinicWisePatientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[fn_ClinicWisePatients] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fnClinicWisePatientPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fnClinicWisePatient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for fn_ClinicWisePatients")
	}

	if len(fnClinicWisePatientAfterDeleteHooks) != 0 {
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
func (o *FNClinicWisePatient) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFNClinicWisePatient(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FNClinicWisePatientSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FNClinicWisePatientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fnClinicWisePatientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[fn_ClinicWisePatients].* FROM [dbo].[fn_ClinicWisePatients] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fnClinicWisePatientPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FNClinicWisePatientSlice")
	}

	*o = slice

	return nil
}

// FNClinicWisePatientExists checks if the FNClinicWisePatient row exists.
func FNClinicWisePatientExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[fn_ClinicWisePatients] where [Id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if fn_ClinicWisePatients exists")
	}

	return exists, nil
}

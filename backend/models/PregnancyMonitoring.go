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

// PregnancyMonitoring is an object representing the database table.
type PregnancyMonitoring struct {
	ID                  int         `boil:"ID" json:"ID" toml:"ID" yaml:"ID"`
	DocumentID          null.String `boil:"DocumentID" json:"DocumentID,omitempty" toml:"DocumentID" yaml:"DocumentID,omitempty"`
	RegistrationID      null.String `boil:"RegistrationID" json:"RegistrationID,omitempty" toml:"RegistrationID" yaml:"RegistrationID,omitempty"`
	BreakfirstRegime    null.String `boil:"BreakfirstRegime" json:"BreakfirstRegime,omitempty" toml:"BreakfirstRegime" yaml:"BreakfirstRegime,omitempty"`
	BreakfirstBefore    null.String `boil:"BreakfirstBefore" json:"BreakfirstBefore,omitempty" toml:"BreakfirstBefore" yaml:"BreakfirstBefore,omitempty"`
	BreakfirstAfter     null.String `boil:"BreakfirstAfter" json:"BreakfirstAfter,omitempty" toml:"BreakfirstAfter" yaml:"BreakfirstAfter,omitempty"`
	BreakfirstAdgRegime null.String `boil:"BreakfirstAdgRegime" json:"BreakfirstAdgRegime,omitempty" toml:"BreakfirstAdgRegime" yaml:"BreakfirstAdgRegime,omitempty"`
	LunchRegime         null.String `boil:"LunchRegime" json:"LunchRegime,omitempty" toml:"LunchRegime" yaml:"LunchRegime,omitempty"`
	LunchBefore         null.String `boil:"LunchBefore" json:"LunchBefore,omitempty" toml:"LunchBefore" yaml:"LunchBefore,omitempty"`
	LunchAfter          null.String `boil:"LunchAfter" json:"LunchAfter,omitempty" toml:"LunchAfter" yaml:"LunchAfter,omitempty"`
	LunchAdgRegime      null.String `boil:"LunchAdgRegime" json:"LunchAdgRegime,omitempty" toml:"LunchAdgRegime" yaml:"LunchAdgRegime,omitempty"`
	DinnerRegime        null.String `boil:"DinnerRegime" json:"DinnerRegime,omitempty" toml:"DinnerRegime" yaml:"DinnerRegime,omitempty"`
	DinnerBefore        null.String `boil:"DinnerBefore" json:"DinnerBefore,omitempty" toml:"DinnerBefore" yaml:"DinnerBefore,omitempty"`
	DinnerAfter         null.String `boil:"DinnerAfter" json:"DinnerAfter,omitempty" toml:"DinnerAfter" yaml:"DinnerAfter,omitempty"`
	DinnerAdgRegime     null.String `boil:"DinnerAdgRegime" json:"DinnerAdgRegime,omitempty" toml:"DinnerAdgRegime" yaml:"DinnerAdgRegime,omitempty"`
	Note                null.String `boil:"Note" json:"Note,omitempty" toml:"Note" yaml:"Note,omitempty"`
	Date                null.Time   `boil:"Date" json:"Date,omitempty" toml:"Date" yaml:"Date,omitempty"`
	POA                 null.String `boil:"POA" json:"POA,omitempty" toml:"POA" yaml:"POA,omitempty"`
	PregnancyID         null.Int    `boil:"PregnancyID" json:"PregnancyID,omitempty" toml:"PregnancyID" yaml:"PregnancyID,omitempty"`

	R *pregnancyMonitoringR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pregnancyMonitoringL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PregnancyMonitoringColumns = struct {
	ID                  string
	DocumentID          string
	RegistrationID      string
	BreakfirstRegime    string
	BreakfirstBefore    string
	BreakfirstAfter     string
	BreakfirstAdgRegime string
	LunchRegime         string
	LunchBefore         string
	LunchAfter          string
	LunchAdgRegime      string
	DinnerRegime        string
	DinnerBefore        string
	DinnerAfter         string
	DinnerAdgRegime     string
	Note                string
	Date                string
	POA                 string
	PregnancyID         string
}{
	ID:                  "ID",
	DocumentID:          "DocumentID",
	RegistrationID:      "RegistrationID",
	BreakfirstRegime:    "BreakfirstRegime",
	BreakfirstBefore:    "BreakfirstBefore",
	BreakfirstAfter:     "BreakfirstAfter",
	BreakfirstAdgRegime: "BreakfirstAdgRegime",
	LunchRegime:         "LunchRegime",
	LunchBefore:         "LunchBefore",
	LunchAfter:          "LunchAfter",
	LunchAdgRegime:      "LunchAdgRegime",
	DinnerRegime:        "DinnerRegime",
	DinnerBefore:        "DinnerBefore",
	DinnerAfter:         "DinnerAfter",
	DinnerAdgRegime:     "DinnerAdgRegime",
	Note:                "Note",
	Date:                "Date",
	POA:                 "POA",
	PregnancyID:         "PregnancyID",
}

var PregnancyMonitoringTableColumns = struct {
	ID                  string
	DocumentID          string
	RegistrationID      string
	BreakfirstRegime    string
	BreakfirstBefore    string
	BreakfirstAfter     string
	BreakfirstAdgRegime string
	LunchRegime         string
	LunchBefore         string
	LunchAfter          string
	LunchAdgRegime      string
	DinnerRegime        string
	DinnerBefore        string
	DinnerAfter         string
	DinnerAdgRegime     string
	Note                string
	Date                string
	POA                 string
	PregnancyID         string
}{
	ID:                  "PregnancyMonitoring.ID",
	DocumentID:          "PregnancyMonitoring.DocumentID",
	RegistrationID:      "PregnancyMonitoring.RegistrationID",
	BreakfirstRegime:    "PregnancyMonitoring.BreakfirstRegime",
	BreakfirstBefore:    "PregnancyMonitoring.BreakfirstBefore",
	BreakfirstAfter:     "PregnancyMonitoring.BreakfirstAfter",
	BreakfirstAdgRegime: "PregnancyMonitoring.BreakfirstAdgRegime",
	LunchRegime:         "PregnancyMonitoring.LunchRegime",
	LunchBefore:         "PregnancyMonitoring.LunchBefore",
	LunchAfter:          "PregnancyMonitoring.LunchAfter",
	LunchAdgRegime:      "PregnancyMonitoring.LunchAdgRegime",
	DinnerRegime:        "PregnancyMonitoring.DinnerRegime",
	DinnerBefore:        "PregnancyMonitoring.DinnerBefore",
	DinnerAfter:         "PregnancyMonitoring.DinnerAfter",
	DinnerAdgRegime:     "PregnancyMonitoring.DinnerAdgRegime",
	Note:                "PregnancyMonitoring.Note",
	Date:                "PregnancyMonitoring.Date",
	POA:                 "PregnancyMonitoring.POA",
	PregnancyID:         "PregnancyMonitoring.PregnancyID",
}

// Generated where

var PregnancyMonitoringWhere = struct {
	ID                  whereHelperint
	DocumentID          whereHelpernull_String
	RegistrationID      whereHelpernull_String
	BreakfirstRegime    whereHelpernull_String
	BreakfirstBefore    whereHelpernull_String
	BreakfirstAfter     whereHelpernull_String
	BreakfirstAdgRegime whereHelpernull_String
	LunchRegime         whereHelpernull_String
	LunchBefore         whereHelpernull_String
	LunchAfter          whereHelpernull_String
	LunchAdgRegime      whereHelpernull_String
	DinnerRegime        whereHelpernull_String
	DinnerBefore        whereHelpernull_String
	DinnerAfter         whereHelpernull_String
	DinnerAdgRegime     whereHelpernull_String
	Note                whereHelpernull_String
	Date                whereHelpernull_Time
	POA                 whereHelpernull_String
	PregnancyID         whereHelpernull_Int
}{
	ID:                  whereHelperint{field: "[dbo].[PregnancyMonitoring].[ID]"},
	DocumentID:          whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[DocumentID]"},
	RegistrationID:      whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[RegistrationID]"},
	BreakfirstRegime:    whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[BreakfirstRegime]"},
	BreakfirstBefore:    whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[BreakfirstBefore]"},
	BreakfirstAfter:     whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[BreakfirstAfter]"},
	BreakfirstAdgRegime: whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[BreakfirstAdgRegime]"},
	LunchRegime:         whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[LunchRegime]"},
	LunchBefore:         whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[LunchBefore]"},
	LunchAfter:          whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[LunchAfter]"},
	LunchAdgRegime:      whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[LunchAdgRegime]"},
	DinnerRegime:        whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[DinnerRegime]"},
	DinnerBefore:        whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[DinnerBefore]"},
	DinnerAfter:         whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[DinnerAfter]"},
	DinnerAdgRegime:     whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[DinnerAdgRegime]"},
	Note:                whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[Note]"},
	Date:                whereHelpernull_Time{field: "[dbo].[PregnancyMonitoring].[Date]"},
	POA:                 whereHelpernull_String{field: "[dbo].[PregnancyMonitoring].[POA]"},
	PregnancyID:         whereHelpernull_Int{field: "[dbo].[PregnancyMonitoring].[PregnancyID]"},
}

// PregnancyMonitoringRels is where relationship names are stored.
var PregnancyMonitoringRels = struct {
}{}

// pregnancyMonitoringR is where relationships are stored.
type pregnancyMonitoringR struct {
}

// NewStruct creates a new relationship struct
func (*pregnancyMonitoringR) NewStruct() *pregnancyMonitoringR {
	return &pregnancyMonitoringR{}
}

// pregnancyMonitoringL is where Load methods for each relationship are stored.
type pregnancyMonitoringL struct{}

var (
	pregnancyMonitoringAllColumns            = []string{"ID", "DocumentID", "RegistrationID", "BreakfirstRegime", "BreakfirstBefore", "BreakfirstAfter", "BreakfirstAdgRegime", "LunchRegime", "LunchBefore", "LunchAfter", "LunchAdgRegime", "DinnerRegime", "DinnerBefore", "DinnerAfter", "DinnerAdgRegime", "Note", "Date", "POA", "PregnancyID"}
	pregnancyMonitoringColumnsWithAuto       = []string{}
	pregnancyMonitoringColumnsWithoutDefault = []string{"DocumentID", "RegistrationID", "BreakfirstRegime", "BreakfirstBefore", "BreakfirstAfter", "BreakfirstAdgRegime", "LunchRegime", "LunchBefore", "LunchAfter", "LunchAdgRegime", "DinnerRegime", "DinnerBefore", "DinnerAfter", "DinnerAdgRegime", "Note", "Date", "POA", "PregnancyID"}
	pregnancyMonitoringColumnsWithDefault    = []string{"ID"}
	pregnancyMonitoringPrimaryKeyColumns     = []string{"ID"}
)

type (
	// PregnancyMonitoringSlice is an alias for a slice of pointers to PregnancyMonitoring.
	// This should almost always be used instead of []PregnancyMonitoring.
	PregnancyMonitoringSlice []*PregnancyMonitoring
	// PregnancyMonitoringHook is the signature for custom PregnancyMonitoring hook methods
	PregnancyMonitoringHook func(context.Context, boil.ContextExecutor, *PregnancyMonitoring) error

	pregnancyMonitoringQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pregnancyMonitoringType                 = reflect.TypeOf(&PregnancyMonitoring{})
	pregnancyMonitoringMapping              = queries.MakeStructMapping(pregnancyMonitoringType)
	pregnancyMonitoringPrimaryKeyMapping, _ = queries.BindMapping(pregnancyMonitoringType, pregnancyMonitoringMapping, pregnancyMonitoringPrimaryKeyColumns)
	pregnancyMonitoringInsertCacheMut       sync.RWMutex
	pregnancyMonitoringInsertCache          = make(map[string]insertCache)
	pregnancyMonitoringUpdateCacheMut       sync.RWMutex
	pregnancyMonitoringUpdateCache          = make(map[string]updateCache)
	pregnancyMonitoringUpsertCacheMut       sync.RWMutex
	pregnancyMonitoringUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var pregnancyMonitoringBeforeInsertHooks []PregnancyMonitoringHook
var pregnancyMonitoringBeforeUpdateHooks []PregnancyMonitoringHook
var pregnancyMonitoringBeforeDeleteHooks []PregnancyMonitoringHook
var pregnancyMonitoringBeforeUpsertHooks []PregnancyMonitoringHook

var pregnancyMonitoringAfterInsertHooks []PregnancyMonitoringHook
var pregnancyMonitoringAfterSelectHooks []PregnancyMonitoringHook
var pregnancyMonitoringAfterUpdateHooks []PregnancyMonitoringHook
var pregnancyMonitoringAfterDeleteHooks []PregnancyMonitoringHook
var pregnancyMonitoringAfterUpsertHooks []PregnancyMonitoringHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PregnancyMonitoring) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PregnancyMonitoring) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PregnancyMonitoring) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PregnancyMonitoring) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PregnancyMonitoring) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PregnancyMonitoring) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PregnancyMonitoring) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PregnancyMonitoring) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PregnancyMonitoring) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pregnancyMonitoringAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPregnancyMonitoringHook registers your hook function for all future operations.
func AddPregnancyMonitoringHook(hookPoint boil.HookPoint, pregnancyMonitoringHook PregnancyMonitoringHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pregnancyMonitoringBeforeInsertHooks = append(pregnancyMonitoringBeforeInsertHooks, pregnancyMonitoringHook)
	case boil.BeforeUpdateHook:
		pregnancyMonitoringBeforeUpdateHooks = append(pregnancyMonitoringBeforeUpdateHooks, pregnancyMonitoringHook)
	case boil.BeforeDeleteHook:
		pregnancyMonitoringBeforeDeleteHooks = append(pregnancyMonitoringBeforeDeleteHooks, pregnancyMonitoringHook)
	case boil.BeforeUpsertHook:
		pregnancyMonitoringBeforeUpsertHooks = append(pregnancyMonitoringBeforeUpsertHooks, pregnancyMonitoringHook)
	case boil.AfterInsertHook:
		pregnancyMonitoringAfterInsertHooks = append(pregnancyMonitoringAfterInsertHooks, pregnancyMonitoringHook)
	case boil.AfterSelectHook:
		pregnancyMonitoringAfterSelectHooks = append(pregnancyMonitoringAfterSelectHooks, pregnancyMonitoringHook)
	case boil.AfterUpdateHook:
		pregnancyMonitoringAfterUpdateHooks = append(pregnancyMonitoringAfterUpdateHooks, pregnancyMonitoringHook)
	case boil.AfterDeleteHook:
		pregnancyMonitoringAfterDeleteHooks = append(pregnancyMonitoringAfterDeleteHooks, pregnancyMonitoringHook)
	case boil.AfterUpsertHook:
		pregnancyMonitoringAfterUpsertHooks = append(pregnancyMonitoringAfterUpsertHooks, pregnancyMonitoringHook)
	}
}

// One returns a single pregnancyMonitoring record from the query.
func (q pregnancyMonitoringQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PregnancyMonitoring, error) {
	o := &PregnancyMonitoring{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for PregnancyMonitoring")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PregnancyMonitoring records from the query.
func (q pregnancyMonitoringQuery) All(ctx context.Context, exec boil.ContextExecutor) (PregnancyMonitoringSlice, error) {
	var o []*PregnancyMonitoring

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PregnancyMonitoring slice")
	}

	if len(pregnancyMonitoringAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PregnancyMonitoring records in the query.
func (q pregnancyMonitoringQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count PregnancyMonitoring rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pregnancyMonitoringQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if PregnancyMonitoring exists")
	}

	return count > 0, nil
}

// PregnancyMonitorings retrieves all the records using an executor.
func PregnancyMonitorings(mods ...qm.QueryMod) pregnancyMonitoringQuery {
	mods = append(mods, qm.From("[dbo].[PregnancyMonitoring]"))
	return pregnancyMonitoringQuery{NewQuery(mods...)}
}

// FindPregnancyMonitoring retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPregnancyMonitoring(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*PregnancyMonitoring, error) {
	pregnancyMonitoringObj := &PregnancyMonitoring{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[PregnancyMonitoring] where [ID]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pregnancyMonitoringObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from PregnancyMonitoring")
	}

	if err = pregnancyMonitoringObj.doAfterSelectHooks(ctx, exec); err != nil {
		return pregnancyMonitoringObj, err
	}

	return pregnancyMonitoringObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PregnancyMonitoring) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PregnancyMonitoring provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pregnancyMonitoringColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pregnancyMonitoringInsertCacheMut.RLock()
	cache, cached := pregnancyMonitoringInsertCache[key]
	pregnancyMonitoringInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pregnancyMonitoringAllColumns,
			pregnancyMonitoringColumnsWithDefault,
			pregnancyMonitoringColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pregnancyMonitoringType, pregnancyMonitoringMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pregnancyMonitoringType, pregnancyMonitoringMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[PregnancyMonitoring] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[PregnancyMonitoring] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into PregnancyMonitoring")
	}

	if !cached {
		pregnancyMonitoringInsertCacheMut.Lock()
		pregnancyMonitoringInsertCache[key] = cache
		pregnancyMonitoringInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PregnancyMonitoring.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PregnancyMonitoring) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pregnancyMonitoringUpdateCacheMut.RLock()
	cache, cached := pregnancyMonitoringUpdateCache[key]
	pregnancyMonitoringUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pregnancyMonitoringAllColumns,
			pregnancyMonitoringPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, pregnancyMonitoringColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update PregnancyMonitoring, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[PregnancyMonitoring] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, pregnancyMonitoringPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pregnancyMonitoringType, pregnancyMonitoringMapping, append(wl, pregnancyMonitoringPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update PregnancyMonitoring row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for PregnancyMonitoring")
	}

	if !cached {
		pregnancyMonitoringUpdateCacheMut.Lock()
		pregnancyMonitoringUpdateCache[key] = cache
		pregnancyMonitoringUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pregnancyMonitoringQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for PregnancyMonitoring")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for PregnancyMonitoring")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PregnancyMonitoringSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pregnancyMonitoringPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[PregnancyMonitoring] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pregnancyMonitoringPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pregnancyMonitoring slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pregnancyMonitoring")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *PregnancyMonitoring) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PregnancyMonitoring provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pregnancyMonitoringColumnsWithDefault, o)

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

	pregnancyMonitoringUpsertCacheMut.RLock()
	cache, cached := pregnancyMonitoringUpsertCache[key]
	pregnancyMonitoringUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pregnancyMonitoringAllColumns,
			pregnancyMonitoringColumnsWithDefault,
			pregnancyMonitoringColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, pregnancyMonitoringColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(pregnancyMonitoringPrimaryKeyColumns, v) && strmangle.ContainsAny(pregnancyMonitoringColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert PregnancyMonitoring, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, pregnancyMonitoringColumnsWithAuto)
		ret = strmangle.SetMerge(ret, pregnancyMonitoringColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			pregnancyMonitoringAllColumns,
			pregnancyMonitoringPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, pregnancyMonitoringColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert PregnancyMonitoring, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[PregnancyMonitoring]", pregnancyMonitoringPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(pregnancyMonitoringPrimaryKeyColumns))
		copy(whitelist, pregnancyMonitoringPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(pregnancyMonitoringType, pregnancyMonitoringMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pregnancyMonitoringType, pregnancyMonitoringMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert PregnancyMonitoring")
	}

	if !cached {
		pregnancyMonitoringUpsertCacheMut.Lock()
		pregnancyMonitoringUpsertCache[key] = cache
		pregnancyMonitoringUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PregnancyMonitoring record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PregnancyMonitoring) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PregnancyMonitoring provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pregnancyMonitoringPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[PregnancyMonitoring] WHERE [ID]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from PregnancyMonitoring")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for PregnancyMonitoring")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pregnancyMonitoringQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pregnancyMonitoringQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from PregnancyMonitoring")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PregnancyMonitoring")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PregnancyMonitoringSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(pregnancyMonitoringBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pregnancyMonitoringPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[PregnancyMonitoring] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pregnancyMonitoringPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pregnancyMonitoring slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PregnancyMonitoring")
	}

	if len(pregnancyMonitoringAfterDeleteHooks) != 0 {
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
func (o *PregnancyMonitoring) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPregnancyMonitoring(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PregnancyMonitoringSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PregnancyMonitoringSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pregnancyMonitoringPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[PregnancyMonitoring].* FROM [dbo].[PregnancyMonitoring] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pregnancyMonitoringPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PregnancyMonitoringSlice")
	}

	*o = slice

	return nil
}

// PregnancyMonitoringExists checks if the PregnancyMonitoring row exists.
func PregnancyMonitoringExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[PregnancyMonitoring] where [ID]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if PregnancyMonitoring exists")
	}

	return exists, nil
}

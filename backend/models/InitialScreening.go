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

// InitialScreening is an object representing the database table.
type InitialScreening struct {
	PresentingComplaint string      `boil:"PresentingComplaint" json:"PresentingComplaint" toml:"PresentingComplaint" yaml:"PresentingComplaint"`
	PastMedicalHx       string      `boil:"PastMedicalHx" json:"PastMedicalHx" toml:"PastMedicalHx" yaml:"PastMedicalHx"`
	ObstreticHx         string      `boil:"ObstreticHx" json:"ObstreticHx" toml:"ObstreticHx" yaml:"ObstreticHx"`
	MedicationSummary   string      `boil:"MedicationSummary" json:"MedicationSummary" toml:"MedicationSummary" yaml:"MedicationSummary"`
	FamilyHistory       string      `boil:"FamilyHistory" json:"FamilyHistory" toml:"FamilyHistory" yaml:"FamilyHistory"`
	AllegicHx           string      `boil:"AllegicHx" json:"AllegicHx" toml:"AllegicHx" yaml:"AllegicHx"`
	SocialHx            string      `boil:"SocialHx" json:"SocialHx" toml:"SocialHx" yaml:"SocialHx"`
	Occupation          string      `boil:"Occupation" json:"Occupation" toml:"Occupation" yaml:"Occupation"`
	Smoking             string      `boil:"Smoking" json:"Smoking" toml:"Smoking" yaml:"Smoking"`
	Alcohole            string      `boil:"Alcohole" json:"Alcohole" toml:"Alcohole" yaml:"Alcohole"`
	GeneralExamination  string      `boil:"GeneralExamination" json:"GeneralExamination" toml:"GeneralExamination" yaml:"GeneralExamination"`
	FGS                 string      `boil:"FGS" json:"FGS" toml:"FGS" yaml:"FGS"`
	FGSImage            null.String `boil:"FGSImage" json:"FGSImage,omitempty" toml:"FGSImage" yaml:"FGSImage,omitempty"`
	ThyroidExam         string      `boil:"ThyroidExam" json:"ThyroidExam" toml:"ThyroidExam" yaml:"ThyroidExam"`
	Eyes                string      `boil:"Eyes" json:"Eyes" toml:"Eyes" yaml:"Eyes"`
	RS                  string      `boil:"RS" json:"RS" toml:"RS" yaml:"RS"`
	ABD                 string      `boil:"ABD" json:"ABD" toml:"ABD" yaml:"ABD"`
	CNS                 string      `boil:"CNS" json:"CNS" toml:"CNS" yaml:"CNS"`
	CVS                 string      `boil:"CVS" json:"CVS" toml:"CVS" yaml:"CVS"`
	ClinicType          null.String `boil:"ClinicType" json:"ClinicType,omitempty" toml:"ClinicType" yaml:"ClinicType,omitempty"`
	RegistrationID      string      `boil:"RegistrationID" json:"RegistrationID" toml:"RegistrationID" yaml:"RegistrationID"`
	ID                  int         `boil:"ID" json:"ID" toml:"ID" yaml:"ID"`
	CreatedUser         string      `boil:"CreatedUser" json:"CreatedUser" toml:"CreatedUser" yaml:"CreatedUser"`
	CreatedDate         time.Time   `boil:"CreatedDate" json:"CreatedDate" toml:"CreatedDate" yaml:"CreatedDate"`

	R *initialScreeningR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L initialScreeningL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InitialScreeningColumns = struct {
	PresentingComplaint string
	PastMedicalHx       string
	ObstreticHx         string
	MedicationSummary   string
	FamilyHistory       string
	AllegicHx           string
	SocialHx            string
	Occupation          string
	Smoking             string
	Alcohole            string
	GeneralExamination  string
	FGS                 string
	FGSImage            string
	ThyroidExam         string
	Eyes                string
	RS                  string
	ABD                 string
	CNS                 string
	CVS                 string
	ClinicType          string
	RegistrationID      string
	ID                  string
	CreatedUser         string
	CreatedDate         string
}{
	PresentingComplaint: "PresentingComplaint",
	PastMedicalHx:       "PastMedicalHx",
	ObstreticHx:         "ObstreticHx",
	MedicationSummary:   "MedicationSummary",
	FamilyHistory:       "FamilyHistory",
	AllegicHx:           "AllegicHx",
	SocialHx:            "SocialHx",
	Occupation:          "Occupation",
	Smoking:             "Smoking",
	Alcohole:            "Alcohole",
	GeneralExamination:  "GeneralExamination",
	FGS:                 "FGS",
	FGSImage:            "FGSImage",
	ThyroidExam:         "ThyroidExam",
	Eyes:                "Eyes",
	RS:                  "RS",
	ABD:                 "ABD",
	CNS:                 "CNS",
	CVS:                 "CVS",
	ClinicType:          "ClinicType",
	RegistrationID:      "RegistrationID",
	ID:                  "ID",
	CreatedUser:         "CreatedUser",
	CreatedDate:         "CreatedDate",
}

var InitialScreeningTableColumns = struct {
	PresentingComplaint string
	PastMedicalHx       string
	ObstreticHx         string
	MedicationSummary   string
	FamilyHistory       string
	AllegicHx           string
	SocialHx            string
	Occupation          string
	Smoking             string
	Alcohole            string
	GeneralExamination  string
	FGS                 string
	FGSImage            string
	ThyroidExam         string
	Eyes                string
	RS                  string
	ABD                 string
	CNS                 string
	CVS                 string
	ClinicType          string
	RegistrationID      string
	ID                  string
	CreatedUser         string
	CreatedDate         string
}{
	PresentingComplaint: "InitialScreening.PresentingComplaint",
	PastMedicalHx:       "InitialScreening.PastMedicalHx",
	ObstreticHx:         "InitialScreening.ObstreticHx",
	MedicationSummary:   "InitialScreening.MedicationSummary",
	FamilyHistory:       "InitialScreening.FamilyHistory",
	AllegicHx:           "InitialScreening.AllegicHx",
	SocialHx:            "InitialScreening.SocialHx",
	Occupation:          "InitialScreening.Occupation",
	Smoking:             "InitialScreening.Smoking",
	Alcohole:            "InitialScreening.Alcohole",
	GeneralExamination:  "InitialScreening.GeneralExamination",
	FGS:                 "InitialScreening.FGS",
	FGSImage:            "InitialScreening.FGSImage",
	ThyroidExam:         "InitialScreening.ThyroidExam",
	Eyes:                "InitialScreening.Eyes",
	RS:                  "InitialScreening.RS",
	ABD:                 "InitialScreening.ABD",
	CNS:                 "InitialScreening.CNS",
	CVS:                 "InitialScreening.CVS",
	ClinicType:          "InitialScreening.ClinicType",
	RegistrationID:      "InitialScreening.RegistrationID",
	ID:                  "InitialScreening.ID",
	CreatedUser:         "InitialScreening.CreatedUser",
	CreatedDate:         "InitialScreening.CreatedDate",
}

// Generated where

var InitialScreeningWhere = struct {
	PresentingComplaint whereHelperstring
	PastMedicalHx       whereHelperstring
	ObstreticHx         whereHelperstring
	MedicationSummary   whereHelperstring
	FamilyHistory       whereHelperstring
	AllegicHx           whereHelperstring
	SocialHx            whereHelperstring
	Occupation          whereHelperstring
	Smoking             whereHelperstring
	Alcohole            whereHelperstring
	GeneralExamination  whereHelperstring
	FGS                 whereHelperstring
	FGSImage            whereHelpernull_String
	ThyroidExam         whereHelperstring
	Eyes                whereHelperstring
	RS                  whereHelperstring
	ABD                 whereHelperstring
	CNS                 whereHelperstring
	CVS                 whereHelperstring
	ClinicType          whereHelpernull_String
	RegistrationID      whereHelperstring
	ID                  whereHelperint
	CreatedUser         whereHelperstring
	CreatedDate         whereHelpertime_Time
}{
	PresentingComplaint: whereHelperstring{field: "[dbo].[InitialScreening].[PresentingComplaint]"},
	PastMedicalHx:       whereHelperstring{field: "[dbo].[InitialScreening].[PastMedicalHx]"},
	ObstreticHx:         whereHelperstring{field: "[dbo].[InitialScreening].[ObstreticHx]"},
	MedicationSummary:   whereHelperstring{field: "[dbo].[InitialScreening].[MedicationSummary]"},
	FamilyHistory:       whereHelperstring{field: "[dbo].[InitialScreening].[FamilyHistory]"},
	AllegicHx:           whereHelperstring{field: "[dbo].[InitialScreening].[AllegicHx]"},
	SocialHx:            whereHelperstring{field: "[dbo].[InitialScreening].[SocialHx]"},
	Occupation:          whereHelperstring{field: "[dbo].[InitialScreening].[Occupation]"},
	Smoking:             whereHelperstring{field: "[dbo].[InitialScreening].[Smoking]"},
	Alcohole:            whereHelperstring{field: "[dbo].[InitialScreening].[Alcohole]"},
	GeneralExamination:  whereHelperstring{field: "[dbo].[InitialScreening].[GeneralExamination]"},
	FGS:                 whereHelperstring{field: "[dbo].[InitialScreening].[FGS]"},
	FGSImage:            whereHelpernull_String{field: "[dbo].[InitialScreening].[FGSImage]"},
	ThyroidExam:         whereHelperstring{field: "[dbo].[InitialScreening].[ThyroidExam]"},
	Eyes:                whereHelperstring{field: "[dbo].[InitialScreening].[Eyes]"},
	RS:                  whereHelperstring{field: "[dbo].[InitialScreening].[RS]"},
	ABD:                 whereHelperstring{field: "[dbo].[InitialScreening].[ABD]"},
	CNS:                 whereHelperstring{field: "[dbo].[InitialScreening].[CNS]"},
	CVS:                 whereHelperstring{field: "[dbo].[InitialScreening].[CVS]"},
	ClinicType:          whereHelpernull_String{field: "[dbo].[InitialScreening].[ClinicType]"},
	RegistrationID:      whereHelperstring{field: "[dbo].[InitialScreening].[RegistrationID]"},
	ID:                  whereHelperint{field: "[dbo].[InitialScreening].[ID]"},
	CreatedUser:         whereHelperstring{field: "[dbo].[InitialScreening].[CreatedUser]"},
	CreatedDate:         whereHelpertime_Time{field: "[dbo].[InitialScreening].[CreatedDate]"},
}

// InitialScreeningRels is where relationship names are stored.
var InitialScreeningRels = struct {
}{}

// initialScreeningR is where relationships are stored.
type initialScreeningR struct {
}

// NewStruct creates a new relationship struct
func (*initialScreeningR) NewStruct() *initialScreeningR {
	return &initialScreeningR{}
}

// initialScreeningL is where Load methods for each relationship are stored.
type initialScreeningL struct{}

var (
	initialScreeningAllColumns            = []string{"PresentingComplaint", "PastMedicalHx", "ObstreticHx", "MedicationSummary", "FamilyHistory", "AllegicHx", "SocialHx", "Occupation", "Smoking", "Alcohole", "GeneralExamination", "FGS", "FGSImage", "ThyroidExam", "Eyes", "RS", "ABD", "CNS", "CVS", "ClinicType", "RegistrationID", "ID", "CreatedUser", "CreatedDate"}
	initialScreeningColumnsWithAuto       = []string{}
	initialScreeningColumnsWithoutDefault = []string{"FGSImage"}
	initialScreeningColumnsWithDefault    = []string{"PresentingComplaint", "PastMedicalHx", "ObstreticHx", "MedicationSummary", "FamilyHistory", "AllegicHx", "SocialHx", "Occupation", "Smoking", "Alcohole", "GeneralExamination", "FGS", "ThyroidExam", "Eyes", "RS", "ABD", "CNS", "CVS", "ClinicType", "RegistrationID", "ID", "CreatedUser", "CreatedDate"}
	initialScreeningPrimaryKeyColumns     = []string{"ID"}
)

type (
	// InitialScreeningSlice is an alias for a slice of pointers to InitialScreening.
	// This should almost always be used instead of []InitialScreening.
	InitialScreeningSlice []*InitialScreening
	// InitialScreeningHook is the signature for custom InitialScreening hook methods
	InitialScreeningHook func(context.Context, boil.ContextExecutor, *InitialScreening) error

	initialScreeningQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	initialScreeningType                 = reflect.TypeOf(&InitialScreening{})
	initialScreeningMapping              = queries.MakeStructMapping(initialScreeningType)
	initialScreeningPrimaryKeyMapping, _ = queries.BindMapping(initialScreeningType, initialScreeningMapping, initialScreeningPrimaryKeyColumns)
	initialScreeningInsertCacheMut       sync.RWMutex
	initialScreeningInsertCache          = make(map[string]insertCache)
	initialScreeningUpdateCacheMut       sync.RWMutex
	initialScreeningUpdateCache          = make(map[string]updateCache)
	initialScreeningUpsertCacheMut       sync.RWMutex
	initialScreeningUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var initialScreeningBeforeInsertHooks []InitialScreeningHook
var initialScreeningBeforeUpdateHooks []InitialScreeningHook
var initialScreeningBeforeDeleteHooks []InitialScreeningHook
var initialScreeningBeforeUpsertHooks []InitialScreeningHook

var initialScreeningAfterInsertHooks []InitialScreeningHook
var initialScreeningAfterSelectHooks []InitialScreeningHook
var initialScreeningAfterUpdateHooks []InitialScreeningHook
var initialScreeningAfterDeleteHooks []InitialScreeningHook
var initialScreeningAfterUpsertHooks []InitialScreeningHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InitialScreening) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InitialScreening) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InitialScreening) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InitialScreening) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InitialScreening) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InitialScreening) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InitialScreening) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InitialScreening) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InitialScreening) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range initialScreeningAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInitialScreeningHook registers your hook function for all future operations.
func AddInitialScreeningHook(hookPoint boil.HookPoint, initialScreeningHook InitialScreeningHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		initialScreeningBeforeInsertHooks = append(initialScreeningBeforeInsertHooks, initialScreeningHook)
	case boil.BeforeUpdateHook:
		initialScreeningBeforeUpdateHooks = append(initialScreeningBeforeUpdateHooks, initialScreeningHook)
	case boil.BeforeDeleteHook:
		initialScreeningBeforeDeleteHooks = append(initialScreeningBeforeDeleteHooks, initialScreeningHook)
	case boil.BeforeUpsertHook:
		initialScreeningBeforeUpsertHooks = append(initialScreeningBeforeUpsertHooks, initialScreeningHook)
	case boil.AfterInsertHook:
		initialScreeningAfterInsertHooks = append(initialScreeningAfterInsertHooks, initialScreeningHook)
	case boil.AfterSelectHook:
		initialScreeningAfterSelectHooks = append(initialScreeningAfterSelectHooks, initialScreeningHook)
	case boil.AfterUpdateHook:
		initialScreeningAfterUpdateHooks = append(initialScreeningAfterUpdateHooks, initialScreeningHook)
	case boil.AfterDeleteHook:
		initialScreeningAfterDeleteHooks = append(initialScreeningAfterDeleteHooks, initialScreeningHook)
	case boil.AfterUpsertHook:
		initialScreeningAfterUpsertHooks = append(initialScreeningAfterUpsertHooks, initialScreeningHook)
	}
}

// One returns a single initialScreening record from the query.
func (q initialScreeningQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InitialScreening, error) {
	o := &InitialScreening{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for InitialScreening")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InitialScreening records from the query.
func (q initialScreeningQuery) All(ctx context.Context, exec boil.ContextExecutor) (InitialScreeningSlice, error) {
	var o []*InitialScreening

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InitialScreening slice")
	}

	if len(initialScreeningAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InitialScreening records in the query.
func (q initialScreeningQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count InitialScreening rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q initialScreeningQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if InitialScreening exists")
	}

	return count > 0, nil
}

// InitialScreenings retrieves all the records using an executor.
func InitialScreenings(mods ...qm.QueryMod) initialScreeningQuery {
	mods = append(mods, qm.From("[dbo].[InitialScreening]"))
	return initialScreeningQuery{NewQuery(mods...)}
}

// FindInitialScreening retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInitialScreening(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*InitialScreening, error) {
	initialScreeningObj := &InitialScreening{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[InitialScreening] where [ID]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, initialScreeningObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from InitialScreening")
	}

	if err = initialScreeningObj.doAfterSelectHooks(ctx, exec); err != nil {
		return initialScreeningObj, err
	}

	return initialScreeningObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InitialScreening) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InitialScreening provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(initialScreeningColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	initialScreeningInsertCacheMut.RLock()
	cache, cached := initialScreeningInsertCache[key]
	initialScreeningInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			initialScreeningAllColumns,
			initialScreeningColumnsWithDefault,
			initialScreeningColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(initialScreeningType, initialScreeningMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(initialScreeningType, initialScreeningMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[InitialScreening] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[InitialScreening] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into InitialScreening")
	}

	if !cached {
		initialScreeningInsertCacheMut.Lock()
		initialScreeningInsertCache[key] = cache
		initialScreeningInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InitialScreening.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InitialScreening) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	initialScreeningUpdateCacheMut.RLock()
	cache, cached := initialScreeningUpdateCache[key]
	initialScreeningUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			initialScreeningAllColumns,
			initialScreeningPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, initialScreeningColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update InitialScreening, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[InitialScreening] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, initialScreeningPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(initialScreeningType, initialScreeningMapping, append(wl, initialScreeningPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update InitialScreening row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for InitialScreening")
	}

	if !cached {
		initialScreeningUpdateCacheMut.Lock()
		initialScreeningUpdateCache[key] = cache
		initialScreeningUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q initialScreeningQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for InitialScreening")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for InitialScreening")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InitialScreeningSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), initialScreeningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[InitialScreening] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, initialScreeningPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in initialScreening slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all initialScreening")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *InitialScreening) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InitialScreening provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(initialScreeningColumnsWithDefault, o)

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

	initialScreeningUpsertCacheMut.RLock()
	cache, cached := initialScreeningUpsertCache[key]
	initialScreeningUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			initialScreeningAllColumns,
			initialScreeningColumnsWithDefault,
			initialScreeningColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, initialScreeningColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(initialScreeningPrimaryKeyColumns, v) && strmangle.ContainsAny(initialScreeningColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert InitialScreening, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, initialScreeningColumnsWithAuto)
		ret = strmangle.SetMerge(ret, initialScreeningColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			initialScreeningAllColumns,
			initialScreeningPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, initialScreeningColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert InitialScreening, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[InitialScreening]", initialScreeningPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(initialScreeningPrimaryKeyColumns))
		copy(whitelist, initialScreeningPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(initialScreeningType, initialScreeningMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(initialScreeningType, initialScreeningMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert InitialScreening")
	}

	if !cached {
		initialScreeningUpsertCacheMut.Lock()
		initialScreeningUpsertCache[key] = cache
		initialScreeningUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single InitialScreening record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InitialScreening) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InitialScreening provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), initialScreeningPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[InitialScreening] WHERE [ID]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from InitialScreening")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for InitialScreening")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q initialScreeningQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no initialScreeningQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from InitialScreening")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InitialScreening")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InitialScreeningSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(initialScreeningBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), initialScreeningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[InitialScreening] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, initialScreeningPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from initialScreening slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InitialScreening")
	}

	if len(initialScreeningAfterDeleteHooks) != 0 {
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
func (o *InitialScreening) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInitialScreening(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InitialScreeningSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InitialScreeningSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), initialScreeningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[InitialScreening].* FROM [dbo].[InitialScreening] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, initialScreeningPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InitialScreeningSlice")
	}

	*o = slice

	return nil
}

// InitialScreeningExists checks if the InitialScreening row exists.
func InitialScreeningExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[InitialScreening] where [ID]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if InitialScreening exists")
	}

	return exists, nil
}

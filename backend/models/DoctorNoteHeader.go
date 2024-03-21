// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/types"
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

// DoctorNoteHeader is an object representing the database table.
type DoctorNoteHeader struct {
	ID              int           `boil:"ID" json:"ID" toml:"ID" yaml:"ID"`
	CompanyCode     string        `boil:"CompanyCode" json:"CompanyCode" toml:"CompanyCode" yaml:"CompanyCode"`
	DocumentID      string        `boil:"DocumentID" json:"DocumentID" toml:"DocumentID" yaml:"DocumentID"`
	EntryDate       time.Time     `boil:"EntryDate" json:"EntryDate" toml:"EntryDate" yaml:"EntryDate"`
	CreatedUser     null.String   `boil:"CreatedUser" json:"CreatedUser,omitempty" toml:"CreatedUser" yaml:"CreatedUser,omitempty"`
	ModifyUser      null.String   `boil:"ModifyUser" json:"ModifyUser,omitempty" toml:"ModifyUser" yaml:"ModifyUser,omitempty"`
	DeletedUser     null.String   `boil:"DeletedUser" json:"DeletedUser,omitempty" toml:"DeletedUser" yaml:"DeletedUser,omitempty"`
	CreatedDate     time.Time     `boil:"CreatedDate" json:"CreatedDate" toml:"CreatedDate" yaml:"CreatedDate"`
	ModifyDate      time.Time     `boil:"ModifyDate" json:"ModifyDate" toml:"ModifyDate" yaml:"ModifyDate"`
	DeletedDate     time.Time     `boil:"DeletedDate" json:"DeletedDate" toml:"DeletedDate" yaml:"DeletedDate"`
	BDelete         bool          `boil:"bDelete" json:"bDelete" toml:"bDelete" yaml:"bDelete"`
	RegistrationNo  string        `boil:"RegistrationNo" json:"RegistrationNo" toml:"RegistrationNo" yaml:"RegistrationNo"`
	BP              null.String   `boil:"BP" json:"BP,omitempty" toml:"BP" yaml:"BP,omitempty"`
	PulseRate       types.Decimal `boil:"PulseRate" json:"PulseRate" toml:"PulseRate" yaml:"PulseRate"`
	Weight          types.Decimal `boil:"Weight" json:"Weight" toml:"Weight" yaml:"Weight"`
	Height          types.Decimal `boil:"Height" json:"Height" toml:"Height" yaml:"Height"`
	BMI             types.Decimal `boil:"BMI" json:"BMI" toml:"BMI" yaml:"BMI"`
	Comment         string        `boil:"Comment" json:"Comment" toml:"Comment" yaml:"Comment"`
	Image1Path      string        `boil:"Image1Path" json:"Image1Path" toml:"Image1Path" yaml:"Image1Path"`
	Image2Path      string        `boil:"Image2Path" json:"Image2Path" toml:"Image2Path" yaml:"Image2Path"`
	Image3Path      string        `boil:"Image3Path" json:"Image3Path" toml:"Image3Path" yaml:"Image3Path"`
	Image4Path      string        `boil:"Image4Path" json:"Image4Path" toml:"Image4Path" yaml:"Image4Path"`
	CommentInternal string        `boil:"CommentInternal" json:"CommentInternal" toml:"CommentInternal" yaml:"CommentInternal"`
	NextClinicIn    string        `boil:"NextClinicIn" json:"NextClinicIn" toml:"NextClinicIn" yaml:"NextClinicIn"`
	AnnualVisit     string        `boil:"AnnualVisit" json:"AnnualVisit" toml:"AnnualVisit" yaml:"AnnualVisit"`

	R *doctorNoteHeaderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L doctorNoteHeaderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DoctorNoteHeaderColumns = struct {
	ID              string
	CompanyCode     string
	DocumentID      string
	EntryDate       string
	CreatedUser     string
	ModifyUser      string
	DeletedUser     string
	CreatedDate     string
	ModifyDate      string
	DeletedDate     string
	BDelete         string
	RegistrationNo  string
	BP              string
	PulseRate       string
	Weight          string
	Height          string
	BMI             string
	Comment         string
	Image1Path      string
	Image2Path      string
	Image3Path      string
	Image4Path      string
	CommentInternal string
	NextClinicIn    string
	AnnualVisit     string
}{
	ID:              "ID",
	CompanyCode:     "CompanyCode",
	DocumentID:      "DocumentID",
	EntryDate:       "EntryDate",
	CreatedUser:     "CreatedUser",
	ModifyUser:      "ModifyUser",
	DeletedUser:     "DeletedUser",
	CreatedDate:     "CreatedDate",
	ModifyDate:      "ModifyDate",
	DeletedDate:     "DeletedDate",
	BDelete:         "bDelete",
	RegistrationNo:  "RegistrationNo",
	BP:              "BP",
	PulseRate:       "PulseRate",
	Weight:          "Weight",
	Height:          "Height",
	BMI:             "BMI",
	Comment:         "Comment",
	Image1Path:      "Image1Path",
	Image2Path:      "Image2Path",
	Image3Path:      "Image3Path",
	Image4Path:      "Image4Path",
	CommentInternal: "CommentInternal",
	NextClinicIn:    "NextClinicIn",
	AnnualVisit:     "AnnualVisit",
}

var DoctorNoteHeaderTableColumns = struct {
	ID              string
	CompanyCode     string
	DocumentID      string
	EntryDate       string
	CreatedUser     string
	ModifyUser      string
	DeletedUser     string
	CreatedDate     string
	ModifyDate      string
	DeletedDate     string
	BDelete         string
	RegistrationNo  string
	BP              string
	PulseRate       string
	Weight          string
	Height          string
	BMI             string
	Comment         string
	Image1Path      string
	Image2Path      string
	Image3Path      string
	Image4Path      string
	CommentInternal string
	NextClinicIn    string
	AnnualVisit     string
}{
	ID:              "DoctorNoteHeader.ID",
	CompanyCode:     "DoctorNoteHeader.CompanyCode",
	DocumentID:      "DoctorNoteHeader.DocumentID",
	EntryDate:       "DoctorNoteHeader.EntryDate",
	CreatedUser:     "DoctorNoteHeader.CreatedUser",
	ModifyUser:      "DoctorNoteHeader.ModifyUser",
	DeletedUser:     "DoctorNoteHeader.DeletedUser",
	CreatedDate:     "DoctorNoteHeader.CreatedDate",
	ModifyDate:      "DoctorNoteHeader.ModifyDate",
	DeletedDate:     "DoctorNoteHeader.DeletedDate",
	BDelete:         "DoctorNoteHeader.bDelete",
	RegistrationNo:  "DoctorNoteHeader.RegistrationNo",
	BP:              "DoctorNoteHeader.BP",
	PulseRate:       "DoctorNoteHeader.PulseRate",
	Weight:          "DoctorNoteHeader.Weight",
	Height:          "DoctorNoteHeader.Height",
	BMI:             "DoctorNoteHeader.BMI",
	Comment:         "DoctorNoteHeader.Comment",
	Image1Path:      "DoctorNoteHeader.Image1Path",
	Image2Path:      "DoctorNoteHeader.Image2Path",
	Image3Path:      "DoctorNoteHeader.Image3Path",
	Image4Path:      "DoctorNoteHeader.Image4Path",
	CommentInternal: "DoctorNoteHeader.CommentInternal",
	NextClinicIn:    "DoctorNoteHeader.NextClinicIn",
	AnnualVisit:     "DoctorNoteHeader.AnnualVisit",
}

// Generated where

type whereHelpertypes_Decimal struct{ field string }

func (w whereHelpertypes_Decimal) EQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_Decimal) NEQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_Decimal) LT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Decimal) LTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Decimal) GT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Decimal) GTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var DoctorNoteHeaderWhere = struct {
	ID              whereHelperint
	CompanyCode     whereHelperstring
	DocumentID      whereHelperstring
	EntryDate       whereHelpertime_Time
	CreatedUser     whereHelpernull_String
	ModifyUser      whereHelpernull_String
	DeletedUser     whereHelpernull_String
	CreatedDate     whereHelpertime_Time
	ModifyDate      whereHelpertime_Time
	DeletedDate     whereHelpertime_Time
	BDelete         whereHelperbool
	RegistrationNo  whereHelperstring
	BP              whereHelpernull_String
	PulseRate       whereHelpertypes_Decimal
	Weight          whereHelpertypes_Decimal
	Height          whereHelpertypes_Decimal
	BMI             whereHelpertypes_Decimal
	Comment         whereHelperstring
	Image1Path      whereHelperstring
	Image2Path      whereHelperstring
	Image3Path      whereHelperstring
	Image4Path      whereHelperstring
	CommentInternal whereHelperstring
	NextClinicIn    whereHelperstring
	AnnualVisit     whereHelperstring
}{
	ID:              whereHelperint{field: "[dbo].[DoctorNoteHeader].[ID]"},
	CompanyCode:     whereHelperstring{field: "[dbo].[DoctorNoteHeader].[CompanyCode]"},
	DocumentID:      whereHelperstring{field: "[dbo].[DoctorNoteHeader].[DocumentID]"},
	EntryDate:       whereHelpertime_Time{field: "[dbo].[DoctorNoteHeader].[EntryDate]"},
	CreatedUser:     whereHelpernull_String{field: "[dbo].[DoctorNoteHeader].[CreatedUser]"},
	ModifyUser:      whereHelpernull_String{field: "[dbo].[DoctorNoteHeader].[ModifyUser]"},
	DeletedUser:     whereHelpernull_String{field: "[dbo].[DoctorNoteHeader].[DeletedUser]"},
	CreatedDate:     whereHelpertime_Time{field: "[dbo].[DoctorNoteHeader].[CreatedDate]"},
	ModifyDate:      whereHelpertime_Time{field: "[dbo].[DoctorNoteHeader].[ModifyDate]"},
	DeletedDate:     whereHelpertime_Time{field: "[dbo].[DoctorNoteHeader].[DeletedDate]"},
	BDelete:         whereHelperbool{field: "[dbo].[DoctorNoteHeader].[bDelete]"},
	RegistrationNo:  whereHelperstring{field: "[dbo].[DoctorNoteHeader].[RegistrationNo]"},
	BP:              whereHelpernull_String{field: "[dbo].[DoctorNoteHeader].[BP]"},
	PulseRate:       whereHelpertypes_Decimal{field: "[dbo].[DoctorNoteHeader].[PulseRate]"},
	Weight:          whereHelpertypes_Decimal{field: "[dbo].[DoctorNoteHeader].[Weight]"},
	Height:          whereHelpertypes_Decimal{field: "[dbo].[DoctorNoteHeader].[Height]"},
	BMI:             whereHelpertypes_Decimal{field: "[dbo].[DoctorNoteHeader].[BMI]"},
	Comment:         whereHelperstring{field: "[dbo].[DoctorNoteHeader].[Comment]"},
	Image1Path:      whereHelperstring{field: "[dbo].[DoctorNoteHeader].[Image1Path]"},
	Image2Path:      whereHelperstring{field: "[dbo].[DoctorNoteHeader].[Image2Path]"},
	Image3Path:      whereHelperstring{field: "[dbo].[DoctorNoteHeader].[Image3Path]"},
	Image4Path:      whereHelperstring{field: "[dbo].[DoctorNoteHeader].[Image4Path]"},
	CommentInternal: whereHelperstring{field: "[dbo].[DoctorNoteHeader].[CommentInternal]"},
	NextClinicIn:    whereHelperstring{field: "[dbo].[DoctorNoteHeader].[NextClinicIn]"},
	AnnualVisit:     whereHelperstring{field: "[dbo].[DoctorNoteHeader].[AnnualVisit]"},
}

// DoctorNoteHeaderRels is where relationship names are stored.
var DoctorNoteHeaderRels = struct {
}{}

// doctorNoteHeaderR is where relationships are stored.
type doctorNoteHeaderR struct {
}

// NewStruct creates a new relationship struct
func (*doctorNoteHeaderR) NewStruct() *doctorNoteHeaderR {
	return &doctorNoteHeaderR{}
}

// doctorNoteHeaderL is where Load methods for each relationship are stored.
type doctorNoteHeaderL struct{}

var (
	doctorNoteHeaderAllColumns            = []string{"ID", "CompanyCode", "DocumentID", "EntryDate", "CreatedUser", "ModifyUser", "DeletedUser", "CreatedDate", "ModifyDate", "DeletedDate", "bDelete", "RegistrationNo", "BP", "PulseRate", "Weight", "Height", "BMI", "Comment", "Image1Path", "Image2Path", "Image3Path", "Image4Path", "CommentInternal", "NextClinicIn", "AnnualVisit"}
	doctorNoteHeaderColumnsWithAuto       = []string{}
	doctorNoteHeaderColumnsWithoutDefault = []string{"CreatedUser", "ModifyUser", "DeletedUser", "BP"}
	doctorNoteHeaderColumnsWithDefault    = []string{"ID", "CompanyCode", "DocumentID", "EntryDate", "CreatedDate", "ModifyDate", "DeletedDate", "bDelete", "RegistrationNo", "PulseRate", "Weight", "Height", "BMI", "Comment", "Image1Path", "Image2Path", "Image3Path", "Image4Path", "CommentInternal", "NextClinicIn", "AnnualVisit"}
	doctorNoteHeaderPrimaryKeyColumns     = []string{"ID"}
)

type (
	// DoctorNoteHeaderSlice is an alias for a slice of pointers to DoctorNoteHeader.
	// This should almost always be used instead of []DoctorNoteHeader.
	DoctorNoteHeaderSlice []*DoctorNoteHeader
	// DoctorNoteHeaderHook is the signature for custom DoctorNoteHeader hook methods
	DoctorNoteHeaderHook func(context.Context, boil.ContextExecutor, *DoctorNoteHeader) error

	doctorNoteHeaderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	doctorNoteHeaderType                 = reflect.TypeOf(&DoctorNoteHeader{})
	doctorNoteHeaderMapping              = queries.MakeStructMapping(doctorNoteHeaderType)
	doctorNoteHeaderPrimaryKeyMapping, _ = queries.BindMapping(doctorNoteHeaderType, doctorNoteHeaderMapping, doctorNoteHeaderPrimaryKeyColumns)
	doctorNoteHeaderInsertCacheMut       sync.RWMutex
	doctorNoteHeaderInsertCache          = make(map[string]insertCache)
	doctorNoteHeaderUpdateCacheMut       sync.RWMutex
	doctorNoteHeaderUpdateCache          = make(map[string]updateCache)
	doctorNoteHeaderUpsertCacheMut       sync.RWMutex
	doctorNoteHeaderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var doctorNoteHeaderBeforeInsertHooks []DoctorNoteHeaderHook
var doctorNoteHeaderBeforeUpdateHooks []DoctorNoteHeaderHook
var doctorNoteHeaderBeforeDeleteHooks []DoctorNoteHeaderHook
var doctorNoteHeaderBeforeUpsertHooks []DoctorNoteHeaderHook

var doctorNoteHeaderAfterInsertHooks []DoctorNoteHeaderHook
var doctorNoteHeaderAfterSelectHooks []DoctorNoteHeaderHook
var doctorNoteHeaderAfterUpdateHooks []DoctorNoteHeaderHook
var doctorNoteHeaderAfterDeleteHooks []DoctorNoteHeaderHook
var doctorNoteHeaderAfterUpsertHooks []DoctorNoteHeaderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DoctorNoteHeader) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DoctorNoteHeader) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DoctorNoteHeader) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DoctorNoteHeader) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DoctorNoteHeader) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DoctorNoteHeader) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DoctorNoteHeader) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DoctorNoteHeader) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DoctorNoteHeader) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorNoteHeaderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDoctorNoteHeaderHook registers your hook function for all future operations.
func AddDoctorNoteHeaderHook(hookPoint boil.HookPoint, doctorNoteHeaderHook DoctorNoteHeaderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		doctorNoteHeaderBeforeInsertHooks = append(doctorNoteHeaderBeforeInsertHooks, doctorNoteHeaderHook)
	case boil.BeforeUpdateHook:
		doctorNoteHeaderBeforeUpdateHooks = append(doctorNoteHeaderBeforeUpdateHooks, doctorNoteHeaderHook)
	case boil.BeforeDeleteHook:
		doctorNoteHeaderBeforeDeleteHooks = append(doctorNoteHeaderBeforeDeleteHooks, doctorNoteHeaderHook)
	case boil.BeforeUpsertHook:
		doctorNoteHeaderBeforeUpsertHooks = append(doctorNoteHeaderBeforeUpsertHooks, doctorNoteHeaderHook)
	case boil.AfterInsertHook:
		doctorNoteHeaderAfterInsertHooks = append(doctorNoteHeaderAfterInsertHooks, doctorNoteHeaderHook)
	case boil.AfterSelectHook:
		doctorNoteHeaderAfterSelectHooks = append(doctorNoteHeaderAfterSelectHooks, doctorNoteHeaderHook)
	case boil.AfterUpdateHook:
		doctorNoteHeaderAfterUpdateHooks = append(doctorNoteHeaderAfterUpdateHooks, doctorNoteHeaderHook)
	case boil.AfterDeleteHook:
		doctorNoteHeaderAfterDeleteHooks = append(doctorNoteHeaderAfterDeleteHooks, doctorNoteHeaderHook)
	case boil.AfterUpsertHook:
		doctorNoteHeaderAfterUpsertHooks = append(doctorNoteHeaderAfterUpsertHooks, doctorNoteHeaderHook)
	}
}

// One returns a single doctorNoteHeader record from the query.
func (q doctorNoteHeaderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DoctorNoteHeader, error) {
	o := &DoctorNoteHeader{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for DoctorNoteHeader")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DoctorNoteHeader records from the query.
func (q doctorNoteHeaderQuery) All(ctx context.Context, exec boil.ContextExecutor) (DoctorNoteHeaderSlice, error) {
	var o []*DoctorNoteHeader

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DoctorNoteHeader slice")
	}

	if len(doctorNoteHeaderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DoctorNoteHeader records in the query.
func (q doctorNoteHeaderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count DoctorNoteHeader rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q doctorNoteHeaderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if DoctorNoteHeader exists")
	}

	return count > 0, nil
}

// DoctorNoteHeaders retrieves all the records using an executor.
func DoctorNoteHeaders(mods ...qm.QueryMod) doctorNoteHeaderQuery {
	mods = append(mods, qm.From("[dbo].[DoctorNoteHeader]"))
	return doctorNoteHeaderQuery{NewQuery(mods...)}
}

// FindDoctorNoteHeader retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDoctorNoteHeader(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DoctorNoteHeader, error) {
	doctorNoteHeaderObj := &DoctorNoteHeader{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[DoctorNoteHeader] where [ID]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, doctorNoteHeaderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from DoctorNoteHeader")
	}

	if err = doctorNoteHeaderObj.doAfterSelectHooks(ctx, exec); err != nil {
		return doctorNoteHeaderObj, err
	}

	return doctorNoteHeaderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DoctorNoteHeader) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no DoctorNoteHeader provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(doctorNoteHeaderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	doctorNoteHeaderInsertCacheMut.RLock()
	cache, cached := doctorNoteHeaderInsertCache[key]
	doctorNoteHeaderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			doctorNoteHeaderAllColumns,
			doctorNoteHeaderColumnsWithDefault,
			doctorNoteHeaderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(doctorNoteHeaderType, doctorNoteHeaderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(doctorNoteHeaderType, doctorNoteHeaderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[DoctorNoteHeader] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[DoctorNoteHeader] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into DoctorNoteHeader")
	}

	if !cached {
		doctorNoteHeaderInsertCacheMut.Lock()
		doctorNoteHeaderInsertCache[key] = cache
		doctorNoteHeaderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DoctorNoteHeader.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DoctorNoteHeader) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	doctorNoteHeaderUpdateCacheMut.RLock()
	cache, cached := doctorNoteHeaderUpdateCache[key]
	doctorNoteHeaderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			doctorNoteHeaderAllColumns,
			doctorNoteHeaderPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, doctorNoteHeaderColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update DoctorNoteHeader, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[DoctorNoteHeader] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, doctorNoteHeaderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(doctorNoteHeaderType, doctorNoteHeaderMapping, append(wl, doctorNoteHeaderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update DoctorNoteHeader row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for DoctorNoteHeader")
	}

	if !cached {
		doctorNoteHeaderUpdateCacheMut.Lock()
		doctorNoteHeaderUpdateCache[key] = cache
		doctorNoteHeaderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q doctorNoteHeaderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for DoctorNoteHeader")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for DoctorNoteHeader")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DoctorNoteHeaderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorNoteHeaderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[DoctorNoteHeader] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, doctorNoteHeaderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in doctorNoteHeader slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all doctorNoteHeader")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DoctorNoteHeader) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no DoctorNoteHeader provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(doctorNoteHeaderColumnsWithDefault, o)

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

	doctorNoteHeaderUpsertCacheMut.RLock()
	cache, cached := doctorNoteHeaderUpsertCache[key]
	doctorNoteHeaderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			doctorNoteHeaderAllColumns,
			doctorNoteHeaderColumnsWithDefault,
			doctorNoteHeaderColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, doctorNoteHeaderColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(doctorNoteHeaderPrimaryKeyColumns, v) && strmangle.ContainsAny(doctorNoteHeaderColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert DoctorNoteHeader, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, doctorNoteHeaderColumnsWithAuto)
		ret = strmangle.SetMerge(ret, doctorNoteHeaderColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			doctorNoteHeaderAllColumns,
			doctorNoteHeaderPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, doctorNoteHeaderColumnsWithAuto)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert DoctorNoteHeader, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[DoctorNoteHeader]", doctorNoteHeaderPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(doctorNoteHeaderPrimaryKeyColumns))
		copy(whitelist, doctorNoteHeaderPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(doctorNoteHeaderType, doctorNoteHeaderMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(doctorNoteHeaderType, doctorNoteHeaderMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert DoctorNoteHeader")
	}

	if !cached {
		doctorNoteHeaderUpsertCacheMut.Lock()
		doctorNoteHeaderUpsertCache[key] = cache
		doctorNoteHeaderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DoctorNoteHeader record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DoctorNoteHeader) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DoctorNoteHeader provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), doctorNoteHeaderPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[DoctorNoteHeader] WHERE [ID]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from DoctorNoteHeader")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for DoctorNoteHeader")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q doctorNoteHeaderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no doctorNoteHeaderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from DoctorNoteHeader")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for DoctorNoteHeader")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DoctorNoteHeaderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(doctorNoteHeaderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorNoteHeaderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[DoctorNoteHeader] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, doctorNoteHeaderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from doctorNoteHeader slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for DoctorNoteHeader")
	}

	if len(doctorNoteHeaderAfterDeleteHooks) != 0 {
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
func (o *DoctorNoteHeader) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDoctorNoteHeader(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DoctorNoteHeaderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DoctorNoteHeaderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorNoteHeaderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[DoctorNoteHeader].* FROM [dbo].[DoctorNoteHeader] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, doctorNoteHeaderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DoctorNoteHeaderSlice")
	}

	*o = slice

	return nil
}

// DoctorNoteHeaderExists checks if the DoctorNoteHeader row exists.
func DoctorNoteHeaderExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[DoctorNoteHeader] where [ID]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if DoctorNoteHeader exists")
	}

	return exists, nil
}
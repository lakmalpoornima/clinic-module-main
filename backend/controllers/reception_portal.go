package controllers

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// DocumentID := r.URL.Query().Get("DocID")
func GetClinicAppointments(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	clinic := r.URL.Query().Get("clinic")
	date := r.URL.Query().Get("date")
	err := queries.Raw("SELECT ( SELECT * FROM fxEndoClinicWisePatients(@clinic,@date) order by AppNo FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('appointments') ) AS result ", sql.Named("clinic", clinic), sql.Named("date", date)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetPatientInfo(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	appno := r.URL.Query().Get("AppNo")
	date := r.URL.Query().Get("ClinicDate")
	clinic := r.URL.Query().Get("clinic")
	reference := r.URL.Query().Get("ReferenceNo")
	// appno := r.URL.Query().Get( "AppNo")
	// date := r.URL.Query().Get( "ClinicDate")
	// clinic := r.URL.Query().Get( "clinic")
	// reference := r.URL.Query().Get( "ReferenceNo")
	err := queries.Raw("SELECT JSON_QUERY((SELECT * from fxGetPatientInfo(@appno,@date,@clinic,@reference) order by AppNo FOR JSON PATH  , ROOT('patient')) , '$.patient[0]')  AS result ", sql.Named("appno", appno), sql.Named("date", date), sql.Named("clinic", clinic), sql.Named("reference", reference)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func SaveReceptionPatient(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	fmt.Println(b, "value GEt", r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
		// input := strings({"referenceNo":"CLI2217052","pr_level":"regular","comment":"jhj","user":"hbt"})
		// fmt.Println(input, "value GEt")
		// st := strings(input)
	sql:
		_, err = db.ExecContext(ctx, "spSaveReceptionPatient",
			sql.Named("input", input), sql.Named("cResult", sql.Out{Dest: &resultDoc}))
		if err == nil {
			if resultDoc == "" {
				if retries < 2 {
					retries++
					goto sql
				}
			} else {
				result = "success"
			}
		} else {
			log.Println(err.Error())
		}
	} else {
		log.Println(err.Error())
	}
	render.JSON(w, r, result)
}

func GetPatientRegistrationInfo(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegistrationNo := r.URL.Query().Get("RegiNo")
	// RegistrationNo := r.URL.Query().Get( "RegiNo")
	err := queries.Raw("SELECT( SELECT * From fxGetPatientRegistrationInfo(@RegiNo) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetPatientRegistrationInfo') ) AS result ", sql.Named("RegiNo", RegistrationNo)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

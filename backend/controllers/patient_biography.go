package controllers

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetLabResultSummary(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	regcode := r.URL.Query().Get("regcode")
	// fmt.Println(regcode, "reg")
	err := queries.Raw("SELECT( SELECT * From fxGetLabResultSummary(@RegCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('LabResultSummary') ) AS result ", sql.Named("RegCode", regcode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}
func GetLabResultDetail(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	regcode := r.URL.Query().Get("regcode")
	showall := r.URL.Query().Get("showall") == "true"
	ClinicCode := r.URL.Query().Get("clinic")
	// fmt.Println(regcode, "reg")
	err := queries.Raw("SELECT( SELECT * From fxGetLabResultDetail(@RegCode,@Clinic,@ShowAll) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('LabResultDetail') ) AS result ", sql.Named("RegCode", regcode), sql.Named("Clinic", ClinicCode), sql.Named("ShowAll", showall)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetLabResultDetailExternal(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	regcode := r.URL.Query().Get("regcode")
	showall := r.URL.Query().Get("showall") == "true"
	ClinicCode := r.URL.Query().Get("clinic")
	// fmt.Println(regcode, "reg")
	err := queries.Raw("SELECT( SELECT * From fxGetLabResultDetailExternal(@RegCode,@Clinic,@ShowAll) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('LabResultDetail') ) AS result ", sql.Named("RegCode", regcode), sql.Named("Clinic", ClinicCode), sql.Named("ShowAll", showall)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

// fxGetLabResultDetailElement

func GetLabResultDetailElement(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	regcode := r.URL.Query().Get("regcode")
	elementCode := r.URL.Query().Get("elementcode")
	// fmt.Println("REG", regcode, "ELEM", elementCode)
	err := queries.Raw("SELECT( SELECT * From fxGetLabResultDetailElement(@RegCode,@ElementCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('RegistrationWiseClinicSessions') ) AS result ",
		sql.Named("RegCode", regcode), sql.Named("ElementCode", elementCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

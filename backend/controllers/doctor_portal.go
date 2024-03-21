package controllers

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetDoctorPortalAppointments(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	clinic := r.URL.Query().Get("clinic")
	err := queries.Raw("SELECT ( SELECT * FROM fxServiceAppoinments(@clinic) where Completed = 'false' FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('appointments') ) AS result ", sql.Named("clinic", clinic)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func POSTDoctorPortalAppointments(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "savespPatientDiseaseofImage",
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

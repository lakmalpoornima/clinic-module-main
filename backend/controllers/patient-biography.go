package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetPatientActivityInfo(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	regcode := r.URL.Query().Get("regcode")
	fmt.Println(regcode, "reg")
	err := queries.Raw("SELECT(SELECT * FROM fxGetPatientActivityDetail(@RegCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('PatientActivityInfo') ) AS result ", sql.Named("RegCode", regcode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

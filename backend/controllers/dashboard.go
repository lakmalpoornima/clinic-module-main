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

func GetDashboardAppointments(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	clinic := r.URL.Query().Get("clinic")
	err := queries.Raw("SELECT ( SELECT * FROM fxDashboardAppoinments(@clinic) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('appointments') ) AS result ", sql.Named("clinic", clinic)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

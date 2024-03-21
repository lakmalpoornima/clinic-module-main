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

func GetSummaryReport(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	DocumentID := r.URL.Query().Get("DocID")
	err := queries.Raw("SELECT * FROM fxGetSummaryReport(@DocID) ", sql.Named("DocID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

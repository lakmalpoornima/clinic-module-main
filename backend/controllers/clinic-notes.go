package controllers

import (
	// "fmt"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	auth "systolic/middleware"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetRegistrationWiseClinicSessions(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegCode := r.URL.Query().Get("regno")
	ClinicCode := r.URL.Query().Get("clinic")
	// fmt.Println("REG", regcode, "ELEM", elementCode)
	err := queries.Raw("SELECT( SELECT * From fxGetRegistrationWiseClinicSessions(@RegCode,@ClinicCode) order by dtpDate desc FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('RegistrationWiseClinicSessions') ) AS result ", sql.Named("RegCode", RegCode), sql.Named("ClinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetLastWiseClinicSessions(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegCode := r.URL.Query().Get("regno")
	ClinicCode := r.URL.Query().Get("clinic")
	// fmt.Println("REG", regcode, "ELEM", elementCode)
	err := queries.Raw("SELECT( SELECT * From fxGetLastDocumentNo(@RegCode,@ClinicCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('RegistrationWiseClinicSessions') ) AS result ", sql.Named("RegCode", RegCode), sql.Named("ClinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetClinicWiseElements(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	ClinicCode := r.URL.Query().Get("clinic")
	// fmt.Println("REG", regcode, "ELEM", elementCode)
	err := queries.Raw("SELECT( SELECT * From fxGetClinicWiseElements(@ClinicCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('ClinicWiseElements') ) AS result ", sql.Named("ClinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentHeaderNotes(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	DocumentID := r.URL.Query().Get("DocID")
	fmt.Println("DOC", DocumentID)
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentHeaderNotes(@DocID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentHeaderNotes') ) AS result ", sql.Named("DocID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

//GetDocumentHeaderProblemList
func GetDocumentHeaderProblemList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	// DocumentID := r.URL.Query().Get( "DocID")
	DocumentID := r.URL.Query().Get("DocID")
	fmt.Println("DOC", DocumentID)
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentHeaderProblemList(@DocID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentHeaderProblemList') ) AS result ", sql.Named("DocID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentHeaderImageList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	// Done
	// date := r.URL.Query().Get("date")
	DocumentID := r.URL.Query().Get("DocID")
	// DocumentID := r.URL.Query().Get("DocID")
	fmt.Println("DOC", DocumentID)
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentHeaderImageList(@DocID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentHeaderImageList') ) AS result ", sql.Named("DocID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

// 2022-03-02

func GetConsultantInfo(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	DocumentID := r.URL.Query().Get("DocID")
	err := queries.Raw("SELECT(SELECT * From GetConsultantInfo(@DocID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('ConsultantInfo') ) AS result ", sql.Named("DocID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDashboardCount(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	ClinicCode := auth.DecodeJWT("location", r).(string)
	userid := auth.DecodeJWT("user_id", r).(string)
	err := queries.Raw("SELECT(SELECT * From DashboardCount(@ClinicCode,@UserID,GETDATE()) FOR JSON PATH ,INCLUDE_NULL_VALUES, WITHOUT_ARRAY_WRAPPER  ) AS result ", sql.Named("ClinicCode", ClinicCode), sql.Named("UserID", userid)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDosageList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	err := queries.Raw("SELECT(SELECT * From DosageList() FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDosageList') ) AS result ").Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDiagnosticList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	SearchText := r.URL.Query().Get("SearchText")
	err := queries.Raw("SELECT(SELECT * From DiagnosticListFilter(@SearchText) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDiagnosticList') ) AS result ", sql.Named("SearchText", SearchText)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetUnitList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	err := queries.Raw("SELECT(SELECT * From UnitList() FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetUnitList') ) AS result ").Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentPrescription(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	DocumentID := r.URL.Query().Get("DocID")
	fmt.Println("DOC", DocumentID)
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentPrescription(@DocID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentPrescription') ) AS result ", sql.Named("DocID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentInitialScreening(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegistrationNo := r.URL.Query().Get("RegiNo")
	ClinicCode := r.URL.Query().Get("ClinicCode")
	// fmt.Println("DOC", DocumentID)
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentInitialScreening(@RegiNo,@ClinicCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentInitialScreening') ) AS result ", sql.Named("RegiNo", RegistrationNo), sql.Named("ClinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentReferralLetterTemplate(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	ClinicCode := r.URL.Query().Get("ClinicCode")
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentReferralLetterTemplate(@ClinicCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentReferralLetterTemplate') ) AS result ", sql.Named("ClinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDoctorList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	err := queries.Raw("SELECT( SELECT * From DoctorList()  FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDoctorNamelist') ) AS result ").Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentReferralLetter(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	DocumentID := r.URL.Query().Get("DocumentID")
	fmt.Println(DocumentID)
	err := queries.Raw("SELECT( SELECT * From  fxGetDocumentReferralLetter(@DocumentID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentReferralLetter') ) AS result ", sql.Named("DocumentID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentCheckedList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	DocumentID := r.URL.Query().Get("DocID")
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentCheckedList(@DocumentID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentCheckedList') ) AS result ", sql.Named("DocumentID", DocumentID)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetDocumentPatientCheckedList(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	ClinicCode := r.URL.Query().Get("Clinic")
	RegistrationNo := r.URL.Query().Get("RegiNo")
	err := queries.Raw("SELECT( SELECT * From fxGetDocumentPatientCheckedList(@RegiNo,@Clinic) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetDocumentPatientCheckedList') ) AS result ", sql.Named("RegiNo", RegistrationNo), sql.Named("Clinic", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func GetPatientDiseaseOfImage(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	// RegistrationNo := r.URL.Query().Get( "RegiNo")
	RegistrationNo := r.URL.Query().Get("RegiNo")
	fmt.Println(RegistrationNo)
	err := queries.Raw("SELECT( SELECT * From fxGetPatientImage(@RegiNo) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetPatientDiseaseofImage') ) AS result ", sql.Named("RegiNo", RegistrationNo)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

// save Image
func Saveimage(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveProfileImage",
			sql.Named("input", input), sql.Named("RegiNo", sql.Out{Dest: &resultDoc}))
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

// SAVE IMAGE

func GetPatientProfileImage(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegistrationNo := r.URL.Query().Get("RegiNo")
	// RegistrationNo := r.URL.Query().Get( "RegiNo")
	fmt.Println(RegistrationNo)
	err := queries.Raw("SELECT( SELECT * From fxGetPatientProfileImage(@RegiNo) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetPatientProfileImage') ) AS result ", sql.Named("RegiNo", RegistrationNo)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

// Filter The Drugs
func GetFilterDrugslist(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	SearchText := r.URL.Query().Get("SearchText")
	err := queries.Raw("SELECT( SELECT * From DrugListWithFilter(@SearchText) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetFilterDrugsList') ) AS result ", sql.Named("SearchText", SearchText)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

// Filter The LabTestWithFilter
func GetFilterLablist(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	SearchText := r.URL.Query().Get("SearchText")
	err := queries.Raw("SELECT( SELECT * From LabTestWithFilter(@SearchText) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetFilterLablist') ) AS result ", sql.Named("SearchText", SearchText)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)

}

// Filter The OtherFilter
func GetFilterOtherInvestigationlist(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	SearchText := r.URL.Query().Get("SearchText")
	err := queries.Raw("SELECT( SELECT * From OtherInvestigation(@SearchText) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetFilterOtherInvestigationlist') ) AS result ", sql.Named("SearchText", SearchText)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)

}

// GET TODAY DOCUMENT NUMBER
func GetTodayDocNumber(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegistrationNo := r.URL.Query().Get("RegiNo")
	ClinicCode := r.URL.Query().Get("ClinicCode")

	err := queries.Raw("SELECT( SELECT * From fxGetTodayDocumentNo(@RegiNo,@ClinicCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('TodayDocument') ) AS result ", sql.Named("RegiNo", RegistrationNo), sql.Named("ClinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)

}

// GetSpecialNotes
func GetSpecialNotes(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	RegID := r.URL.Query().Get("regID")
	ClinicCode := r.URL.Query().Get("clinicCode")
	err := queries.Raw("SELECT( SELECT * From GetSpecialNotes(@regID,@clinicCode) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('GetSpecialNotes') ) AS result ", sql.Named("regID", RegID), sql.Named("clinicCode", ClinicCode)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func DrugAutoComplete(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	GenericName := r.URL.Query().Get("GenericName")
	err := queries.Raw("SELECT( SELECT * From AutoCompleteforDrug(@GenericName) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('DrugAutoCompleteI') ) AS result ", sql.Named("GenericName", GenericName)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

func SaveClinicHeader(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicHeader",
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

func SaveClinicPrescription(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicPrescription",
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

func SaveClinicInitScreening(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicInitScreening",
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

func SaveClinicCheckList(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicCheckList",
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

func SaveClinicImageDisease(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicImageDisease",
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

func SaveProfileImage(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveProfileImage",
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

func SaveClinicExternalLabData(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicExternalLabData",
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

func SaveClinicReferral(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicReferral",
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

func SaveClinicReferralTemplate(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicReferralTemplate",
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

func SaveClinicSpecialNotes(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spSaveClinicSpecialNotes",
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

func AddNewProblem(w http.ResponseWriter, r *http.Request) {
	result := "failed"
	resultDoc := ""
	retries := 0

	b, err := io.ReadAll(r.Body)
	if err == nil {
		input := strings.Replace(string(b[:]), `\"`, `"`, -1)
	sql:
		_, err = db.ExecContext(ctx, "spAddNewProblem",
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

package controllers

/*
import (
	"encoding/json"
	"net/http"
	"strconv"
	models "systolic/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetAllDesignations
func GetAllDesignations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindDesignation
func FindDesignation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteDesignation
func DeleteDesignation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateDesignation
func CreateDesignation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateDesignation
func UpdateDesignation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllSpecialities
func GetAllSpecialities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindSpeciality
func FindSpeciality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteSpeciality
func DeleteSpeciality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateSpeciality
func CreateSpeciality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateSpeciality
func UpdateSpeciality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllAgents
func GetAllAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindAgent
func FindAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteAgent
func DeleteAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateAgent
func CreateAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateAgent
func UpdateAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllConsultants
func GetAllConsultants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindConsultant
func FindConsultant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteConsultant
func DeleteConsultant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateConsultant
func CreateConsultant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateConsultant
func UpdateConsultant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllDebtors
func GetAllDebtors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindDebtor
func FindDebtor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteDebtor
func DeleteDebtor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateDebtor
func CreateDebtor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateDebtor
func UpdateDebtor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllQuotedPrice
func GetAllQuotedPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindQuotedPrice
func FindQuotedPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteQuotedPrice
func DeleteQuotedPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateQuotedPrice
func CreateQuotedPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateQuotedPrice
func UpdateQuotedPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllOutsourceCompanies
func GetAllOutsourceCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindOutsourceCompany
func FindOutsourceCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteOutsourceCompany
func DeleteOutsourceCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateOutsourceCompany
func CreateOutsourceCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateOutsourceCompany
func UpdateOutsourceCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}
*/

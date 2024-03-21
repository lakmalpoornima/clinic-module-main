package controllers

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	models "systolic/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// var result string
// type StringResult struct {
// 	Result string `boil:"result"`
// }

// type StringResult struct {
// 	Result string `boil:"result"`
// }

// GetAllDepartments
func GetAllDepartment(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	queries.Raw(`SELECT get_all_departments() AS result;`).Bind(ctx, db, &result)
	render.PlainText(w, r, result.Result)

}

// FindDepartment
func FindDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var department models.MastDepartment
	param := r.URL.Query().Get( "department_code")
	list := queries.Raw("select * from mast_users where department_code=$1", param).Bind(ctx, db, &department)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, department)
}

// DeleteDepartment
func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := r.URL.Query().Get( "department_code")
	fmt.Println(param, "Department_code ")
	// list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	// json.NewDecoder(r.Body).Decode(&list)
	// render.JSON(w, r, user)
}

// CreateDepartment
func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	// var result StringResult
	department := &models.MastDepartment{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&department)
	department.Insert(ctx, db, boil.Infer())
	render.JSON(w, r, department)
	json.NewEncoder(w).Encode(err)

	if err != nil {
		fmt.Println(err)
	}
	// something went to wrong
}

// UpdateDepartment
func UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := r.URL.Query().Get( "department_code")
	fmt.Println(param, "Department_code")
	//" Here have to get department_code and "

}

//GetAllDivisions
func GetAllDivisions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindDivision
func FindDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteDivision
func DeleteDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateDivision
func CreateDivision(w http.ResponseWriter, r *http.Request) {
	MastCompany := &models.MastCompany{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&MastCompany)
	MastCompany.Insert(ctx, db, boil.Infer())
	render.JSON(w, r, MastCompany)
	json.NewEncoder(w).Encode(err)

	if err != nil {
		fmt.Println(err)
	}
	// something went to wrong
}

// UpdateDivision
func UpdateDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllAreas
func GetAllAreas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindArea
func FindArea(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteArea
func DeleteArea(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateArea
func CreateArea(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateArea
func UpdateArea(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllCompanyInfo
func GetAllCompanyInfo(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	queries.Raw(`SELECT get_all_campanies() AS result;`).Bind(ctx, db, &result)
	render.PlainText(w, r, result.Result)
}

// FindCompanyInfo
func FindCompanyInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastCompany
	param := r.URL.Query().Get( "company_code")
	// user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_company where company_code=$1", param).Bind(ctx, db, &user)

	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteCompanyInfo
func DeleteCompanyInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateCompanyInfo
func CreateCompanyInfo(w http.ResponseWriter, r *http.Request) {
	MastCompany := &models.MastCompany{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&MastCompany)
	MastCompany.CreateID = "Admin"
	MastCompany.Insert(ctx, db, boil.Infer())
	render.JSON(w, r, MastCompany)
	json.NewEncoder(w).Encode(err)
	// insert
	if err != nil {
		fmt.Println(err)
	}
	// something went to wrong
}

// UpdateCompanyInfo
func UpdateCompanyInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser

	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)

}
*/

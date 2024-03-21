package controllers

//"encoding/json"
//"net/http"
//"strconv"
//models "systolic/models"

//"github.com/go-chi/chi"
//"github.com/go-chi/render"
//_ "github.com/lib/pq"
//"github.com/volatiletech/sqlboiler/v4/queries"
// . "github.com/volatiletech/sqlboiler/v4/queries/qm"

/*
// GetAllAnalyzers
func GetAllAnalyzers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindAnalyzer
func FindAnalyzer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteAnalyzer
func DeleteAnalyzer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateAnalyzer
func CreateAnalyzer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateAnalyzer
func UpdateAnalyzer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllAntibiotics
func GetAllAntibiotics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindAntibiotics
func FindAntibiotics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteAntibiotics
func DeleteAntibiotics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateAntibiotics
func CreateAntibiotics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateAntibiotics
func UpdateAntibiotics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllOrganism
func GetAllOrganism(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindOrganism
func FindOrganism(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteOrganism
func DeleteOrganism(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateOrganism
func CreateOrganism(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateOrganism
func UpdateOrganism(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllContainerType
func GetAllContainerType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindContainerType
func FindContainerType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteContainerType
func DeleteContainerType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateContainerType
func CreateContainerType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateContainerType
func UpdateContainerType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllTestGroup
func GetAllTestGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindTestGroup
func FindTestGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteTestGroup
func DeleteTestGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateTestGroup
func CreateTestGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateTestGroup
func UpdateTestGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllElements
func GetAllElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindElements
func FindElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteElements
func DeleteElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateElements
func CreateElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateElements
func UpdateElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllSpecimans
func GetAllSpecimans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindSpecimans
func FindSpecimans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteSpecimans
func DeleteSpecimans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateSpecimans
func CreateSpecimans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateSpecimans
func UpdateSpecimans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}
*/

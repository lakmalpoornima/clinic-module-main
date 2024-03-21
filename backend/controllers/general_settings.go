package controllers

import (
	_ "github.com/lib/pq"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

/*
// GetAllBillingCategories
func GetAllBillingCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindBillingCategory
func FindBillingCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteBillingCategory
func DeleteBillingCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateBillingCategory
func CreateBillingCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateBillingCategory
func UpdateBillingCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllPriceCategories
func GetAllPriceCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindPriceCategory
func FindPriceCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeletePriceCategory
func DeletePriceCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreatePriceCategory
func CreatePriceCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdatePriceCategory
func UpdatePriceCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllUnitofMeasurments
func GetAllUnitofMeasurments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindUnitofMeasurment
func FindUnitofMeasurment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteUnitofMeasurment
func DeleteUnitofMeasurment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateUnitofMeasurment
func CreateUnitofMeasurment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateUnitofMeasurment
func UpdateUnitofMeasurment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllItem
func GetAllItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindItem
func FindItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteItem
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateItem
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateItem
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// GetAllLocation
func GetAllLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// FindLocation
func FindLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// DeleteLocation
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// CreateLocation
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

// UpdateLocation
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")
	user_code, _ := strconv.Atoi(param)
	list := queries.Raw("select * from mast_users where user_code=$1", user_code).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}
*/

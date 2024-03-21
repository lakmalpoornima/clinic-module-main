package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	auth "systolic/middleware"
	"systolic/utils"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/queries"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LoginData struct {
	Userid   string `json:"Userid"`
	Password string `json:"Password"`
	Location string `json:"Location"`
}

type LableAndCode struct {
	Label string `json:"label" boil:"Label"`
	Code  string `json:"code" boil:"Code"`
}

type LoginUser struct {
	UserName     string `boil:"vcUserName"`
	PasswordHash string `boil:"vcPassword"`
	Userid       string `boil:"vcLoginID"`
}

func GetLoginLocationForUser(w http.ResponseWriter, r *http.Request) {
	var result []LableAndCode
	userid := r.URL.Query().Get("userid")
	err := queries.Raw("SELECT COALESCE(cLocationCode,'') as Code , COALESCE(cLocationName,'') as Label FROM GetLoginLocationForUser(@id,'0001')", sql.Named("id", userid)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.JSON(w, r, result)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var result string
	var login LoginData
	result = "failed"
	json.NewDecoder(r.Body).Decode(&login)
	var loginuser LoginUser
	err := queries.Raw("SELECT * FROM GetUserLoginData(@id)", sql.Named("id", login.Userid)).Bind(ctx, db, &loginuser)
	if err != nil {
		log.Println(err.Error())
	}
	if (login.Password) != "" && (login.Userid) != "" {
		validuser := utils.ComparePasswordVb6(login.Password, loginuser.PasswordHash)
		if validuser {
			tokenstring := auth.EncodeJWT(login.Userid, login.Location)
			expiration := time.Now().Add(24 * time.Hour)
			cookie := http.Cookie{Name: "jwt", Value: tokenstring, Expires: expiration, HttpOnly: true}
			http.SetCookie(w, &cookie)
			result = "success"
			log.Println(fmt.Sprintf("DEBUG: login token:%s username:%s password:%s location:%s", tokenstring, loginuser.UserName, login.Password, login.Location))
		} else {
			result = "failed"
			expiration := time.Now()
			cookie := http.Cookie{Name: "jwt", Value: "", Expires: expiration, HttpOnly: true}
			http.SetCookie(w, &cookie)
			log.Println(fmt.Sprintf("DEBUG: login failed username:%s password:%s location:%s", loginuser.UserName, login.Password, login.Location))
		}
	}
	render.JSON(w, r, result)
}

type StringResult struct {
	Result string `boil:"result"`
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	cookie := http.Cookie{Name: "jwt", Value: "", Expires: expiration, HttpOnly: true}
	http.SetCookie(w, &cookie)
}

func UserName(w http.ResponseWriter, r *http.Request) {
	userid := auth.DecodeJWT("user_id", r).(string)
	var loginuser LoginUser
	queries.Raw("SELECT * FROM GetUserLoginData(@id)", sql.Named("id", userid)).Bind(ctx, db, &loginuser)
	render.JSON(w, r, loginuser.UserName)
}

func UserCode(w http.ResponseWriter, r *http.Request) {
	userid := auth.DecodeJWT("user_id", r).(string)
	var loginuser LoginUser
	queries.Raw("SELECT * FROM GetUserLoginData(@id)", sql.Named("id", userid)).Bind(ctx, db, &loginuser)
	render.JSON(w, r, loginuser.Userid)
}
func GetUserPrivilages(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	userid := auth.DecodeJWT("user_id", r).(string)
	err := queries.Raw("SELECT( SELECT * From GetUserPrivilages(@UserID) FOR JSON PATH ,INCLUDE_NULL_VALUES, ROOT('Privilages') ) AS result ", sql.Named("UserID", userid)).Bind(ctx, db, &result)
	if err != nil {
		log.Println(err.Error())
	}
	render.PlainText(w, r, result.Result)
}

/*
func GetDepartmentCode(w http.ResponseWriter, r *http.Request) {
	var user models.MastUser
	userid := auth.DecodeJWT(r).(string)
	models.MastUsers(Select(models.MastUserColumns.DepartmentCode), models.MastUserWhere.LoginID.EQ(userid)).Bind(ctx, db, &user)
	render.JSON(w, r, user.DepartmentCode)
}

func UserPrivilages(w http.ResponseWriter, r *http.Request) {
	userid := auth.DecodeJWT(r).(string)
	var result StringResult
	var user models.MastUser
	models.MastUsers(Select(models.MastUserColumns.UserCode), models.MastUserWhere.LoginID.EQ(userid)).Bind(ctx, db, &user)
	queries.Raw(`SELECT get_security_rights($1,$2) AS result;`, strings.TrimSpace(user.UserCode), "MAST").Bind(ctx, db, &result)
	render.JSON(w, r, result.Result)
}

//  For get all users
func SPGetAllUsers(w http.ResponseWriter, r *http.Request) {
	var result StringResult
	queries.Raw(`SELECT get_all_users() AS result;`).Bind(ctx, db, &result)
	render.PlainText(w, r, result.Result)
}

// Find a user
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.MastUser
	param := r.URL.Query().Get( "user_code")

	list := queries.Raw("select * from mast_users where user_code=$1", param).Bind(ctx, db, &user)
	json.NewDecoder(r.Body).Decode(&list)
	render.JSON(w, r, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// var user models.MastUser
	// models.MastUsers(Select(models.MastUserColumns.UserCode), models.MastUserWhere.LoginID.EQ(userid)).Bind(ctx, db, &user)
	// queries.Raw(`SELECT get_security_rights($1,$2) AS result;`, strings.TrimSpace(user.UserCode), "MAST").Bind(ctx, db, &result)
	// render.JSON(w, r, result.Result)
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// For Create a New user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	User := &models.MastUser{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&User)
	// User.CreateID = "Admin"
	// Here get password and generate hash key
	gethashpassword, err := GeneratehashPassword(User.PasswordHash)
	// Here we have to push to Hash password to DB
	User.PasswordHash = gethashpassword

	fmt.Println(err, gethashpassword, "Hello Passaworx")

	User.Insert(ctx, db, boil.Infer())

	render.JSON(w, r, User)
	json.NewEncoder(w).Encode(err)
	// insert
	if err != nil {
		fmt.Println(err)
	}

	// var user models.MastUser
	// // var result string
	// // models.MastUsers(Select(models.MastUserColumns.UserCode), models.MastUserWhere.LoginID.EQ(userid)).Bind(ctx, db, &user)
	// queries.Raw(`INSERT create_user($1,$2,$3,$4,$5) AS result;`, strings.TrimSpace(user.Email), strings.TrimSpace(user.UserName), strings.TrimSpace(user.CompanyCode), strings.TrimSpace(user.PasswordHash), "MAST").Bind(ctx, db, &result)
	// // render.JSON(w, r, result.Result) login_id,user_name,password_hash,company_code,email , ,
	// // login_id,user_name,password_hash,company_code,email

}

//Update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
*/

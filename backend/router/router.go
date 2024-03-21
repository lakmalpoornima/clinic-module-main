package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	controllers "systolic/controllers"
	auth "systolic/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

func SetupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(auth.EncodeJWT("user_id", "Hi")))
	})

	r.Route("/api", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Get("/login", controllers.UserLogin)
			r.Post("/login", controllers.UserLogin)
			r.Get("/logout", controllers.UserLogout)
			r.Post("/logout", controllers.UserLogout)
			r.Get("/user/locations", controllers.GetLoginLocationForUser)
			r.Route("/display", func(r chi.Router) {
				r.Get("/videos/{file}", func(w http.ResponseWriter, r *http.Request) {
					file := chi.URLParam(r, "file")
					http.ServeFile(w, r, "./videos/"+file+".mp4")
				})
				r.Get("/videolist", func(w http.ResponseWriter, r *http.Request) {
					files, err := ioutil.ReadDir("./videos")
					if err != nil {
						log.Fatal(err)
					}
					var videofiles []string
					for _, f := range files {
						name := strings.Split(strings.TrimSpace(f.Name()), ".")
						ext := name[len(name)-1]
						if ext == "mp4" {
							videofiles = append(videofiles, f.Name())
						}
					}
					if err != nil {
						log.Panicln(err.Error())
					}
					render.JSON(w, r, videofiles)
				})
				r.Get("/appointments", controllers.GetDashboardAppointments)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.Jwt.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/verify", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(fmt.Sprintf("hi %s", auth.DecodeJWT("user_id", r))))
			})
			// r.Get("/securityprivilages", controllers.UserPrivilages)

			r.Get("/username", controllers.UserName)
			r.Get("/usercode", controllers.UserCode)

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(fmt.Sprintf("%s", auth.DecodeJWT("user_id", r))))
			})
			r.Route("/security", func(r chi.Router) {
				r.Get("/GetUserPrivilages", controllers.GetUserPrivilages)
			})
			r.Route("/reception", func(r chi.Router) {
				r.Get("/patients", controllers.GetPatientInfo)
				// Url {AppNo}/{ClinicDate}/{clinic}/{ReferenceNo}
				r.Get("/GetClinicappointments", controllers.GetClinicAppointments)
				r.Post("/patient", controllers.SaveReceptionPatient)
				r.Get("/GetPatientRegistrationInfo", controllers.GetPatientRegistrationInfo)
			})
			// GetPatientRegistrationInfo
			r.Route("/doctorportal", func(r chi.Router) {
				r.Get("/appointments", controllers.GetDoctorPortalAppointments)
				r.Get("/patientactivityinfo", controllers.GetPatientActivityInfo)
			})

			r.Route("/patient-biography/lab", func(r chi.Router) {
				r.Get("/LabResultSummary", controllers.GetLabResultSummary)
				r.Get("/LabResultDetail", controllers.GetLabResultDetail)
				r.Get("/LabResultDetailExternal", controllers.GetLabResultDetailExternal)
				r.Get("/LabResultDetailElement", controllers.GetLabResultDetailElement)
			})
			r.Route("/report", func(r chi.Router) {
				r.Get("/SummaryReport", controllers.GetSummaryReport)
			})
			r.Route("/dashboard", func(r chi.Router) {
				r.Get("/GetDashboardCount", controllers.GetDashboardCount)
			})

			r.Route("/Clinic-notes", func(r chi.Router) {
				r.Get("/RegistrationWiseClinicSessions", controllers.GetRegistrationWiseClinicSessions)
				r.Get("/LastWiseClinicSessions", controllers.GetLastWiseClinicSessions)
				r.Get("/GetDrugAutoComplete", controllers.DrugAutoComplete)
				r.Get("/GetDocumentHeaderNotes", controllers.GetDocumentHeaderNotes)
				r.Get("/GetDoctorList", controllers.GetDoctorList)
				r.Get("/GetClinicWiseElements", controllers.GetClinicWiseElements)
				r.Get("/GetDocumentHeaderProblemList", controllers.GetDocumentHeaderProblemList)
				r.Get("/GetDocumentHeaderImageList", controllers.GetDocumentHeaderImageList)
				// 2022-03-02 GetDocumentHeaderProblemList
				r.Get("/GetDocumentPrescription", controllers.GetDocumentPrescription)
				r.Get("/GetDocumentInitialScreening", controllers.GetDocumentInitialScreening)
				r.Get("/GetDocumentReferralLetterTemplate", controllers.GetDocumentReferralLetterTemplate)
				r.Get("/GetDocumentReferralLetter", controllers.GetDocumentReferralLetter)
				r.Get("/GetDocumentCheckedList", controllers.GetDocumentCheckedList)
				r.Get("/GetDocumentPatientCheckedList", controllers.GetDocumentPatientCheckedList)
				r.Get("/GetPatientDiseaseOfImage", controllers.GetPatientDiseaseOfImage)
				r.Get("/GetTodayDocumentNumber", controllers.GetTodayDocNumber)
				// r.Post("/SaveProfileImage", controllers.Saveimage)
				r.Post("/SaveClinicHeader", controllers.SaveClinicHeader)
				r.Post("/SaveClinicSpecialNotes", controllers.SaveClinicSpecialNotes)
				r.Post("/SaveClinicInitScreening", controllers.SaveClinicInitScreening)
				r.Post("/SaveClinicCheckList", controllers.SaveClinicCheckList)
				r.Post("/SaveClinicReferral", controllers.SaveClinicReferral)
				r.Post("/SaveClinicReferralTemplate", controllers.SaveClinicReferralTemplate)
				r.Post("/SaveClinicPrescription", controllers.SaveClinicPrescription)
				r.Post("/SaveClinicImageDisease", controllers.SaveClinicImageDisease)
				r.Post("/SaveProfileImage", controllers.SaveProfileImage)
				r.Post("/AddNewProblem", controllers.AddNewProblem)
				r.Post("/SaveClinicExternalLabData", controllers.SaveClinicExternalLabData)
				r.Get("/GetPatientProfileImage", controllers.GetPatientProfileImage)
				// Dosage /GetPatientDiseaseOfImage
				r.Get("/GetConsultantInfo", controllers.GetConsultantInfo)
				r.Get("/GetDosageList", controllers.GetDosageList)
				r.Get("/GetUnitList", controllers.GetUnitList)
				r.Get("/GetDiagnosticList", controllers.GetDiagnosticList)
				//Druglist
				r.Get("/GetFilterDruglist", controllers.GetFilterDrugslist)
				// Filter The LabTestWithFilter
				r.Get("/GetFilterLabInvestlist", controllers.GetFilterLablist)
				// Filter The OtherInvest
				r.Get("/GetFilterOtherlist", controllers.GetFilterOtherInvestigationlist)
				r.Get("/GetSpecialNotes", controllers.GetSpecialNotes)

			})
			// Clinic notes
		})

	})

}

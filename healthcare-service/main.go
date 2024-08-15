package main

import (
	"context"
	"fmt"
	"healthcare/clients"
	"healthcare/handlers"
	"healthcare/middleware"
	"healthcare/repository"
	"healthcare/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// env
	port := os.Getenv("PORT")
	universityHost := "university-service"
	universityPort := "8080"
	authServiceURL := fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVICE_URL"), os.Getenv("AUTH_SERVICE_PORT"))
	// client
	customClient := http.DefaultClient
	universityClient := clients.NewUniversityClient(universityHost, universityPort, customClient)
	// MongoService initialization
	mongoService, err := services.NewMongoService(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	healthcareRepository, err := repository.NewHealthcareRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	healthcareService, err := services.NewHealthcareService(healthcareRepository, universityClient)
	if err != nil {
		log.Fatalln(err)
	}
	departmentRepository, err := repository.NewDepartmentRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	departmentService, err := services.NewDepartmentService(departmentRepository, universityClient, healthcareService)
	if err != nil {
		log.Fatalln(err)
	}
	healthcareHandler, err := handlers.NewHealthcareHandler(healthcareService, departmentService)
	if err != nil {
		log.Fatalln(err)
	}

	// ROUTING
	router := mux.NewRouter()

	router.HandleFunc("/ping", healthcareHandler.Ping).Methods("GET")
	router.HandleFunc("/createRecord/{id}", middleware.ValidateJWT(middleware.ValidateRole(healthcareHandler.CreateRecordForUser, "Doctor"), authServiceURL)).Methods("POST")
	router.HandleFunc("/createRecords/{id}", healthcareHandler.CreateRecordForUser).Methods("POST")
	router.HandleFunc("/getCertificate/{id}", healthcareHandler.GetCertificateForUser).Methods("GET")

	//
	router.HandleFunc("/appointments/{id}", healthcareHandler.GetAppointmentsByDoctorID).Methods("GET")
	router.HandleFunc("/departments", healthcareHandler.GetAllDepartments).Methods("GET")
	router.HandleFunc("/department/create/{name}", healthcareHandler.CreateDepartment).Methods("POST")
	router.HandleFunc("/department/{name}", healthcareHandler.GetDepartmentByName).Methods("GET")
	router.HandleFunc("/department/{name}/docSchedule/{date}/free", healthcareHandler.GetFreeDoctorSlots).Methods("GET")
	router.HandleFunc("/department/{name}/docSchedule/{date}/add", healthcareHandler.AddDoctorToSchedule).Methods("POST")
	router.HandleFunc("/department/{name}/schedule/{date}/free", healthcareHandler.GetFreeSlots).Methods("GET")
	router.HandleFunc("/department/{name}/schedule/{date}/add", healthcareHandler.AddPatientToSchedule).Methods("POST")
	router.HandleFunc("/records", healthcareHandler.GetAllRecords).Methods("GET")
	router.HandleFunc("/records/{id}", healthcareHandler.GetRecordForUser).Methods("GET")
	router.HandleFunc("/records/{id}/createCertificate", healthcareHandler.CreateCertificateForUser).Methods("POST")
	router.HandleFunc("/records/{id}/certificate", healthcareHandler.GetCertificateForUser).Methods("GET")
	router.HandleFunc("/records/{id}/appointments", healthcareHandler.GetAppointmentsByPatientID).Methods("GET")
	router.HandleFunc("/records/{id}/appointments/{appId}/get", healthcareHandler.GetAppointmentById).Methods("GET")
	router.HandleFunc("/records/{id}/appointments/{appId}/update", healthcareHandler.UpdateAppointment).Methods("POST")
	router.HandleFunc("/records/{id}/referrals", healthcareHandler.GetReferralsByPatientID).Methods("GET")
	router.HandleFunc("/records/{id}/referrals/createReferral", healthcareHandler.CreateReferral).Methods("POST")
	router.HandleFunc("/records/{id}/referrals/{refId}", healthcareHandler.GetReferralById).Methods("GET")
	router.HandleFunc("/records/{id}/prescriptions", healthcareHandler.GetPrescriptionsByPatientID).Methods("GET")
	router.HandleFunc("/records/{id}/prescriptions/createPrescription", healthcareHandler.CreatePrescription).Methods("POST")
	router.HandleFunc("/records/{id}/prescriptions/{presId}", healthcareHandler.GetPrescriptionById).Methods("GET")
	router.HandleFunc("/records/{id}/prescriptions/{presId}/{status}", healthcareHandler.UpdatePrescription).Methods("POST")

	// CORS
	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	server := http.Server{
		Addr:         ":" + port,
		Handler:      gorillaHandlers.CORS(headersOk, methodsOk)(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Println("Server listening on port", port)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Panicf("PANIC FROM AUTH-SERVICE ON LISTENING")
		}
	}()
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	log.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		log.Fatalf("Cannot gracefully shutdown...")
	}
	log.Println("Server stopped")
}

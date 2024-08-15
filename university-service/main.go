package main

import (
	"context"
	"fakultet-service/client"
	"fakultet-service/config"
	"fakultet-service/handlers"
	"fakultet-service/middleware"
	"fakultet-service/repository"
	"fakultet-service/services"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// env
	port := os.Getenv("PORT")

	healthCareServiceURL := os.Getenv("HEALTHCARE_SERVICE_URL")
	healthCareServicePort := os.Getenv("HEALTHCARE_SERVICE_PORT")
	//	authServicePort := os.Getenv("AUTH_SERVICE_PORT")
	authServiceURL := fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVICE_URL"), os.Getenv("AUTH_SERVICE_PORT"))

	// client
	customHttpClient := http.DefaultClient
	healthCareClient := client.NewHealthCareClient(healthCareServiceURL, healthCareServicePort, customHttpClient)
	authServiceClient := client.NewAuthServiceClient(authServiceURL, customHttpClient)
	// MongoService initialization
	mongoService, err := services.NewMongoService(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	universityRepository, err := repository.NewUniversityRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	universityService, err := services.NewUniversityService(universityRepository, healthCareClient, authServiceClient)
	if err != nil {
		log.Fatalln(err)
	}
	universityHandler, err := handlers.NewUniversityHandler(universityService)
	if err != nil {
		log.Fatalln(err)
	}

	logger := config.NewLogger("./logs/log.log")
	router := mux.NewRouter()

	router.HandleFunc("/ping", universityHandler.Ping).Methods("GET")
	router.HandleFunc("/", universityHandler.CreateUniversity).Methods("POST")
	router.HandleFunc("/", universityHandler.FindAllUniversities).Methods("GET")
	router.HandleFunc("/student", universityHandler.CreateStudent).Methods("POST")
	router.HandleFunc("/student/{id}", universityHandler.FindStudentById).Methods("GET")
	router.HandleFunc("/student/budget/{id}", universityHandler.CheckBudget).Methods("GET")
	router.HandleFunc("/student/status/{id}", middleware.ValidateJWT(middleware.ValidateRole(universityHandler.ExtendStatus, "Student"), authServiceURL)).Methods("PUT")
	router.HandleFunc("/student/status1/{id}", universityHandler.ExtendStatus).Methods("PUT")
	router.HandleFunc("/professor", universityHandler.CreateProfessor).Methods("POST")
	router.HandleFunc("/professor/{id}", universityHandler.FindProfessorById).Methods("GET")
	router.HandleFunc("/scholarship", universityHandler.CreateScholarship).Methods("POST")
	router.HandleFunc("/professor/{id}", universityHandler.DeleteProfessor).Methods("DELETE")
	router.HandleFunc("/student/{id}", universityHandler.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/scholarship/{id}", universityHandler.DeleteScholarship).Methods("DELETE")
	router.HandleFunc("/stateApplication", universityHandler.CreateStateExamApplication).Methods("POST")
	router.HandleFunc("/entranceExam", universityHandler.CreateEntranceExam).Methods("POST")
	router.HandleFunc("/entranceExam", universityHandler.FindAllEntranceExams).Methods("GET")
	router.HandleFunc("/extendStatusApplication", universityHandler.CreateExtendStatusApplication).Methods("POST")
	router.HandleFunc("/extendStatusApplication", universityHandler.FindAllExtendStatusApplications).Methods("GET")
	router.HandleFunc("/scholarshipApplication", universityHandler.CreateScholarshipApplication).Methods("POST")
	router.HandleFunc("/scholarshipApplication", universityHandler.FindAllScholarshipApplications).Methods("GET")
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

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal("Error while server is listening and serving requests", log.Fields{
				"module": "server-main",
				"error":  err.Error(),
			})
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)
	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Error during graceful shutdown", log.Fields{
			"module": "server-main",
		})
	}
	logger.LogInfo("server-main", "Server shut down")
}

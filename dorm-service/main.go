package main

import (
	"context"
	"dorm-service/client"
	"dorm-service/handlers"
	"dorm-service/repositories"
	"dorm-service/services"
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
	healthCareServiceURL := os.Getenv("HEALTHCARE_SERVICE_URL")
	healthCareServicePort := os.Getenv("HEALTHCARE_SERVICE_PORT")

	// client
	customHttpClient := http.DefaultClient
	healthCareClient := client.NewHealthCareClient(healthCareServiceURL, healthCareServicePort, customHttpClient)
	// services
	mongoService, err := services.NewMongoService(timeoutContext)
	if err != nil {
		log.Fatalln(err)
	}
	dormRepository, err := repositories.NewDormRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	dormService, err := services.NewDormService(dormRepository)
	if err != nil {
		log.Fatalln(err)
	}
	dormHandler, err := handlers.NewDormHandler(dormService)
	if err != nil {
		log.Fatalln(err)
	}
	admissionsRepository, err := repositories.NewAdmissionsRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	admissionsService, err := services.NewAdmissionsServices(admissionsRepository)
	if err != nil {
		log.Fatalln(err)
	}
	admissionsHandler, err := handlers.NewAdmissionsHandler(admissionsService)
	if err != nil {
		log.Fatalln(err)
	}
	applicationsRepository, err := repositories.NewApplicationsRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	applicationsServices, err := services.NewApplicationsService(applicationsRepository, healthCareClient)
	if err != nil {
		log.Fatalln(err)
	}
	applicationsHandler, err := handlers.NewDormApplicationHandler(applicationsServices)
	if err != nil {
		log.Fatalln(err)
	}
	// ROUTING
	router := mux.NewRouter()
	router.HandleFunc("/ping", dormHandler.Ping).Methods("POST")
	router.HandleFunc("/", dormHandler.CreateNewDorm).Methods("POST")
	router.HandleFunc("/{id}", dormHandler.FindDormById).Methods("GET")
	router.HandleFunc("/{id}", dormHandler.DeleteDormById).Methods("DELETE")
	router.HandleFunc("/{id}", dormHandler.UpdateDormById).Methods("PUT")
	router.HandleFunc("/admissions", admissionsHandler.CreateNewAdmission).Methods("POST")
	router.HandleFunc("/admissions/{id}", admissionsHandler.GetAdmissionsByID).Methods("GET")
	router.HandleFunc("/admissions/{id}", admissionsHandler.DeleteAdmissionById).Methods("GET")
	router.HandleFunc("/admissions/dorm/{id}", admissionsHandler.GetAdmissionByDormId).Methods("GET")
	router.HandleFunc("/applications", applicationsHandler.CreateNewApplication).Methods("POST")
	// CORS
	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"})
	server := http.Server{
		Addr:         ":" + port,
		Handler:      gorillaHandlers.CORS(headersOk, methodsOk, originsOk)(router),
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

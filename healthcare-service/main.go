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
	universityPort := "8000"
	authServiceURL := fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVICE_URL"), os.Getenv("AUTH_SERVICE_PORT"))
	// client
	customClient := http.DefaultClient
	universityClient := clients.NewUnivesityClient(universityHost, universityPort, customClient)
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
	healthcareHandler, err := handlers.NewHealthcareHandler(healthcareService)
	if err != nil {
		log.Fatalln(err)
	}

	// ROUTING
	router := mux.NewRouter()
	router.HandleFunc("/ping", healthcareHandler.Ping).Methods("GET")
	router.HandleFunc("/createRecord/{id}", middleware.ValidateJWT(middleware.ValidateRole(healthcareHandler.CreateRecordForUser, "Doctor"), authServiceURL)).Methods("POST")
	router.HandleFunc("/createRecords/{id}", healthcareHandler.CreateRecordForUser).Methods("POST")
	router.HandleFunc("/getRecord/{id}", healthcareHandler.GetRecordForUser).Methods("GET")
	router.HandleFunc("/createCertificate", healthcareHandler.CreateCertificateForUser).Methods("POST")
	router.HandleFunc("/getCertificate/{id}", healthcareHandler.GetCertificateForUser).Methods("GET")

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

package main

import (
	"context"
	"fakultet-service/config"
	"fakultet-service/handlers"
	"fakultet-service/repository"
	"fakultet-service/services"
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
	// MongoService initialization
	mongoService, err := services.NewMongoService(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	universityRepository, err := repository.NewUniversityRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	universityService, err := services.NewUniversityService(universityRepository)
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
	router.HandleFunc("/student", universityHandler.CreateStudent).Methods("POST")
	router.HandleFunc("/student/{id}", universityHandler.FindStudentById).Methods("GET")
	router.HandleFunc("/student/budget/{id}", universityHandler.CheckBudget).Methods("GET")

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

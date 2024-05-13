package main

import (
	"context"
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
	// ROUTING
	router := mux.NewRouter()
	router.HandleFunc("/ping", dormHandler.Ping).Methods("POST")
	router.HandleFunc("/", dormHandler.CreateNewDorm).Methods("POST")
	router.HandleFunc("/{id}", dormHandler.FindDormById).Methods("GET")
	router.HandleFunc("/{id}", dormHandler.DeleteDormById).Methods("DELETE")
	router.HandleFunc("/{id}", dormHandler.UpdateDormById).Methods("PUT")
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

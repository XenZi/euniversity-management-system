package main

import (
	"context"
	"food/client"
	"food/handlers"
	"food/repository"
	"food/services"
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
	// MongoService initialization
	universityServiceURL := os.Getenv("UNIVERSITY_SERVICE_URL")
	universityServicePort := os.Getenv("UNIVERSITY_SERVICE_PORT")

	log.Println("UNI SERVICE URL: " + universityServiceURL)
	log.Println("UNI SERVICE PORT: " + universityServicePort)

	//Client
	customHttpClient := http.DefaultClient
	universityClient := client.NewUniversityClient(universityServiceURL, universityServicePort, customHttpClient)
	mongoService, err := services.NewMongoService(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	foodRepository, err := repository.NewFoodRepository(mongoService.GetCLI())
	if err != nil {
		log.Fatalln(err)
	}
	foodService, err := services.NewFoodCardService(foodRepository, universityClient)
	if err != nil {
		log.Fatalln(err)
	}
	foodHandler, err := handlers.NewFoodHandler(foodService)
	if err != nil {
		log.Fatalln(err)
	}

	// ROUTING
	router := mux.NewRouter()
	router.HandleFunc("/ping", foodHandler.Ping).Methods("GET")
	router.HandleFunc("/deleteSupplier/{id}", foodHandler.DeleteSupplierById).Methods("POST")
	router.HandleFunc("/allSuppliers", foodHandler.GetAllSuppliers).Methods("GET")
	router.HandleFunc("/createSupplier", foodHandler.CreateSupplier).Methods("POST")
	router.HandleFunc("/updateMess/{id}", foodHandler.UpdateMessRoom).Methods("POST")
	router.HandleFunc("/deleteFoodCard/{id}", foodHandler.DeleteFoodCard).Methods("POST")
	router.HandleFunc("/deleteMessRoom/{id}", foodHandler.DeleteMessRoom).Methods("POST")
	router.HandleFunc("/allMessRooms", foodHandler.GetAllMessRooms).Methods("GET")
	router.HandleFunc("/createMessRoom", foodHandler.CreateMessRoom).Methods("POST")
	router.HandleFunc("/createFoodCard", foodHandler.CreateFoodCard).Methods("POST")
	router.HandleFunc("/allFoodCards", foodHandler.GetAllFoodCards).Methods("GET")
	router.HandleFunc("/createPayment", foodHandler.CreatePayment).Methods("POST")
	router.HandleFunc("/payForMeal/{id}", foodHandler.PayForMeal).Methods("POST")

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
			log.Panicf("PANIC FROM FOOD-SERVICE ON LISTENING")
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

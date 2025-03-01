package main

import (
	"goCrud/goControllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Initalize the Gorilla Mux router
	router := mux.NewRouter()

	//Call services to define routes
	goControllers.SetupUserRoutes(router)

	// Enable CORS middleware
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"()"}))

	log.Fatal(http.ListenAndServe(":8080", cors(router)))
	/*
		// Load data from the JSON file
		dataPath := "../database/data.json"
		users, err := goHandlers.LoadData(dataPath)
		if err != nil {
			log.Fatal(err)
		}

		// Print the loaded data
		fmt.Println("Loaded users:")
		for _, user := range users {
			fmt.Printf("%d: %s (%s)\n", user.Id, user.Name, user.Email)
		}

		// You can also modify the data and save it back to the JSON file
		newUser := goHandlers.User{
			Id:    3,
			Name:  "Alice Johnson",
			Email: "alice@example.com",
		}
		users = append(users, newUser)

		// Save the modified data back to the JSON file
		if err := goHandlers.SaveData(dataPath, users); err != nil {
			log.Fatal(err)
		}
	*/
}

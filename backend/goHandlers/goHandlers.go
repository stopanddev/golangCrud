package goHandlers

import (
	"encoding/json"
	"os"
)

type User struct {
	Id    int    `json:id`
	Name  string `json:name`
	Email string `json:color`
}

// Function to load data from a JSON file
func LoadDataHandler(filename string) ([]User, error) {
	// Read the file content using os.ReadFile (instead of ioutil.ReadAll)
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of User structs
	var users []User
	if err := json.Unmarshal(fileContent, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// Function to save data to a JSON file
func SaveDataHandler(filename string, users User) error {
	// Marshal the users slice into JSON format
	fileContent, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	// Write the JSON data to the file using os.WriteFile (instead of ioutil.WriteFile)
	err = os.WriteFile(filename, fileContent, 0644)
	if err != nil {
		return err
	}
	return nil
}

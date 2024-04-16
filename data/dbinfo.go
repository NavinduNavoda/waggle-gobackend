package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// DBConfig represents the JSON structure for database configuration
type DBConfig struct {
	ConnectionURL string `json:"connection_url"`
	Password      string `json:"password"`
	Username      string `json:"username"`
}

const dbInfoFilePath = "./data/dbinfo.json"

func TestDbInfo() {
	connectionInfo, err := ReadDBInfoFromFile()
	if err != nil {
		fmt.Println("Error reading DB info:", err)
		return
	}

	fmt.Println("Connection URL:", connectionInfo.ConnectionURL)
	fmt.Println("Password:", connectionInfo.Password)
	fmt.Println("Username:", connectionInfo.Username)
}

// ReadDBInfoFromFile reads the database connection info from the JSON file
func ReadDBInfoFromFile() (DBConfig, error) {
	file, err := os.Open(dbInfoFilePath)
	if err != nil {
		return DBConfig{}, err
	}
	defer file.Close()

	var config DBConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return DBConfig{}, err
	}

	return config, nil
}

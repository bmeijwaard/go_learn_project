package config

import (
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
	Server    string
	Port      int
	User      string
	Pass      string
	Database  string
}

func GetConfig() Configuration {
	file, _ := os.Open("./conf.json")
	//defer file.Close()

	decoder := json.NewDecoder(file)
	fmt.Println("File: ", file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("error:", err)
	}
	return configuration
}


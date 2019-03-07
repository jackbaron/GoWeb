package main

import (
	"encoding/json"
	"os"
	"projects/blog/dataservice"
	"projects/blog/helpers"
	"projects/blog/routers"
)

func main() {
	// Conect DB here
	// Load the configuration file
	helpers.Load("dataservice"+string(os.PathSeparator)+"configDB.json", config)
	// Configure the session cookie store
	helpers.Configure(config.Session)
	// Connect to database
	dataservice.InitDb(config.Database)
	// Run server
	start := routers.GetRouter()
	if !start {
		panic("error. Cannot start server")
	}

}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database dataservice.Info `json:"Database"`
	Session  helpers.Session  `json:"Session"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

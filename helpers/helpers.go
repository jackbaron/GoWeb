package helpers

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
* Write file log
 */
func WiteLog(prefix string, text string) {
	f, err := os.OpenFile("log/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, prefix, log.LstdFlags)
	logger.Printf("<========= Start log " + prefix + "=========>")
	logger.Printf(text)
	logger.Printf("<========= End log " + prefix + "=========>")
}

// Handling json file

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file
func Load(configFile string, p Parser) {
	var err error
	var absPath string
	var input = io.ReadCloser(os.Stdin)
	if absPath, err = filepath.Abs(configFile); err != nil {
		log.Fatalln(err)
	}

	if input, err = os.Open(absPath); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}

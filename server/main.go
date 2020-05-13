package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"os/user"
	"strings"

	"gopkg.in/yaml.v2"
)

// Configuration - the configuration file structure.
type Configuration struct {
	Shell string
}

// The global configuration.
var configuration = Configuration {
	Shell: "bash", // default is bash
}

// APIResponse - The API response to request.
type APIResponse struct {
	Message string `json:"message"`
}

// Returns the JSON of the response.
func (response APIResponse) json() string {
	responseAsJSON, _ := json.Marshal(&response)
	return string(responseAsJSON)
}

// RunCommand - runs a command inside a terminal emulator.
func RunCommand(command string, terminal string, shell string) bool {
	currentUser, error := user.Current()
	errored := true

	if error == nil {
		terminalArguments := strings.Fields("--working-directory=" + currentUser.HomeDir + " -- " + shell + " -c ")
		terminalArguments = append(terminalArguments, command + ";" + shell)
		
		error := exec.Command(terminal, terminalArguments...).Run()
		errored = error != nil
	}

	return errored
}

// HandleCommandRun - Handles GET /run requests. 
func HandleCommandRun(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		commandToRun := request.URL.Query().Get("command")
		message := "SUCCESS"
		
		if commandToRun != "" {
			fmt.Printf("$ %s\n", commandToRun)
			errored := RunCommand(commandToRun, "gnome-terminal", configuration.Shell)
		
			if errored {
				message = "ERROR"
			}
		} else {
			message = "MISSING_COMMAND_QUERY"
		}
		
		response := APIResponse { 
			Message: message,
		}
	
		fmt.Println("> " + response.json())
		fmt.Fprintf(responseWriter, response.json())
	} else if request.Method == "OPTIONS" {
		responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		responseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

// RaiseErrorAndUseDefaultConfiguration - raises an error and returns the configuration file.
func RaiseErrorAndUseDefaultConfiguration(filePath string, error error) Configuration {
	fmt.Println("Cannot load configuration file at " + filePath + ".")
	fmt.Println("    " + error.Error())
	fmt.Println("")
	fmt.Println("Using default configuration...")

	return configuration
}

// LoadConfigurationFile - loads the configuration file.
func LoadConfigurationFile(filePath string) Configuration {
	fileContents, readError := ioutil.ReadFile(filePath)

	if readError != nil {
		return RaiseErrorAndUseDefaultConfiguration(filePath, readError)
	}

	yamlError := yaml.Unmarshal(fileContents, &configuration)

	if yamlError != nil {
		return RaiseErrorAndUseDefaultConfiguration(filePath, yamlError)
	}

	return configuration
}

func main() {
	configuration = LoadConfigurationFile("shellrun.config.yml")

	fmt.Println("Shell: " + configuration.Shell)

	fmt.Println("The server will be listening on port 2727.")
	http.HandleFunc("/run", HandleCommandRun)
	http.ListenAndServe(":2727", nil)
}
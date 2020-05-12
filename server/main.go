package main

import (
	"fmt"
	"strings"
	"os/exec"
	"os/user"
	"net/http"
	"encoding/json"
)

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
			errored := RunCommand(commandToRun, "gnome-terminal", "zsh")
		
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

func main() {
	fmt.Println("The server will be listening on port 2727.")

	http.HandleFunc("/run", HandleCommandRun)
	http.ListenAndServe(":2727", nil)
}
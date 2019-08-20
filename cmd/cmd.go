package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// frontendConnection stores all data necessary for a successful connection to a Frontend
type frontendConnection struct {
	FrontendIP string
	Username   string
	Password   string
	Csrftoken  http.Cookie
	SessionID  http.Cookie
}

// {"API Error": "Command Not Found"}
// successful unmarshalling into this struct indicates an error from the Stacki API
type stackiError struct {
	Error string `json:"API Error"`
}

// successful unmarshalling into this struct indicates an empty response from the Stacki API
type emptyOutput struct {
	Output string `json:"output"`
}

// frontendAuth will authenticate with a Stacki Frontend and return a frontendConnection struct
func frontendAuth() (frontendConnection, error) {

	// create a connection
	connection := frontendConnection{}
	connection.FrontendIP = os.Getenv("STACKILACKEY_FRONTEND_IP")
	if connection.FrontendIP == "" {
		log.Fatal("No Frontend IP available. Please run `export STACKILACKEY_FRONTEND_IP=<the ip of your Frontend>`")
	}
	connection.Username = os.Getenv("STACKILACKEY_USERNAME")
	if connection.Username == "" {
		log.Fatal("No Frontend username available. Please run `export STACKILACKEY_USERNAME=<username of your api user>`")
	}
	connection.Password = os.Getenv("STACKILACKEY_PASSWORD")
	if connection.Password == "" {
		log.Fatal("No Frontend password available. Please run `export STACKILACKEY_PASSWORD=<password of your api user>`")
	}

	// hit this URL first to get a valid csrftoken to use for this session
	initialURL := "http://" + connection.FrontendIP + "/stack"

	// begin forming a GET request
	initialRequest, err := http.NewRequest("GET", initialURL, nil)
	if err != nil {
		return frontendConnection{}, err
	}

	// perform the request
	initialResponse, err := http.DefaultClient.Do(initialRequest)
	if err != nil {
		return frontendConnection{}, err
	}

	// retrieve the csrftoken from the Set-Cookie Header
	Token := initialResponse.Header.Get("Set-Cookie")
	Token = strings.Split(Token, "csrftoken=")[1]
	Token = strings.Split(Token, ";")[0]

	// TODO: validate token
	// this URL will give us a session cookie that we can use to perform authenticated API calls
	loginURL := "http://" + connection.FrontendIP + "/stack/login"

	// the csrf token from the previous request needs to be used twice, first as a cookie
	connection.Csrftoken = http.Cookie{
		Name:  "csrftoken",
		Value: Token,
	}

	// the Frontend expects the username and password to be in the body of the POST
	payload := strings.NewReader("USERNAME=" + connection.Username + "&PASSWORD=" + connection.Password)

	// begin forming the login post
	loginPost, err := http.NewRequest("POST", loginURL, payload)
	if err != nil {
		return frontendConnection{}, err
	}

	// add the required cookies and headers
	loginPost.AddCookie(&connection.Csrftoken)
	loginPost.Header.Add("X-CSRFToken", Token)
	loginPost.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// perform the post
	loginResponse, err := http.DefaultClient.Do(loginPost)
	if err != nil {
		return frontendConnection{}, err
	}

	// if we don't get a 200, something when wrong
	if loginResponse.StatusCode != 200 {
		return frontendConnection{}, errors.New("Unable to authenticate with the Frontend, http error " + fmt.Sprint(loginResponse.StatusCode))
	}

	// we have successfully authenticated with the Stacki Frontend, return the sessionid and csrftoken
	// retrieve the csrftoken from the Set-Cookie Header
	SessionID := fmt.Sprint(loginResponse.Cookies())
	SessionID = strings.Split(SessionID, "sessionid=")[1]
	SessionID = strings.Split(SessionID, ";")[0]

	// add the session cookie to the connection struct
	connection.SessionID = http.Cookie{
		Name:  "sessionid",
		Value: SessionID,
	}
	return connection, nil
}

// RunCommand will run a command on the Stacki Frontend and return the json response as a []byte
func RunCommand(command string) ([]byte, error) {

	// authenticate with the Frontend
	connection, err := frontendAuth()
	if err != nil {
		log.Fatal(err)
	}

	APIEndpoint := "http://" + connection.FrontendIP + "/stack"

	// the command to send to the Frontend, for example {"cmd":"list host"}
	payload := strings.NewReader("{\"cmd\": \"" + command + "\"}")

	// begin crafting the request
	APICall, err := http.NewRequest("POST", APIEndpoint, payload)
	if err != nil {
		return nil, err
	}

	// add the required cookies and headers
	APICall.AddCookie(&connection.Csrftoken)
	APICall.AddCookie(&connection.SessionID)
	APICall.Header.Add("X-CSRFToken", connection.Csrftoken.Value)
	APICall.Header.Add("Content-Type", "application/json")

	// make the request
	APICallResponse, err := http.DefaultClient.Do(APICall)
	if err != nil {
		return nil, err
	}

	// get the json response out of the body
	body, err := ioutil.ReadAll(APICallResponse.Body)
	if err != nil {
		return nil, err
	}

	apiError := checkError(body)

	return body, apiError

}

// checkError will check the Frontend API response for errors
func checkError(APIResponse []byte) error {

	// first check if it is empty (not an error)
	emptyOutput := emptyOutput{}
	err := json.Unmarshal(APIResponse, &emptyOutput)
	if err != nil {
		// this is not an empty response
		return nil
	}

	// occasionally we get unpredictable garbage. for example, a successful "remove host" returns "{}"
	if string(APIResponse) == "{}" {
		return nil
	}

	// if it can be unmarshalled into our error struct, then it is an error
	apiError := stackiError{}
	err = json.Unmarshal(APIResponse, &apiError)
	if err != nil {
		// this is not an API error, return the sucessful response
		return nil
	}
	return errors.New(string(APIResponse))
}

// ArgsExpander will generate a stack command from a list of arg keys and values
func ArgsExpander(command string, argKeys []string, argValues []interface{}) (string, error) {
	// if we do not have the same number of args keys and values, were in trouble
	if len(argKeys) != len(argValues) {
		return "", fmt.Errorf("unable to expand args for %s: incorrect number of args keys and values", command)
	}
	for i, value := range argValues {
		// we cannot assign "nil" to a string or an int in Go, and that presents us with some issues
		// we can make the assumption that there will never be a valid case to run a stack command with an empty string
		valueString := fmt.Sprintf("%s", value)
		if valueString != "" {
			command = fmt.Sprintf("%s %s='%v'", command, argKeys[i], value)
		}
	}
	return command, nil
}

// old retry code
/*
	var connection frontendConnection
	err := retry(120, func() (err error) {
		connection, err = frontendAuth()
		return
	})
	if err != nil {
		log.Fatal(err)
	}
*/
/*
func retry(timeoutSeconds int, f func() error) (err error) {
	err = f()
	if err == nil {
		return
	}

	timer := time.NewTicker(time.Second * time.Duration(30))
	defer timer.Stop()

	timeout := time.After(time.Second * time.Duration(timeoutSeconds))

	for {
		select {
		case <-timer.C:
			err = f()
			if err == nil {
				return
			}
			//t.Log("attempt to connect to the Frontend failed... retrying")
		case <-timeout:
			return
		}
	}
}
*/

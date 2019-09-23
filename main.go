package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	APIURL = "https://dynamicdisplay.recurse.com/matrix/api/message"
)

// DynamicDisplayRequest contains the data needed to POST a message
type DynamicDisplayRequest struct {
	Name    string `json:"author"`
	Message string `json:"message"`
}

// Displays a message on RC Floor 5
// See also:
//   https://dynamicdisplay.recurse.com
func main() {

	messageToDisplay := "hello world!"
	if len(os.Args) > 1 {
		messageToDisplay = strings.Join(os.Args[1:], " ")
	}

	ddr := &DynamicDisplayRequest{
		Name:    "golang",
		Message: messageToDisplay,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(ddr)

	request, err := http.NewRequest("POST", APIURL, buffer)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("problem while posting; error: %v", err)
		return
	}

	defer response.Body.Close()

	fmt.Printf("Response Status: %v\n", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("problem while reading response; error: %v", err)
		return
	}
	fmt.Printf("Response Body: %v\n", string(body))
}

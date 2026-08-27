package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	const Username = "USERNAME"
	const Password = "PASSWORD"

	payload := map[string]interface{}{
		"source":       "gemini",
		"prompt":       "best supplements for better sleep",
		"parse":        true,
		"geo_location": "United States",
		"callback_url": "https://your-server.com/oxylabs-callback",
	}

	jsonValue, _ := json.Marshal(payload)

	client := &http.Client{}
	request, _ := http.NewRequest("POST",
		"https://data.oxylabs.io/v1/queries",
		bytes.NewBuffer(jsonValue),
	)

	request.SetBasicAuth(Username, Password)
	response, _ := client.Do(request)

	responseText, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseText))
}

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-server/internal/models"
)

func SendPostRequest() {
	url := "http://localhost:8080/people"

	data := models.Person{
		Name: "Kami",
		Age:  00,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Response:", result)
}

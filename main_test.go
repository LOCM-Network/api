package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

//create http request to get the data from the server
func TestGetData(t *testing.T) {
	//create a new http request
	req, err := http.NewRequest("GET", "http://localhost:8080/player/phuongaz", nil)
	if err != nil {
		t.Fatal(err)
	}
	//create a new http client
	client := &http.Client{}
	//make the request
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	//check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code was %d expected %d", resp.StatusCode, http.StatusOK)
	}
}

//test post data with json format
func TestPostData(t *testing.T) {
	data := map[string]string{"name": "phuongaz", "join_date": "2020-01-01", "coin": "100"}
	json_data, err := json.Marshal(data)
	//create a new http request
	resp, err := http.Post("http://localhost:8080/player/phuongaz", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
}

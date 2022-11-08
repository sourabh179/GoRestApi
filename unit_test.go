package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Testcase to test the status code for GET Method
func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/statistics", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Statistics)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	//expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	//	expected :=`[{"avg":"100.250000","count":"1","max":"100.250000","min":"100.250000","max":"100.250000","sum":"100.250000"},{"avg":"0.000000","count":"0","min":"0.000000","max":"0.000000","sum":"0.000000"}]`
	//expected := '[{"avg":"100.250000","count":"1","max":"100.250000","min":"100.250000","max":"100.250000","sum":"100.250000"}]'
	/*	expected := "{\"avg\": \"100.250000\",	\"count\": \"1\",\"max\": \"100.250000\",\"min\": \"100.250000\",\"sum\": \"100.250000\"}"
		expected2 := "{\"avg\":\"0.000000\",\"count\":\"0\",\"max\":\"0.000000\",\"min\":\"0.000000\",\"sum\":\"0.000000\"}"
		out:=rr.Body.String()
		if out != expected2 || out != expected {
			t.Errorf("handler returned unexpected body: got %v want %v or %v",
				rr.Body.String(), expected,expected2)
		}*/
	if rr.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, rr.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, rr.Code)
	}

	//expected := "{\"avg\": \"100.250000\",	\"count\": \"1\",\"max\": \"100.250000\",\"min\": \"100.250000\",\"sum\": \"100.250000\"}"
	/*expected := "{\"avg\":\"0.000000\",\"count\":\"0\",\"max\":\"0.000000\",\"min\":\"0.000000\",\"sum\":\"0.000000\"}"
	out := rr.Body.String()
	if out == expected {
		t.Logf("Expected to get status %s is same ast %s\n",rr.Body.String(), expected)

	}else {
		t.Errorf("handler returned unexpected body: got %v want %v ",
			rr.Body.String(), expected)

	}*/
}

//Test Case to test transaction
func TestCreateTransaction(t *testing.T) {

	var jsonStr = []byte(`{"amount":"100.25","timestamp":"2021-07-17T09:59:51.312Z"}`)

	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Transactions)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

//Testcase to test statistic after transaction
func TestGetStatistics(t *testing.T) {
	req, err := http.NewRequest("GET", "/statistics", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Statistics)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"avg\":\"100.250000\",\"count\":\"1\",\"max\":\"100.250000\",\"min\":\"100.250000\",\"sum\":\"100.250000\"}"
	//expected := "{\"avg\":\"0.000000\",\"count\":\"0\",\"max\":\"0.000000\",\"min\":\"0.000000\",\"sum\":\"0.000000\"}"
	out := rr.Body.String()
	if out == expected {
		t.Logf("Expected to get status %s is same ast %s\n",rr.Body.String(), expected)

	}else {
		t.Errorf("handler returned unexpected body: got %v want %v ",
			rr.Body.String(), expected)

	}
}

//Testcase to delete statistic
func TestDeleteEntries(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Delete)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

//Testcase to test invalid json
func TestInvalidJson(t *testing.T) {
	var jsonStr = []byte(`{"amount":100.25","timestamp":"2021-07-17T09:59:51.312Z"}`)
	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Transactions)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}
}

//To test 400 bad request
func TestBadRequest(t *testing.T) {
	//var jsonStr = []byte(`{"amount":100.25","timestamp":"2021-07-17T09:59:51.312Z"}`)
	req, err := http.NewRequest("POST", "/transaction", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Transactions)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

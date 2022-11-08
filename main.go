package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"net"
	"net/http"
	"time"
)

type ApiResponseStruct struct {
	Status     string
	Result     string
	StatusCode int
}
type ApiRequest struct {
	amount    float64
	timestamp time.Time
}

var statResp ApiResponseStruct

var (
	sum     float64
	avg     float64
	max     float64
	min     float64
	count   int
	reqTime time.Time
)

func main() {
	reqTime = time.Now()
	http.HandleFunc("/transactions", Transactions)
	http.HandleFunc("/statistics", Statistics)
	http.HandleFunc("/delete", Delete)

	er := http.ListenAndServe(net.JoinHostPort("localhost", "9090"), nil)
	if er != nil {
		fmt.Println("Error initializing Http rest Api" + er.Error())
	}

	fmt.Println("Client initialized successfully")

}

func Transactions(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed, For Deleting use /Delete in url instead /transaction"))
		fmt.Println("error in request method ", request.Method)
		return
	}
	log.Println("Body in post method for transaction is ", request.Body)
	SubmitTransaction(request.Body, writer)

	//fmt.Println("Before writing ",jsonResp)
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, _ = writer.Write([]byte("transaction proceeded successfully"))
	fmt.Println("Transaction Proceeded successfully")
	return
}

func SubmitTransaction(request io.ReadCloser, writer http.ResponseWriter) {

	bytes, err := ioutil.ReadAll(request)
	defer request.Close()
	if err != nil {
		fmt.Println("erron in reading ", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid Json"))
		return
	}
	var req ApiRequest
	var jsonMap map[string]interface{}
	err = json.Unmarshal(bytes, &jsonMap)
	if err != nil {
		fmt.Println("error while marshaling json ")
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte("fields are not parsable: marshaling failed"))
		return
	}

	req.amount, err = strconv.ParseFloat(jsonMap["amount"].(string), 64)
	req.timestamp, err = time.Parse("2021-07-17T09:59:51.312Z", jsonMap["timestamp"].(string))
	if err != nil {
		fmt.Println("Error in timestamp conversion", err)
	}
	fmt.Println("Data after Marshaling ", req)
	//Todo write code here to store data

	sum = sum + req.amount
	count++
	avg = sum / (float64(count))
	if max < req.amount || count==1{
		max = req.amount
	}
	if min > req.amount || count==1{
		min = req.amount
	}

}

func Statistics(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = writer.Write([]byte("Method not allowed"))
		log.Println("Error in Request method type, for Statistics use GET method", request.Method)
		return
	}
	log.Println("Get Params were: ", request.URL.Query())
	var resp = make(map[string]string)
	//resp["sum"] = strconv.FormatFloat(sum, 'E', -1, 32)
	resp["sum"] = fmt.Sprintf("%f", sum)
	resp["avg"] = fmt.Sprintf("%f", avg)
	resp["max"] = fmt.Sprintf("%f", max)
	resp["min"] = fmt.Sprintf("%f", min)
	resp["count"] = fmt.Sprintf("%d",count)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error while converting map to response json format")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Status Internal Server Error"))
		return
	}
	fmt.Println("jsonResponse is ", jsonResp)
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResp)
	/*	param1 := request.URL.Query().Get("a")
		param2 := request.URL.Query().Get("b")
		fmt.Println("Received param are ", param1, " ", param2)
		//u can push into channel here
		//get response from server and write output using below.
		_, _ = writer.Write([]byte("Result is "))*/

	return
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "DELETE" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
		fmt.Println("error in request menthod ", request.Method)
		return
	}

	sum=0
	avg=0
	min=0
	max=0
	count=0
//	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}

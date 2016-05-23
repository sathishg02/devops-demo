package main

import (
	"net/http"
	"time"
	"encoding/json"
	"strconv"
)

func IngestContent(w http.ResponseWriter, r *http.Request) {

	log("Ingestion Service request received")
	response := make(map[string]string)
	response["status"] = "success"
	response["code"] = "200"

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	time.Sleep(100)

	w.Write(js)
	log("Completed service")
}

func log(data string) {
	resctime := time.Now().UnixNano() / int64(time.Millisecond)
	println(time.Now().String() + " [" + strconv.FormatInt(resctime, 10) + "] " + data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ingest", IngestContent)
	log("Starting Ingestion Service")
	http.ListenAndServe(":8090", mux)
}

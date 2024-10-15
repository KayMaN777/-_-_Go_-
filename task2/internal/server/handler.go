package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"task2/internal/model"
	"time"
)

const version = "v1.0.0"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, version)
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	var req model.DecodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(req.InputString)
	if err != nil {
		http.Error(w, "invalid base64 string", http.StatusBadRequest)
		return
	}

	res := model.DecodeResponse{
		OutputString: string(decoded),
	}
	json.NewEncoder(w).Encode(res)
}

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	sleepTime := rand.Intn(11) + 10 // from 10 to 20 seconds
	time.Sleep(time.Duration(sleepTime) * time.Second)

	if rand.Intn(2) == 0 {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Operation completed successfully")
	}
}

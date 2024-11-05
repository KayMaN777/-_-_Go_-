package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const url = "http://localhost:8080"

func main() {
	client := &http.Client{}

	// Запрос на версию
	resp, err := client.Get(url + "/version")
	if err != nil {
		log.Fatalf("Failed to get version: %v", err)
	}
	defer resp.Body.Close()
	version, _ := io.ReadAll(resp.Body)
	fmt.Println(string(version))

	// Запрос на decode
	decodeReq := `{"inputString": "U29mdHdhcmUgRW5naW5lZXJpbmc="}`
	resp, err = client.Post(url+"/decode", "application/json", io.NopCloser(bytes.NewBufferString(decodeReq)))
	if err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}
	defer resp.Body.Close()
	decodeRes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(decodeRes))

	// Запрос на hard-op с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url+"/hard-op", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Request canceled due to timeout or other error:", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Request successful with status:", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response:", string(body))
	}
}

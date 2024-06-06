package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type Request struct {
	Array  []int `json:"array"`
	Target int   `json:"target"`
}

type Response struct {
	Result int `json:"result"`
}

type Data struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func ternarySearch(l, r, x int, ar []int) int {
	if r >= l {
		mid1 := l + (r-l)/3
		mid2 := r - (r-l)/3
		if ar[mid1] == x {
			return mid1
		}
		if ar[mid2] == x {
			return mid2
		}
		if x < ar[mid1] {
			return ternarySearch(l, mid1-1, x, ar)
		} else if x > ar[mid2] {
			return ternarySearch(mid2+1, r, x, ar)
		} else {
			return ternarySearch(mid1+1, mid2-1, x, ar)
		}
	}
	return -1
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON-данных из тела запроса
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := ternarySearch(0, len(req.Array)-1, req.Target, req.Array)
	resp := Response{Result: result}
	data := Data{
		Request:  req,
		Response: resp,
	}

	var dataArray []Data
	dataFile, err := os.Open("result.json")
	if err == nil {
		err = json.NewDecoder(dataFile).Decode(&dataArray)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	dataFile.Close()

	dataArray = append(dataArray, data)

	file, _ := json.MarshalIndent(dataArray, "", " ")
	_ = os.WriteFile("result.json", file, 0644)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/search", searchHandler)
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Server is listening on port 8080")

}

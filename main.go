package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handleMemeSearch(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("jmhReJfMFvoKM6psA2wfNUYMc6DRot4Z") 
	searchQuery := r.URL.Query().Get("q")

	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("http://api.giphy.com/v1/gifs/search?q=%s&api_key=%s", searchQuery, apiKey)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch memes", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read meme data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/search-memes", handleMemeSearch)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is running at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

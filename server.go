package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/go-chi/chi"
)



var data = map[string]string{}

func main() {
	// Get port from env variables or set to 8080.
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Get("/key/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
	
		data, err := Get(r.Context(), key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			JSON(w, map[string]string{"error": err.Error()})
			return
		}
	
		w.Write([]byte(data))
	})

	r.Delete("/key/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
	
		err := Delete(r.Context(), key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			JSON(w, map[string]string{"error": err.Error()})
			return
		}
	
		JSON(w, map[string]string{"status": "success"})
	})

	r.Post("/key/{key}", func (w http/ResponseWriter, r *http.Request){
		key := chi.URLParam(r, "key")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			JSON(w, map[string]string{"error": err.Error()})
			return
		}


		err = Set(r.Content(), key, string(body))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			JSON(w, map[string]string{"error": err.Error()})
			return
		}
	JSON(w, map[string]string{"status", "success"})
	})
	log.Fatal(http.ListenAndServe(":"+port, r))
}


func JSON (w http.ResponseWriter, data interface {}) {
	w.Header().Set("Content-Type", "application/json: charset=utf-8")
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		JSON(w , map[string]string{"error": err.Error()})
		return
	}
	w.Write(b)
}


func Set(ctx context.Context, key, value string) error {
	data[key] = value
	return nil
}

func Get(ctx context.Context, key string) (string, error) {
	return data[key], nil
}

func Delete(ctx context.Context, key string) error {
	delete(data, key)

	return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)
var cacheMutex sync.RWMutex

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/users", handleUsers)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Extract path after "/users"
	path := strings.TrimPrefix(r.URL.Path, "/users")

	switch r.Method {
	case http.MethodPost:
		if path == "" { // POST /users
			postUser(w, r)
			return
		}
	case http.MethodGet:
		if strings.HasPrefix(path, "/") { // GET /users/{id}
			id, err := strconv.Atoi(strings.TrimPrefix(path, "/"))
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			getUser(w, r, id)
			return
		}
	case http.MethodDelete:
		if strings.HasPrefix(path, "/") { // DELETE /users/{id}
			id, err := strconv.Atoi(strings.TrimPrefix(path, "/"))
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			deleteUser(w, r, id)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if _, ok := userCache[id]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(userCache, id)
	w.WriteHeader(http.StatusNoContent)
}

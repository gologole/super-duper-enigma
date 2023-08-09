package htp

import (
	"encoding/json"
	"fmt"
	service "forummodule/internal/service"
	"log"
	"net/http"
)

//запуск через горутину

func Handlereg() {
	fmt.Println("server is working")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/registration.html")
	})
	http.HandleFunc("/handleClick", handleJSONRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("", err)
	}

}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	// Check the request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Decode the JSON packet from the request body
	var data service.AccountData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	loginresult := service.LogIn(data)
	var response map[string]interface{}
	if loginresult == true {
		response = map[string]interface{}{
			"account LogIN": "successfully",
		}
	} else {
		response = map[string]interface{}{
			"account LogIN": " not successfully",
		}
	}
	// Send a response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

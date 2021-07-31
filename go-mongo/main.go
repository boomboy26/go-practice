package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main (){
	fmt.Println("Hello vue go mongo")
	r := mux.NewRouter()
	
	r.HandleFunc("/api/expense",getExpense).Methods("GET")
	r.HandleFunc("/api/expense/create",createExpense).Methods("POST")
	r.HandleFunc("/api/expense/update",updateExpense).Methods("PUT")
	r.HandleFunc("/api/expense/delete",deleteExpense).Methods("DELETE")
	// r.HandleFunc("/api/",createExpense).Methods("POST")
	// r.HandleFunc("/api/profile",createExpense).Methods("POST")
	// r.HandleFunc("/api/profile",createExpense).Methods("POST")
	db()

	log.Fatal(http.ListenAndServe(":8000", r))

	
}


// // Get all expense
// func getExpense(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("get")
// }

// // Create expense 

// func createExpense(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("create")
// }

// // Update Expense 
// func updateExpense(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("update")
// }
// // Delete Expense 
// func deleteExpense(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("delete")
// }
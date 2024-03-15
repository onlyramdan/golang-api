package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Ramdan", Age: 20, Division: "Marketing"},
	{ID: 3, Name: "Mailo", Age: 20, Division: "IT"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)

	http.HandleFunc("/employee", createEmplyee)

	fmt.Println("Aplication is listening on port ", PORT)
	http.ListenAndServe(PORT, nil)
}
func getEmployees(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		tpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//json.NewEncoder(w).Encode(employees)

		tpl.Execute(w, employees)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmplyee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

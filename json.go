package main

import (
	"encoding/json"
	"log"
	"os"
)

type Employee struct {
	ID       uint8   `json:"id"`
	Name     string `json:"name"`
	Age      uint8    `json:"age"`
	Position string `json:"position"`
}

func main() {
	f, err := os.OpenFile("employees.json",os.O_RDONLY,0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var jData []Employee
	err = json.NewDecoder(f).Decode(&jData)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range jData {
		if e.ID == 3 {
			e.Age = 50
			break
		}
	}

	
	newE := Employee{ID: 6, Name: "Dilshodbek", Age: 19, Position: "Intern"}
	jData = append(jData, newE)

	updated, err := json.MarshalIndent(jData,""," ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("employees_new.json", updated, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

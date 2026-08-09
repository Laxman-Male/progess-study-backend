package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"progress_tracker/service"
)

// get plan from DB for ther user
func HandleGetPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("-------------------------------------------------------------")
		var id int
		fromdata, err := service.GetPlan(id)
		if err != nil {
			fmt.Println("error in getting plan from db")
		}
		fmt.Println("handler data formated------------------------------------ ", fromdata)

		var rawJSON json.RawMessage

		// / 2. Convert string to raw JSON and check if it's valid
		// JSON → Go	Read JSON into struct ->In-Memory	json.Unmarshal()
		// What Unmarshal does here: When Unmarshal is given a json.RawMessage as the target, it doesn't try to interpret the JSON structure (like finding "title", "introduction", etc.). Instead, it simply reads the entire JSON payload from fromdata and stores those exact bytes directly into rawJSON, but only if the input is valid JSON.

		err = json.Unmarshal([]byte(fromdata), &rawJSON)
		if err != nil {
			http.Error(w, "failed to parse", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// []byte(rawJSON): Since rawJSON is already a []byte (because json.RawMessage is a []byte alias), this conversion is redundant but harmless. It essentially says "take these JSON bytes."

		w.Write([]byte(rawJSON))

	}
}

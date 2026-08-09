package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	// allstruct "progress_tracker/allStruct"
	"progress_tracker/service"
	"strconv"
	"strings"

	// "progress_tracker/service"
	"time"
)

// to get plan
func HandleTimeTable(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// var StudyPlan allstruct.UserAskForPlan
		// err := json.NewDecoder(r.Body).Decode(&StudyPlan)

		subject := r.URL.Query().Get("subject")
		topic := r.URL.Query().Get("topic")

		daysStr := r.URL.Query().Get("days")
		hoursStr := r.URL.Query().Get("hours")

		// Trim any leading/trailing whitespace from string inputs
		subject = strings.TrimSpace(subject)
		topic = strings.TrimSpace(topic)

		// Convert days and hours to integers
		days, err := strconv.Atoi(daysStr)
		if err != nil {
			http.Error(w, "Invalid 'days' parameter: must be a number", http.StatusBadRequest)
			return
		}
		hours, err := strconv.Atoi(hoursStr)
		if err != nil {
			http.Error(w, "Invalid 'hours' parameter: must be a number", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, "invalid request error", http.StatusBadRequest)
			return
		}

		// This is the crucial part:
		// Create a new context with a timeout, derived from the incoming request's context (r.Context()).
		// The 60-second timeout applies to the entire operation triggered by this handler,
		// including the Gemini API call.
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Minute)
		defer cancel()
		// Call the Gemini API function, passing the newly created 'ctx'.

		response, err := service.GenerateStudyPlan(ctx, subject, topic, days, hours)

		if err != nil {
			fmt.Println("error returned form study plan ", err)
			http.Error(w, "failed to generate study plan", http.StatusInternalServerError)
			return
		}
		fmt.Println("response--> in plan->", response)

		// cleaned := strings.TrimSpace(response)
		// cleaned = strings.TrimPrefix(cleaned, "```json")
		// cleaned = strings.TrimSuffix(cleaned, "```")
		// But that string is already valid JSON.

		// So you want to parse it just to confirm it's valid JSON (Unmarshal), and then write it directly.
		// 1. Declare a variable to hold raw JSON

		var rawJSON json.RawMessage
		// var rawJSON map[string]interface{}

		// / 2. Convert string to raw JSON and check if it's valid
		// JSON → Go	Read JSON into struct ->In-Memory	json.Unmarshal()
		err = json.Unmarshal([]byte(response), &rawJSON)
		if err != nil {
			http.Error(w, "failed to parse", http.StatusInternalServerError)
			return
		}

		// 3. Set response headers and send back raw JSON
		// Set the response type to JSON, and write the raw JSON (already validated) directly to the HTTP response.”
		w.Header().Set("Content-Type", "application/json")
		w.Write(rawJSON)

		//not using
		// json.NewEncoder().Encode()
		// because already i am getting JSON data and i dont want to make it again
	}
}

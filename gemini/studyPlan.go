package gemini

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/google/generative-ai-go/genai"
// 	"google.golang.org/api/option"
// )

// func GenerateStudyPlan(ctx context.Context, subject string, days, dailyHour int) (string, error) {
// 	apiKey := os.Getenv("GeminiApi")
// 	if apiKey == "" {
// 		fmt.Println("failed to connect with API")
// 		return "", fmt.Errorf("GEMINI_API_KEY set  environment variable")
// 	}
// 	promtText := fmt.Sprintf("i want to learn %s subject or you can say topic in %d days and want to become good in it, and i can spend %d hours daily for this. Create a proper study plan for this", subject, days, dailyHour)
// 	fmt.Println("prompt text", promtText)

// 	// ctx := context.Background()

// 	// --- 4. Initialize Gemini Client ---
// 	// This sets up the connection to the Gemini API using your API key.
// 	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create gemini client")
// 	}
// 	defer client.Close()

// 	fmt.Println("bwtween client and res")
// 	// --- 5. Select the Model and Set System Instruction
// 	// Choose a suitable Gemini model. "gemini-1.5-flash" is often a good balance of speed and capability.
// 	model := client.GenerativeModel("gemini-1.5-flash")

// 	//this will set the ai model to assume himself  like a expert teacher and give me the response
// 	model.SystemInstruction = &genai.Content{
// 		Parts: []genai.Part{
// 			genai.Text("You are an expert tutor and study planner. " +
// 				"Your task is to create a detailed, actionable study plan for the given subject, duration, and daily hours. " +
// 				"Provide weekly and  daily breakdown, listing key topics, resources, and suggested activities. " +
// 				"Include an introduction and a concluding motivational message. " + "Use clear headings and bullet points for readability."),
// 		},
// 	}

// 	//here use the setted model description with my prompt
// 	response, err := model.GenerateContent(ctx, genai.Text(promtText))
// 	if err != nil {
// 		fmt.Println("failed to get res")
// 		//get the actual error
// 		log.Printf("Error calling model.GenerateContent: %v", err)
// 		return "", fmt.Errorf("failed to get response from  gemini client")
// 	}
// 	fmt.Println(" ai response1 ", response)
// 	fmt.Println(" ai response2 ", response.Candidates)
// 	fmt.Println(" ai response3 ", response.Candidates[0])
// 	fmt.Println(" ai response4 ", response.Candidates[0].Content)
// 	fmt.Println(" ai response 5", response.Candidates[0].Content.Parts)
// 	var responseToStr string
// 	for index, value := range response.Candidates[0].Content.Parts {
// 		fmt.Print("index--", index)
// 		if value, ok := value.(genai.Text); ok {
// 			responseToStr += string(value)
// 		} else {
// 			fmt.Println("error in converting")
// 		}
// 	}
// 	// var r = string(response.Candidates[0].Content.Parts);
// 	// fmt.Println(" ai response 5", response.Candidates[1].Content.Parts)

// 	fmt.Println(" ai response6 ", response.Candidates[0].Content.Role)

// 	return responseToStr, nil
// }

//

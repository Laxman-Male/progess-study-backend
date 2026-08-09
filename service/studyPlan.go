package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"progress_tracker/repo"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateStudyPlan(ctx context.Context, subject string, topic string, days, dailyHour int) (string, error) {
	apiKey := os.Getenv("GeminiApi")
	if apiKey == "" {
		fmt.Println("failed to connect with API")
		return "", fmt.Errorf("GEMINI_API_KEY set  environment variable")
	}
	promtText := " "

	if topic != "" {
		promtText = fmt.Sprintf("I want to learn ONLY the  topic %s in %s as an absolute beginner in %d days, with %d hours per day. DO NOT include advanced or adjacent topics unless necessary for the core understanding. The plan should be focused and simple, with basic examples, visual understanding, daily revision, and resources for beginners.”  Make sure the plan is complete within %d days, not beyond.Limit the number of major learning goals to 1 or 2 per day. FOCUS MORE ON DEPTH, NOT COVERAGE.  Create a proper study plan for this which should executable.", topic, subject, days, dailyHour, days)
		fmt.Println("prompt text in topic", promtText)
	}

	if topic == "" {
		promtText = fmt.Sprintf("I want to learn ONLY  %s subject  as an absolute beginner in %d days, with %d hours per day. DO NOT include advanced or adjacent topics unless necessary for the core understanding. The plan should be focused and simple, with basic examples, visual understanding, daily revision, and resources for beginners.”  Make sure the plan is complete within %d days, not beyond.Limit the number of major learning goals to 1 or 2 per day. FOCUS MORE ON DEPTH, NOT COVERAGE.  Create a proper study plan for this which should executable.", subject, days, dailyHour, days)
		fmt.Println("prompt text in topic------------------------------------------------------------------------------------------------", promtText)
	}
	// if subject != "" {
	// 	promtText = fmt.Sprintf("I want to learn %s subject, Create a structured study plan to master %s topic from basics to advanced level in  %d day, studying %d days daily. Make sure the plan is complete within %d days, not beyond. Create a proper study plan for this which should executable.", subject,topic, days, dailyHour,days)
	// }
	// promtText := fmt.Sprintf("i want to learn %s subject or you can say topic in %d days and want to become Best into it, and i can spend %d hours daily for this. Create a proper study plan for this which should executable.", subject, days, dailyHour)
	// fmt.Println("prompt text", promtText)

	const (
		dailyBreakdown = `
		{
		"title": "string",
		"introduction": "string",
		"overallStrategy":
		 "weeklyBreakdown": [
   		 {
     		 "week": "string",
     		 "topic": ["string"],
     		 "resources": ["string"],
     		 "activities": ["string"],
     		 "dailyBreakdown": [
      		  {
        		  "day": "string",
         		 "description": ["string"],
				  "whyToLearn": ["string"],
         		 "revision": "string"
     	   }
     	 ]
  	  }
 	 ],
	 "forNextTopic":"string",
 	 "finalReview": "string",
	 "further": "string",
 	 "motivationMessage": "string",
	 "mcqs": [
	   {
	     "question": "string",
	     "options": {"A": "string", "B": "string", "C": "string", "D": "string"},
	     "correctAns": "A"
	   }
	 ]
	}`
		weeklyBreakdown = `{
  			"title": "string",
 			 "introduction": "string",
 			 "overallStrategy": "string",
 			 "weeklyBreakdown": [
   				 {
     				 "weekRange": "string",
     				 "topic": ["string"],
					 "description": ["string"],
					 "whyToLearn": ["string"],
    				  "resources": ["string"],
    				  "activities": ["string"]
  				  }
 		 ],
		  "forNextTopic":"string",
 	 "finalReview": "string",
	 "further": "string",
 	 "motivationMessage": "string",
	 "mcqs": [
	   {
	     "question": "string",
	     "options": {"A": "string", "B": "string", "C": "string", "D": "string"},
	     "correctAns": "A"
	   }
	 ]
	}`
		MoreweeklyBreakdown = `{
  			"title": "string",
 			 "introduction": "string",
 			 "overallStrategy": "string",
 			 "weeklyBreakdown": [
   				 {
     				 "weekRange": "string",
     				 "topic": ["string"],
					 "description": ["string"],
					 "whyToLearn": ["string"],
    				  "resources": ["string"],
    				  "activities": ["string"]
  				  }
 		 ],
		  "forNextTopic":"string",
 	 "finalReview": "string",
	 "further": "string",
 	 "motivationMessage": "string",
	 "mcqs": [
	   {
	     "question": "string",
	     "options": {"A": "string", "B": "string", "C": "string", "D": "string"},
	     "correctAns": "A"
	   }
	 ]
	}`
	)

	mcqInstruction := " Additionally, generate exactly 10 multiple-choice questions (MCQs) covering the ENTIRE plan end-to-end (across all days/weeks), to be used as a final test once the learner completes the plan. Each MCQ must have exactly 4 options keyed 'A','B','C','D', and 'correctAns' must be exactly one of 'A','B','C','D'. Questions must range from easy to hard and test conceptual understanding, not rote lookup: avoid questions answerable by pasting them directly into a search engine or another AI, use unique numbers/scenarios rather than standard textbook examples, require multi-step reasoning where relevant, and make incorrect options (distractors) plausible, often reflecting common mistakes or partial reasoning. Put these 10 MCQs in the top-level 'mcqs' array, matching the schema below."

	sysFixInstruction := "You are an expert tutor, study planner and a skilled programmer . Your task is to create a detailed, actionable study plan for the given subject, topic if mentioned , duration, and daily hours. The goal is to help a learner to  master a specific topic within a given timeframe and daily time allocation to a absolute beginner learner . The plan must progress logically from fundamentals to '70%' concepts of it (or related it).  Provide weekly and  daily breakdown, listing key topics and also list why to learn that topic where it can be used. Resources from reputed website and one with less reputed, and suggested activities.  As you are expert tutor ['If student ask for more time than expected then please recommend to focus on next topic with proper message in this case you can give simple text message to focus on next topic in FURTHER JSON  & here you  can skip JSON format ']. DONT ASSUME THAT STUDENT KNOW ANY RELATED TOPIC TAKE HIM FROM BASICS. **IF STUDENT SPECIFIED THE STANDARD IN WHICH HE/SHE IS LEARNING THEN PLEASE PROVIDE ONLY THAT MUCH PLAN DON'T INCLUDE ADVANCE TOPIC OR UPPER CLASS TOPICS**. Include an introduction and a concluding motivational message. " + mcqInstruction
	var finalInstruction string

	if days <= 45 {
		forFiveDays := sysFixInstruction + "the 'day' field MUST be formatted as 'Day X' field (e.g.,Week1 into it 'Day 1', 'Day 2', 'Day 3', etc.) for each individual day  plan, " + "The 'dailyBreakdown' array MUST contain exactly 7 objects for each week. STRICTLY ensure each week has 7 days (Days 1 to Days 7). DO NOT generate more than the  7 days per week" + "Only include topics necessary for understanding the requested topic. Do not introduce unrelated or advanced concepts unless they are prerequisites. STICK STRICTLY TO THE TOPIC UNLESS SPECIFIED BY THE USER/ STUDENT. Provide actual link for recourses array" + dailyBreakdown
		finalInstruction = forFiveDays

		// "The 'dailyBreakdown' array MUST contain exactly 7 objects for each week. STRICTLY ensure each week has 7 days (Day 1 to Day 7). Do NOT generate more or fewer than 7 days per week."

	} else if days > 45 && days < 100 {
		gapFor4days := sysFixInstruction + "the 'day' field MUST be formatted as 'Day X-Y' with difference of 4 days  (e.g., 'Day 1-Day 4', 'Day 5-Day 8', 'Day 9-Day 12', etc.), for each range days there should be a proper plan with listed points" + " Stricly follow the difference of days DO NOT CREATE MORE GAP JUST DISTRIBUTE THE TOPICS. STICK STRICTLY TO THE TOPIC UNLESS SPECIFIED BY THE USER/ STUDENT. Provide actual link for recourses array  " + weeklyBreakdown
		finalInstruction = gapFor4days
	} else if days > 100 {

		gapForWeek := sysFixInstruction + " Provide with Week breakdown like, Week X-Y (eg. week 1- week 2, week 3- week 4) but it should be detailed plan not a only heading, follow the JSON FORMAT AND DON'T SKIP ANY WEEK, **IF THE STUDY PLAN IS COMPLETED OVER WITH ALL PROPER PLAN THEN YOU CAN SKIP THE WEEKLY RANGE JSON AND CAN GO TO FURTHER JSON FORMAT WHERE YOU CAN MENTION, ABOUT SUBJECT STATUS AND WHAT NEED TO DO FURTHER**.  Stricly follow the difference of days DO NOT CREATE MORE GAP JUST DISTRIBUTE THE TOPICS. STICK STRICTLY TO THE TOPIC UNLESS SPECIFIED BY THE USER/ STUDENT.  Provide actual link for recourses array" + MoreweeklyBreakdown
		finalInstruction = gapForWeek
	}
	// ctx := context.Background()

	// --- 4. Initialize Gemini Client ---
	// This sets up the connection to the Gemini API using your API key.
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("failed to create gemini client")
	}
	defer client.Close()

	fmt.Println("bwtween client and res")
	// --- 5. Select the Model and Set System Instruction
	// Choose a suitable Gemini model. "gemini-1.5-flash" is often a good balance of speed and capability.
	model := client.GenerativeModel("gemini-1.5-flash")

	//this will set the ai model to assume himself  like a expert teacher and give me the response
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(finalInstruction),
		},
	}

	//here use the setted model description with my prompt
	response, err := model.GenerateContent(ctx, genai.Text(promtText))
	if err != nil {
		fmt.Println("failed to get res")
		//get the actual error
		log.Printf("Error calling model.GenerateContent: %v", err)
		return "", fmt.Errorf("failed to get response from  gemini client")
	}
	fmt.Println(" ai response1 ", response)
	fmt.Println(" ai response2 ", response.Candidates)
	fmt.Println(" ai response3 ", response.Candidates[0])
	fmt.Println(" ai response4 ", response.Candidates[0].Content)
	fmt.Println(" ai response 5", response.Candidates[0].Content.Parts)
	var responseToStr string

	// responseToStr= DataCleanUp( response)

	for index, value := range response.Candidates[0].Content.Parts {
		fmt.Print("index--", index)
		if value, ok := value.(genai.Text); ok {
			responseToStr += string(value)
		} else {
			fmt.Println("error in converting")
		}
	}
	// var r = string(response.Candidates[0].Content.Parts);
	// fmt.Println(" ai response 5", response.Candidates[1].Content.Parts)

	fmt.Println(" ai response6 ", response.Candidates[0].Content.Role)

	cleaned := strings.TrimSpace(responseToStr)
	cleaned = strings.TrimPrefix(cleaned, "```json")
	cleaned = strings.TrimSuffix(cleaned, "```")
	// err = repo.ExecutePlanQuery(cleaned)
	// if err != nil {
	// 	return " ", nil
	// }
	return cleaned, nil
}

func GetPlan(id int) (string, error) {

	repoReturn, err := repo.GetPlanrepo(id)
	if err != nil {
		return " ", err
	}
	fmt.Println("service reporeturn  ", repoReturn)
	return repoReturn.Jsondata, nil
}

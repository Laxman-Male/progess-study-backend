package routes

import (
	"fmt"
	"net/http"
	"progress_tracker/handler"
	mcqtest "progress_tracker/handler/mcqTest"
)

func SetUpRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Todo API! Visit /todos to see todos."))
	})
	fmt.Println("route call reg")
	r.HandleFunc("/register", handler.HandleRegister)
	r.HandleFunc("/login", handler.HandleLogin)
	r.HandleFunc("/logged-out", handler.HandleLoggedOut)
	// r.HandleFunc("/study-plan", handler.HandleTimeTable)
	r.Handle("/study-plan", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.HandleTimeTable))))
	// r.HandleFunc("/getPlan", handler.HandleGetPlan)
	r.Handle("/savePlan", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.SavePlanOfUser))))
	r.Handle("/getCount", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.GetPlanForProfile))))

	// r.HandleFunc("/savePlan", handler.SavePlanOfUser)
	// r.HandleFunc("/profile", handler.Profile)
	r.Handle("/profile", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.Profile))))
	r.Handle("/OwnPlanDescription", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.GetOwnPlanDescription))))
	r.Handle("/getWeekCount_view_plan", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.GetWeekPlanViewPlan))))
	r.Handle("/week_completed", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(handler.WeekCompleted))))
	r.Handle("/SubmQuestion", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(mcqtest.SubmitQuestion))))
	r.Handle("/firstQ", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(mcqtest.FirstQuestion))))
	r.Handle("/quizReview", handler.CorsMiddleware(handler.JWTAuthMiddleware(http.HandlerFunc(mcqtest.QuizReview))))
	// r.HandleFunc("/firstQ", mcqtest.FirstQuestion)
	// r.HandleFunc("/question", mcqtest.SubmitQuestion)
	// r.HandleFunc("/quiz/start", handler.CreateQuestionsAfterPlanComplete)
	// r.HandleFunc("/api/OwnPlanDescription", handler.GetOwnPlanDescription)
	// r.HandleFunc("/api/isLogin", handler.HandleLoginState)
	return r
}

package allstruct

type DBconnect struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
}
type LoginFromUser struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}
type GetProfileDetails struct {
	// Id    string `db:"id" json:""`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}
type LoginInfo struct {
	Id       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
type RegInfo struct {
	Id       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
type UserAskForPlan struct {
	Subject    string `json:"subject"`
	Topic      string `json:"topic"`
	Days       int    `json:"days"`
	DailyHours int    `json:"hours"`
}
type DailyBreakdown struct {
	Day         string   `json:"day"`
	Description []string `json:"description"`
	WhyToLearn  []string `json:"whyToLearn"`
	Revision    string   `json:"revision"`
}

type Week struct {
	DayRange       string           `json:"week"`
	Topic          []string         `json:"topic"`
	Resources      []string         `json:"resources"`
	Activities     []string         `json:"activities"`
	DailyBreakdown []DailyBreakdown `json:"dailyBreakdown"`
}

type GoStudyPlanWeekGapOne struct {
	Title             string `json:"title"`
	Introduction      string `json:"introduction"`
	OverallStrategy   string `json:"overallStrategy"`
	WeeklyBreakdown   []Week `json:"weeklyBreakdown"`
	ForNextTopic      string `json:"forNextTopic"`
	FinalReview       string `json:"finalReview"`
	Further           string `json:"further"`
	MotivationMessage string `json:"motivationMessage"`
}

type TitleArray struct {
	Title []byte `db:"title"`
}
type TitleAndCount struct {
	Title     string `db:"title" json:"title"`
	WeekCount int    `db:"weekCount" json:"count"`
}

type SaveUserPlanToDB struct {
	WeekCount int                   `json:"count"`
	Title     string                `json:"title"`
	Plan      GoStudyPlanWeekGapOne `json:"userPlan2"`
	MCQs      []GeneratedMCQ        `json:"mcqs"`
}

// GeneratedMCQ is one of the 10 MCQs Gemini returns alongside the study plan
// itself (single call), persisted to mcqtbl once the plan is saved.
type MCQOptions struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
}

type GeneratedMCQ struct {
	Question   string     `json:"question"`
	Options    MCQOptions `json:"options"`
	CorrectAns string     `json:"correctAns"`
}

type TimeDifference struct {
	CreatedAt int64 `db:"CreatedAt"`
	NowTime   int64 `db:"userClickNowTime"`
}

type MCQFormat struct {
	MCQID       int    `db:"mcqID"`
	MCQ         string `db:"mcqs"`
	Options     string `db:"options"`
	IsAttempted bool   `db:"isAttempted"`
}

// MCQCounts is the plan's overall MCQ progress, shown at the top of the quiz screen.
type MCQCounts struct {
	Total     int `db:"total"`
	Attempted int `db:"attempted"`
}

type OptTitle struct {
	Title  string `json:"Title"`
	MCQID  int    `json:"mcqID"`
	Option string `json:"option"`
}

// MCQReview is used for the post-quiz review screen: all 10 MCQs for a plan
// with the user's selected option and whether it was correct.
type MCQReview struct {
	MCQID          int    `db:"mcqID" json:"mcqID"`
	MCQ            string `db:"mcqs" json:"mcq"`
	Options        string `db:"options" json:"options"`
	CorrectAns     string `db:"correctAns" json:"correctAns"`
	OptionSelected string `db:"optionSelected" json:"optionSelected"`
	IsCorrect      bool   `db:"isCorrect" json:"isCorrect"`
}
